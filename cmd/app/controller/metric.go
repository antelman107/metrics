package controller

import (
	"net/http"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/echo"
)

type (
	aggRequest struct {
		SiteId int64 `form:"site_id" query:"site_id" validate:"required,gt=0"`
	}
)

type metric struct {
	metricRepo domain.MetricRepository
}

func NewMetric(metricRepo domain.MetricRepository) (s *metric, err error) {
	s = &metric{
		metricRepo: metricRepo,
	}

	return s, nil
}

func (h *metric) Serve(e *echo.Echo) {
	e.GET("/api/v1/agg", h.GetAgg)
}

func (h *metric) GetAgg(c echo.Context) (err error) {
	var req = new(aggRequest)
	if err = c.Bind(req); err != nil {
		return
	}

	if err = c.Validate(req); err != nil {
		return
	}

	var min, max, avg, count int64
	min, max, avg, count, err = h.metricRepo.GetStatBySiteId(req.SiteId)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"min":   min,
			"max":   max,
			"avg":   avg,
			"count": count,
		})
}
