package domain

type Response interface {
	StatusCode() int
	Error() error
}
