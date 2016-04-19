var webpack = require('webpack');
var path = require('path');

var BUILD_DIR = path.resolve(__dirname, 'dist');
var APP_DIR = path.resolve(__dirname, 'src/app');

var config = {
  entry: APP_DIR + '/index.jsx',
  output:  {
    path: BUILD_DIR,
    filename: 'helix.js'
  },
  module: {
    loaders: [
      {
	test:  /\.jsx?/,
	include: APP_DIR,
	loader: 'babel'
      }
    ]
  }
};

if (process.env.NODE_ENV == 'production') {
  config.plugins = [
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'development')
      }
    })
  ];
}

module.exports = config;
