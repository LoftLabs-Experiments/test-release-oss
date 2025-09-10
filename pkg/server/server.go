package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// TestServer represents a simple test server
type TestServer struct {
	port int
	srv  *http.Server
}

// NewTestServer creates a new test server
func NewTestServer() *TestServer {
	return &TestServer{
		port: 8080,
	}
}

// Start starts the test server
func (s *TestServer) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})
	
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "test-release-oss-dev")
	})

	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}

	logrus.Infof("Starting test server on port %d", s.port)
	
	go func() {
		<-ctx.Done()
		logrus.Info("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.srv.Shutdown(shutdownCtx)
	}()

	return s.srv.ListenAndServe()
}
