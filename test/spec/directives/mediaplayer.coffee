'use strict'

describe 'Directive: mediaplayer', ->

  # load the directive's module
  beforeEach module 'podcasterApp'

  scope = {}

  beforeEach inject ($controller, $rootScope) ->
    scope = $rootScope.$new()

  it 'should make hidden element visible', inject ($compile) ->
    element = angular.element '<mediaplayer></mediaplayer>'
    element = $compile(element) scope
    expect(element.text()).toBe 'this is the mediaplayer directive'
