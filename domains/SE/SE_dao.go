package SE

import "fmt"

const (
	queryDSLSearch     = "{\"query\": {\"match\": {\"data.data\": \"%s\"}}, \"size\": %d, \"from\": %d}"
	queryDSLSuggestion = "{\"query\": {\"bool\": {\"should\": [{\"prefix\": {\"Description\": \"%s\"}},{\"prefix\": {\"Title\": \"%s\"}},{\"prefix\": {\"SEOTitle\": \"%s\"}},{\"prefix\": {\"SEODescription\": \"%s\"}},{\"prefix\": {\"AboutMe\": \"%s\"}},{\"prefix\": {\"FullName\": \"%s\" }},{\"prefix\": {\"question\": \"%s\"}},{\"prefix\": {\"abstract\": \"%s\" }}, {\"prefix\": {\"text\": \"%s\" }}, { \"prefix\": {\"breadcrumb_titles\": \"%s\"}},{\"prefix\": {\"description\": \"%s\"}}, {\"prefix\": {\"answer\": \"%s\"}}, {\"prefix\": {\"title\": \"%s\"}}, {\"prefix\": {\"abstract\": \"%s\"}}, {\"prefix\": {\"text\": \"%s\"}}],\"minimum_should_match\": 1}},\"indices_boost\": [{ \"users\": 1.0 },{ \"products\": 1.0 },{ \"light\": 1.0 },{ \"blog\": 1.0 }], \"from\": %d, \"size\": %d}"
	queryDSLActivity   = "{\"size\": 5, \"query\": {\"match\": {\"user_id\": %d }},\"_source\": [\"_seq_no\", \"_primary_term\", \"query\"], \"sort\": [{\"_seq_no\": {\"order\": \"desc\"}}]}"
	queryDSLCount      = "{\"query\": {\"bool\": {\"should\": [{\"prefix\": {\"Description\": \"%s\"}},{\"prefix\": {\"Title\": \"%s\"}},{\"prefix\": {\"SEOTitle\": \"%s\"}},{\"prefix\": {\"SEODescription\": \"%s\"}},{\"prefix\": {\"AboutMe\": \"%s\"}},{\"prefix\": {\"FullName\": \"%s\" }},{\"prefix\": {\"question\": \"%s\"}},{\"prefix\": {\"abstract\": \"%s\" }}, {\"prefix\": {\"text\": \"%s\" }}, { \"prefix\": {\"breadcrumb_titles\": \"%s\"}},{\"prefix\": {\"description\": \"%s\"}}, {\"prefix\": {\"answer\": \"%s\"}}, {\"prefix\": {\"title\": \"%s\"}}, {\"prefix\": {\"abstract\": \"%s\"}}, {\"prefix\": {\"text\": \"%s\"}}],\"minimum_should_match\": 1}}}"
)

func (se *SearchEngineReq) GenerateSearchQuery() (*string, error) {

	queryDSL := fmt.Sprintf(queryDSLSearch, se.Query, *se.Size, *se.From)

	return &queryDSL, nil
}

func (se *SearchEngineReq) GenerateSuggestionQuery() (*string, error) {
	queryDSL := fmt.Sprintf(queryDSLSuggestion, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, *se.From, *se.Size)

	return &queryDSL, nil
}

func (se *SearchEngineReq) GenerateQueryDSLCount() (*string, error) {
	queryDSL := fmt.Sprintf(queryDSLCount, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query, se.Query)
	return &queryDSL, nil
}
