package pengikis

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0",
}

func Kikis(urlTarget string, pemilih []string, formatKeluaran string) (map[string][]string, error) {
	hasil := make(map[string][]string)
	var mu sync.Mutex
	var kesalahanPengikisan error

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
	})

	for _, p := range pemilih {
		pemilihLokal := p
		c.OnHTML(pemilihLokal, func(e *colly.HTMLElement) {
			mu.Lock()
			defer mu.Unlock()
			var konten string
			if formatKeluaran == "Teks Bersih Saja" {
				konten = strings.TrimSpace(e.Text)
			} else {
				html, err := goquery.OuterHtml(e.DOM)
				if err != nil {
					kesalahanPengikisan = fmt.Errorf("tidak bisa mendapatkan HTML untuk pemilih %s: %w", pemilihLokal, err)
					return
				}
				konten = strings.TrimSpace(html)
			}

			if konten != "" {
				hasil[pemilihLokal] = append(hasil[pemilihLokal], konten)
			}
		})
	}

	c.OnError(func(r *colly.Response, err error) {
		mu.Lock()
		defer mu.Unlock()
		kesalahanPengikisan = fmt.Errorf("permintaan ke %s gagal dengan status %d: %w", r.Request.URL, r.StatusCode, err)
	})

	if err := c.Visit(urlTarget); err != nil {
		return nil, fmt.Errorf("gagal memulai kunjungan ke %s: %w", urlTarget, err)
	}

	c.Wait()

	return hasil, kesalahanPengikisan
}
