'use strict'

angular.module('podcasterApp')
  .controller 'PodcastoneCtrl', ($scope, podcast, media,$rootScope) ->
  	$scope.podcast = podcast
  	$scope.media = media
  	$rootScope.title = "Podcaster - #{podcast.name}"
  	$rootScope.description = podcast.description