FROM golang:latest as builder

WORKDIR /app
COPY . . 
RUN go mod download 



RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine 

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/api /api

EXPOSE 8080 

ENTRYPOINT ["/api"]