(function() {
	'use strict';

	angular.module('myApp', ['ngRoute', 'ngAnimate'])
		.config(['$routeProvider', '$httpProvider', function($routeProvider, $http) {
			$routeProvider.when('/', {
				controller: 'TodoCtl',
				templateUrl: 'views/todo.html'
			});
		}]);
}());
