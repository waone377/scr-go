package utilitas

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SimpanKeJSON(data interface{}, namaBerkas string) error {
	// Memastikan direktori untuk file keluaran ada.
	dir := filepath.Dir(namaBerkas)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	berkas, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(namaBerkas, berkas, 0644)
}
