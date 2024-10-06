package api

import (
	"NumismaticClubApi/pkg/api/handler"
	"NumismaticClubApi/pkg/api/middlewares"
	"NumismaticClubApi/pkg/api/utils"
	"NumismaticClubApi/pkg/service/coin"
	"context"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

const (
	maxHeaderBytes = 1 << 20 // 1 MB
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(ctx utils.MyContext) *Server {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	wrappedRouter := middlewares.RecoveryMiddleware(ctx, router)

	return &Server{
		httpServer: &http.Server{
			Addr:           viper.GetString("db"),
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			Handler:        wrappedRouter,
		},
		router: router,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) HandleCoins(ctx utils.MyContext, service coin.CoinService) {
	s.router.HandleFunc("/api/v1/coins/", handler.Create(ctx, service)).Methods(http.MethodPost)
	s.router.HandleFunc("/api/v1/coins/", handler.GetAll(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/coins/{id}/", handler.GetById(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/coins/{id}/", handler.Update(ctx, service)).Methods(http.MethodPut)
	s.router.HandleFunc("/api/v1/coins/{id}/", handler.Delete(ctx, service)).Methods(http.MethodDelete)
}
