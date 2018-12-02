package controller

import (
	"net/http"
	"strconv"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/echo"
)

type metric struct {
	metricRepo domain.MetricRepository
	siteRepo   domain.SiteRepository
}

func NewMetric(
	metricRepo domain.MetricRepository,
	siteRepo domain.SiteRepository,
) (s *metric, err error) {
	s = &metric{
		metricRepo: metricRepo,
		siteRepo:   siteRepo,
	}

	return s, nil
}

func (h *metric) Serve(e *echo.Echo) {
	e.GET("/api/agg", h.GetAgg)
	e.GET("/api/sites", h.GetSites)
}

func (h *metric) GetAgg(c echo.Context) (err error) {
	var siteId int
	siteId, err = strconv.Atoi(c.FormValue("site_id"))
	if err != nil {
		return err
	}

	var min, max, avg, count int64
	min, max, avg, count, err = h.metricRepo.GetStatBySiteId(int64(siteId))
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

func (h *metric) GetSites(c echo.Context) (err error) {
	var sites domain.Sites
	sites, err = h.siteRepo.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		sites,
	)
}
