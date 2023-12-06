package dto

type MatchStatus string

const (
	Finished    MatchStatus = "finished"
	NotStarted  MatchStatus = "not_started"
	HalfTime    MatchStatus = "half_time"
	Ongoing     MatchStatus = "ongoing"
	Penalties   MatchStatus = "penalties"
	Interrupted MatchStatus = "interrupted"
)

type MatchDTO struct {
	ID       string      `json:"id"`
	Elapsed  int         `json:"elapsed"`
	Status   MatchStatus `json:"status"`
	HomeTeam TeamDTO     `json:"homeTeam"`
	AwayTeam TeamDTO     `json:"awayTeam"`
	Stadium  StadiumDTO  `json:"stadium"`
}

type TeamDTO struct {
	Name  string `json:"name"`
	Logo  string `json:"logo"`
	Score int    `json:"score"`
}

type StadiumDTO struct {
	Name string `json:"name"`
	City string `json:"city"`
}
