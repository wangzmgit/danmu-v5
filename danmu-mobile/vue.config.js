const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  productionSourceMap: false,
  transpileDependencies: true,
  publicPath: process.env.NODE_ENV === "production" ? "/mobile/" : "/",
  outputDir: "mobile"
})
