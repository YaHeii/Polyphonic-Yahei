package tokenx

import "fmt"

const (
	jwtAccessPrefix  = "access"
	jwtRefreshPrefix = "refresh"
	signPrefix       = "sign"
)

func JwtAccessKey(uid string) string {
	return fmt.Sprintf("%s:%s", jwtAccessPrefix, uid)
}

func JwtRefreshKey(uid string) string {
	return fmt.Sprintf("%s:%s", jwtRefreshPrefix, uid)
}

func SignKey(uid string) string {
	return fmt.Sprintf("%s:%s", signPrefix, uid)
}
