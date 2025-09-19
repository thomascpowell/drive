package models

type ShareRequest struct {
	FileID string `json:"fileid"`
	TTL    string `json:"ttl"`
}
