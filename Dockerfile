FROM alpine:latest
RUN apk update
RUN apk add git go
RUN go install github.com/gin-gonic/gin@latest
RUN go install github.com/gin-contrib/sessions@latest
COPY . /app
WORKDIR /app
RUN go build -o main /app
EXPOSE 8080
CMD ["/app/main"]
