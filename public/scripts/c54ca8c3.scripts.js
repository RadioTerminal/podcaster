(function () {
  'use strict';
  var app;
  app = angular.module('podcasterApp', [
    'ngSanitize',
    'ngRoute',
    'mgcrea.ngStrap',
    'ngAnimate',
    'chieffancypants.loadingBar',
    'restangular',
    'angulartics',
    'angulartics.google.analytics'
  ]).config([
    '$routeProvider',
    function ($routeProvider) {
      return $routeProvider.when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl',
        resolve: {
          popular: function (Restangular) {
            return Restangular.all('popular').getList();
          }
        }
      }).when('/media', {
        templateUrl: 'views/media.html',
        controller: 'MediaCtrl',
        resolve: {
          media: function (Restangular) {
            return Restangular.all('media').getList();
          }
        }
      }).when('/podcasts', {
        templateUrl: 'views/podcasts.html',
        controller: 'PodcastsCtrl',
        resolve: {
          groups: function (Restangular) {
            return Restangular.all('groups').getList();
          }
        }
      }).when('/media/:mediaId', {
        templateUrl: 'views/mediaone.html',
        controller: 'MediaoneCtrl',
        resolve: {
          data: function ($route, Restangular) {
            return Restangular.one('media', $route.current.params.mediaId).get();
          }
        }
      }).when('/podcast/:slug', {
        templateUrl: 'views/podcastone.html',
        controller: 'PodcastoneCtrl',
        resolve: {
          podcast: function ($route, Restangular) {
            return Restangular.one('group', $route.current.params.slug).get();
          },
          media: function ($route, Restangular) {
            return Restangular.one('group', $route.current.params.slug).one('media').get();
          }
        }
      }).otherwise({ redirectTo: '/' });
    }
  ]);
  app.run([
    'Restangular',
    '$alert',
    '$location',
    '$rootScope',
    function (Restangular, $alert, $location, $rootScope) {
      $rootScope.$watch('description', function () {
        return angular.element('head meta[name=description]').attr('content', $rootScope.description);
      });
      $rootScope.$watch('title', function () {
        return angular.element('head title').text($rootScope.title);
      });
      Restangular.setBaseUrl('/api');
      return Restangular.setErrorInterceptor(function (res) {
        if (res.status = 404) {
          $location.path('/');
          $alert({
            title: res.statusText,
            content: res.data.error,
            placement: 'top-right',
            type: 'warning',
            duration: 5,
            show: true
          });
        }
        return false;
      });
    }
  ]);
}.call(this));
/*
//@ sourceMappingURL=app.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').controller('MainCtrl', [
    '$scope',
    'popular',
    '$rootScope',
    function ($scope, popular, $rootScope) {
      $scope.popular = popular;
      $rootScope.title = 'Podcaster - Popular';
      return $rootScope.description = 'Podcaster, the RESTfull app';
    }
  ]);
}.call(this));
/*
//@ sourceMappingURL=main.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').controller('MediaCtrl', [
    '$scope',
    'media',
    '$rootScope',
    function ($scope, media, $rootScope) {
      $scope.media = media;
      $rootScope.title = 'Podcaster - Media';
      return $rootScope.description = 'Podcaster, the RESTfull app';
    }
  ]);
}.call(this));
/*
//@ sourceMappingURL=media.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').controller('PodcastsCtrl', [
    '$scope',
    'groups',
    '$rootScope',
    function ($scope, groups, $rootScope) {
      $scope.groups = groups;
      $rootScope.title = 'Podcaster - Podcasts';
      return $rootScope.description = 'Podcaster, the RESTfull app';
    }
  ]);
}.call(this));
/*
//@ sourceMappingURL=podcasts.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').controller('PodcastoneCtrl', [
    '$scope',
    'podcast',
    'media',
    '$rootScope',
    function ($scope, podcast, media, $rootScope) {
      $scope.podcast = podcast;
      $scope.media = media;
      $rootScope.title = 'Podcaster - ' + podcast.name;
      return $rootScope.description = podcast.description;
    }
  ]);
}.call(this));
/*
//@ sourceMappingURL=podcastone.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').directive('bootswatch', function () {
    return {
      template: '<form><select>\n  <option ng-repeat="theme in [\'lumen\',\'cerulean\',\'simplex\',\'united\',\'spacelab\',\'slate\',\'readable\',\'journal\',\'flatly\',\'cosmo\',\'cyborg\',\'amelia\',\'darkly\',\'yeti\']" value="{{theme}}">{{theme}}</option>\n</select></form>',
      restrict: 'E',
      link: function (scope, element, attrs) {
        return element.bind('change', function (event) {
          return angular.element('head link:nth(0)').attr('href', '//netdna.bootstrapcdn.com/bootswatch/3.1.1/' + event.target.value + '/bootstrap.min.css');
        });
      }
    };
  });
}.call(this));
/*
//@ sourceMappingURL=bootswatch.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').directive('waveformplayer', function () {
    return {
      restrict: 'E',
      scope: { item: '=' },
      templateUrl: '/views/player.html',
      link: function (scope, element, attr) {
        var canvas, clear, context, createCanvas, data, gradient, height, innerColor, interpolateArray, linearInterpolate, outerColor, patchCanvasForIE, redraw, toHHMMSS, width, _this = this;
        scope.audio = new Audio();
        scope.preload = 'metadata';
        scope.now = 0;
        scope.duration = 0;
        scope.cursor = 0;
        scope.cursor_over = false;
        patchCanvasForIE = function (canvas) {
          var oldGetContext;
          if (typeof window.G_vmlCanvasManager !== 'undefined') {
            canvas = window.G_vmlCanvasManager.initElement(canvas);
            oldGetContext = canvas.getContext;
            return canvas.getContext = function (a) {
              var ctx;
              ctx = oldGetContext.apply(canvas, arguments);
              canvas.getContext = oldGetContext;
              return ctx;
            };
          }
        };
        createCanvas = function (container, width, height) {
          var canvas;
          canvas = document.createElement('canvas');
          container[0].appendChild(canvas);
          canvas.width = width || 370;
          canvas.height = height || 80;
          return canvas;
        };
        toHHMMSS = function (times) {
          var hours, minutes, sec_num, seconds, time;
          sec_num = parseInt(times, 10);
          hours = Math.floor(sec_num / 3600);
          minutes = Math.floor((sec_num - hours * 3600) / 60);
          seconds = sec_num - hours * 3600 - minutes * 60;
          if (hours < 10) {
            hours = '0' + hours;
          }
          if (minutes < 10) {
            minutes = '0' + minutes;
          }
          if (seconds < 10) {
            seconds = '0' + seconds;
          }
          time = hours + ':' + minutes + ':' + seconds;
          return time;
        };
        linearInterpolate = function (before, after, atPoint) {
          return before + (after - before) * atPoint;
        };
        interpolateArray = function (data, fitCount) {
          var after, atPoint, before, i, newData, springFactor, tmp, x, _i, _len;
          newData = new Array();
          springFactor = new Number((data.length - 1) / (fitCount - 1));
          for (x = _i = 0, _len = data.length; _i < _len; x = ++_i) {
            i = data[x];
            data[x] = Number(i);
          }
          newData[0] = Number(data[0]);
          i = 1;
          while (i < fitCount - 1) {
            tmp = i * springFactor;
            before = new Number(Math.floor(tmp)).toFixed();
            after = new Number(Math.ceil(tmp)).toFixed();
            atPoint = tmp - before;
            newData[i] = linearInterpolate(data[before], data[after], atPoint);
            i++;
          }
          newData[fitCount - 1] = data[data.length - 1];
          return newData;
        };
        redraw = function () {
          var d, i, middle, t, _i, _len, _results;
          clear();
          if (typeof innerColor === 'function') {
            context.fillStyle = innerColor();
          } else {
            context.fillStyle = innerColor;
          }
          middle = height / 2;
          i = 0;
          _results = [];
          for (_i = 0, _len = data.length; _i < _len; _i++) {
            d = data[_i];
            t = width / data.length;
            if (typeof innerColor === 'function') {
              context.fillStyle = innerColor(i / width, d);
            }
            context.clearRect(t * i, middle - middle * d, t, middle * d * 2);
            context.fillRect(t * i, middle - middle * d, t, middle * d * 2);
            _results.push(i++);
          }
          return _results;
        };
        clear = function () {
          context.fillStyle = outerColor;
          context.clearRect(0, 0, width, height);
          return context.fillRect(0, 0, width, height);
        };
        scope.playpause = function () {
          if (scope.audio.paused) {
            if (scope.audio.src === '/api/media/head/' + scope.item.id) {
              scope.audio.src = '/api/media/play/' + scope.item.id;
            }
            return scope.audio.play();
          } else {
            return scope.audio.pause();
          }
        };
        scope.audio.addEventListener('ended', function () {
          return scope.$apply(function () {
            scope.audio.paused = true;
            scope.audio.currentTime = 0;
            return redraw();
          });
        }, false);
        canvas = createCanvas(element, element.parent().clientWidth, element.clientHeight);
        patchCanvasForIE(canvas);
        context = canvas.getContext('2d');
        width = parseInt(context.canvas.width, 10);
        height = parseInt(context.canvas.height, 10);
        outerColor = 'transparent';
        gradient = context.createLinearGradient(0, 0, 0, height);
        gradient.addColorStop(0, '#f60');
        gradient.addColorStop(1, '#ff1b00');
        innerColor = function (x, y) {
          var buffered, e;
          try {
            buffered = scope.audio.buffered.end(0);
          } catch (_error) {
            e = _error;
            buffered = 0;
          }
          if (x < scope.cursor / width && scope.cursor_over) {
            return '#f60';
          } else if (x < scope.audio.currentTime / scope.audio.duration) {
            return gradient;
          } else if (x < buffered / scope.audio.duration) {
            return 'rgba(0, 0, 0, 0.7)';
          } else {
            return 'rgba(0, 0, 0, 0.5)';
          }
        };
        data = [];
        canvas.addEventListener('mousedown', function (event) {
          var oncanvas;
          oncanvas = 100 / width * event.offsetX;
          if (scope.audio.paused) {
            scope.playpause();
          }
          return scope.audio.currentTime = scope.audio.duration / 100 * oncanvas;
        }, false);
        canvas.addEventListener('mouseover', function (event) {
          scope.cursor_over = true;
          scope.cursor = event.offsetX;
          return redraw();
        }, false);
        canvas.addEventListener('mouseout', function (event) {
          scope.cursor_over = false;
          scope.cursor = event.offsetX;
          return redraw();
        }, false);
        canvas.addEventListener('mousemove', function (event) {
          scope.cursor = event.offsetX;
          return redraw();
        }, false);
        scope.audio.addEventListener('timeupdate', function () {
          redraw();
          return scope.$apply(function () {
            scope.duration = toHHMMSS(scope.audio.duration);
            return scope.now = toHHMMSS(scope.audio.currentTime);
          });
        }, false);
        scope.$watch('item', function (item) {
          data = interpolateArray(item.wave.split(','), width);
          scope.item = item;
          scope.audio.src = '/api/media/head/' + item.id;
          return redraw();
        });
        return scope.$on('$destroy', function () {
          scope.audio.pause();
          scope.audio.src = '';
          scope.audio.removeAttribute('src');
          return scope.audio = null;
        });
      }
    };
  });
}.call(this));
/*
//@ sourceMappingURL=waveformplayer.js.map
*/
(function () {
  'use strict';
  angular.module('podcasterApp').directive('mediaplayer', function () {
    return {
      restrict: 'E',
      scope: { item: '=' },
      templateUrl: '/views/mediaone.html',
      link: function (scope, element, attr) {
        return scope.$watch('item', function (item) {
          return scope.item = item;
        });
      }
    };
  });
}.call(this));  /*
//@ sourceMappingURL=mediaplayer.js.map
*/