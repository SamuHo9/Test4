package entity

type User struct {
	Name  string `valid:"required~Error Name is required"`
	Email string `valid:"email~Error Email is invalid,required"`
	Age   uint    `valid:"range(18|150)~Error Age must be between 18 and 150"`

}