package methods

import (
	"context"
	"io"
	"log/slog"
	"time"
)

func SendALog(w io.Writer) {
	logger := slog.New(slog.NewJSONHandler(w, nil))
	logger.Info("Well, Hello There!")
}

func TickEveryT(ctx context.Context, t time.Duration, c chan bool) {
	ticker := time.NewTicker(t)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// fmt.Println("Tick Time: ", tick)
			c <- true
		case <-ctx.Done():
			c <- false
			// fmt.Println("Context done, returning!")
			return
		}
	}
}
