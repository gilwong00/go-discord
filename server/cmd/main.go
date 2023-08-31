package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/gen/proto/v1/user/userv1connect"
	"github.com/gilwong00/go-discord/internal/userservice"
	"github.com/gilwong00/go-discord/pkg/cors"
	"github.com/joho/godotenv"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting env vars", err)
		os.Exit(1)
	}
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
		os.Exit(1)
	}
	store := db.NewStore(conn)
	mux := http.NewServeMux()
	userservice := userservice.NewUserService(store)
	if err != nil {
		log.Fatalln("failed to create user service")
	}
	userPath, userHandler := userv1connect.NewUserServiceHandler(userservice)
	mux.Handle(userPath, userHandler)
	// new server
	srv := &http.Server{
		Addr: "localhost:8080",
		Handler: h2c.NewHandler(
			cors.NewCORS().Handler(mux),
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v", err)
		}
	}()
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err)
	}
}
