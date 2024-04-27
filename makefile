run:
	export PORT=8080 && export DATABASE_URL={REPLACE_ME} && export ADMIN_USERNAME=adminTax && export ADMIN_PASSWORD=admin! && go run main.go
test:
	go test ./...