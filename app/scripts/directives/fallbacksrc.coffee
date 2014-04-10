'use strict'

angular.module('podcasterApp')
  .directive('fallbackSrc', ->
    link: (scope, iElement, iAttrs) ->
      iElement.bind "error", ->
        angular.element(this).attr "src", iAttrs.fallbackSrc