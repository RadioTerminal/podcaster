'use strict'

describe 'Controller: PodcastsCtrl', ->

  # load the controller's module
  beforeEach module 'podcasterApp'

  PodcastsCtrl = {}
  scope = {}

  # Initialize the controller and a mock scope
  beforeEach inject ($controller, $rootScope) ->
    scope = $rootScope.$new()
    PodcastsCtrl = $controller 'PodcastsCtrl', {
      $scope: scope
    }

  it 'should attach a list of awesomeThings to the scope', ->
    expect(scope.awesomeThings.length).toBe 3
