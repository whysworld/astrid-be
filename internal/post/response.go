package post

import (
	storage "internal/pkg/storage"
	"time"
)

type createResponse struct {
	ID string `json:"id"`
}

type findResponse struct {
	ID         string           `json:"id"`
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   storage.MetaData `json:"metadata"`
	Spec       storage.Spec     `json:"spec"`
	CreatedAt  time.Time        `json:"created_at"`
}
