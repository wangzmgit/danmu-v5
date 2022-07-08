# 弹幕网站前端

打包前配置(/src/config.js)
```js
//网站标题
const title = "弹幕网站";
//后端地址(docker部署 url = "/")
const url = "http://localhost:9000/";
//移动端地址
const mobile = "/mobile/";
const icp = "icp备案信息";
const security = "公网安备信息"
//上传文件大小限制
const maxImgSize = 5;//上传图片最大大小(单位M)
const maxVideoSize = 500;//上传视频最大大小(单位M)
```

## 项目初始化
```
npm install
```
### 编译运行
```
npm run serve
```

### 编译打包
```
npm run build
```