FROM golang:1.18-buster AS builder

RUN mkdir /coreostools
WORKDIR /coreostools
COPY cmd/bu2ign/go.mod .
COPY cmd/bu2ign/go.sum .

RUN go mod download
COPY cmd/bu2ign/main.go .

RUN GOARCH=wasm GOOS=js go build -o bu2ign.wasm main.go

FROM nginx:1

COPY assets/ /usr/share/nginx/html
COPY --from=builder /coreostools/bu2ign.wasm /usr/share/nginx/html/bu2ign/main.wasm

COPY nginx.conf /etc/nginx/nginx.conf
