# 弹幕网站后台管理端

打包前配置(/src/config.js)
```js
//网站标题
const title = "弹幕网站后台管理";
//后端地址(docker部署 url = "/")
const url = "http://localhost:9000/";
//上传图片大小限制
const maxImgSize = 5;//上传图片最大大小(单位M)
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