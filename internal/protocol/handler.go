package protocol

import (
	"fmt"
	"net"
	"strings"

	"github.com/SybbotaS/mini-redis/internal/storage"
)

type Handler struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Execute(conn net.Conn, line string) {
	line = strings.TrimSpace(line)

	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}

	command := strings.ToUpper(parts[0])

	switch command {

	case "PING":
		fmt.Fprintln(conn, "PONG")

	case "SET":
		if len(parts) != 3 {
			fmt.Fprintln(conn, "ERR usage: SET <key> <value>")
			return
		}

		h.storage.Set(parts[1], parts[2])
		fmt.Fprintln(conn, "OK")

	case "GET":
		if len(parts) != 2 {
			fmt.Fprintln(conn, "ERR usage: GET <key>")
			return
		}

		value, ok := h.storage.Get(parts[1])

		if !ok {
			fmt.Fprintln(conn, "(nil)")
			return
		}

		fmt.Fprintln(conn, value)

	case "DEL":
		if len(parts) != 2 {
			fmt.Fprintln(conn, "ERR usage: DEL <key>")
			return
		}

		if h.storage.Delete(parts[1]) {
			fmt.Fprintln(conn, "OK")
		} else {
			fmt.Fprintln(conn, "(nil)")
		}

	case "EXISTS":
		if len(parts) != 2 {
			fmt.Fprintln(conn, "ERR usage: EXISTS <key>")
			return
		}

		if h.storage.Exists(parts[1]) {
			fmt.Fprintln(conn, "1")
		} else {
			fmt.Fprintln(conn, "0")
		}

	default:
		fmt.Fprintln(conn, "ERR unknown command")
	}
}
