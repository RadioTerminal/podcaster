'use strict'

angular.module('podcasterApp')
  .directive('mediaplayer', ->
    restrict: "E"
    scope: { item: '=' },
    templateUrl: "/views/mediaone.html"
    link: (scope, element, attr) ->
      scope.$watch 'item', (item)->
      	scope.item = item
  )
