package example

type User struct {
	ID        int64  `json:"id" db:"id" validate:"-"`
	Email     string `json:"email" db:"email" validate:"required,email"`
	FirstName string `json:"first_name" db:"first_name" validate:"-"`
	Age       int    `json:"age" db:"age" validate:"gte=0,lte=130"`
}
