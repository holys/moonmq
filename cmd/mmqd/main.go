package main

import (
	"flag"
	"io/ioutil"

	"github.com/siddontang/moonmq/broker"
)

var configFile = flag.String("config", "", "config file")

func main() {
	flag.Parse()

	if len(*configFile) == 0 {
		panic("config file must set")
	}

	buf, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	var app *broker.App
	app, err = broker.NewApp(buf)
	if err != nil {
		panic(err)
	}

	app.Run()
}
