package godata

import (
	"github.com/gorilla/mux"
	"github.com/qwerty-iot/godata/edmx"
)

type EntityHandler interface {
	List(s *Server, scope string, options *ListOptions) ([]any, string, error)
}

type MetadataHandler interface {
	Metadata(s *Server, scope string) (*edmx.Model, error)
}

type Server struct {
	Fqdn         string
	Model        *edmx.Model
	ModelHandler MetadataHandler
	Handlers     map[string]EntityHandler
	router       *mux.Router
}

func NewServer(fqdn string, model *edmx.Model) *Server {
	return &Server{Fqdn: fqdn, Model: model, Handlers: make(map[string]EntityHandler)}
}

func (s *Server) SetMetadataHandler(handler MetadataHandler) {
	s.ModelHandler = handler
}
