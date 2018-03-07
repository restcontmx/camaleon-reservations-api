package models

import "time"

// WaitListModel will be the model for wait list
type WaitListModel struct {
	ID         int
	Location   LocationModel
	Status     WaitListStatusModel
	ClientInfo ClientInfoModel
	TimeLimit  int
	Guests     int
	AlertTime  time.Time
	Timestamp  time.Time
	Updated    time.Time
}
