package postgres

import (
	"database/sql"
	"github.com/iqoption/nap"
	"go.uber.org/zap"

	"github.com/antelman107/metrics/cmd/app/domain"
)

type metric struct {
	db     *nap.DB
	logger *zap.Logger
}

func NewMetric(
	db *nap.DB,
	logger *zap.Logger,
) domain.MetricRepository {
	return &metric{
		db:     db,
		logger: logger,
	}
}

func (repo *metric) Add(metric *domain.Metric) (err error) {
	_, err = repo.db.Master().Exec(
		`
INSERT INTO "metric" ("site_id", "duration", "error", "is_success") 
VALUES ($1, $2, $3, $4)
RETURNING "id"`,
		metric.SiteID,
		metric.Duration,
		metric.Error,
		metric.IsSuccess,
	)

	return err
}

func (repo *metric) GetStatBySiteId(siteId int64) (min, max, avg, count int64, err error) {
	err = repo.db.Slave().QueryRow(
		`SELECT MIN(duration), MAX(duration), ROUND(AVG(duration)), COUNT(1) FROM "metric" 
		WHERE site_id = $1`,
		siteId,
	).Scan(&min, &max, &avg, &count)

	if err == sql.ErrNoRows {
		return 0, 0, 0, 0, nil
	}

	return min, max, avg, count, err
}

func (repo *metric) GetGroupStat() (metrics domain.GroupMetrics, err error) {
	var rows *sql.Rows
	rows, err = repo.db.Slave().Query(
		`SELECT 
       				MIN(m.duration) AS min_duration, 
       				MAX(m.duration) AS max_duration, 
       				ROUND(AVG(m.duration)) AS avg_duration, 
       				COUNT(1) AS count, 
       				m.site_id,
				FROM "metric" AS m
				GROUP BY site_id
				ORDER BY site_id`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	metrics = make(domain.GroupMetrics, 0)

	for rows.Next() {
		var metric = domain.GroupMetric{}

		err = rows.Scan(
			&metric.MinDuration,
			&metric.MaxDuration,
			&metric.AvgDuration,
			&metric.Count,
			&metric.SiteID,
		)

		metrics = append(metrics, &metric)
	}

	return metrics, nil
}
