(function() {
	'use strict';

	angular.module('myApp')
		.controller('CustomersCtrl', [
      '$scope', 'customers',
      function($scope, customers) {

        var load = function() {
          customers.read()
            .then(function(response) {
              if(response !== "null") {
                $scope.customers = response;
              }
            });
        };

        $scope.save = function() {
          customers.create($scope.newCustomer)
            .then(function() {
              load();
            });
        };

        load();
      }
    ]);

}());
