(function() {
  'use strict';

  angular.module('myApp')
    .factory('customers', [
      '$http', '$q',
      function($http, $q) {
        return {
          read: function(id) {
            var deferred = $q.defer();

            if(id !== undefined) {
              $http.get('/api/customers/' + id + '/')
                .success(function(response) {
                  deferred.resolve(response);
                })
                .error(function(data, status) {
                  deferred.reject(status);
                });
            } else {
              $http.get('/api/customers/')
                .success(function(response) {
                  deferred.resolve(response);
                })
                .error(function(data, status) {
                  deferred.reject(status);
                });
            }

            return deferred.promise;
          },
          create: function(customer) {
            var deferred = $q.defer();

            $http.post('/api/customers/', customer)
              .success(function(response) {
                deferred.resolve(response);
              });

            return deferred.promise;
          },
          delete: function(id) {
            var deferred = $q.defer();

            $http.delete('/api/customers/' + id+ '/')
              .success(function(response) {
                deferred.resolve(response);
              })
              .error(function(data, status) {
                deferred.reject(status);
              });

            return deferred.promise;
          }
        };
      }
    ]);

}());
