package models

type Team struct {
	Id      int    `json:"id"`
	MlbId   int64  `json:"mlb_id"`
	Abbr    string `json:"abbr"`
	Hashtag string `json:"hashtag"`
}
