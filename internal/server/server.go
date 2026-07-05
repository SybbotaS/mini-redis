package server

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/SybbotaS/mini-redis/internal/protocol"
	"github.com/SybbotaS/mini-redis/internal/storage"
)

type Server struct {
	address string
	storage *storage.Storage
	handler *protocol.Handler
}

func New(address string) *Server {

	store := storage.New()

	return &Server{
		address: address,
		storage: store,
		handler: protocol.New(store),
	}
}

func (s *Server) Start() error {

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	defer listener.Close()

	log.Printf("Server is listening on %s\n", s.address)

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {

	defer conn.Close()

	log.Printf("Client connected: %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {

		line, err := reader.ReadString('\n')

		if err != nil {

			if err != io.EOF {
				log.Println(err)
			}

			return
		}

		s.handler.Execute(conn, line)
	}
}
