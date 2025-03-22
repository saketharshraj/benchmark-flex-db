package main

import (
	"fmt"
)

func main() {
    address := "localhost:9000"

    config, err := loadConfig("benchmark_config.json")
    if err != nil {
        fmt.Println("Failed to load config:", err)
        return
    }

    single := runSingleClientBenchmark(address, config.SingleOps)
    multi := runMultiClientBenchmark(address, config.MultiClients, config.OpsPerClient)

    summary := []BenchmarkResult{single, multi}

    for _, result := range summary {
        fmt.Printf("%s\nTotal Ops: %d\nDuration: %s\nOps/sec: %.2f\n\n",
            result.Name, result.TotalOps, result.Duration, result.OpsPerSecond)
    }

    filename := fmt.Sprintf("benchmark_report_%d_%d_%d.json",
        config.SingleOps, config.MultiClients, config.OpsPerClient)

    err = writeReport(summary, filename)
    if err != nil {
        fmt.Println("Failed to write report:", err)
    } else {
        fmt.Printf("Benchmark report saved to %s\n", filename)
    }
}

