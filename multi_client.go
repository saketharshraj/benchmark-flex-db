package main
  
import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)


func runMultiClientBenchmark(address string, clients int, opsPerClient int) BenchmarkResult {
	start := time.Now()
	var wg sync.WaitGroup
	totalOps := clients * opsPerClient

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Println("Client", id, "failed to connect:", err)
				return
			}
			defer conn.Close()

			reader := bufio.NewReader(conn)
			writer := bufio.NewWriter(conn)
			reader.ReadString('>') // Prompt

			for j := 0; j < opsPerClient; j++ {
				key := "key" + strconv.Itoa(j)
				value := "client" + strconv.Itoa(id)
				command := fmt.Sprintf("SET %s %s\n", key, value)
				writer.WriteString(command)
				writer.Flush()
				reader.ReadString('\n')
			}
		} (i)
	}

	wg.Wait()
	duration := time.Since(start)
	return BenchmarkResult{
		Name:         "Multi-Client Concurrent Benchmark",
		TotalOps:     totalOps,
		Duration:     duration,
		OpsPerSecond: float64(totalOps) / duration.Seconds(),
	}
}
