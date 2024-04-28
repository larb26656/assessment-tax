run:
	export PORT=8080 && export DATABASE_URL=postgres://postgres:postgres@127.0.0.1:5432/ktaxes?sslmode=disable && export ADMIN_USERNAME=adminTax && export ADMIN_PASSWORD=admin! && go run main.go
test:
	go test ./...