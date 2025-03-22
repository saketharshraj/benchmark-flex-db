package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

func runSingleClientBenchmark(address string, numOps int) BenchmarkResult {
	start := time.Now()

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return BenchmarkResult{Error: err.Error()}
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	reader.ReadString('>') // Wait for prompt

	for i := 0; i < numOps; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		command := fmt.Sprintf("SET %s %s\n", key, value)
		writer.WriteString(command)
		writer.Flush()
		reader.ReadString('\n') // Read response
	}

	duration := time.Since(start)
	return BenchmarkResult{
		Name:         "Single Client Benchmark",
		TotalOps:     numOps,
		Duration:     duration,
		OpsPerSecond: float64(numOps) / duration.Seconds(),
	}
}

 