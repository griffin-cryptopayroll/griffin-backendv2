package service

type LoginInformation struct {
	Username string
	Password string
}
type LoginService interface {
}

//func generateJWT() (string, error) {
//	token := jwt.New(jwt.SigningMethodEdDSA)
//
//	claims
//}
