'use strict'

angular
  .module('podcasterApp', [
    'ngCookies',
    'ngResource',
    'ngSanitize',
    'ngRoute',
    'mgcrea.ngStrap',
    'ngAnimate',
  ])
  .config ($routeProvider) ->
    $routeProvider
      .when '/',
        templateUrl: 'views/main.html'
        controller: 'MainCtrl'
      .when '/media',
        templateUrl: 'views/media.html'
        controller: 'MediaCtrl'
      .when '/media/:id',
        templateUrl: 'views/mediaone.html'
        controller: 'MediaoneCtrl'
      .when '/podcasts',
        templateUrl: 'views/podcasts.html'
        controller: 'PodcastsCtrl'
      .when '/podcast/:slug',
        templateUrl: 'views/podcastone.html'
        controller: 'PodcastoneCtrl'
      .otherwise
        redirectTo: '/'

