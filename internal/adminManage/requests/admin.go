package requests

type AdminLoginRequest struct {
	AdminID  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
