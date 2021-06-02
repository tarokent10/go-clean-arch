module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      '/v1': {
        target: 'http://localhost:3000', 
        changeOrigin: true
      }
    }
  },
  configureWebpack: {
    devtool: 'source-map'
  }

}
