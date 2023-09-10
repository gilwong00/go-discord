package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	db "github.com/gilwong00/go-discord/db/sqlc"
	"github.com/gilwong00/go-discord/gen/proto/v1/member/memberv1connect"
	"github.com/gilwong00/go-discord/gen/proto/v1/message/messagev1connect"
	"github.com/gilwong00/go-discord/gen/proto/v1/server/serverv1connect"
	"github.com/gilwong00/go-discord/gen/proto/v1/user/userv1connect"
	"github.com/gilwong00/go-discord/internal/memberservice"
	"github.com/gilwong00/go-discord/internal/messageservice"
	"github.com/gilwong00/go-discord/internal/serverservice"
	"github.com/gilwong00/go-discord/internal/userservice"
	"github.com/gilwong00/go-discord/pkg/config"
	"github.com/gilwong00/go-discord/pkg/cors"
	"github.com/gilwong00/go-discord/pkg/interceptors"
	"github.com/gilwong00/go-discord/pkg/token"
	"github.com/gilwong00/go-discord/pkg/websocket"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ApiServer struct {
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewApiServer(
	config config.Config,
	store db.Store,
	tokenMaker token.Maker,
) *ApiServer {
	return &ApiServer{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
}

func (server *ApiServer) StartAPIServer() {
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		interceptors.NewAuthInterceptor(server.tokenMaker),
		interceptors.NewUnaryLoggingInteceptor(),
	)
	manager := websocket.NewManager(server.store)
	manager.SetupEventHandlers()
	userservice := userservice.NewUserService(server.store)
	serverservice := serverservice.NewServerService(server.store)
	messageservice := messageservice.NewMessageService(server.store)
	memberservice := memberservice.NewMemberService(server.store)
	userPath, userHandler := userv1connect.NewUserServiceHandler(userservice)
	serverPath, serverHandler := serverv1connect.NewServerServiceHandler(serverservice, interceptors)
	messagePath, messageHandler := messagev1connect.NewServerServiceHandler(messageservice, interceptors)
	memberPath, memberHandler := memberv1connect.NewMemberServiceHandler(memberservice, interceptors)
	mux.Handle(userPath, userHandler)
	mux.Handle(serverPath, serverHandler)
	mux.Handle(messagePath, messageHandler)
	mux.Handle(memberPath, memberHandler)
	http.HandleFunc("/ws", manager.ServeWS)
	// new server
	srv := &http.Server{
		Addr: server.config.ServerAddress,
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
		log.Printf("Starting server on: %s", server.config.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start http server: %v", err)
		}
	}()
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v", err)
	}
}
