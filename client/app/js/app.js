(function() {
	'use strict';

	angular.module('myApp', ['ngRoute'])
		.config(['$routeProvider', '$httpProvider', function($routeProvider, $http) {
			$routeProvider.when('/', {
				controller: 'TodoCtl',
				templateUrl: 'views/todo.html'
			});
		}]);
}());
