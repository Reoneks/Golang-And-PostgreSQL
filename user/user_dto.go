package user

type UserDto struct {
	Id        int64  `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func FromUserDto(userDto UserDto) User {
	return User(userDto)
}

func FromUserDtos(UserDtos []UserDto) (users []User) {
	for _, dto := range UserDtos {
		users = append(users, User(dto))
	}
	return
}

func ToUserDto(user User) UserDto {
	return UserDto(user)
}

func ToUserDtos(users []User) (userDtos []UserDto) {
	for _, dto := range users {
		userDtos = append(userDtos, UserDto(dto))
	}
	return
}

func (UserDto) TableName() string {
	return "users"
}
