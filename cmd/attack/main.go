package main

// export GRAFANA_TOKEN=token && \
// export GRAFANA_HOST=http://127.0.0.1:3000/api/ds/query && \
// export ATTACK_FREQ=100 && \
// go run cmd/attack/main.go

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

const (
	ENV_GRAFANA_TOKEN = "GRAFANA_TOKEN"
	ENV_GRAFANA_HOST  = "GRAFANA_HOST"
	ENV_ATTACK_FREQ   = "ATTACK_FREQ"
	FILE_NAME         = "request.json"
)

func main() {
	// set header
	header := http.Header{}
	header.Set("Authorization", "Bearer "+os.Getenv(ENV_GRAFANA_TOKEN))
	header.Set("Content-Type", "application/json")

	// get json
	content, err := os.ReadFile(FILE_NAME)
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(os.Getenv(ENV_ATTACK_FREQ))
	if err != nil {
		panic(err)
	}

	rate := vegeta.Rate{Freq: num, Per: time.Second}
	duration := 3 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    os.Getenv(ENV_GRAFANA_HOST),
		Header: header,
		Body:   content,
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Grafana attack testing") {
		metrics.Add(res)
	}
	metrics.Close()
	fmt.Printf("%v\n", metrics.StatusCodes)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("maximum observed request latency: %s\n", metrics.Latencies.Max)
}
