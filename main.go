package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func getXPByPlayerID(w http.ResponseWriter, r *http.Request) {
	// request: GET /xp/:playerID

	w.Header().Set("Content-Type", "application/json")

	s := strings.Split(r.URL.Path, "/")
	playerId := s[len(s)-1]

	response := map[string]interface{}{
		"PlayerId": playerId,
		"XP":       1,
	}

	json.NewEncoder(w).Encode(response)

	return
}

func main() {
	http.HandleFunc("/xp/", getXPByPlayerID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
