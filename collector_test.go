package sqlstats

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestMultipleDB(t *testing.T) {
	db1, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	db2, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	c1 := NewStatsCollector("db1", db1)
	c2 := NewStatsCollector("db2", db2)

	assert.NoError(t, prometheus.Register(c1))
	assert.NoError(t, prometheus.Register(c2))
}
