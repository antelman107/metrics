package cache

import (
	"encoding/json"
	"fmt"

	"github.com/antelman107/metrics/cmd/app/domain"
	"github.com/antelman107/metrics/nosql"
)

type metric struct {
	kv nosql.KeyValueInterface
	db domain.MetricRepository
}

func NewStat(
	kv nosql.KeyValueInterface,
	db domain.MetricRepository,
) domain.MetricRepository {
	return &metric{
		kv: kv,
		db: db,
	}
}

func (repo *metric) Add(metric *domain.Metric) (err error) {
	err = repo.delCache(metric.SiteID)
	if err != nil {
		return err
	}

	return repo.db.Add(metric)
}

func (repo *metric) GetStatBySiteId(siteID int64) (min, max, avg, count int64, err error) {
	min, max, avg, count, err = repo.getCache(siteID)
	if err != nil {
		return min, max, avg, count, err
	}

	if count != 0 {
		return min, max, avg, count, nil
	}

	return repo.storeCache(siteID)
}

func (repo *metric) GetGroupStat() (metrics domain.GroupMetrics, err error) {
	return repo.db.GetGroupStat()
}

func (repo *metric) delCache(siteID int64) (err error) {
	return repo.kv.Del(repo.getCacheKey(siteID))
}

func (repo *metric) getCache(siteID int64) (min, max, avg, count int64, err error) {
	var data []byte

	data, err = repo.kv.Get(repo.getCacheKey(siteID))
	if err != nil {
		return 0, 0, 0, 0, err
	}

	var metrics = make(map[string]int64)

	if data != nil {
		err = json.Unmarshal(data, &metrics)
		if err != nil {
			return min, max, avg, count, err
		}

		return metrics["min"], metrics["max"], metrics["avg"], metrics["count"], err
	}

	return min, max, avg, count, nil
}

func (repo *metric) storeCache(siteID int64) (min, max, avg, count int64, err error) {
	var data = make([]byte, 0)

	min, max, avg, count, err = repo.db.GetStatBySiteId(siteID)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	var metrics = map[string]int64{
		"min":   min,
		"max":   max,
		"avg":   avg,
		"count": count,
	}

	data, err = json.Marshal(metrics)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return min, max, avg, count, repo.kv.Set(repo.getCacheKey(siteID), 0, data)
}

func (repo *metric) getCacheKey(siteID int64) string {
	return fmt.Sprintf(
		"stat_%d",
		siteID,
	)
}
