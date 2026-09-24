package skio

import (
	"context"
	"fmt"
	"log"
	"sync"

	"food_delivery/component/token_provider/jwt"
	userstore "food_delivery/modules/user/store"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(appCtx AppContext, r *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (e *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	e.locker.Lock()

	if v, ok := e.storage[userId]; ok {
		e.storage[userId] = append(v, appSck)
	} else {
		e.storage[userId] = []AppSocket{appSck}
	}

	e.locker.Unlock()
}

func (e *rtEngine) getAppSocket(userId int) []AppSocket {
	e.locker.RLock()
	defer e.locker.RUnlock()
	return e.storage[userId]
}

func (e *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	e.locker.Lock()
	defer e.locker.Unlock()

	if v, ok := e.storage[userId]; ok {
		for i := range v {
			e.storage[userId] = append(v[:i], v[i+1:]...)
			break
		}
	}
}

func (e *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket

	if scks, ok := e.storage[userId]; ok {
		return scks
	}

	return sockets
}

func (e *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	e.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (e *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := e.getAppSocket(userId)

	for _, s := range sockets {
		s.Emit(key, data)
	}

	return nil
}

func (e *rtEngine) Run(appCtx AppContext, r *gin.Engine) error {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	e.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), "IP:", s.RemoteAddr())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("error:", e.Error())
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnected:", s.ID(), reason)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetMainDBConnection()
		store := userstore.NewSQLStore(db)

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			s.Emit("authentication_failed", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("authentication_failed", "you has been banned/deleted")
			s.Close()
			return
		}

		user.Mask(false)

		appSck := NewAppSocket(s, user)
		e.saveAppSocket(user.Id, appSck)

		s.Emit("authentication_success", user)
	})

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	if err := r.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}

	return nil
}
