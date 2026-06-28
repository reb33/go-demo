package stat

import (
	"adv_demo/configs"
	"adv_demo/pkg/middleware"
	"adv_demo/pkg/response"
	"net/http"
	"time"
)

const (
	GrouperByDay   = "day"
	GrouperByMonth = "month"
)

type StatHandlerDeps struct {
	*StatRepository
	*configs.Config
}

type StatHandler struct {
	*StatRepository
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		layout := "2006-01-02"

		fromStr := r.URL.Query().Get("from")
		from, err := time.Parse(layout, fromStr)
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		toStr := r.URL.Query().Get("to")
		to, err := time.Parse(layout, toStr)
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GrouperByDay && by != GrouperByMonth {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		statsFromRepo := handler.StatRepository.GetStats(by, from, to)
		response.Json(w, http.StatusOK, ConvertRepoStatsToPayload(statsFromRepo))
	}
}
