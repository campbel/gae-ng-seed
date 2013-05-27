(function() {
'use strict';

// Declare app level module which depends on filters, and services
angular.module('myApp', ['myApp.filters', 'myApp.services', 'myApp.directives', 'myApp.controllers'])
  .config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/', {templateUrl: 'partials/home.html', controller: 'HomeController'});
    $routeProvider.when('/settings', {templateUrl: 'partials/settings.html', controller: 'SettingsController'});
    $routeProvider.otherwise({redirectTo: '/'});
  }]);
}());