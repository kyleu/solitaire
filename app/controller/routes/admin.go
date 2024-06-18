package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/solitaire/app/controller/clib"
)

func adminRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/admin", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/sandbox", clib.SandboxList)
	makeRoute(r, http.MethodGet, "/admin/sandbox/{key}", clib.SandboxRun)
	makeRoute(r, http.MethodGet, "/admin/{path:.*}", clib.Admin)
	makeRoute(r, http.MethodPost, "/admin/{path:.*}", clib.Admin)
}
