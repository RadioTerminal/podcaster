'use strict'

describe 'Controller: MediaoneCtrl', ->

  # load the controller's module
  beforeEach module 'podcasterApp'

  MediaoneCtrl = {}
  scope = {}

  # Initialize the controller and a mock scope
  beforeEach inject ($controller, $rootScope) ->
    scope = $rootScope.$new()
    MediaoneCtrl = $controller 'MediaoneCtrl', {
      $scope: scope
    }

  it 'should attach a list of awesomeThings to the scope', ->
    expect(scope.awesomeThings.length).toBe 3
