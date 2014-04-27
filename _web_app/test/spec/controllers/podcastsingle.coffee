'use strict'

describe 'Controller: PodcastsingleCtrl', ->

  # load the controller's module
  beforeEach module 'podcasterApp'

  PodcastsingleCtrl = {}
  scope = {}

  # Initialize the controller and a mock scope
  beforeEach inject ($controller, $rootScope) ->
    scope = $rootScope.$new()
    PodcastsingleCtrl = $controller 'PodcastsingleCtrl', {
      $scope: scope
    }

  it 'should attach a list of awesomeThings to the scope', ->
    expect(scope.awesomeThings.length).toBe 3
