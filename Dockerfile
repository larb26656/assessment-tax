FROM golang:1.21.9-alpine AS build

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o /bin/app

FROM golang:1.21.9-alpine

COPY --from=build /bin /bin

# Set initial env
ENV PORT=8080
ENV DATABASE_URL="host=host.docker.internal port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"
ENV ADMIN_USERNAME=adminTax
ENV ADMIN_PASSWORD=admin!

EXPOSE 8080

CMD /bin/app