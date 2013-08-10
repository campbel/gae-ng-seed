module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    watch: {
      scripts: {
        files: ['../app/css/app.scss'],
        tasks: ['sass'],
        options: {
          spawn: false,
        },
      },
    },
    sass: {
      dev: {
        options: {
          style: 'expanded'
        },
        files: {
          '../app/css/app.css' : '../app/css/app.scss'
        }
      }
    }
  });

  // Load the plugin that provides the "uglify" task.
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['watch']);

};