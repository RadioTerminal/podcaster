'use strict'

angular.module('podcasterApp')
  .controller 'MainCtrl', ($scope, latest, popular) ->
  	$scope.latest = latest
  	$scope.popular = popular