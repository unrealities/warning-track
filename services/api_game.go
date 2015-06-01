package services

import "regexp"

func MlbApiMlbTvLinkToUrl(l string) string {
	calEventId := regexp.MustCompile(`[0-9-]+`)
	c := calEventId.FindString(l)
	return "http://m.mlb.com/tv/e" + c + "/v113652483/?&media_type=video&clickOrigin=Media%20Grid&team=mlb"
}
