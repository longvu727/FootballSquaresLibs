package squareservices


type Square struct {
	SquareID     int    `json:"square_id"`
	SquareGUID   string `json:"square_guid"`
	SquareSize   int    `json:"square_size"`
	RowPoints    string `json:"row_points"`
	ColumnPoints string `json:"column_points"`
}
