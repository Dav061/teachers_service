package models

type AudienceRequest struct {
	AccessKey int64  `json:"access_key"`
	Audience    string `json:"audience"`
}
