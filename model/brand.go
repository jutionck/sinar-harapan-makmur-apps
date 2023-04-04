package model

type Brand struct {
	BaseModel
	Name    string    `json:"name"`
	Vehicle []Vehicle `json:"vehicle,omitempty"`
}

func (Brand) TableName() string {
	return "mst_brand"
}
