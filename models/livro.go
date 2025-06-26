package models

type Livro struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Titulo string `json:"titulo" form:"titulo" bson:"titulo" binding:"required"`
	Autor  string `json:"autor" form:"autor" bson:"autor" binding:"required"`
	Desc   string `json:"desc" form:"desc" bson:"desc" binding:"required"`
}
