// App
(function() {
'use strict';

	angular.module('myApp', ['ngRoute', 'ngAnimate'])
		.config(['$routeProvider', '$httpProvider', function($routeProvider, $http) {
			$routeProvider.when('/', {templateUrl: 'views/home.html', controller: 'HomeCtrl'});
			$routeProvider.otherwise({redirectTo: '/'});

      /* Setting up XSRF
			var token = document.getElementById("xsrftoken").value;
			$http.defaults.headers.common["X-XSRF-token"] = token;
      */
		}]);
}());

