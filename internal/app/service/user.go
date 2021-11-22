package service

// UserRegisterParam 用户注册
type UserRegisterParam struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
	Email    string `json:"email" validate:"required"`    // 邮箱
}

func (svc *Service) CreateUser(param *UserRegisterParam) error {
	return svc.dao.CreateUser(param.Username, param.Password, param.Email)
}
