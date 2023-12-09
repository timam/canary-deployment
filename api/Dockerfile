FROM golang:1.20

WORKDIR /app
COPY . .

RUN go build -o /bin/api -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v .

EXPOSE 8080
CMD ["/bin/app"]
