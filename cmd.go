package main

import (
	"os"

	"github.com/urfave/cli"
)

const (
	VERSION = "0.0.7"
)

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "sandNb"
	app.Usage = "cnb develop version"
	app.EnableBashCompletion = true //TODO
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "app",
			Usage:  "start an application",
			Action: CmdApp,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input data tsv/xls/google sheet id",
					Value: "",
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8080,
				},
				cli.StringFlag{
					Name:  "mode,m",
					Usage: "start web or desktop application (w/d)",
					Value: "w",
				},
				cli.StringFlag{
					Name:  "root,r",
					Usage: "root directory",
					Value: "",
				},
				cli.StringFlag{
					Name:  "cred,c",
					Usage: "cred.json config file",
					Value: "./creds.json",
				},
				cli.StringFlag{
					Name:  "ext,e",
					Usage: "chrome extenstion id, Default is Nucleome Browswer Extension in chrome web store",
					Value: "djcdicpaejhpgncicoglfckiappkoeof",
				},
			},
		},
	}
	app.Run(os.Args)
}
