package storage

import (
	"context"
	"time"
)

type PostStorer interface {
	Insert(ctx context.Context, post Post) error
	Update(ctx context.Context, post Post) error
	Delete(ctx context.Context, id string) error
	FindOne(ctx context.Context, id string) (Post, error)
}

// type Post struct {
// 	ID        string     `json:"id"`
// 	Title     string     `json:"title"`
// 	Text      string     `json:"text"`
// 	Tags      []string   `json:"tags"`
// 	CreatedAt *time.Time `json:"created_at,omitempty"`
// }

type MetaData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Links       string `json:"links"`
	Annotations string `json:"annotations"`
}

type Spec struct {
	Type      string `json:"type"`
	Owner     string `json:"owner"`
	Lifecycle string `json:"liefcycle"`
}

type Post struct {
	ID         string     `json:"ID"`
	ApiVersion string     `json:"apiVersion"`
	Kind       string     `json:"kind"`
	Metadata   MetaData   `json:"metadata"`
	Spec       Spec       `json:"spec"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}
