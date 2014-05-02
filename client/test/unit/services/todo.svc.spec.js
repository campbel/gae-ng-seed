(function() {
  'use strict';

  describe('Todo Service', function(){
    var httpBackend, todo;

    beforeEach(module('myApp'));

    beforeEach(inject(function($httpBackend, Todo) {
      httpBackend = $httpBackend;
      todo = Todo;
    }));

    afterEach(function() {
      httpBackend.verifyNoOutstandingExpectation();
      httpBackend.verifyNoOutstandingRequest();
    });

    describe('get', function() {
      it('correctly returns the list of tasks', function() {
        httpBackend.expectGET('/api/todo').respond({tasks:['foo','bar','baz']});

        todo.get().then(function(response) {
          expect(response.data).toEqual({tasks:['foo','bar','baz']});
        });

        httpBackend.flush();
      });

      it('correctly returns an error', function() {
        httpBackend.expectGET('/api/todo').respond(500);

        todo.get().then(null, function(response) {
          expect(response.status).toEqual(500);
        });

        httpBackend.flush();
      });
    });

    describe('put', function() {
      it('correctly adds the task', function() {
        httpBackend.expectPUT('/api/todo', {task: 'foo'}).respond({task:'foo'});

        todo.put('foo').then(function(response) {
          expect(response.data).toEqual({task: 'foo'});
        });

        httpBackend.flush();
      });

      it('correctly returns an error', function() {
        httpBackend.expectPUT('/api/todo').respond(500);

        todo.put().then(null, function(response) {
          expect(response.status).toEqual(500);
        });

        httpBackend.flush();
      });
    });

    describe('del', function() {
      it('correctly deletes the task', function() {
        httpBackend.expectDELETE('/api/todo/1').respond({task:'foo'});

        todo.del(1).then(function(response) {
          expect(response.data).toEqual({task: 'foo'});
        });

        httpBackend.flush();
      });

      it('correctly returns an error', function() {
        httpBackend.expectDELETE('/api/todo/1').respond(500);

        todo.del(1).then(null, function(response) {
          expect(response.status).toEqual(500);
        });

        httpBackend.flush();
      });
    });
  });
}());
