const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  productionSourceMap: false,
  transpileDependencies: true,
  lintOnSave: false,
  publicPath: process.env.NODE_ENV === "production" ? "/manage/" : "/",
  outputDir: "manage"
})
