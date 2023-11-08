package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soshika/sample-search/app"
)

func main() {
	// TODO: should change to release mode!
	gin.SetMode(gin.DebugMode)
	app.StartApplication()
}

//func main() {
//	f, err := excelize.OpenFile("test.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer func() {
//		if err := f.Close(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//
//	// could have multiple sheets
//	sheets := f.GetSheetList()
//	for _, sheetName := range sheets {
//		d, err := f.GetRows(sheetName)
//		if err != nil {
//			fmt.Println("error reading sheet", sheetName, ":", err)
//			return
//		}
//
//		saveAsJSON(d, sheetName+".json")
//		saveAsJSONWithHeaders(d, sheetName+"_with_headers.json")
//	}
//
//}
//
//func saveAsJSONWithHeaders(rows [][]string, filename string) error {
//	data := make([]map[string]string, len(rows)-1)
//	headers := rows[0]
//	// excluding header row
//	for i, row := range rows[1:] {
//		data[i] = make(map[string]string)
//		for j, cellValue := range row {
//			data[i][headers[j]] = cellValue
//		}
//	}
//
//	return saveAsJSON(data, filename)
//}
//
//func saveAsJSON(data interface{}, filename string) error {
//	file, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	encoder := json.NewEncoder(file)
//	if err := encoder.Encode(data); err != nil {
//		return err
//	}
//	return nil
//}
