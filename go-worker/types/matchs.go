package types

import "github.com/cblokkeel/footies/dto"

type Match struct {
	Informations     *Fixture          `json:"fixture"`
	League           *League           `json:"league"`
	TeamInformations *TeamInformations `json:"teams"`
	Score            *Score            `json:"goals"`
}

type Fixture struct {
	ID     string  `json:"id"`
	Venue  *Venue  `json:"venue"`
	Status *Status `json:"status"`
}

type TeamInformations struct {
	Home *Team `json:"home"`
	Away *Team `json:"away"`
}

type Team struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Logo           string  `json:"logo"`
	Winner         bool    `json:"winner"`
	WinProbability float32 `json:"winProbability"`
}

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Venue struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed int    `json:"elapsed"`
}

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func (m *Match) getStatus() dto.MatchStatus {
	switch m.Informations.Status.Short {
	case "NS", "TBD":
		return dto.NotStarted
	case "1H", "2H", "ET":
		return dto.Ongoing
	case "HT", "BT":
		return dto.HalfTime
	case "P":
		return dto.Penalties
	case "INT":
		return dto.Interrupted
	case "FT", "AET", "AWD", "CANC", "WO":
		return dto.Finished
	default:
		return dto.NotStarted // maybe unknown type
	}
}

func (m *Match) getTeamInformation(home bool) *dto.TeamDTO {
	if home {
		return &dto.TeamDTO{
			Name:           m.TeamInformations.Home.Name,
			Logo:           m.TeamInformations.Home.Logo,
			Score:          m.Score.Home,
			WinProbability: m.TeamInformations.Home.WinProbability,
		}
	}
	return &dto.TeamDTO{
		Name:           m.TeamInformations.Away.Name,
		Logo:           m.TeamInformations.Away.Logo,
		Score:          m.Score.Away,
		WinProbability: m.TeamInformations.Away.WinProbability,
	}
}

func (m *Match) ToDTO() *dto.MatchDTO {
	return &dto.MatchDTO{
		ID:       m.Informations.ID,
		Elapsed:  m.Informations.Status.Elapsed,
		Status:   m.getStatus(),
		HomeTeam: *m.getTeamInformation(true),
		AwayTeam: *m.getTeamInformation(false),
		Stadium: dto.StadiumDTO{
			Name: m.Informations.Venue.Name,
			City: m.Informations.Venue.City,
		},
	}
}

func ToMatchDTOs(matchs []*Match) []*dto.MatchDTO {
	var result []*dto.MatchDTO
	for _, m := range matchs {
		result = append(result, m.ToDTO())
	}
	return result
}
