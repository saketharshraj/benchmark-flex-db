package main

import (
	"fmt"
)

func main() {
	address := "localhost:9000"

	// Configurable load values
	singleOps := 1000
	multiClients := 10
	opsPerClient := 100

	// Run benchmarks
	single := runSingleClientBenchmark(address, singleOps)
	multi := runMultiClientBenchmark(address, multiClients, opsPerClient)

	summary := []BenchmarkResult{single, multi}

	// Print summary
	for _, result := range summary {
		fmt.Printf("%s\nTotal Ops: %d\nDuration: %s\nOps/sec: %.2f\n\n",
			result.Name, result.TotalOps, result.Duration, result.OpsPerSecond)
	}

	// Generate file name: <singleOps>_<multiClients>_<opsPerClient>.json
	filename := fmt.Sprintf("benchmark_report_%d_%d_%d.json", singleOps, multiClients, opsPerClient)

	// Save report
	err := writeReport(summary, filename)
	if err != nil {
		fmt.Println("Failed to write report:", err)
	} else {
		fmt.Printf("Benchmark report saved to %s\n", filename)
	}
}
