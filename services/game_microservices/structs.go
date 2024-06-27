package gameservices

type Game struct {
	GameGUID string `json:"game_guid"`
	GameID   int64  `json:"game_id"`
	Sport    string `json:"sport"`
	TeamA    string `json:"team_a"`
	TeamB    string `json:"team_b"`
}
