module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      '/v1': {
        target: 'http://localhost:8080', 
        changeOrigin: true
      }
    }
  },
  configureWebpack: {
    devtool: 'source-map'
  }

}
