package services

func MlbApiMlbTvLinkToUrl(g, c string) string {
	return "http://www.mlb.com/tv/g" + g + "/v" + c
}
