FROM bigrocs/golang-gcc:1.12 as builder

# 安装 odbc 依赖
RUN apk add --no-cache gcc git musl-dev make zip unixodbc unixodbc-dev freetds

WORKDIR /go/src/github.com/gomsa/old-sql
COPY . .

RUN GO111MODULE=off
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/service

FROM bigrocs/alpine:ca-data
# 安装 odbc 依赖
RUN apk add --update unixodbc unixodbc-dev freetds

ADD docker/etc_odbcinst.ini /etc/odbcinst.ini

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/github.com/gomsa/old-sql/bin/service /usr/local/bin/
CMD ["service"]
