package repository

type User struct {
	Id   string
	Name string
}

type UserRepository struct {
	users map[string]User
}

func (u UserRepository) AddUser(user User) {
	u.users[user.Id] = user
}

func (u UserRepository) GetUsers() []User {
	users := make([]User, len(u.users))
	id := 0
	for _, v := range u.users {
		users[id] = v
		id++
	}
	return users
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]User),
	}
}
