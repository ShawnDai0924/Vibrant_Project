const { defineConfig } = require('@vue/cli-service')

module.exports = {
  devServer: {
    proxy: {
      '/graphql': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: { '^/graphql': '' },
      },
    },
  },
};