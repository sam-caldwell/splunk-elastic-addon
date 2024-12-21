package main

import (
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/data"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/elastic"
	"github.com/sam-caldwell/splunk-elastic-addon/pkg/input"
	"log"
	"os"
	"sync"
)

const (
	// maxInputSize - this is the maximum size of a splunk query we will consume
	maxInputSize = 5 * 1048576 // 5MB
)

var traceId uuid.UUID

// init - initialize logging
func init() {
	// traceId - a UUID representing the current runtime used to track executions for logging.
	traceId = uuid.New()
	log.SetOutput(os.Stderr)
	log.SetPrefix(fmt.Sprintf("[traceId:%s]", traceId.String()))
}

func main() {
	var (
		err                error
		rawInput           []byte
		stream             data.Stream
		streamWorkingGroup sync.WaitGroup
	)

	// Splunk expects to pass all inputs via the stdin device as an XML structure
	if rawInput, err = input.ReadStdin(maxInputSize); err != nil {
		log.Fatalf("Failed to read input or exceeded size limit: %v", err)
	}

	// Given a raw input from stdin, we must unmarshal the XML
	if err = xml.Unmarshal(rawInput, &stream); err != nil {
		log.Fatalf("Failed to parse XML input: %v", err)
	}

	// For every Elastic 'item' representing a query, spawn a go routine to run the query asynchronously.
	for itemId, item := range stream.Items {
		streamWorkingGroup.Add(1)
		go elastic.ProcessItem(traceId, itemId, item, &streamWorkingGroup)
	}

	// Wait until all go routines are finished processing the queries.
	streamWorkingGroup.Wait()
}
