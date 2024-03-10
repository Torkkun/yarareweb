FROM golang:latest as builder
WORKDIR /workdir
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV HTTPS_PROXY=wwwproxy.kanazawa-it.ac.jp:8080
ENV HTTP_PROXY=wwwproxy.kanazawa-it.ac.jp:8080
# Build
RUN go build -o app ./cmd/main.go

#RUN cd app
#RUN go run ./cmd/main.go

# Runtime Container
# FROM alpine
# RUN apk add --no-cache ca-certificates
# COPY --from=builder /workdir/app /app
#  ENTRYPOINT [ "/app" ]