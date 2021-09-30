module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  configureWebpack: {
    output: {
      filename: '[name].[hash].js'
    }
  },
  outputDir: '../sub/static',
  devServer: {
    host: '0.0.0.0',
    port: 19528,
    disableHostCheck: true,
    proxy: {
      '^/api': {
        target: 'http://127.0.0.1:4001',
        ws: true,
        changeOrigin: true
      },
      '^/media': {
        target: 'http://127.0.0.1:4001',
        ws: true,
        changeOrigin: true
      }
    }
  }
}
