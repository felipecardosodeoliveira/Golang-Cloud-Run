FROM golang:1.22.3 AS builder 

WORKDIR /app 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/ -o golang_cloud_run

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/golang_cloud_run ./
ENTRYPOINT ["./golang_cloud_run"]
