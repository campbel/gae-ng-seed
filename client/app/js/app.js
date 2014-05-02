(function() {
	'use strict';

	angular.module('myApp', ['ngRoute', 'ngAnimate'])
		.config(['$routeProvider', '$httpProvider', function($routeProvider, $http) {
			$routeProvider.when('/', {templateUrl: 'views/todo.html', controller: 'TodoCtl'});
      $routeProvider.otherwise({redirectTo: '/'});
		}]);
}());
