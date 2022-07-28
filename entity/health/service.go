package health

type Service interface {
	Check() error
}
