package ports

import (
	"encoding/json"
	storagelib "gasprice-oracle/internal/adapter/storage"
	"go.uber.org/zap"
	"net/http"
)

type RootHandler struct {
	Logger  *zap.Logger
	Storage storagelib.IStorage
}

type RootResponse struct {
	BlockNumber int64 `json:"blockNumber"`

	Slow    int64 `json:"slow"`
	Normal  int64 `json:"normal"`
	Fast    int64 `json:"fast"`
	Fastest int64 `json:"fastest"`
}

func (handler *RootHandler) ServeHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.Logger.Debug("serving root request")

		storage := handler.Storage
		blockNumber, _ := storage.Get(storagelib.BlockNumber)
		slow, _ := storage.Get(storagelib.P40)
		normal, _ := storage.Get(storagelib.P60)
		fast, _ := storage.Get(storagelib.P75)
		fastest, _ := storage.Get(storagelib.P95)

		response := RootResponse{
			BlockNumber: blockNumber,
			Slow:        slow,
			Normal:      normal,
			Fast:        fast,
			Fastest:     fastest,
		}

		data, err := json.Marshal(response)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal server error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}
