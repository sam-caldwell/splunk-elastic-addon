package data

import "github.com/google/uuid"

// MetaData - Observability data used in debugging a query result
type MetaData struct {
	Time     int64     `json:"time"`
	TraceId  uuid.UUID `json:"traceId"`
	ItemId   int       `json:"itemId"`
	BatchId  int       `json:"batchId"`
	ResultId int       `json:"resultId"`
	WorkerId int       `json:"workerId"`
}
