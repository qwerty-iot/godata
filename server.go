package godata

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"
)

func (s *Server) Start(port string, tlsConfig *tls.Config) {

	httpServer := &http.Server{
		Addr:         port,
		Handler:      s,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			logError(nil, err, "error starting http server.")
			os.Exit(-1)
		}
	}()
	logDebug(nil, nil, "server started on port: %s", port)
}

func (s *Server) SetHandler(entity string, handler EntityHandler) {
	s.Handlers[entity] = handler
}
