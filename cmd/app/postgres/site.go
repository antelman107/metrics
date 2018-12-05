package postgres

import (
	"database/sql"

	"github.com/iqoption/nap"
	"go.uber.org/zap"

	"github.com/antelman107/metrics/cmd/app/domain"
)

type site struct {
	db     *nap.DB
	logger *zap.Logger
}

func NewSite(
	db *nap.DB,
	logger *zap.Logger,
) domain.SiteRepository {
	return &site{
		db:     db,
		logger: logger,
	}
}

func (repo *site) GetAll() (sites domain.Sites, err error) {
	var rows *sql.Rows

	rows, err = repo.db.Master().Query(
		`SELECT
			id,
       		url
		 FROM "site"
         ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sites = make(domain.Sites, 0)

	for rows.Next() {
		var site = &domain.Site{}

		err = rows.Scan(
			&site.ID,
			&site.Url,
		)
		if err != nil {
			return nil, err
		}

		sites = append(sites, site)
	}

	return sites, nil
}
