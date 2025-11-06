test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!" 

sqlc:
	sqlc generate

templ:
	templ generate
	
db_schema:
	dbml2sql --postgres -o scripts/dbdoc/schema.sql scripts/dbdoc/schema.dbml

swag:
	swag init -g main.go -d cmd/api,internal && swag fmt

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