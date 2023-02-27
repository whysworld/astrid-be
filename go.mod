module whysworld.com

go 1.20

require (
	github.com/elastic/go-elasticsearch/v7 v7.17.7
	github.com/google/uuid v1.3.0
	github.com/julienschmidt/httprouter v1.3.0
)

require internal/pkg/storage/elasticsearch v1.0.0
replace internal/pkg/storage/elasticsearch => ./internal/pkg/storage/elasticsearch

require internal/pkg/domain v1.0.0
replace internal/pkg/domain => ./internal/pkg/domain

require internal/pkg/storage v1.0.0
replace internal/pkg/storage => ./internal/pkg/storage

require internal/post v1.0.0
replace internal/post => ./internal/post