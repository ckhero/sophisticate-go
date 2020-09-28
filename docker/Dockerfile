FROM golang:latest as builder
RUN mkdir /app
COPY ./ /app
WORKDIR /app
#EXPOSE 1314
#RUN go build main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o main
#FROM alpine:3.7
FROM amd64/alpine:3.7
COPY --from=builder /app/main /app/
CMD  "/app/main"
