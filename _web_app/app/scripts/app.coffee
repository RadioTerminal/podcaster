'use strict'

app = angular
  .module('podcasterApp', [
    'ngSanitize',
    'ngRoute',
    'mgcrea.ngStrap',
    'ngAnimate',
    'chieffancypants.loadingBar',
    'restangular',
    'angulartics',
    'angulartics.google.analytics'
  ])
  .config ($routeProvider,$locationProvider) ->
    $locationProvider.hashPrefix('!')
    $routeProvider
      .when '/',
        templateUrl: 'views/main.html'
        controller: 'MainCtrl'
        resolve: 
          popular: (Restangular)->
              Restangular.all('popular').getList()
      .when '/media',
        templateUrl: 'views/media.html'
        controller: 'MediaCtrl'
        resolve:
          media: (Restangular)->
              Restangular.all('latest').getList()
      .when '/podcasts',
        templateUrl: 'views/podcasts.html'
        controller: 'PodcastsCtrl'
        resolve:
          groups: (Restangular)->
              Restangular.all('groups').getList()
      .when '/podcast/:slug',
        templateUrl: 'views/podcastone.html'
        controller: 'PodcastoneCtrl'
        resolve:
          podcast: ($route, Restangular)->
              Restangular.one('group', $route.current.params.slug).get()
          media: ($route, Restangular)->
              Restangular.one('group', $route.current.params.slug).one("media").get()
      .when '/media/:media_slug',
        templateUrl: 'views/podcastsingle.html'
        controller: 'PodcastsingleCtrl'
        resolve:
          media: ($route, Restangular)->
              Restangular.one("media", $route.current.params.media_slug).get()
      .otherwise
        redirectTo: '/'

app.run (Restangular, $alert, $location, $rootScope)->
  $rootScope.$watch 'description', () ->
    angular.element("head meta[name=description]").attr("content", $rootScope.description)
  $rootScope.$watch 'title', () ->
    angular.element("head title").text($rootScope.title)  
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