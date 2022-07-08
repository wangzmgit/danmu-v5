FROM golang:1.16-alpine
WORKDIR /danmu/go-danmu
COPY . .

# #构建后端和安装环境
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o app main.go \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add ffmpeg  

#获取前端文件，并配置Nginx
COPY nginx.conf /etc/nginx/nginx.conf
RUN apk add nginx unzip \
    && unzip danmuFE.zip -d danmaku \
    && mkdir /var/www/danmaku \
    && mv danmaku /var/www \
    && rm danmuFE.zip 

CMD nginx -c /etc/nginx/nginx.conf && ./app
