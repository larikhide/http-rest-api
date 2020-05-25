package apiserver //here apiserver package

// APIServer ...
type APIServer struct {
	config *Config
}
// New returns already configure instance of APIServer
func New(config *Config) *APIServer {
return &APIServer{
	config: config,
}
}

// Start strarts http server Bd and etc
func (s *APIServer) Start() error  {
	return nil
}