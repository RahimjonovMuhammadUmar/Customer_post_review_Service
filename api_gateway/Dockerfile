FROM golang:1.17-alpine
RUN mkdir api_gateway
COPY . /api_gateway
WORKDIR /api_gateway
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 8800
