package repository

type Address interface {
	GetAddressByIP(ip string) string
}
