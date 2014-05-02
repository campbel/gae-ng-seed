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
            $scope.tasks = [];
          }
        );
      };
      loadTasks();

      $scope.add = function() {
        if(!$scope.newTask) return;
        $scope.tasks.push($scope.newTask);
        $scope.newTask = '';
      };
    }
  ]);
