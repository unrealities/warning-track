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
	router.ServeFiles("/www/*filepath", http.Dir("/tv"))

	// TODO:
	// From "github.com/julienschmidt/httprouter"
	// this approach sidesteps the strict core rules of this router to avoid routing problems.
	// A cleaner approach is to use a distinct sub-path for serving files, like /static/*filepath or /files/*filepath.
	router.NotFound = http.FileServer(http.Dir("www/"))

	return router
}
