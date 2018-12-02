package cache

import (
	"encoding/json"
	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/nosql"
)

type site struct {
	kv nosql.KeyValueInterface
	db domain.SiteRepository
}

func NewSite(
	kv nosql.KeyValueInterface,
	db domain.SiteRepository,
) domain.SiteRepository {
	return &site{
		kv: kv,
		db: db,
	}
}

func (repo *site) GetAll() (sites domain.Sites, err error) {
	sites, err = repo.getCache()
	if err != nil {
		return nil, err
	}

	if sites != nil {
		return sites, nil
	}

	return repo.storeCache()
}

func (repo *site) delCache() (err error) {
	return repo.kv.Del(repo.getCacheKey())
}

func (repo *site) getCache() (sites domain.Sites, err error) {
	var data []byte

	data, err = repo.kv.Get(repo.getCacheKey())
	if err != nil {
		return nil, err
	}

	if data != nil {
		err = json.Unmarshal(data, &sites)
		if err != nil {
			return nil, err
		}

		return sites, err
	}

	return nil, nil
}

func (repo *site) storeCache() (sites domain.Sites, err error) {
	var data = make([]byte, 0)

	sites, err = repo.db.GetAll()
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(sites)
	if err != nil {
		return nil, err
	}

	return sites, repo.kv.Set(repo.getCacheKey(), 0, data)
}

func (repo *site) getCacheKey() string {
	return "sites"
}
