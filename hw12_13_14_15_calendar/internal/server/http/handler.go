package internalhttp

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergey-yabloncev/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Handler struct {
	logger Logger
	app    Application
}

type Application interface {
	CreateEvent(context.Context, storage.Event) error
	Index(context.Context) map[string]string
	UpdateEvent(context.Context, storage.Event) error
	DeleteEvent(context.Context, storage.Event) error
	GetEvents(context.Context) ([]storage.Event, error)
}

func NewHandler(logger Logger, app Application) *Handler {
	return &Handler{
		logger,
		app,
	}
}

func (h *Handler) Router() http.Handler {
	router := gin.Default()
	router.Use(loggingMiddleware(h.logger))

	router.GET("/hello", h.helloHandler)
	router.GET("/events", h.getAllEvents)
	router.GET("/events/:id", h.getEvent)
	router.POST("/events", h.createEvent)

	return router
}

func (h *Handler) helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.app.Index(c))
}

func (h *Handler) createEvent(c *gin.Context) {
	var event storage.Event
	err := c.ShouldBind(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = h.app.CreateEvent(c, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *Handler) getEvent(c *gin.Context) {
	events, err := h.app.GetEvents(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *Handler) getAllEvents(c *gin.Context) {
	events, err := h.app.GetEvents(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, events)
}
