package token

type Generator interface {
	CreateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}
