package responses

type ClassInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewClassInfo(id uint, name string) *ClassInfo {
	return &ClassInfo{
		ID:   id,
		Name: name,
	}
}
