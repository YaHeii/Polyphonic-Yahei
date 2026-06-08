package tokenx

type TokenStore interface {
	Set(key string, value string, expireSeconds int) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) (bool, error)
	SetExpire(key string, expireSeconds int) error
}
