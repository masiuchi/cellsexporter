package excelpicker

import (
	"encoding/json"
	"os"
)

// PrintAsJSON ...
func PrintAsJSON(headers []string, data [][]string) error {
	jsonData := makeJSONData(headers, data)
	encoder := json.NewEncoder(os.Stdout)
	for _, d := range jsonData {
		if err := encoder.Encode(d); err != nil {
			return err
		}
	}
	return nil
}

func makeJSONData(headers []string, data [][]string) []map[string]string {
	jsonData := make([]map[string]string, len(data))
	for i, d := range data {
		mapData := map[string]string{}
		for i, header := range headers {
			mapData[header] = d[i]
		}
		jsonData[i] = mapData
	}
	return jsonData
}
