// App
(function() {
'use strict';

	angular.module('myApp', ['myApp.filters', 'myApp.services', 'myApp.directives', 'myApp.controllers', 'ngRoute', 'ngAnimate'])
		.config(['$routeProvider', function($routeProvider) {
		  $routeProvider.when('/', {templateUrl: 'partials/home.html', controller: 'HomeCtrl'});
		  $routeProvider.otherwise({redirectTo: '/'});
		}]);

}());

