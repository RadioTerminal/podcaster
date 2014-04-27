'use strict'

angular.module('podcasterApp')
  .directive('bootswatch', ->
    template: '''<form><select>
                  <option ng-repeat="theme in ['lumen','cerulean','simplex','united','spacelab','slate','readable','journal','flatly','cosmo','cyborg','amelia','darkly','yeti']" value="{{theme}}">{{theme}}</option>
                </select></form>'''
    restrict: 'E'
    link: (scope, element, attrs) ->
      element.bind "change", (event)->
        angular.element("head link:nth(0)").attr("href", "//netdna.bootstrapcdn.com/bootswatch/3.1.1/#{event.target.value}/bootstrap.min.css")          
  )
