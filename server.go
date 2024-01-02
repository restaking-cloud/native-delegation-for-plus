package k2

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (k2 *K2Service) startServer() error {
	k2.server = &http.Server{
		Addr:    k2.cfg.ListenAddress.Host,
		Handler: k2.getRouter(),
	}

	k2.log.WithField("listenAddr", k2.cfg.ListenAddress.String()).Info("Started K2 server")

	return k2.server.ListenAndServe()
}

func (k2 *K2Service) stopServer() error {
	if k2.server == nil {
		return nil
	}

	err := k2.server.Close()
	if err != nil {
		return err
	}

	k2.server = nil

	return nil
}

func (k2 *K2Service) ListenAddress() string {
	return k2.cfg.ListenAddress.String()
}

func (k2 *K2Service) getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc(pathRoot, k2.handleRoot).Methods(http.MethodGet)
	r.HandleFunc(pathExit, k2.handleExit).Methods(http.MethodPost)
	r.HandleFunc(pathClaim, k2.handleClaim).Methods(http.MethodPost)
	r.HandleFunc(pathRegister, k2.handleRegister).Methods(http.MethodPost)

	r.Use(mux.CORSMethodMiddleware(r))
	loggedRouter := LoggingMiddleware(k2.log, r)
	return loggedRouter
}
