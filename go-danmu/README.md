# 弹幕网站后端

## 配置文件
在config文件夹内创建application.yml文件
```yml
admin:
  email: danmu@admin.com
  password: e10adc3949ba59abbe56e057f20f883e
datasource:
  database: 数据库名
  host: 127.0.0.1
  password: 数据库密码
  port: 3306
  username: 数据库用户名
file:
  domain: 
  https: false
  max_img_size: 5
  max_video_size: 500
  oss: false
register:
  prefix: user_
  verify_email: false
server:
  fe_path: 前端文件目录
  jwt_secret: jwt秘钥(随便写点字母数字)
  port: 9100
transcoding:
  max_res: 0

```

## 运行
```
go run main.go
```
### 编译打包
```
go build
```

