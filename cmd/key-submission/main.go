package main

import (
	"github.com/Shopify/goose/logger"
	"github.com/Shopify/goose/safely"

	"github.com/CovidShield/backend/pkg/app"
)

var log = logger.New("main")

func main() {
	defer safely.Recover() // panics -> bugsnag

	log(nil, nil).Info("starting")

	mainApp := app.NewBuilder().WithSubmission().Build()

	err := mainApp.RunAndWait()
	defer log(nil, err).Info("final message before shutdown")
}
