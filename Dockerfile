FROM golang:1.21

WORKDIR /pokemon
COPY . /pokemon

WORKDIR /pokemon/cmd/api
RUN go build -o /bin/pokemon-api -ldflags '-w -s' -tags netgo -a -installsuffix cgo -v .

EXPOSE 8080
CMD ["/bin/pokemon-api"]