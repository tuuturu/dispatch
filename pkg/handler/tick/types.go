package tick

import "github.com/0xAX/notificator"

type tickHandler struct {
	notificator *notificator.Notificator
}

type Quote struct {
	Symbol string `json:"symbol"`

	Timestamp    int64   `json:"timestamp"`
	OpeningPrice float64 `json:"opening_price"`
	CurrentPrice float64 `json:"current_price"`
}
