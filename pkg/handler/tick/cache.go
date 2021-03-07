package tick

import (
	"encoding/json"
	"fmt"
	"os"
)

const cachedQuotePath = "/tmp/cached-quote"

func getCachedQuote() (cachedQuote Quote) {
	rawCache, err := os.ReadFile(cachedQuotePath)
	if err != nil {
		return Quote{}
	}

	err = json.Unmarshal(rawCache, &cachedQuote)
	if err != nil {
		return Quote{}
	}

	return cachedQuote
}

func setCachedQuote(quote Quote) error {
	f, err := os.Create(cachedQuotePath)
	if err != nil {
		return fmt.Errorf("creating cached quote file: %w", err)
	}

	defer func() {
		_ = f.Close()
	}()

	rawQuote, err := json.Marshal(quote)
	if err != nil {
		return fmt.Errorf("marshalling quote: %w", err)
	}

	_, err = f.Write(rawQuote)
	if err != nil {
		return fmt.Errorf("writing cached quote: %w", err)
	}

	return nil
}
