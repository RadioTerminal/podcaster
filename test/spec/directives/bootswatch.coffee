'use strict'

describe 'Directive: bootswatch', ->

  # load the directive's module
  beforeEach module 'podcasterApp'

  scope = {}

  beforeEach inject ($controller, $rootScope) ->
    scope = $rootScope.$new()

  it 'should make hidden element visible', inject ($compile) ->
    element = angular.element '<bootswatch></bootswatch>'
    element = $compile(element) scope
    expect(element.text()).toBe 'this is the bootswatch directive'
