package models

type UserModel struct {
	BaseModel
	Name           string
	Email          string `gorm:"uniqueIndex:users_email_idx"`
	PasswordDigest string
}

func (UserModel) TableName() string {
	return "users"
}
