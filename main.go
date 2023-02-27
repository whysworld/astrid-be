package main

import (
	"log"
	"net/http"

	"internal/pkg/storage/elasticsearch"
	"internal/post"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Bootstrap elasticsearch.
	elastic, err := elasticsearch.New([]string{"http://0.0.0.0:9200"})
	if err != nil {
		log.Fatalln(err)
	}
	if err := elastic.CreateIndex("post"); err != nil {
		log.Fatalln(err)
	}

	// Bootstrap storage.
	storage, err := elasticsearch.NewPostStorage(*elastic)
	if err != nil {
		log.Fatalln(err)
	}

	// Bootstrap API.
	postAPI := post.New(storage)

	// Bootstrap HTTP router.
	router := httprouter.New()
	router.HandlerFunc("POST", "/api/v1/kinds", postAPI.Create)
	router.HandlerFunc("PATCH", "/api/v1/kinds/:id", postAPI.Update)
	router.HandlerFunc("DELETE", "/api/v1/kinds/:id", postAPI.Delete)
	router.HandlerFunc("GET", "/api/v1/kinds/:id", postAPI.Find)
	// router.HandlerFunc("GET", "/api/v1/posts", postAPI.All)

	// Start HTTP server.
	log.Fatalln(http.ListenAndServe(":3000", router))
}
