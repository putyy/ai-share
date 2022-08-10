const DefinePlugin = require('@wepy/plugin-define');
const path = require('path')
var prod = process.env.NODE_ENV === 'production'
let config = {
  dev: {
    apiUrl: 'https://ai-share-api.xxx.com'
  },
  online: {
    apiUrl: 'https://ai-share-api.xxx.com'
  }
}
let env = process.env.NODE_ENV === 'production' ? 'online' : 'dev'

module.exports = {
  wpyExt: '.wpy',
  eslint: true,
  cliLogs: !prod,
  static: ['static'],
  build: {},
  resolve: {
    alias: {
      '@': path.join(__dirname, 'src')
    },
    aliasFields: ['wepy', 'weapp'],
    modules: ['node_modules']
  },
  compilers: {
    less: {
      compress: prod
    },
    babel: {
      sourceMap: true,
      presets: [
        '@babel/preset-env'
      ],
      plugins: [
        '@wepy/babel-plugin-import-regenerator'
      ]
    }
  },
  plugins: [
    DefinePlugin({
      API_URL: JSON.stringify(config[env].apiUrl)
    })
  ],
  appConfig: {
    noPromiseAPI: ['createSelectorQuery']
  }
}
