package types

type CreateFamilyBody struct {
	Name      string `json:"name" binding:"required"`
	Patriarch string `json:"patriarch" binding:"required"`
	Members   string `json:"members" binding:"required"`
}
