## build stage
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
#COPY --from=builder /app/api /app/
#EXPOSE 9527
#ENTRYPOINT ["/app/api"]

FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY . .

#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 9527
ENTRYPOINT ["/app/api"]