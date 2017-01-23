exports.config = {
  files: {
    javascripts: { joinTo: 'helix.js' },
    stylesheets: { joinTo: 'helix.css' }
  },

  plugins: {
    babel: {
      presets: ['es2015', 'react'],
      plugins: ['transform-es2015-destructuring', 'transform-object-rest-spread']
    },
    sass: {
      options: {
        includePaths: ['node_modules/bootstrap-sass/assets/stylesheets'],
        precision: 8
      }
    },
    copycat: {
      "fonts": ["node_modules/bootstrap-sass/assets/fonts"]
    }
  },

  npm: {
    enabled: true,
    globals: {
      $: 'jquery',
      jQuery: 'jquery'
    }
  }
};
