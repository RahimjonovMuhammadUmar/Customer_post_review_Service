FROM golang:1.19.3-alpine
RUN mkdir customer_service
COPY . /customer_service
WORKDIR /customer_service
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9000
