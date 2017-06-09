package main

import (
	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/video-captions-api/database"
	"github.com/NYTimes/video-captions-api/providers"
	"github.com/NYTimes/video-captions-api/service"
)

func main() {
	var cfg service.Config
	var providersList []providers.Provider
	db, err := database.NewDatastoreClient("nyt-video-dev")
	if err != nil {
		server.Log.Fatal("Unable to create Datastore client", err)
	}
	config.LoadJSONFile("./config.json", &cfg)

	// TODO: remove the list from the service constructor and
	// add support for service.AddProvider(provider)
	providersList = append(providersList, providers.NewThreePlay(cfg.APIKey, cfg.APISecret))

	server.Init("video-captions-api", cfg.Server)
	err = server.Register(service.NewSimpleService(&cfg, providersList, db))
	if err != nil {
		server.Log.Fatal("Unable to register service: ", err)
	}

	err = server.Run()
	if err != nil {
		server.Log.Fatal("Server encountered a fatal error: ", err)
	}

}
