package main

import (
	"encoding/gob"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"

	"github.com/KaushikiAnand/go-bootstrap-demo/project-templates/core/application"
	"github.com/KaushikiAnand/go-bootstrap-demo/project-templates/mysql/models"
)

func init() {
	gob.Register(&models.UserRow{})
}

func newConfig() (*viper.Viper, error) {
	defaultDSN := strings.Replace("root:@tcp(localhost:3306)/$GO_BOOTSTRAP_PROJECT_NAME?parseTime=true", "-", "_", -1)

	c := viper.New()
	c.SetDefault("dsn", defaultDSN)
	c.SetDefault("cookie_secret", "$GO_BOOTSTRAP_COOKIE_SECRET")
	c.SetDefault("http_addr", ":8888")
	c.SetDefault("http_cert_file", "")
	c.SetDefault("http_key_file", "")
	c.SetDefault("http_drain_interval", "1s")

	c.AutomaticEnv()

	return c, nil
}

func main() {
	os.Getenv("GO_BOOTSTRAP_REPO_NAME")
	os.Getenv("GO_BOOTSTRAP_REPO_USER")
	os.Getenv("GO_BOOTSTRAP_PROJECT_NAME")
	config, err := newConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	app, err := application.New(config)
	if err != nil {
		logrus.Fatal(err)
	}

	middle, err := app.MiddlewareStruct()
	if err != nil {
		logrus.Fatal(err)
	}

	serverAddress := config.Get("http_addr").(string)

	certFile := config.Get("http_cert_file").(string)
	keyFile := config.Get("http_key_file").(string)
	drainIntervalString := config.Get("http_drain_interval").(string)

	drainInterval, err := time.ParseDuration(drainIntervalString)
	if err != nil {
		logrus.Fatal(err)
	}

	srv := &graceful.Server{
		Timeout: drainInterval,
		Server:  &http.Server{Addr: serverAddress, Handler: middle},
	}

	logrus.Infoln("Running HTTP server on " + serverAddress)

	if certFile != "" && keyFile != "" {
		err = srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
