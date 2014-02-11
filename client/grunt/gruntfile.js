module.exports = function(grunt) {
  'use strict';

  // Project configuration.
  grunt.initConfig({
    watch: {
      options: {
        atBegin: true
      },
      css: {
        files: ['../app/css/app.scss'],
        tasks: ['sass:dev'],
        options: {
          livereload: {
            port: 8090
          }
        },
      },
      js: {
        files: ['../app/js/app.js', '../app/js/**/*.js', '../test/unit/*.js'],
        tasks: ['jshint'],
        options: {
          livereload: {
            port: 8090
          },
        }
      }
    },
    uglify: {
      dist: {
        files: {
          '../app/dist/app.min.js': ['../app/js/app.js', '../app/js/**/*.js']
        }
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
      all: ['gruntfile.js', '../app/js/app.js', '../app/js/**/*.js', '../test/unit/*.js']
    },
    sass: {
      dev: {
        options: {
          style: 'expanded'
        },
        files: {
          '../app/css/app.css' : '../app/css/app.scss'
        }
      },
      dist: {
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
      },
      dist: {
        options: {
          files: [
            'app/lib/angular/angular.js',
            'app/lib/angular/angular-route.js',
            'app/lib/angular/angular-animate.js',
            'test/lib/angular/angular-mocks.js',
            'app/dist/app.min.js',
            'test/unit/*.js'
          ],
        },
        configFile: 'karma.conf.js',
        singleRun: true,
      }
    }
  });

  // Load the plugins.
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-karma');

  // Default task(s).
  grunt.registerTask('default', ['watch']);
  grunt.registerTask('dist', ['jshint', 'uglify:dist', 'sass:dist', 'karma:dist']);

};
