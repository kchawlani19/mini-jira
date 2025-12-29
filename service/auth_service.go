package service

func Login(email string, password string) string {
	if email == "" || password == "" {
		return "Email or password missing"
	}

	return "login successful (dummy)"
}
