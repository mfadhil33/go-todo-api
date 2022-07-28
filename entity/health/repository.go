package health

type Repository interface {
	Ping() error
}
