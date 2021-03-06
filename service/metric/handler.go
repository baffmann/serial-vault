package metric

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server is an http Metrics server.
type Server struct {
	metrics http.Handler
}

var doOnce sync.Once

// NewServer returns a new metrics server.
func NewServer() *Server {
	doOnce.Do(func() {
		InitMetrics()
	})
	return &Server{
		metrics: promhttp.Handler(),
	}
}

// ServeHTTP returns a metrics server.
// Use it to set up the prometheus endpoint in your web service
// router := mux.NewRouter()
// router.Handle("/_status/metrics", metric.NewServer()).Methods("GET")
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.metrics.ServeHTTP(w, r)
}
