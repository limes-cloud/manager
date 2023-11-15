FROM golang:alpine AS build
ARG APP_VERSION
ARG APP_NAME
ARG CONFIG_HOST
ARG CONFIG_TOKEN

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${APP_VERSION} -X main.Name=${APP_NAME} -X main.ConfigHost=${CONFIG_HOST}  -X main.ConfigToken=${CONFIG_TOKEN}" -installsuffix cgo -o manager cmd/manager/main.go

FROM alpine
EXPOSE 7002
EXPOSE 8002
WORKDIR /go/build
COPY ./static/ /go/build/static/
COPY --from=build /go/build/manager /go/build/manager
CMD ["./manager"]
