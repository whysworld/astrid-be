module whysworld.com

go 1.20

require (
	github.com/elastic/go-elasticsearch/v8 v8.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0
)

require internal/pkg/storage/elasticsearch v1.0.0

replace internal/pkg/storage/elasticsearch => ./internal/pkg/storage/elasticsearch

require internal/pkg/domain v1.0.0 // indirect

replace internal/pkg/domain => ./internal/pkg/domain

require (
	github.com/elastic/elastic-transport-go/v8 v8.0.0-20211216131617-bbee439d559c // indirect
	internal/pkg/storage v1.0.0 // indirect
)

replace internal/pkg/storage => ./internal/pkg/storage

require internal/post v1.0.0

replace internal/post => ./internal/post
