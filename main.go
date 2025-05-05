package main

import (
	"yadro_golang_task/config"
	"bufio"
    //"fmt"
    "log"
    "os"
    //"strings"
	//"yadro_golang_task/api/model"
	"yadro_golang_task/api/handlers"
	"yadro_golang_task/parser"
)

const (
	CONFIG_PATH = "./config/config.json"
	EVENTS_PATH = "./sunny_5_skiers/events"
)

func main() {
    config.LoadConfig(CONFIG_PATH)

	file, err := os.Open(EVENTS_PATH)
    if err != nil {
        log.Fatalf("error : %v", err)
    }
    defer file.Close()

    requests := make([]*biathlon.Request, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        req, err := parser.ParseLine(scanner.Text())
        if err != nil {
            log.Fatalf("error : %v", err)
        }
        requests = append(requests, req)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Обработка запросов
    handler := biathlon.NewBiathlonHandler()
    for _, req := range requests {
        biathlon.Dispatch(req, handler)
    }
}