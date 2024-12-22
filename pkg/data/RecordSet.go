package data

// RecordSet - A record set representation for elastic query results with tracing
type RecordSet struct {
	ItemId  int
	BatchId int
	Hit     any
}
