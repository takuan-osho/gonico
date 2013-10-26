package gonico

type NicoUser struct {
	Mail            string
	Password        string
	IsAuthenticated bool
}

func Login() bool {
	return true
}
