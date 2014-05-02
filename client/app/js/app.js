(function() {
	'use strict';

	angular.module('myApp', ['ngRoute', 'ngAnimate'])
		.config(['$routeProvider', '$httpProvider', function($routeProvider, $http) {
			$routeProvider.when('/', {templateUrl: 'views/home.html', controller: 'HomeCtrl'});
      $routeProvider.when('/customers', {templateUrl: 'views/customers.html', controller: 'CustomersCtrl'});
      $routeProvider.when('/customers/:customerId', {templateUrl: 'views/customer.html', controller: 'CustomerCtrl'});
			$routeProvider.otherwise({redirectTo: '/'});
		}]);
}());
