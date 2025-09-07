package server

type Server struct {
	Id []byte
}

func NewServer(id []byte) *Server {
	return &Server{Id: id}
}
