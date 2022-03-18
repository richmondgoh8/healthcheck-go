run:
	go run cmd/main.go

swagger:
	which swagger || go get github.com/go-swagger/go-swagger/cmd/swagger
	SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o ./swagger.yaml --scan-models
	swagger serve ./swagger.yaml -F swagger --no-open

gen-mocks:
	echo "Generating mock services"
	mockgen -source=./pkg/client/client.go -destination=./test/mocks/mock_http.go

coverage:
	bash test.sh