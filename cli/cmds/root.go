package cmds

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	debug       bool
	Kubeconfig  string
	CommonFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        "kubeconfig",
			EnvVars:     []string{"KUBECONFIG"},
			Usage:       "Kubeconfig path",
			Destination: &Kubeconfig,
		},
	}
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "cattle-drive"
	app.Usage = "Tool for migrating rancher objects for RKE downstream clusters"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Turn on debug logs",
			Destination: &debug,
			EnvVars:     []string{"CATTLE_DRIVE_DEBUG"},
		},
	}

	app.Before = func(clx *cli.Context) error {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	return app
}
