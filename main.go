package main

import (
	"log"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"myapp/mylib"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		mylib.AddRouteHello(e)
		mylib.CreateCollection(e, "example")

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
