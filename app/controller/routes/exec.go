package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/solitaire/app/controller/clib"
)

func execRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/admin/exec", clib.ExecList)
	makeRoute(r, http.MethodGet, "/admin/exec/new", clib.ExecForm)
	makeRoute(r, http.MethodPost, "/admin/exec/new", clib.ExecNew)
	makeRoute(r, http.MethodGet, "/admin/exec/{key}/{idx}", clib.ExecDetail)
	makeRoute(r, http.MethodGet, "/admin/exec/{key}/{idx}/connect", clib.ExecSocket)
	makeRoute(r, http.MethodGet, "/admin/exec/{key}/{idx}/kill", clib.ExecKill)
}
