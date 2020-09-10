package schemas

type UserRegister struct {
	Email 		string 		`json:"email" 		binding:"required"`
	Password 	string 		`json:"password" 	binding:"required"`
}