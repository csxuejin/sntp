//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/csxuejin/sntp/netapp"
	"github.com/csxuejin/sntp/netevent"
	"github.com/csxuejin/sntp/sntp"
	"github.com/dolab/logger"
	"github.com/golib/cli"
)

type Config struct {
	ServerIP string `json:"server_ip"`
}

var (
	log    *logger.Logger
	config *Config
)

const (
	VERSION = "1.0.0"
)

func init() {
	log, _ = logger.New("stdout")
	log.SetColor(true)
	log.SetFlag(3)

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Panicf("ioutil.ReadFile(config.json): %v \n", err)
		return
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Panicf("json.Unmarshal(%v, %v): %v\n", data, config, err)
		return
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gontp"
	app.Usage = "gontp -h"
	app.Version = VERSION
	app.Authors = []cli.Author{
		{
			Name:  "Xue Jin",
			Email: "csxuejin@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "run as ntp server",
			Action: ntpServer(log),
		},
		{
			Name:   "client",
			Usage:  "run as ntp client",
			Action: ntpClient(log),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("ok!")
	}
}

func ntpServer(log *logger.Logger) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		var handler = netapp.GetHandler()
		netevent.Reactor.ListenUdp(9090, handler)
		netevent.Reactor.Run()

		return nil
	}
}

func ntpClient(log *logger.Logger) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		for {
			t, err := sntp.Client(config.ServerIP)
			if err != nil {
				log.Errorf("sntp.Client(%v): %v\n", config.ServerIP, err)
				break
			}

			cmd := exec.Command("date", "--set", t.Format("01/02/2006 15:04:05.999999999"))
			if err := cmd.Run(); err != nil {
				log.Errorf("cmd.Run(): %v", err)
				break
			}

			log.Infof("set time to %v\n", t)

			time.Sleep(time.Second * 3)
		}
		return nil
	}
}
