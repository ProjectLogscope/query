.PHONY: build
build: clean
	@go build -C ./cmd/api -v -race -buildvcs=false -trimpath -buildmode=exe \
	-o ./../../release/service-darwin-arm64 > ./release/service-darwin-arm64.build.log 2>&1

.PHONY: run
run:
	./release/service-darwin-arm64

.PHONY: clean
clean:
	@rm -f ./release/service-darwin-arm64

.PHONY: doc
doc:
	@rm -rf ./api/v1/openapi
	@swag init --generalInfo description.go --dir ./api/v1/definition --propertyStrategy camelcase --output ./api/v1/openapi --outputTypes json,yaml 
