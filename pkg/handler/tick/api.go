package tick

import (
	"encoding/json"
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/tuuturu/dispatch/pkg/core"
	"math"
)

func (t tickHandler) Handle(payload []byte) (err error) {
	var quote Quote

	err = json.Unmarshal(payload, &quote)
	if err != nil {
		return fmt.Errorf("unmarshalling quote: %w", err)
	}

	cachedQuote := getCachedQuote()

	if math.Abs(quote.CurrentPrice-cachedQuote.CurrentPrice) >= 10 {
		var title string

		if quote.CurrentPrice > cachedQuote.CurrentPrice {
			title = "GME: Rocketing 🚀"
		} else {
			title = "GME: Falling 💤"
		}

		err := t.notificator.Push(
			title,
			fmt.Sprintf("Currently %f", quote.CurrentPrice),
			"",
			"critical",
		)
		if err != nil {
			return fmt.Errorf("notifying about big difference in price: %w", err)
		}
	}

	err = setCachedQuote(quote)
	if err != nil {
		return fmt.Errorf("setting cached quote: %w", err)
	}

	return nil
}

func NewTickHandler() core.Handler {
	return &tickHandler{
		notificator: notificator.New(notificator.Options{
			AppName: "dispatcher",
		}),
	}
}
