basePath = '../';

files = [
  JASMINE,
  JASMINE_ADAPTER,
  'lib/angular/angular.js',
  'lib/angular/angular-*.js',
  'test/lib/angular/angular-mocks.js',
  'src/js/*.js',
  'test/unit/**/*.js'
];

autoWatch = true;

browsers = ['Chrome'];

junitReporter = {
  outputFile: 'test_out/unit.xml',
  suite: 'unit'
};
