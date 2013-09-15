// App
(function() {
'use strict';

	angular.module('myApp', ['ngRoute', 'ngAnimate'])
		.config(['$routeProvider', function($routeProvider) {
		  $routeProvider.when('/', {templateUrl: 'views/home.html', controller: 'HomeCtrl'});
		  $routeProvider.otherwise({redirectTo: '/'});
		}]);

}());

