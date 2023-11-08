package services

import (
	"encoding/json"
	"fmt"
	"github.com/soshika/sample-search/domains/SE"
	"github.com/soshika/sample-search/logger"
	"github.com/soshika/sample-search/utils/requests"
)

var (
	SEService SEServiceInterface = &seService{}
)

type seService struct{}
type SEServiceInterface interface {
	IndexExcel(*SE.Excel) (map[string]interface{}, error)
}

func (s *seService) StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}

func (s *seService) IndexExcel(req *SE.Excel) (map[string]interface{}, error) {
	logger.Info("Enter to IndexExcel service successfully")

	result, err := req.Save()
	if err != nil {
		return nil, err
	}

	final := make(map[string]interface{})
	final["data"] = result

	req.Data = final

	ESReq := requests.Req{
		URL:    fmt.Sprintf("http://localhost:9200/%s/_doc", req.Index),
		Method: "POST",
	}

	reqMap, _ := s.StructToMap(req)

	body, err := ESReq.POST(reqMap)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	if marshalErr := json.Unmarshal(body, &ret); marshalErr != nil {
		logger.Error("error when trying to marshal response to struct", marshalErr)
		return nil, marshalErr
	}

	logger.Info("Close from IndexExcel service successfully")
	return ret, nil
}

func (s *seService) Search() (map[string]interface{}, error) {
	logger.Info("Enter to Search service successfully")
	logger.Info("Close from Search service successfully")
	return nil, nil
}
