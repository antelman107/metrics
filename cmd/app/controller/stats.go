package controller

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

func (s *Stats) Serve(e *echo.Echo) {
	e.GET("/stats", s.GetStats)
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// track only api requests
			if !strings.Contains(c.Request().URL.Path, "/api") {
				return next(c)
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			s.mutex.Lock()
			defer s.mutex.Unlock()

			s.RequestCount++
			status := strconv.Itoa(c.Response().Status)
			s.Statuses[status]++

			return nil
		}
	}
}

// GetStats is the endpoint to get stats.
func (s *Stats) GetStats(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return c.JSON(http.StatusOK, s)
}
