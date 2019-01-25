package main

import (
	"log"
	"net/http"
	"runtime"
		"github.com/dimfeld/httptreemux"
	"github.com/volcano822/journey/common/configuration"
	"github.com/volcano822/journey/common/database"
	"github.com/volcano822/journey/common/plugins"
	"github.com/volcano822/journey/routes"
	"github.com/volcano822/journey/common/structure/methods"
	"github.com/volcano822/journey/common/templates"
	)

func main() {
	// Setup
	var err error

	// GOMAXPROCS - Maybe not needed
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Configuration is read from config.json by loading the configuration package

	// Database
	if err = database.Initialize(); err != nil {
		log.Fatal("Error: Couldn't initialize database:", err)
		return
	}

	// Global blog data
	if err = methods.GenerateBlog(); err != nil {
		log.Fatal("Error: Couldn't generate blog data:", err)
		return
	}

	// Templates
	if err = templates.Generate(); err != nil {
		log.Fatal("Error: Couldn't compile templates:", err)
		return
	}

	// Plugins
	if err = plugins.Load(); err == nil {
		// Close LuaPool at the end
		defer plugins.LuaPool.Shutdown()
		log.Println("Plugins loaded.")
	}

	// HTTP Server
	httpPort := configuration.Config.HttpHostAndPort
	httpRouter := httptreemux.New()
	// Blog and pages as http
	routes.InitializeBlog(httpRouter)
	routes.InitializePages(httpRouter)
	// Admin as http
	routes.InitializeAdmin(httpRouter)
	// Start http server
	log.Println("Starting http server on port " + httpPort + "...")
	if err := http.ListenAndServe(httpPort, httpRouter); err != nil {
		log.Fatal("Error: Couldn't start the HTTP server:", err)
	}
}
