FROM golang:1.20.0-bullseye as build
ENV TZ=Asia/Tokyo

WORKDIR /go/src/go-layered-architecture
RUN go install github.com/cosmtrek/air@v1.29.0
# RUN go install github.com/google/wire/cmd/wire@latest
COPY . .
# RUN wire /go/src/go-layered-architecture/internal/di/wire.go
RUN go build -o /go/bin/go-layered-architecture ./cmd/api/main.go

FROM gcr.io/distroless/base-debian11
ENV TZ=Asia/TokyoCOPY
COPY --from=build /go/bin/go-layered-architecture /
CMD ["/go-layered-architecture"]