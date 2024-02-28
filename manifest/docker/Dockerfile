FROM    golang:alpine AS stage1
ENV     RUN_PATH=/hrapi
RUN     mkdir -p $RUN_PATH
WORKDIR $RUN_PATH
ENV     GO111MODULE=on
# 複製所有檔案到/hrapi
COPY .. .
RUN go mod tidy
# 下載make
RUN apk update && apk add make
# 編譯
RUN make build

FROM alpine
USER root
# 複製字體
COPY ./manifest/docker/fonts/msjh.ttc /usr/share/fonts/
#安裝依賴庫和Chrome
RUN apk update && apk upgrade && \
    apk add --no-cache \
        chromium \
        fontconfig \
        && \
    rm -rf /var/cache/apk/*
# 設置字體
RUN fc-cache -f -v && \
    fc-list :lang=zh \
# 主程式操作
RUN mkdir /hrapi
WORKDIR /hrapi
# 從stage1編譯好的二進制檔案複製到/hrapi裡
COPY   --from=stage1 /hrapi/bin .
# expose port
EXPOSE 9000
# start server
ENV GB_CFG_FILE=config.prod.yaml
ENTRYPOINT ["./hrapi"]