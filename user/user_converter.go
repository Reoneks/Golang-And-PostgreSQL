package user

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
