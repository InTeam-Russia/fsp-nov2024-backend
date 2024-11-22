package api

import (
	"net/http"

	"github.com/InTeam-Russia/go-backend-template/internal/apierr"
	"github.com/InTeam-Russia/go-backend-template/internal/events/event"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(
	r *gin.Engine,
	logger *zap.Logger,
	eventRepo event.EventRepository,
) {
	e := r.Group("/events")

	e.GET("", func(c *gin.Context) {
		filter, err := BuildEventFilter(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, apierr.InvalidQueryParams)
			return
		}

		events, err := eventRepo.GetEventsByFilter(filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}

		c.JSON(http.StatusOK, events)
	})

	e.GET("/indicators", func(c *gin.Context) {
		filter, err := BuildEventFilter(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, apierr.InvalidQueryParams)
			return
		}

		indicators, err := eventRepo.GetIndicatorsByFilter(filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}

		c.JSON(http.StatusOK, indicators)
	})
}
