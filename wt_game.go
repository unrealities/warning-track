package main

type wtGame struct {
	Id       int    `json:"id"`
	Teams    teams  `json:"teams"`
	DateTime string `json:"date_time"`
	Status   status `json:"status"`
	Links    links  `json:"links"`
}
