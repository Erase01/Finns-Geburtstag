FROM archlinux:latest
COPY . /app
WORKDIR /app
RUN pacman -Syu --noconfirm go
ENV GIN_MODE=release
RUN go get
RUN go build -o main /app
EXPOSE 8080
CMD ["/app/main"]
