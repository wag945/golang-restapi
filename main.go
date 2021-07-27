package main

import (
       "fmt"
       "log"
	   "net/http"
	   "encoding/json"
)

type Track struct {
	TrackNum string `json:"TrackNumber"`
	Category string `json:"Category"`
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
	fmt.Println("Endpoint Hit: tracks")
	json.NewEncoder(w).Encode(Tracks)
}

func main() {
	Tracks = []Track{
		Track{TrackNum: "1", Category: "Surface", Id: "Friend", Lat: "37.0", Lon: "-74.0"},
		Track{TrackNum: "2", Category: "Air", Id: "Unknown", Lat: "37.5", Lon: "-74.3"},
		Track{TrackNum: "3", Category: "Air", Id: "Unknown", Lat: "36.2", Lon: "-74.9"},
	}
	handleRequests()
}