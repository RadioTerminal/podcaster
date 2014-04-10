'use strict'

angular.module('podcasterApp')
  .controller 'PodcastsCtrl', ($scope, groups, $log) ->
    $scope.groups = groups
    $log.info groups