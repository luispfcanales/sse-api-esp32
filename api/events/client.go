package events

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type client struct {
	ID             string
	messageChannel chan SSEvent
}

// NewClient return instance of client
func NewClient(id string) *client {
	return &client{
		ID:             id,
		messageChannel: make(chan SSEvent),
	}
}

func (c *client) Online(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case <-ctx.Done():
			close(c.messageChannel)
			return
		case message := <-c.messageChannel:
			data, err := json.Marshal(message.Data)
			if err != nil {
				log.Println(err)
			}
			messageFormat := "event:%s\ndata:%s\n\n"
			fmt.Fprintf(w, messageFormat, message.EventName, string(data))
			flusher.Flush()
		}
	}
}
