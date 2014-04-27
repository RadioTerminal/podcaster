'use strict'

describe 'Controller: PodcastoneCtrl', ->

  # load the controller's module
  beforeEach module 'podcasterApp'

  PodcastoneCtrl = {}
  scope = {}

  # Initialize the controller and a mock scope
  beforeEach inject ($controller, $rootScope) ->
    scope = $rootScope.$new()
    PodcastoneCtrl = $controller 'PodcastoneCtrl', {
      $scope: scope
    }

  it 'should attach a list of awesomeThings to the scope', ->
    expect(scope.awesomeThings.length).toBe 3
