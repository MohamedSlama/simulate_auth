package main

func main() {
	var dbUser = getDbUser()
	var lnUser = User{userName: "ahmed", password: []byte("ahmed password")}
	dbUser.auth(lnUser)

	lnUser.userName = "mohamed"
	dbUser.auth(lnUser)

	lnUser.password = []byte("mohamed password")
	dbUser.auth(lnUser)
}
