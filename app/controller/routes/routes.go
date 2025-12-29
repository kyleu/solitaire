package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller"
	"github.com/kyleu/solitaire/app/controller/clib"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
)

func makeRoute(x *mux.Router, method string, path string, f http.HandlerFunc) {
	cutil.AddRoute(method, path)
	x.HandleFunc(path, f).Methods(method)
}

func AppRoutes(as *app.State, logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, "/", controller.Home)
	makeRoute(r, http.MethodGet, "/healthcheck", clib.Healthcheck)
	makeRoute(r, http.MethodGet, "/about", clib.About)

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.Profile)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)
	makeRoute(r, http.MethodGet, cutil.DefaultSearchPath, clib.Search)

	themeRoutes(r)

	// $PF_SECTION_START(routes)$
	makeRoute(r, http.MethodGet, "/game", controller.Game)
	makeRoute(r, http.MethodGet, "/game/test/json", controller.GameTestJSON)
	makeRoute(r, http.MethodGet, "/game/test/html", controller.GameTestHTML)
	makeRoute(r, http.MethodGet, "/game/test/wasm", controller.GameTestWASM)
	// $PF_SECTION_END(routes)$

	makeRoute(r, http.MethodGet, "/graphql", clib.GraphQLIndex)
	makeRoute(r, http.MethodGet, "/graphql/{key}", clib.GraphQLDetail)
	makeRoute(r, http.MethodPost, "/graphql/{key}", clib.GraphQLRun)

	execRoutes(r)
	adminRoutes(r)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/robots.txt", clib.RobotsTxt)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	r.Methods(http.MethodOptions).HandlerFunc(controller.Options)
	r.Methods(http.MethodHead).HandlerFunc(controller.Head)

	return cutil.WireRouter(r, defaultHandler, logger)
}
