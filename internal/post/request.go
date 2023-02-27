package post

import (
	storage "internal/pkg/storage"
)

type createRequest struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   storage.MetaData `json:"metadata"`
	Spec       storage.Spec     `json:"spec"`
}

type updateRequest struct {
	ID string

	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   storage.MetaData `json:"metadata"`
	Spec       storage.Spec     `json:"spec"`
}

type deleteRequest struct {
	ID string
}

type findRequest struct {
	ID string
}
