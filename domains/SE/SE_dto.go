package SE

import (
	"encoding/json"
	"fmt"
	"github.com/soshika/sample-search/logger"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	Index    string   `json:"index"`
	FileName string   `json:"file_name"`
	Data     []string `json:"data"`
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

type SearchEngineReq struct {
	Query  string `json:"query"`
	From   *int   `json:"from"`
	Size   *int   `json:"size"`
	UserId int64  `json:"user_id"`
	Index  string `json:"index"`
}

func (se *SearchEngineReq) Init() {
	zero := 0
	five := 5
	if se.From == nil {
		se.From = &zero
	}
	if se.Size == nil {
		se.Size = &five
	}
}

type SearchEngineRes struct {
	Index       string `json:"_index"`
	Id          string `json:"_id"`
	Version     string `json:"_version"`
	Result      string `json:"_result"`
	SeqNo       int    `json:"_seq_no"`
	PrimaryTerm int    `json:"_primary_term"`
}

type SearchEngineResShards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
}

type SearchEngineResponse struct {
	Hits  SearchEngineResponseHits `json:"hits"`
	Count map[string]interface{}   `json:"count"`
}

type SearchEngineResponseHits struct {
	Hits     []SearchEngineResponseHitsHits `json:"hits"`
	MaxScore float64                        `json:"max_score"`
}

type SearchEngineResponseHitsHits struct {
	Id     string                 `json:"id"`
	Index  string                 `json:"_index"`
	Score  float64                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}
type SearchEngineResponseHitsTotal struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}

type CountResponse struct {
	Count  int64                 `json:"count"`
	Shards SearchEngineResShards `json:"_shards"`
}
