package checks

import "database/sql"

type SqlCheck struct {
	Sql *sql.DB
}

func (s SqlCheck) Pass() bool {
	if s.Sql == nil {
		return false
	}
	err := s.Sql.Ping()
	if err != nil {
		return false
	}
	return true
}

func (s SqlCheck) Name() string {
	return "sql"
}
