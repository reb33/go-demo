package user

type User struct {
	Name  string 
	Password string 
	Email string `gorm:"index"`
}