'use strict'

angular.module('podcasterApp')
  .controller 'PodcastsingleCtrl', ($scope, media, $rootScope) ->
  	$scope.media = media
  	$rootScope.title = "Podcaster - #{media.name}"
  	$rootScope.description = media.text