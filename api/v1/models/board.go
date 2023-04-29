package models

type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   string `json:"createdOn"`
}
