package events

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Sensor is data model that arduino sent
type Sensor struct {
	Datetime    string `json:"datetime,omitempty"`
	Temperature string `json:"temperature,omitempty"`
}

// SSEvent is model to send events
type SSEvent struct {
	EventName string
	Data      interface{}
}

// HandlerEvent is handler to managment events to clients
type HandlerEvent struct {
	signal  chan struct{}
	clients map[string]*client
}

// NewHandlerEvent return instance of event handler
func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		signal:  make(chan struct{}, 1),
		clients: make(map[string]*client),
	}
}

// Listen is method to listen event to request http
func (hdl *HandlerEvent) Listen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	flush, ok := w.(http.Flusher)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	c := NewClient(id)
	hdl.register(c)
	c.Online(r.Context(), w, flush)
	hdl.delete(id)

}

// CreateEvent send event to broadcast
func (hdl *HandlerEvent) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	s := &Sensor{}

	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 39

	s.Datetime = time.Now().Format("15:04:05")
	s.Temperature = fmt.Sprintf("%v", rand.Intn((max-min+1)+min))

	//json.NewDecoder(r.Body).Decode(s)

	hdl.broadcast(SSEvent{
		EventName: "arduino",
		Data:      s,
	})
}

func (hdl *HandlerEvent) delete(id string) {
	hdl.signal <- struct{}{}
	delete(hdl.clients, id)
	<-hdl.signal
}
func (hdl *HandlerEvent) register(c *client) {
	hdl.signal <- struct{}{}
	hdl.clients[c.ID] = c
	<-hdl.signal
}

// broadcast send message all clients
func (hdl *HandlerEvent) broadcast(event SSEvent) {
	hdl.signal <- struct{}{}
	for _, client := range hdl.clients {
		client.messageChannel <- event
	}
	<-hdl.signal
}
