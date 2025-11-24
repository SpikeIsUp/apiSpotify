package router

import (
	"net/http"

	"github.com/SpikeIsUp/apiSpotify/siteWeb/controller"
)

func NewRouter() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/", controller.Home)
    mux.HandleFunc("/damso", controller.Damso)
    mux.HandleFunc("/laylow", controller.Laylow)

    return mux
}