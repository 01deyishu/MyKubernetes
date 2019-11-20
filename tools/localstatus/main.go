package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type RES struct {
	Host string `json:"host"`
	Time string `json:"time"`
}

type MRES struct {
	Res     RES
	Env     string `json:"env"`
	Version string `json:"version"`
}

func main() {
	http.HandleFunc("/status", statusRespons)
	http.HandleFunc("/more", moreResponse)
	http.ListenAndServe(":8080", nil)
}

func moreResponse(w http.ResponseWriter, req *http.Request) {
	var mRes MRES
	mRes.Res = getLocal()
	mRes.Version = "v0.2.0"
	mRes.Env = os.Getenv("GOMAXPROC")
	r, _ := json.Marshal(mRes)
	w.Write(r)
}

func statusRespons(w http.ResponseWriter, req *http.Request) {
	r, _ := json.Marshal(getLocal())
	w.Write(r)
}

func getLocal() (localstatus RES) {
	localstatus.Host, _ = os.Hostname()
	localstatus.Time = time.Now().String()
	return
}
