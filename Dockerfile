# 构建打包镜像
FROM golang:alpine AS build
#ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
WORKDIR /go/cache
ADD api/go.mod api/go.mod
ADD api/go.sum api/go.sum
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o manager main.go

# 构建执行镜像
FROM alpine
WORKDIR /go/build
COPY ./static/ /go/build/static/
COPY ./deploy/ /go/build/deploy/
COPY ./conf/*.yaml /go/build/conf/

COPY --from=build /go/build/manager /go/build/manager
CMD ["./manager"]