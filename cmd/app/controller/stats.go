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
		mutex sync.RWMutex

		Uptime       time.Time
		RequestCount uint64
		Sites        map[int64]int64
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime: time.Now(),
		Sites:  map[int64]int64{},
	}
}

func (s *Stats) Serve(e *echo.Echo) {
	e.GET("/stats", s.GetStats)
}

// Process is the middleware function.
func (s *Stats) Process() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// track only api requests
			if !strings.Contains(c.Request().URL.Path, "/api/v1/agg") {
				return next(c)
			}

			var siteIDInt = 0
			siteIDInt, err = strconv.Atoi(c.FormValue("site_id"))
			if err != nil {
				c.Logger().Error(err)

				return next(c)
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			s.mutex.Lock()
			defer s.mutex.Unlock()

			s.RequestCount++
			s.Sites[int64(siteIDInt)]++

			return next(c)
		}
	}
}

// GetStats is the endpoint to get stats.
func (s *Stats) GetStats(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"uptime": s.Uptime,
		"sites":  s.Sites,
	})
}
