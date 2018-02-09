const path = require('path');
const webpack = require('webpack');
const ExtractText = require('extract-text-webpack-plugin');

const extractCSS = new ExtractText('css/style.css');

module.exports = {
  context: path.join(__dirname, 'client'),
  entry: [
    'whatwg-fetch',
    './main.js'
  ],
  devServer: {
    contentBase: 'static',
    compress: true,
    port: 3000
  },
  output: {
    path: path.join(__dirname, 'static'),
    filename: 'js/bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [
              [
                'env',
                {
                  useBuiltIns: true,
                  targets: {
                    safari: 8
                  }
                }
              ]
            ],
            'plugins': [
              'transform-runtime'
            ]
          }
        },
      },
      {
        test: /\.sass$/,
        use: extractCSS.extract({
          fallback: 'style-loader',
          use: ['css-loader', 'sass-loader']
        })
      }
    ],
  },
  resolve: {
    modules: [path.join(__dirname, 'client'), 'node_modules'],
    descriptionFiles: ['package.json'],
    extensions: ['.js', '.ts']
  },
  node: {
   fs: "empty"
  },
  plugins: [
    new webpack.DefinePlugin({
      config: {
        master: {
          server: JSON.stringify('https://prod.service')
        },
        stage: {
          server: JSON.stringify('https://stage.service')
        },
        local: {
          server: JSON.stringify('http://localhost:9000')
        }
      },
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'local')
    }),
    extractCSS  
  ]
};
