package main

import (
	"yadro_golang_task/config"
	"bufio"
    //"fmt"
    "log"
    "os"
    //"strings"
	"yadro_golang_task/model"
	//"yadro_golang_task/handler"
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

    requests := make([]*model.Request, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        req, err := parser.ParseLine(scanner.Text())
        if err != nil {
            log.Fatalf("error : %v", err)
        }
        requests = append(requests, req)
        //handler.Dispatch(req)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}