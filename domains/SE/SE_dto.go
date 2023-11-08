package SE

import (
	"encoding/json"
	"fmt"
	"github.com/soshika/sample-search/logger"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	Index    string                 `json:"index"`
	FileName string                 `json:"file_name"`
	Data     map[string]interface{} `json:"data"`
}

func (excel *Excel) Save() ([]string, error) {
	f, err := excelize.OpenFile(excel.FileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// could have multiple sheets
	sheets := f.GetSheetList()
	final := []string{}
	for _, sheetName := range sheets {
		d, err := f.GetRows(sheetName)
		if err != nil {
			logger.Error(fmt.Sprintf("error reading sheet %s", sheetName), err)
			return nil, err
		}

		result := saveAsJSONWithHeaders(d, sheetName+"_with_headers.json")
		for i := 0; i < len(result); i++ {
			jsonString, _ := json.Marshal(result[i])
			final = append(final, string(jsonString))
		}

	}
	return final, nil
}

func saveAsJSONWithHeaders(rows [][]string, filename string) []map[string]string {
	data := make([]map[string]string, len(rows)-1)
	headers := rows[0]
	// excluding header row
	for i, row := range rows[1:] {
		data[i] = make(map[string]string)
		for j, cellValue := range row {
			data[i][headers[j]] = cellValue
		}
	}

	return data
}
