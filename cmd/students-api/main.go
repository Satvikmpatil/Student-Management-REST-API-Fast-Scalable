package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/config"
	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/http/handlier/student"
	"github.com/Satvikmpatil/Student-Management-REST-API-Fast-Scalable/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	storage, errr:=sqlite.New(cfg)

	if errr != nil{
		log.Fatal(errr)
	}

	slog.Info("Storage Done",slog.String("env",cfg.Env))


	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}",student.GetById(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("addr", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Error", err)
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed ", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown")

}
