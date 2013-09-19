module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    watch: {
      options: {
        atBegin: true
      },
      css: {
        files: ['../app/css/app.scss'],
        tasks: ['sass'],
        options: {
          livereload: true,   
        },
      },
      js: {
        files: ['../app/js/*.js', '../test/unit/*.js'],
        tasks: ['karma:dev:run', 'jshint']
      }
    },
    jshint: {
      options: {
        globals: {
          describe: true,
          it: true,
          module: true,
          angular: true,
          expect: true,
          beforeEach: true,
          afterEach: true
        }
      },
      all: ['gruntfile.js', '../app/js/*.js', '../test/unit/*.js']
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
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-karma');

  // Default task(s).
  grunt.registerTask('default', ['karma:dev', 'watch']);

};