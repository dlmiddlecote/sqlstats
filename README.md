# sqlstats

[![GoDoc](https://godoc.org/github.com/dlmiddlecote/sqlstats?status.svg)](http://godoc.org/github.com/dlmiddlecote/sqlstats)
[![License](https://img.shields.io/github/license/dlmiddlecote/sqlstats.svg)](https://github.com/dlmiddlecote/sqlstats/blob/master/LICENSE)

A Go library for collecting [sql.DBStats](https://golang.org/pkg/database/sql/#DBStats) and exporting them in Prometheus format.

## Installation

```bash
go get github.com/dlmiddlecote/sqlstats
```

## Example

```go
package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/dlmiddlecote/sqlstats"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
    // Open connection to a DB (could also use the https://github.com/jmoiron/sqlx library)
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return err
	}

    // Create a new collector, the name will be used as a label on the metrics
    collector := sqlstats.NewStatsCollector("db_name", db)

    // Register it with Prometheus
	prometheus.MustRegister(collector)

    // Register the metrics handler
	http.Handle("/metrics", promhttp.Handler())

    // Run the web server
	return http.ListenAndServe(":8080", nil)
}
```

## Exposed Metrics

| Name                                         | Description                                                       | Labels  |
|----------------------------------------------|-------------------------------------------------------------------|---------|
| go_sql_stats_connections_max_open            | Maximum number of open connections to the database.               | db_name |
| go_sql_stats_connections_open                | The number of established connections both in use and idle.       | db_name |
| go_sql_stats_connections_in_use              | The number of connections currently in use.                       | db_name |
| go_sql_stats_connections_idle                | The number of idle connections.                                   | db_name |
| go_sql_stats_connections_waited_for          | The total number of connections waited for.                       | db_name |
| go_sql_stats_connections_blocked_seconds     | The total time blocked waiting for a new connection.              | db_name |
| go_sql_stats_connections_closed_max_idle     | The total number of connections closed due to SetMaxIdleConns.    | db_name |
| go_sql_stats_connections_closed_max_lifetime | The total number of connections closed due to SetConnMaxLifetime. | db_name |
