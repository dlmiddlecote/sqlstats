package sqlstats

import "database/sql"

// StatsGetter ...
type StatsGetter interface {
	Stats() sql.DBStats
}

// StatsProvider ...
type StatsProvider interface {
	StatsGetter
	DBName() string
}

// DefaultStatsProvider ...
type DefaultStatsProvider struct {
	dbName string
	db     StatsGetter
}

// NewStatsProvider ...
func NewStatsProvider(dbName string, db StatsGetter) DefaultStatsProvider {
	return DefaultStatsProvider{
		dbName: dbName,
		db:     db,
	}
}
