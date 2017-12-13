package main

import (
	"github.com/gorilla/mux"
	"github.com/nbrowser/sand"
	"github.com/urfave/cli"
)

func CmdApp(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	mode := c.String("mode")
	root := c.String("root")
	router := mux.NewRouter()
	/* TODO Add Personal V2 */
	addBindata(router)
	/* End */
	cred := c.String("cred")

	libs := []string{
		"/lib/three.min.js",
		"/local/lib/TrackballControls.js",
		"/local/lib/CanvasRenderer.js",
		"/local/lib/Projector.js",
		"/local/lib/init3d.js",
		"/local/lib/snow.min.js",
	}
	tail := []string{
		"/local/lib/onload.js",
	}
	styles := []string{"/local/style/snow.css"}
	s := sand.Sand{
		"CMU Nucleome Browser",
		root,
		".cnb",
		VERSION,
		libs,
		tail,
		styles,
		"snow.render",
	}
	idxRoot := s.InitIdxRoot(root) //???
	sand.InitCred(cred)
	addDataServer(uri, router, idxRoot) //TODO
	s.InitRouter(router)
	s.InitHome(root)
	s.Start(mode, port, router)
	return nil
}
