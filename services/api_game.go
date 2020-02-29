package services

import "fmt"

func MlbApiMlbTvLinkToUrl(g int, c string) string {
	return fmt.Sprintf("http://www.mlb.com/tv/g%v/v%v", g, c)
}
