package healthcheck

func New() HealthCheck {
	return &healthcheck{}
}

type healthcheck struct{}
