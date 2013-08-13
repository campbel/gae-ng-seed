module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    watch: {
      sass: {
        files: ['../app/css/app.scss'],
        tasks: ['sass']
      },
      karma: {
      	files: ['../app/js/*.js', '../test/unit/*.js'],
      	tasks: ['karma:dev:run']
      }
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
    },
    karma: {
      dev: {
      	configFile: 'karma.conf.js',
      	background: true
      }
    }
  });

  // Load the plugins.
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-karma');

  // Default task(s).
  grunt.registerTask('default', ['karma:dev', 'watch']);

};