package api

import (
	"net/http"
	"os"

	"github.com/luispfcanales/sse-api-esp32/api/events"
	"github.com/luispfcanales/sse-api-esp32/pkg/mdl"
)

// Run start api
func Run() error {

	hdlEvent := events.NewHandlerEvent()

	http.HandleFunc("/listen-event", mdl.Cors(hdlEvent.Listen))
	http.HandleFunc("/arduino-data", mdl.Cors(hdlEvent.CreateEvent))

	return http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	return ":" + PORT
}
