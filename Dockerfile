# docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:latest sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o proj'
# docker run --rm -it -p 8080:8080 docker1

FROM proj/go
WORKDIR /app
ADD proj /app
ENTRYPOINT ["./proj"]
EXPOSE 7070