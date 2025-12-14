DB_URL=
APP_BINARY_NAME=app
API_BINARY_NAME=api
 	
build_app:
	@echo "Building APP..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP_BINARY_NAME} ./cmd/app
	@echo "APP built!"

build_api:
	@echo "Building API..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${API_BINARY_NAME} ./cmd/api
	@echo "API built!"

build: build_app build_api

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!" 

atlas_schema_apply:
	atlas schema apply \
	-u "$(DB_URL)" \
	--to file://scripts/dbdoc/schema.sql \
	--dev-url "docker://postgres/15" \ 

atlas_schema_inspect:
	atlas schema inspect \
	-u "$(DB_URL)" \
	--web

atlas_migrate_diff:
	atlas migrate diff initial \
  	--dir "file://scripts/migrations" \
  	--to "file://scripts/dbdoc/schema.sql" \
  	--dev-url "docker://postgres/15" \
	--format '{{ sql . "  " }}'

atlas_migrate_push:
	atlas migrate push $(name) \
	--dev-url "docker://postgres/15/dev"

atlas_migrate_apply:
	atlas migrate apply --env local

models:
	go tool sql-to-go -db "$(DB_URL)" -mode separate -package model -output "./internal/model"

sqlc:
	sqlc generate

templ:
	templ generate
	
db_schema:
	dbml2sql --postgres -o scripts/dbdoc/schema.sql scripts/dbdoc/schema.dbml

swag:
	swag init -g main.go -d $(ENTRIES),internal && swag fmt

ca:
	bin/crypto create  --algo=$(algo) ca \
	--key-out $(out).key \
	--cert-out $(out).crt

cert:
	bin/crypto --algo=$(algo) create cert \
		--name $(host) \
		--ca-key $(ca).key \
		--ca-cert $(ca).crt \
		--cert-out $(out).crt \
		--key-out $(out).key
pkcs12:
	openssl pkcs12 -export \
	-in $(name).crt \
	-inkey $(name).key \
	-certfile $(ca).crt \
	-out $(name).p12 \
	-name "$(name) Certificate"