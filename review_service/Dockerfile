FROM golang:1.17-alpine
RUN mkdir review_service
COPY . /review_service
WORKDIR /review_service
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 9900/tcp
