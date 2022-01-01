package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BridgeBuilderDao/bbd-app1/pkg/config"
	"github.com/BridgeBuilderDao/bbd-app1/pkg/handlers"
	"github.com/BridgeBuilderDao/bbd-app1/pkg/render"
)

const portNumber = ":8086"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Printf("Starting applicationon port on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
