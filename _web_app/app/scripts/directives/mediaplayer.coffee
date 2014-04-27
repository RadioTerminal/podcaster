'use strict'

angular.module('podcasterApp')
  .directive('mediaplayer', ->
    restrict: "AE"
    scope: { item: '=', podcast: '=' },
    templateUrl: "/views/mediaone.html"
    link: (scope, element, attr) ->
      scope.$watch 'podcast', (podcast)->
      	scope.podcast = podcast
      scope.$watch 'item', (item)->
      	scope.item = item
  )
