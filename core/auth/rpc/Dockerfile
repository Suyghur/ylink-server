FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY core/auth/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/auth core/auth/rpc/auth.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/auth /app/auth
COPY --from=builder /app/etc /app/etc

CMD ["./auth", "-f", "etc/auth.yaml"]
