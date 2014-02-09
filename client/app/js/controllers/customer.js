(function() {
	'use strict';

	angular.module('myApp')
		.controller('CustomerCtrl', [
      '$scope', '$routeParams', 'customers',
      function($scope, $routeParams, customers) {
        var load = function() {
          customers.read($routeParams.customerId)
            .then(function(response) {
              if(response !== 'null') {
                $scope.customer = response;
              }
            });
        };

        $scope.save = function() {
          
        };

        load();
      }
    ]);

}());
