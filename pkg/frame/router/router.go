package router

import (
	"context"
	"errors"
	"github.com/charmbracelet/log"
	"github.com/dida-social/dida-server/pkg/frame"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Router struct {
	name   string
	port   int
	router *gin.Engine
	serv   *http.Server
}

func (r *Router) RegisterHandler(h frame.Handler) {
	h.RegisterRouter(r.router)
}

func (r *Router) Name() string {
	return r.name
}

func (r *Router) Run(ctx context.Context) error {

	go func() {
		// 服务连接
		if err := r.serv.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return nil
}

func (r *Router) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.serv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Info("server exiting")
	return nil
}

var _ frame.RouterModule = (*Router)(nil)

func NewRouter(name string, port int) *Router {
	router := gin.Default()
	serv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}
	return &Router{
		name:   name,
		port:   port,
		router: router,
		serv:   serv,
	}
}
