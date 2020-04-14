package excelpicker

import (
	"encoding/csv"
	"os"
)

// PrintAsCsv ...
func PrintAsCsv(headers []string, data [][]string) error {
	writer := csv.NewWriter(os.Stdout)

	if err := writer.Write(headers); err != nil {
		return err
	}
	if err := writer.WriteAll(data); err != nil {
		return err
	}

	return nil
}
