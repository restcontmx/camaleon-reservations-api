package models

//
// TableModel is the tables on the locations
// @param ID : int - the unique id of the table
// @param Name : string - the name of the table
// @param Description : string - the additional description of the table
//
type TableModel struct {
	ID          int
	Name        string
	Description string
	ImgURL      string
	MaxGuests   int
	Area        AreaModel
}
