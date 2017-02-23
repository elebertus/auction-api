package routes

import (
	"net/http"

	"github.com/elebertus/auction-api/handlers"
	"github.com/gorilla/mux"
)

/*
   Not going to lie, this is mostly copy-pasta. I haven't
   even looked at it because it's functional. It feels a
   little dirty though.
*/
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"GetAuctionDataFiles",
		"GET",
		"/api/getdatafile",
		handlers.GetAuctionDataFile,
	},
	Route{
		"ShowAuctionField",
		"GET",
		"/api/showauctionfield/{auctionField}",
		handlers.ShowAuctionField,
	},
	Route{
		"GetAuctionData",
		"GET",
		"/api/getauctiondata",
		handlers.GetAuctionData,
	},
}
