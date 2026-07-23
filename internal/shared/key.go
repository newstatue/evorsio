package shared

const (
	KeyAuthCodePrefix = "auth:code:"
)

func KeyAuthCode(email string) string {
	return KeyAuthCodePrefix + email
}
