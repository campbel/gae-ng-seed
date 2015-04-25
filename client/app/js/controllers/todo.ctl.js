(function() {
  'use strict';

  angular.module('myApp')
    .controller('TodoCtl', [
      '$scope',
      function($scope) {
        $scope.tasks = [];
        $scope.add = function(task) {
          $scope.newTask = "";
          $scope.tasks.push(task);
        };
        $scope.remove = function(index) {
          $scope.tasks.splice(index, 1);
        }
      }
    ])
}());
