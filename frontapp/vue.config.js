module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      '/v1': {
        target: 'http://go-app:3000', 
        changeOrigin: true
      }
    }
  },
  configureWebpack: {
    devtool: 'source-map'
  }

}
