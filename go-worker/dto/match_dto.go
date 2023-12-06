package dto

type MatchStatus int

const (
	Finished MatchStatus = iota
	NotStarted
	HalfTime
	Ongoing
	Penalties
	Interrupted
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
