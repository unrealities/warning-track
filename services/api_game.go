package services

import "fmt"

func MlbApiMlbTvLinkToUrl(g int) string {
	return fmt.Sprintf("http://www.mlb.com/tv/g%v", g)
}
