package domain

type (
	Site struct {
		ID  int64
		Url string
	}

	Sites []*Site
)

type SiteRepository interface {
	GetAll() (sites Sites, err error)
}
