(function() {
  'use strict';

  describe('Todo Controller', function(){
    var scope, todo, dm;

    beforeEach(module('myApp'));

    beforeEach(inject(function($controller, $q, $rootScope, Todo) {
      scope = $rootScope.$new();
      todo = Todo;

      dm = $q.defer();
      spyOn(todo, 'get').andReturn(dm.promise);
      dm.resolve({tasks:['foo','bar','baz']});

      $controller('TodoCtl', {
        $scope: scope,
        Todo: todo
      });

      $rootScope.$digest();
    }));

    describe('on load', function() {
      it('correctly loads the tasks', function() {
        expect(scope.tasks).toEqual(['foo','bar','baz']);
      });
    });
  });

}());
