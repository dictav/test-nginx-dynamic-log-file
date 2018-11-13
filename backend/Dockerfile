FROM golang:1.11-alpine AS build

COPY main.go /main.go
RUN env CGO_ENABLED=0 go build -o /backend /main.go

FROM scratch
COPY --from=build /backend /backend

ENTRYPOINT ["/backend"]
