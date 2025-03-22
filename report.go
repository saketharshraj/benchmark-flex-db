package main

import (
	"encoding/json"
	"os"
	"time"
)

func writeReport(results []BenchmarkResult, filename string) error {
	report := struct {
		Timestamp string            `json:"timestamp"`
		Results   []BenchmarkResult `json:"results"`
	}{
		Timestamp: time.Now().Format(time.RFC3339),
		Results:   results,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(report)
}

