package main

import "time"

type BenchmarkResult struct {
	Name         string
	TotalOps     int
	Duration     time.Duration
	OpsPerSecond float64
	Error        string
}
