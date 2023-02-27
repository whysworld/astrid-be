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
 
type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	Tags      []string   `json:"tags"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// type MetaData struct {
// 	name: string `json:"name"`
// 	description: string `json:"description"`
// 	links: string `json:"links"`
// 	annotations: string `json:"annotations"`
// }

// type Spec struct {
// 	type: string `json:"type"`
// 	owner: string `json:"owner"`
// 	lifecycle: string `json:"liefcycle"`
// }

// type Post struct {
// 	ID: string `json:"ID"`
// 	apiVersion: string `json:"apiVersion"`
// 	kind: string `json:"kind"`
// 	metadata: MetaData `json:"metadata"`
// 	spec: Spec `json:"spec"`
// }