angular.module('myApp')
  .factory('Todo', [
    '$http', '$q',
    function($http, $q) {
      return {
        get: function() {
          return $http.get('/api/todo').then(
            function(response) {
              return response;
            },
            function(response) {
              return $q.reject(response);
            }
          );
        },
        put: function(task) {
          return $http.put('/api/todo', {task:task}).then(
            function(response) {
              return response;
            },
            function(response) {
              return $q.reject(response);
            }
          );
        },
        del: function(id) {
          return $http.delete('/api/todo/' + id).then(
            function(response) {
              return response;
            },
            function(response) {
              return $q.reject(response);
            }
          );
        }
      };
    }
  ]);
