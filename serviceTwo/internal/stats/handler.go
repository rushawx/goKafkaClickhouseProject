package stats

import (
	"serviceTwo/config"

	"github.com/gin-gonic/gin"
)

// HandlerDeps structure
type HandlerDeps struct {
	Config *config.Config
}

// Handler  structure
type Handler struct {
	Config *config.Config
}

// NewStatsHandler func
func NewStatsHandler(router gin.IRouter, deps HandlerDeps) {
	handler := &Handler{
		Config: deps.Config,
	}

	router.GET("/clickStats", handler.getClickStats())
}

func (h *Handler) getClickStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		var stats []Stats
		result := h.Config.DbConf.Db.Table("clicks_daily").Find(&stats)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(200, stats)
	}
}
