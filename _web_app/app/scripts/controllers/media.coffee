'use strict'

angular.module('podcasterApp')
  .controller 'MediaCtrl', ($scope, media, $rootScope) ->
    $scope.media = media
    $rootScope.title = "Podcaster - Media"
    $rootScope.description = "Podcaster, the RESTfull app"