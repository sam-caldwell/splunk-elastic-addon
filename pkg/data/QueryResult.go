package data

// QueryResult - a structure representing a query
type QueryResult struct {
	Metadata MetaData `json:"metadata"`
	Results  any      `json:"data"`
}
