package main

import (
	"yadro_golang_task/config"
	"bufio"
    //"fmt"
    "log"
    "os"
    //"strings"
	"yadro_golang_task/api/model"
    managers "yadro_golang_task/managers"
	"yadro_golang_task/api/handlers"
	"yadro_golang_task/parser"
)

const (
	CONFIG_PATH = "./config/config.json"
	EVENTS_PATH = "./sunny_5_skiers/events"
)

func main() {
    cfg := config.LoadConfig(CONFIG_PATH)

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
    var competitors = make(map[uint32]*model.Competitor)

    manager := managers.NewBiathlonManager(competitors, cfg)
    handler := biathlon.NewBiathlonHandler(manager)
    for _, req := range requests {
        biathlon.Dispatch(req, handler)
    }
    status := handler.HandleGetStatus()
    if status.Err != nil {
        log.Fatal(status.DeveloperMessage)
    } else {
        log.Println(status.UserMessage)
    }
}