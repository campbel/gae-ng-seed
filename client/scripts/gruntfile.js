module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    watch: {
      scripts: {
        files: ['../app/css/app.scss'],
        tasks: ['sass.dev'],
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
    karma: {
      dev: {
        basePath: '../',
        frameworks: ['jasmine'],
        autowatch: true,
        browsers: ['Chrome' /*, 'IE', 'Firefox'*/],
        files: [
        'app/lib/angular/angular.min.js',
        'test/lib/angular/angular-mocks.js',
        'app/js/*.js',
        'test/unit/**/*.js'
        ]
      }
    }
  });

  // Load the plugins.
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['watch', 'karma']);

};