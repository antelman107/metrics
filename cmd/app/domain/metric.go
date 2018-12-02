package domain

type (
	Metric struct {
		SiteID    int64
		Error     string
		IsSuccess bool
		Duration  int64
	}

	GroupMetric struct {
		SiteID      int64
		MaxDuration int64
		MinDuration int64
		AvgDuration int64
		Count       int64
	}

	GroupMetrics []*GroupMetric
)

type MetricRepository interface {
	Add(metric *Metric) (err error)
	GetStatBySiteId(siteId int64) (min, max, count, avg int64, err error)
	GetGroupStat() (metrics GroupMetrics, err error)
}
