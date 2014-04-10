'use strict'

angular
  .module('podcasterApp', [
    'ngCookies',
    'ngResource',
    'ngSanitize',
    'ngRoute',
    'mgcrea.ngStrap',
    'ngAnimate',
    'chieffancypants.loadingBar',
    'restangular'
  ])
  .config ($routeProvider, RestangularProvider) ->
    RestangularProvider.setBaseUrl('/api')
    $routeProvider
      .when '/',
        templateUrl: 'views/main.html'
        controller: 'MainCtrl'
        resolve: 
          data: (Restangular)->
              Restangular.all('latest').getList()
      .when '/media',
        templateUrl: 'views/media.html'
        controller: 'MediaCtrl'
        resolve:
          data: (Restangular)->
              Restangular.all('media').getList()
      .when '/podcasts',
        templateUrl: 'views/podcasts.html'
        controller: 'PodcastsCtrl'
        resolve:
          data: (Restangular)->
              Restangular.all('groups').getList()
      .when '/media/:mediaId',
        templateUrl: 'views/mediaone.html'
        controller: 'MediaoneCtrl'
        resolve:
          data: ($route, Restangular)->
              Restangular.one('media', $route.current.params.mediaId).get()
      .when '/podcast/:slug',
        templateUrl: 'views/podcastone.html'
        controller: 'PodcastoneCtrl'
        resolve:
          data: ($route, Restangular)->
              Restangular.one('group', $route.current.params.slug).get()
      .otherwise
        redirectTo: '/'

