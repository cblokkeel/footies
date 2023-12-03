package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cblokkeel/footies/service"
)

type Api struct {
	matchService *service.MatchService
}

func NewApi(matchService *service.MatchService) *Api {
	return &Api{
		matchService,
	}
}

func (a *Api) Start() {
	http.HandleFunc("/matchs", a.handleGetMatchs)
	http.HandleFunc("/match", a.handleGetMatch)
	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func (a *Api) handleGetMatchs(w http.ResponseWriter, r *http.Request) {
	leagueID := r.URL.Query().Get("league")
	season := r.URL.Query().Get("season")
	date := r.URL.Query().Get("date")

	matchs, err := a.matchService.GetMatchByLeague(r.Context(), date, leagueID, season)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matchs)
}

func (a *Api) handleGetMatch(w http.ResponseWriter, r *http.Request) {
	matchID := r.URL.Query().Get("id")

	match, err := a.matchService.GetMatchByID(r.Context(), matchID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}
