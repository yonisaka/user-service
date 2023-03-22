package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/log.LogService/SaveHttpLog":       true,
		"/log.LogService/SaveStreamHttpLog": true,
		"/user.UserService/GetUserList":     true,
		"/user.UserService/GetUser":         true,
		"/user.UserService/CreateUser":      false,
		"/user.UserService/UpdateUser":      true,
		"/user.UserService/DeleteUser":      true,
	}
}
