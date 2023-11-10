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
	Search(*SE.SearchEngineReq) (*SE.SearchEngineResponse, error)
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

func (s *seService) Search(req *SE.SearchEngineReq) (*SE.SearchEngineResponse, error) {
	logger.Info("Enter to Search service successfully")

	// just to make sure
	// we fill the $from & $size with proper values.
	req.Init()

	// now we need to generate elasticsearch query
	query, err := req.GenerateSearchQuery()
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}

	// Unmarshal the JSON string into the map
	err = json.Unmarshal([]byte(*query), &body)
	if err != nil {
		return nil, err
	}

	esREQ := requests.Req{
		URL:    "http://localhost:9200/excel/_search",
		Method: "POST",
	}

	resBody, err := esREQ.POST(body)
	if err != nil {
		return nil, err
	}

	ret := SE.SearchEngineResponse{}
	if marshalErr := json.Unmarshal(resBody, &ret); marshalErr != nil {
		return nil, marshalErr
	}

	//count := make(map[string]interface{})
	//count["users"], _ = s.GetCount(req, "users")
	//ret.Count = count

	logger.Info("Close from Search service successfully")
	return &ret, nil
}

func (s *seService) Count(req *SE.SearchEngineReq, index string) (*int64, error) {
	logger.Info("Enter to GetUserCount service successfully")

	// now we need to generate elasticsearch query
	query, err := req.GenerateQueryDSLCount()
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}

	// Unmarshal the JSON string into the map
	err = json.Unmarshal([]byte(*query), &body)
	if err != nil {
		return nil, err
	}

	esREQ := requests.Req{
		URL:    fmt.Sprintf("http://localhost:9200/%s/_search", index),
		Method: "POST",
	}

	resBody, err := esREQ.POST(body)
	if err != nil {
		return nil, err
	}

	ret := SE.CountResponse{}
	if marshalErr := json.Unmarshal(resBody, &ret); marshalErr != nil {
		return nil, marshalErr
	}

	fmt.Println(ret)

	logger.Info("Close from GetUserCount service successfully")
	return &ret.Count, nil
}
