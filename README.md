# 弹幕视频网站

## 项目介绍
项目Web端使用Vue3 + NaiveUI ,后端使用Golang +Gin +Gorm进行开发。文件存储支持本地存储和阿里云OSS存储，支持视频处理，部署可以前后端分离部署也可以使用docker部署。

## 项目结构
```
danmu_v5
|--danmu-admin 后台管理
|--go-danmu 后端
|--vue-danmu 前端
```

## 项目文档

部署文档 正在编写中

接口文档 正在编写中

开发交流群 909847398

旧版本项目 https://gitee.com/wzmgit/go-danmu

旧版本演示视频 https://www.bilibili.com/video/BV1TA411F7xz


## 使用教程
使用前需安装MySQL、Redis、Docker。安装完成后，进入后端目录，按照 README.md 文件所示创建配置文件，下载前端发行版。使用以下命令构建并启动项目。
```
docker build -t "danmu" .

docker run -itd --name danmuV5 -v /usr/danmaku:/danmu/go-danmu/file -p 9000:9000 danmu
```
启动成功后，进入后台管理系统, 使用默认账号`danmu@admin.com`密码`123456` 进行登录，登录成功后可对网站进行配置。


## 项目截图

暂无

