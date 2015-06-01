package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unrealities/warning-track/controllers"
)

func Routes() http.Handler {
	router := httprouter.New()

	router.GET("/games", controllers.GameJSON)
	router.GET("/fetchGames", controllers.SetGames)
	router.GET("/fetchStatuses", controllers.SetStatuses)
	router.GET("/fetchAllStatuses", controllers.SetAllStatuses)
	router.GET("/setTwitterCredentials", controllers.SetTwitterCredentials)
	router.GET("/deleteGamesStatuses", controllers.DeleteGamesStatuses)
	router.NotFound = http.FileServer(http.Dir("www/")).ServeHTTP
	router.ServeFiles("/www/*filepath", http.Dir("/tv"))

	return router
}
