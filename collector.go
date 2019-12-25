package sqlstats

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "go_sql_stats"
	subsystem = "connections"
)

// StatsCollector implements the prometheus.Collector interface.
type StatsCollector struct {
	sp StatsProvider

	// descriptions of exported metrics
	maxOpenDesc           *prometheus.Desc
	openDesc              *prometheus.Desc
	inUseDesc             *prometheus.Desc
	idleDesc              *prometheus.Desc
	waitedForDesc         *prometheus.Desc
	blockedTimeDesc       *prometheus.Desc
	closedMaxIdleDesc     *prometheus.Desc
	closedMaxLifetimeDesc *prometheus.Desc
}

// NewStatsCollector creates a new StatsCollector.
func NewStatsCollector(sp StatsProvider) *StatsCollector {
	return &StatsCollector{
		sp: sp,
		maxOpenDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "max_open"),
			"Maximum number of open connections to the database.",
			[]string{"db_name"},
			nil,
		),
		openDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "open"),
			"The number of established connections both in use and idle.",
			[]string{"db_name"},
			nil,
		),
		inUseDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "in_use"),
			"The number of connections currently in use.",
			[]string{"db_name"},
			nil,
		),
		idleDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "idle"),
			"The number of idle connections.",
			[]string{"db_name"},
			nil,
		),
		waitedForDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "waited_for"),
			"The total number of connections waited for.",
			[]string{"db_name"},
			nil,
		),
		blockedTimeDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "blocked_time"),
			"The total time blocked waiting for a new connection.",
			[]string{"db_name"},
			nil,
		),
		closedMaxIdleDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "closed_max_idle"),
			"The total number of connections closed due to SetMaxIdleConns.",
			[]string{"db_name"},
			nil,
		),
		closedMaxLifetimeDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "closed_max_lifetime"),
			"The total number of connections closed due to SetConnMaxLifetime.",
			[]string{"db_name"},
			nil,
		),
	}
}

// Describe implements the prometheus.Collector interface.
func (c StatsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.maxOpenDesc
	ch <- c.openDesc
	ch <- c.inUseDesc
	ch <- c.idleDesc
	ch <- c.waitedForDesc
	ch <- c.blockedTimeDesc
	ch <- c.closedMaxIdleDesc
	ch <- c.closedMaxLifetimeDesc
}

// Collect implements the prometheus.Collector interface.
func (c StatsCollector) Collect(ch chan<- prometheus.Metric) {
	//dbName := c.sp.DBName()
	stats := c.sp.Stats()

	for _, dbName := range []string{c.sp.DBName(), "foo"} {
		ch <- prometheus.MustNewConstMetric(
			c.maxOpenDesc,
			prometheus.GaugeValue,
			float64(stats.MaxOpenConnections),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.openDesc,
			prometheus.GaugeValue,
			float64(stats.OpenConnections),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.inUseDesc,
			prometheus.GaugeValue,
			float64(stats.InUse),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.idleDesc,
			prometheus.GaugeValue,
			float64(stats.Idle),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.waitedForDesc,
			prometheus.CounterValue,
			float64(stats.WaitCount),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.blockedTimeDesc,
			prometheus.CounterValue,
			stats.WaitDuration.Seconds(),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.closedMaxIdleDesc,
			prometheus.CounterValue,
			float64(stats.MaxIdleClosed),
			dbName,
		)
		ch <- prometheus.MustNewConstMetric(
			c.closedMaxLifetimeDesc,
			prometheus.CounterValue,
			float64(stats.MaxLifetimeClosed),
			dbName,
		)
	}
}
