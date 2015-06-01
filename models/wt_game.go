package models

type WtGame struct {
	Id       int    `json:"id"`
	Teams    teams  `json:"teams"`
	DateTime string `json:"date_time"`
	Status   Status `json:"status"`
	Links    links  `json:"links"`
}
