package apiserver //here apiserver package

// APIServer ...
type APIServer struct {

}
// New returns already configure instance of APIServer
func New() *APIServer {
return &APIServer{}
}

// Start strarts http server Bd and etc
func (s *APIServer) Start() error  {
	return nil
}