package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nbrowser/sand"
	"github.com/nbrowser/cnb/sandNb"
	"github.com/urfave/cli"
)

func addBindata(router *mux.Router) {
	router.PathPrefix("/local").Handler(sandNb.BindataServer("app"))
}
func CmdApp(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	mode := c.String("mode")
	root := c.String("root")
	router := mux.NewRouter()
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := Asset("favicon.ico")
		if err == nil {
			w.Write(bytes)
		}
	})
	/* TODO Add Personal V2 */
	addBindata(router)
	/* End */
	cred := c.String("cred")
	extID := c.String("ext")

	libs := []string{
		"/local/lib/handsontable.min.js",
		"/lib/three.min.js",
		"/lib/ml.min.js",
		"/lib/numeric.min.js",
		"/lib/plotly.min.js",
		"/lib/jspdf.min.js",
		"/local/lib/galleria-1.5.7.min.js",
		"/local/lib/TrackballControls.js",
		"/local/lib/CanvasRenderer.js",
		"/local/lib/Projector.js",
		"/local/lib/init3d.js",
		"/local/lib/snow.min.js",
	}
	tail := []string{
		"/local/lib/onload.js",
	}
	styles := []string{
		"/local/style/handsontable.min.css",
		"/local/style/snow.css",
	}
	modes := map[string]string{}
	s := sand.Sand{
		"CMU Nucleome Browser",
		root,
		".cnb",
		VERSION,
		libs,
		tail,
		styles,
		"snow.render",
		modes,
		extID,
	}
	idxRoot := s.InitIdxRoot(root) //???
	sand.InitCred(cred)
	addDataServer(uri, router, idxRoot) //TODO
	s.InitRouter(router)
	s.InitHome(root)
	s.Start(mode, port, router)
	return nil
}
