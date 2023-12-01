# 构建打包镜像
FROM golang:alpine AS build
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${APP_VERSION} -X main.Name=${APP_NAME} -X main.ConfigHost=${CONFIG_HOST}  -X main.ConfigToken=${CONFIG_TOKEN}" -installsuffix cgo -o manager cmd/manager/main.go

# 构建执行镜像
FROM alpine
WORKDIR /go/build

ARG APP_VERSION
ARG APP_NAME
ARG CONF_HOST
ARG CONF_TOKEN
ENV CONF_HOST=$CONF_HOST
ENV CONF_TOKEN=$CONF_TOKEN
ENV APP_NAME=$APP_NAME
ENV APP_VERSION=$APP_VERSION
RUN echo ${APP_NAME} $CONF_HOST

COPY ./static/ /go/build/static/
COPY --from=build /go/build/manager /go/build/manager
CMD ["./manager"]
