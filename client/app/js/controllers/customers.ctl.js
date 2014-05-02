(function() {
	'use strict';

	angular.module('myApp')
		.controller('CustomersCtrl', [
      '$scope', 'customers',
      function($scope, customers) {

        $scope.customers = [];

        var load = function() {
          customers.read()
            .then(function(response) {
              if(response !== 'null') {
                $scope.customers = response;
              }
            });
        };

        load();

        $scope.save = function() {
          customers.create($scope.newCustomer)
            .then(function(customer) {
              $scope.customers.push(customer);
              $scope.newCustomer = {};
            });
        };
      }
    ]);

}());
