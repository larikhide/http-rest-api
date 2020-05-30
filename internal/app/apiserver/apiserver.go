package apiserver //here apiserver package

import (
	"github.com/gorilla/mux"
	"github.com/larikhide/http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New returns already configure instance of APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start strarts http server Bd and etc
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

//configureStore try to open store, if ok - write to APIServer.store config store
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (a *APIServer) handleHello() http.HandlerFunc {
	// allow to init some variables which use only this handler
	// and code will execute once
	// for example
	/*type request struct {
		name string
	}*/
	return func(w http.ResponseWriter, r *http.Request) {
		// business logic here
		io.WriteString(w, "Hello! from web-server?")
	}
}
