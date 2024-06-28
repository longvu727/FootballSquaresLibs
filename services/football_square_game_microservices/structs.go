package footballsquaregameservices

type FootballSquareGameElement struct {
	FootballSquaresGameID int  `json:"football_square_game_id"`
	ColumnIndex           int  `json:"column_index"`
	RowIndex              int  `json:"row_index"`
	WinnerQuaterNumber    int  `json:"winner_quater_number"`
	Winner                bool `json:"winner"`
	UserID                int  `json:"user_id"`
	SquareID              int  `json:"square_id"`
	GameID                int  `json:"game_id"`

	UserName  string `json:"user_name"`
	UserAlias string `json:"user_alias"`
}
