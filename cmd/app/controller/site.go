package controller

import (
	"net/http"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/echo"
)

type (
	siteResponse struct {
		ID  int64
		Url string
	}

	sitesResponse []*siteResponse
)

type site struct {
	siteRepo domain.SiteRepository
}

func NewSite(siteRepo domain.SiteRepository) (s *site, err error) {
	s = &site{
		siteRepo: siteRepo,
	}

	return s, nil
}

func (h *site) Serve(e *echo.Echo) {
	e.GET("/api/v1/sites", h.GetList)
}

func (h *site) GetList(c echo.Context) (err error) {
	var sites domain.Sites

	sites, err = h.siteRepo.GetAll()
	if err != nil {
		return err
	}

	var response = make(sitesResponse, 0)
	for _, s := range sites {
		response = append(response, &siteResponse{
			ID:  s.ID,
			Url: s.Url,
		})
	}

	return c.JSON(
		http.StatusOK,
		response,
	)
}
