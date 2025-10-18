package utilitas

import (
	"encoding/json"
	"os"
)

func SimpanKeJSON(data interface{}, namaBerkas string) error {
	berkas, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(namaBerkas, berkas, 0644)
}
