package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Time struct {
	Time string `json:"time,omitempty"`
	NanoTime string `json:"nanoTime,omitempty"`
}

func main() {
	log.Print("Начало работы сервиса time-service");

	http.HandleFunc("/time", serverTime)
	http.HandleFunc("/nanotime", serverNanoTime)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func serveTime(w http.ResponseWriter, r *http.Request) {
	log.Print("Вызов функции serveTime()")
	var serverTime Time
	serverTime.Time = time.Now().Format("02 Jan 2006")
	json.NewEncoder(w).Encode(serverTime)
}

func serveNanoTime(w http.ResponseWriter, r *http.Request) {
	log.Print("Вызов функции serveNanoTime()")
	var nanoTime Time
	nanoTime.NanoTime = strconv.FormatInt(time.Now().UnixNano(), 10)
	json.NewEncoder(w).Encode(nanoTime)
}
