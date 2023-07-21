package checks

import (
	"database/sql"
	"fmt"
	"go-clean-architecture-example/pkg/health"
	"sync"
	"time"
)

type Sql struct {
	db *sql.DB
}

func NewSQLChecker(db *sql.DB) *Sql {
	return &Sql{db: db}
}
func (s Sql) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	if s.db == nil {
		myStatus = false
		result.Status = false
		errorMessage = "connection is nil"
	}
	err := s.db.Ping()
	if err != nil {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("ping failed: %s", err)
	}
	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "sql",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
