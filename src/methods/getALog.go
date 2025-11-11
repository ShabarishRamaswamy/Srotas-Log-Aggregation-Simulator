package methods

import (
	"log/slog"
	"net/http"
)

func GetALog(w http.ResponseWriter) {
	logger := slog.New(slog.NewJSONHandler(w, nil))
	logger.Info("Well, Hello There!")
}
