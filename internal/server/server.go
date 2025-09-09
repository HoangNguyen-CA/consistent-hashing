package server

type Server struct {
	Id string
}

func NewServer(id string) *Server {
	return &Server{Id: id}
}
