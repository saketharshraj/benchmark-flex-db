package main

import (
	"encoding/json"
	"os"
	"time"
)

type BenchmarkResult struct {
	Name         string
	TotalOps     int
	Duration     time.Duration
	OpsPerSecond float64
	Error        string
}

type BenchmarkConfig struct {
	SingleOps    int `json:"single_ops"`
	MultiClients int `json:"multi_clients"`
	OpsPerClient int `json:"ops_per_client"`
}

func loadConfig(path string) (BenchmarkConfig, error) {
	var config BenchmarkConfig
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
