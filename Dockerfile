FROM golang:1.17-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy the go source
COPY main.go main.go
COPY db/ db/
COPY models/ models/
COPY routes/ routes/
RUN go build -o /go-chi-api

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /go-chi-api /go-chi-api
COPY names.json /
EXPOSE 8000
USER nonroot:nonroot
ENTRYPOINT ["/go-chi-api"]