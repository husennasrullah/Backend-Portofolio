FROM golang:1.20-alpine
WORKDIR /app
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN go build -o crudproduct
EXPOSE 8080
CMD ["/app/crudproduct"]