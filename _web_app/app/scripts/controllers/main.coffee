'use strict'

angular.module('podcasterApp')
  .controller 'MainCtrl', ($scope, popular, $rootScope) ->
  	$scope.popular = popular
  	$rootScope.title = "Podcaster - Popular"
  	$rootScope.description = "Podcaster, the RESTfull app"