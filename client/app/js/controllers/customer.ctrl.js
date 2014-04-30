(function() {
	'use strict';

	angular.module('myApp')
		.controller('CustomerCtrl', [
      '$scope', '$routeParams', '$location', 'customers',
      function($scope, $routeParams, $location, customers) {

        var load = function() {
          customers.read($routeParams.customerId)
            .then(function(response) {
              $scope.customer = response;
            }, function() {
              $location.path('/customers');
            });
        };

        load();

        $scope.delete = function() {
          customers.delete($routeParams.customerId)
            .then(function() {
              $location.path('/customers');
            });
        };
      }
    ]);

}());
