package main

import (
       "fmt"
       "log"
	   "net/http"
	   "encoding/json"
)

type Track struct {
	TrackNum string `json:"TrackNumber"`
	Id string `json:"Id"`
	Lat string `json:"Lat"`
	Lon string `json:"Lon"`
}

var Tracks []Track

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/tracks",getTracks)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func getTracks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getTracks")
	json.NewEncoder(w).Encode(Tracks)
}

func main() {
	Tracks = []Track{
		Track{TrackNum: "1", Id: "Friend", Lat: "37.0", Lon: "-74.0"},
	}
	handleRequests()
}