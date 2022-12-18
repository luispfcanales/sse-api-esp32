package api

import (
	"net/http"
	"os"

	"github.com/luispfcanales/sse-api/api/events"
	"github.com/luispfcanales/sse-api/pkg/mdl"
)

// Run start api
func Run() error {

	hdlEvent := events.NewHandlerEvent()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/listen-event", hdlEvent.Listen)
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
