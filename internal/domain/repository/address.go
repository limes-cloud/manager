package repository

type Address interface {
	GetIPAddress(ip string) string
}
