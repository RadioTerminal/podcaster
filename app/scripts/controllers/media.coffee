'use strict'

angular.module('podcasterApp')
  .controller 'MediaCtrl', ($scope, media) ->
    $scope.media = media