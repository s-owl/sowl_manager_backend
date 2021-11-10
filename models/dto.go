package models

type InfoDTO struct {
	Info string `json:"info" binding:"required"`
}

type ErrorDTO struct {
	Error string `json:"error" binding:"required"`
}
