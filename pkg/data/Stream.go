package data

// Stream - XML structure for Splunk input
type Stream struct {
	Items []Item `xml:"item"`
}
