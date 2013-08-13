module.exports = function(config) {
  config.set({
    basePath: '../',
    frameworks: ['jasmine'],
    autowatch: true,
    browsers: ['Chrome' /*, 'IE', 'Firefox'*/],
    files: [
		'app/lib/angular/angular.min.js',
		'test/lib/angular/angular-mocks.js',
		'app/js/*.js',
		'test/unit/**/*.js'
  	],
  });
};