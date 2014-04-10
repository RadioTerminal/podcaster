'use strict'

app = angular
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
  .config ($routeProvider) ->
    $routeProvider
      .when '/',
        templateUrl: 'views/main.html'
        controller: 'MainCtrl'
        resolve: 
          latest: (Restangular)->
              Restangular.all('latest').getList() 
          popular: (Restangular)->
              Restangular.all('popular').getList()
      .when '/media',
        templateUrl: 'views/media.html'
        controller: 'MediaCtrl'
        resolve:
          media: (Restangular)->
              Restangular.all('media').getList()
      .when '/podcasts',
        templateUrl: 'views/podcasts.html'
        controller: 'PodcastsCtrl'
        resolve:
          groups: (Restangular)->
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
          podcast: ($route, Restangular)->
              Restangular.one('group', $route.current.params.slug).get()
          media: ($route, Restangular)->
              Restangular.one('group', $route.current.params.slug).one("media").get()
      .otherwise
        redirectTo: '/'

app.run (Restangular, $alert, $location)->
  Restangular.setBaseUrl('/api')
  Restangular.setErrorInterceptor (res)->
    if res.status = 404
      $location.path('/')
      $alert
        title: res.statusText
        content: res.data.error
        placement: "top-right"
        type: "warning"
        duration: 5
        show: true
    return false