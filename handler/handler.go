package handler

import "github.com/tj/nsq_to_postgres/client"
import "github.com/segmentio/go-stats"
import "github.com/bitly/go-nsq"
import "time"

// Handler.
type Handler struct {
	stats *stats.Stats
	db    *client.Client
}

// New Handler with the given db client.
func New(db *client.Client) *Handler {
	stats := stats.New()
	go stats.TickEvery(10 * time.Second)
	return &Handler{
		stats: stats,
		db:    db,
	}
}

// HandleMessage inserts the message body into postgres.
func (h *Handler) HandleMessage(msg *nsq.Message) error {
	h.stats.Incr("inserts")
	return h.db.Insert(msg.Body)
}
