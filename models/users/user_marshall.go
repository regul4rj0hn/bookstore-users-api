package users

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	Status    string `json:"status"`
	CreatedOn string `json:"created_on"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedOn string `json:"created_on"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:        user.Id,
			CreatedOn: user.CreatedOn,
			Status:    user.Status,
		}
	}

	jsonUser, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(jsonUser, &privateUser)
	return privateUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
