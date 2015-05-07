package main

type wtGame struct {
	Id     int    `json:"id"`
	Teams  teams  `json:"teams"`
	Status status `json:"status"`
	Links  links  `json:"links"`
}
