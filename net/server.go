package net

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fmt"

	"github.com/alissonperez/api-foo/infra/plog"
	"github.com/alissonperez/api-foo/infra/teardown"
	"github.com/alissonperez/api-foo/net/v1"
	"github.com/alissonperez/api-foo/config"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "So far! So good!")
}

func SetupServer(container *dig.Container) {
	err := container.Invoke(func(r *mux.Router, t *teardown.TearDown, logger plog.Log, config config.Config) {
		r.HandleFunc("/healthz", Healthz).Name("healthz")

		v1.Setup(r.PathPrefix("/v1").Subrouter(), container)

		done := make(chan os.Signal, 1)

		port := config.GetInt("server_port")

		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		srv := &http.Server{
			Handler: r,
			Addr:    fmt.Sprintf("0.0.0.0:%d", port),
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		}()

		logger.Infof("Server started... http://localhost:%d", port)

		<-done

		logger.Info("Server stoped...")

		// TODO: Continue, check how to use contexts...
		defer func() {
			logger.Info("TearDown cancel...")
			t.Cancel()
			logger.Info("TearDown cancel done... bye!")
		}()

		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Fatalf("Server Shutdown Failed:%+v", err)
		}

		logger.Info("Server Exited Properly")
	})

	if err != nil {
		log.Fatal(err)
	}
}
