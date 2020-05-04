package server

import (
	"context"
	"net/http"
	"path"
	"time"

	"github.com/observatorium/observatorium/internal/proxy"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

// gracePeriod is duration the server gracefully shuts down.
const gracePeriod = 2 * time.Minute

// DefaultRequestTimeout is the default value of the timeout duration per request.
const DefaultRequestTimeout = 2 * time.Minute

// DefaultReadTimeout is the default value of the maximum duration for reading the entire request, including the body.
const DefaultReadTimeout = 2 * time.Minute

// DefaultWriteTimeout is the default value of the maximum duration before timing out writes of the response.
const DefaultWriteTimeout = 2 * time.Minute

// Server defines parameters for running an HTTP server.
type Server struct {
	logger log.Logger
	srv    *http.Server

	opts options
}

// New creates a new Server.
func New(logger log.Logger, reg *prometheus.Registry, opts ...Option) Server {
	options := options{}

	for _, o := range opts {
		o.apply(&options)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(options.requestTimeout))

	ins := newInstrumentationMiddleware(reg)

	{
		// Legacy endpoints
		r.Handle("/api/v1/query",
			ins.newHandler("query_legacy", proxy.New(logger, "/api/v1", options.metricsReadEndpoint, options.proxyOptions...)),
		)
		r.Handle("/api/v1/query_range",
			ins.newHandler("query_range_legacy", proxy.New(logger, "/api/v1", options.metricsReadEndpoint, options.proxyOptions...)),
		)
	}

	if options.metricsUIEndpoint != nil {
		uiPath := "/ui/metrics/v1"

		r.Get(path.Join(uiPath, "*"),
			ins.newHandler("ui", proxy.New(logger, uiPath, options.metricsUIEndpoint, options.proxyOptions...)))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, path.Join(uiPath, "graph"), http.StatusMovedPermanently)
		})

		// NOTICE: Following redirects added to be compatible with existing Read UI.
		// Paths are explicitly specified to prevent unnecessary request to read handler.
		for _, p := range []string{
			"graph",
			"stores",
			"status",
		} {
			p := p
			r.Get(path.Join("/", p), func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, path.Join(uiPath, p), http.StatusMovedPermanently)
			})
		}
	}

	namespace := "/api/metrics/v1"
	r.Route(namespace, func(r chi.Router) {
		if options.metricsReadEndpoint != nil {
			r.Get("/api/v1/query",
				ins.newHandler("query", proxy.New(logger, path.Join(namespace, "api/v1"), options.metricsReadEndpoint, options.proxyOptions...)))

			r.Get("/api/v1/query_range",
				ins.newHandler("query_range", proxy.New(logger, path.Join(namespace, "api/v1"), options.metricsReadEndpoint, options.proxyOptions...)))

			r.Get("/api/v1/*",
				ins.newHandler("read", proxy.New(logger, path.Join(namespace, "api/v1"), options.metricsReadEndpoint, options.proxyOptions...)))
		}

		writePath := "/write"
		r.Post(writePath,
			ins.newHandler("write", proxy.New(logger, path.Join(namespace, writePath), options.metricsWriteEndpoint, options.proxyOptions...)))
	})

	return Server{
		logger: logger,
		srv: &http.Server{
			Addr:         options.listen,
			Handler:      r,
			TLSConfig:    options.tlsConfig,
			ReadTimeout:  options.readTimeout,
			WriteTimeout: options.writeTimeout,
		},
		opts: options,
	}
}

// ListenAndServe listens on the TCP network address and handles connections with given server configuration.
func (s *Server) ListenAndServe() error {
	level.Info(s.logger).Log("msg", "starting the HTTP server", "address", s.opts.listen)

	if s.opts.tlsConfig != nil {
		// certFile and keyFile passed in TLSConfig at initialization.
		return s.srv.ListenAndServeTLS("", "")
	}

	return s.srv.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(err error) {
	if err == http.ErrServerClosed {
		level.Warn(s.logger).Log("msg", "internal server closed unexpectedly")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	level.Info(s.logger).Log("msg", "shutting down internal server")

	if err := s.srv.Shutdown(ctx); err != nil {
		level.Error(s.logger).Log("msg", "shutting down failed", "err", err)
	}
}