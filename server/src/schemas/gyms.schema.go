package schemas

type CreateGym struct {
	Name string `json:"name" binding:"required"`
}

