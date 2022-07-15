# 弹幕网站移动端

打包前配置(/src/config.ts)
```js
//网站标题
const title = "弹幕网站";
//后端地址(docker部署 url = "/")
const url = "http://localhost:9000/";
//桌面端地址
const desktop = "/";
const icp = "icp备案信息";
const security = "公网安备信息"
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