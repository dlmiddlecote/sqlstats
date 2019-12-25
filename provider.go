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

// DBName implements the StatsProvider interface.
func (p DefaultStatsProvider) DBName() string {
	return p.dbName
}

// Stats implements the StatsGetter interface.
func (p DefaultStatsProvider) Stats() sql.DBStats {
	return p.db.Stats()
}
