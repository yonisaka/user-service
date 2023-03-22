package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/log.LogService/SaveHttpLog":       false,
		"/log.LogService/SaveStreamHttpLog": false,
		"/log.LogService/FindHttpLog":       true,
		"/log.LogService/GetHttpLog":        true,
		"/user.UserService/GetUserList":     true,
	}
}
