package response

import "myexample/go-gin/features/users"

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromCore(userCore users.Core) user {
	userResponse := user{
		ID:    userCore.UserID,
		Name:  userCore.Name,
		Email: userCore.Email,
	}
	return userResponse
}
