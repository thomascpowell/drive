package models

type ShareRequest struct {
	FileID uint `json:"FileID"`
	TTL    uint `json:"TTL"`
}
