'use strict'

angular.module('podcasterApp')
  .controller 'MainCtrl', ($scope, latest, popular, $rootScope) ->
  	$scope.latest = latest
  	$scope.popular = popular
  	$rootScope.title = "Podcaster - Latest"
  	$rootScope.description = "Podcaster, the RESTfull app"