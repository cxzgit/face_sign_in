package responses

type AdminInfo struct {
	AdminID string `json:"admin_id"`
	Name    string `json:"name"`
}

func NewAdminInfo(adminID, name string) *AdminInfo {
	return &AdminInfo{
		AdminID: adminID,
		Name:    name,
	}
}
