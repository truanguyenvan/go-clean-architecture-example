package healthcheck

type Healthcheck interface {
	Pass() bool
	Name() string
}

type Config struct {
	HealthPath  string
	Method      string
	StatusOK    int
	StatusNotOK int
}
