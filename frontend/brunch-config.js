exports.config = {
  files: {
    javascripts: { joinTo: 'helix.js' },
    stylesheets: { joinTo: 'helix.css' }
  },

  plugins: {
    babel: { presets: ['es2015', 'react'] }
  }
};
