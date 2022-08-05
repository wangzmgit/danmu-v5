# 弹幕视频网站

## 项目介绍
项目Web端使用Vue3 + NaiveUI ,后端使用Golang +Gin +Gorm进行开发。文件存储支持本地存储和阿里云OSS存储，支持视频处理，部署可以前后端分离部署也可以使用docker部署。

## 项目结构
```
danmu_v5
|--images 项目截图
|--danmu-admin 后台管理
|--go-danmu 后端
|--vue-danmu 前端
|--danmu-mobile 移动端
```

## 项目文档

`开发交流群 909847398`

新版本使用go + vue3 + ts + NaiveUI开发，旧版本使用go + vue2 + antd 开发，相关链接在下方表格：

|内容|新版本|
|:-----  |:-----|
|演示视频 |[新版]( https://www.bilibili.com/video/BV1DT411E7oV) - [旧版](https://www.bilibili.com/video/BV1TA411F7xz)  |
|新版项目地址 |[gitee](https://gitee.com/wzmgit/danmu-v5) - [github](https://github.com/wangzmgit/danmu-v5)  |
|部署文档 |[新版](http://www.kuukaa.fun/quick_use/config.html) - [旧版](http://old.kuukaa.fun/quick_use/config.html)  |
|接口文档 |[新版](http://www.kuukaa.fun/api/explain.html) - [旧版](http://old.kuukaa.fun/api/status.html)  |
|旧版前端地址 |[gitee](https://gitee.com/wzmgit/vue-danmu) - [github](https://github.com/wangzmgit/vue-danmu)  |
|旧版后端地址 |[gitee](https://gitee.com/wzmgit/go-danmu) - [github](https://github.com/wangzmgit/go-danmu)  |
|旧版移动端地址 |[gitee](https://gitee.com/wzmgit/danmu-mobile) - [github](https://github.com/wangzmgit/danmu-mobile)  |
|旧版管理端地址 |[gitee](https://gitee.com/wzmgit/danmu-admin) - [github](https://github.com/wangzmgit/danmu-admin)  |


## 使用教程
使用前需安装MySQL、Redis、Docker。安装完成后，进入后端目录，按照 README.md 文件所示创建配置文件，下载前端发行版。使用以下命令构建并启动项目。
```
docker build -t "danmu" .

docker run -itd --name danmuV5 -v /usr/danmaku:/danmu/go-danmu/file -p 9000:9000 danmu
```
启动成功后，进入后台管理系统, 使用默认账号`danmu@admin.com`密码`123456` 进行登录，登录成功后可对网站进行配置。


## 项目截图

![首页](http://www.kuukaa.fun/img/home.png)

![个人中心](http://www.kuukaa.fun/img/space.png)

![投稿](http://www.kuukaa.fun/img/upload.png)

![视频](http://www.kuukaa.fun/img/video.png)

![收藏夹](http://www.kuukaa.fun/img/collection.png)

![后台管理](http://www.kuukaa.fun/img/manage_about.png)

## 移动端
![移动端](http://www.kuukaa.fun/img/mobile.jpeg)