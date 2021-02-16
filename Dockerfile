FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix cgo -ldflags '-w -extldflags "-static"' -o redirect2url .

FROM scratch

WORKDIR /app
COPY --from=builder /build/redirect2url .

EXPOSE 4000

CMD ["./redirect2url"]
