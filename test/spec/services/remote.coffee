'use strict'

describe 'Service: Remote', ->

  # load the service's module
  beforeEach module 'podcasterApp'

  # instantiate service
  Remote = {}
  beforeEach inject (_Remote_) ->
    Remote = _Remote_

  it 'should do something', ->
    expect(!!Remote).toBe true
