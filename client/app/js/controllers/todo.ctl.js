angular.module('myApp')
  .controller('TodoCtl', [
    '$scope', 'Todo',
    function($scope, Todo) {
      var loadTasks = function() {
        Todo.get().then(
          function(response) {
            $scope.tasks = response.tasks;
          },
          function(response) {
            $scope.error = true;
          }
        );
      };
      loadTasks();
    }
  ]);
