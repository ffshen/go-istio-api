# build stage
#FROM golang as builder
#
#ENV GO111MODULE=on
#
#WORKDIR /app
#
#COPY go.mod .
#COPY go.sum .
#
#RUN go mod download
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
#
## final stage
#FROM scratch
#COPY --from=builder /app/repo /app/
#EXPOSE 9528
#ENTRYPOINT ["/app/repo"]

FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY . .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 9528
ENTRYPOINT ["/app/repo"]