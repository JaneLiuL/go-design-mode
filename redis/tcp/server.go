package tcp

import (
	"github.com/HDT3213/godis/lib/sync/atomic"
	database2 "github.com/hdt3213/godis/database"
	"sync"
	//"github.com/janeliul/go-design-mode/redis/database"
	"github.com/hdt3213/godis/interface/database"
)

type Handler struct {
	activeConn sync.Map
	db database.DB
	closing    atomic.Boolean
}

func NewHandler() *Handler {
	var db database.DB
	db = database2.NewStandaloneServer()
	return &Handler{
		db: db,
	} 
}

func ListenAndServeWithSignal(addr string, handler Handler)  {
	closeChan := make(chan struct{})
}