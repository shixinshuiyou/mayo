package bo

type UserRegisterBo struct {
	Phone      string `form:"phone" binding:"required,phone"`
	Name       string `form:"name" binding:"required"`
	Icon       string `form:"icon" `
	BriefIntro string `form:"intro" `
	Sex        int    `form:"sex" `
	BrithDay   string `form:"brith" `
	School     string `form:"school" `
	Business   int    `form:"business" `
}
