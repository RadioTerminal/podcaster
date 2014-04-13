'use strict'

angular.module('podcasterApp')
  .controller 'PodcastsCtrl', ($scope, groups, $rootScope) ->
    $scope.groups = groups
    $rootScope.title = "Podcaster - Podcasts"
    $rootScope.description = "Podcaster, the RESTfull app"