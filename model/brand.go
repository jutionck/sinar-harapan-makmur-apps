package model

type Brand struct {
	BaseModel
	Name    string
	Vehicle []Vehicle
}

func (Brand) TableName() string {
	return "mst_brand"
}
