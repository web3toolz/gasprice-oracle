package rest

import (
	"encoding/json"
	storagelib "gasprice-oracle/internal/adapter/storage"
	"go.uber.org/zap"
	"net/http"
)

type RHandler struct {
	Logger   *zap.Logger
	Storage  storagelib.IStorage
	Networks []string
}

type responseItem struct {
	UpdatedAt int64 `json:"updatedAt"`

	Slow    int64 `json:"slow"`
	Normal  int64 `json:"normal"`
	Fast    int64 `json:"fast"`
	Fastest int64 `json:"fastest"`
}

func (handler *RHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.Logger.Debug("serving root request")

	storage := handler.Storage

	response := make(map[string]responseItem)

	for _, networkName := range handler.Networks {
		updatedAt, _ := storage.Get(networkName, storagelib.UpdatedAt)
		slow, _ := storage.Get(networkName, storagelib.P40)
		normal, _ := storage.Get(networkName, storagelib.P60)
		fast, _ := storage.Get(networkName, storagelib.P75)
		fastest, _ := storage.Get(networkName, storagelib.P95)
		response[networkName] = responseItem{
			UpdatedAt: updatedAt,
			Slow:      slow,
			Normal:    normal,
			Fast:      fast,
			Fastest:   fastest,
		}
	}

	data, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal server error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
