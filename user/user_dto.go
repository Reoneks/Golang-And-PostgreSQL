package user

//^ Status codes
const (
	Active   = 1
	UnActive = 2
	Banned   = 3
	Deleted  = 4
)

type UserDto struct {
	Id        int64  `gorm:"column:id"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Status    int64  `gorm:"column:status"`
}

type UserFilter struct {
	FirstName string
	LastName  string
	Email     string
	Status    int64
}

func (UserDto) TableName() string {
	return "users"
}
