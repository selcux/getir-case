package main

import (
	"context"
	"fmt"
	"getir-case/internal/router"
	"getir-case/internal/util"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(util.GetLogLevelFromEnv())
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Debugln("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("error in reading port\n: %s", err.Error())
	}
	appAddress := fmt.Sprintf("0.0.0.0:%d", port)

	r := router.RegisterRoutes()

	srv := &http.Server{
		Addr:    appAddress,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	log.Infof("server started... %s", appAddress)

	shutdownGracefully(srv)

}

func shutdownGracefully(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infoln("server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
