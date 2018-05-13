package http

type AirfareHttpInterface interface {
	Get(url string) ([]byte, error)
}
