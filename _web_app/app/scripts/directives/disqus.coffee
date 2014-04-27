'use strict'

angular.module('podcasterApp')
  .directive('disqus', ($window) ->
      restrict: "AE"
      scope:
        disqus_shortname: "@disqusShortname"
        disqus_identifier: "@disqusIdentifier"
        disqus_title: "@disqusTitle"
        disqus_url: "@disqusUrl"
        disqus_category_id: "@disqusCategoryId"
        disqus_disable_mobile: "@disqusDisableMobile"
        readyToBind: "@"

      template: "<div id=\"disqus_thread\"></div><a href=\"http://disqus.com\" class=\"dsq-brlink\">comments powered by <span class=\"logo-disqus\">Disqus</span></a>"
      link: (scope) ->
        
        # ensure that the disqus_identifier and disqus_url are both set, otherwise we will run in to identifier conflicts when using URLs with "#" in them
        # see http://help.disqus.com/customer/portal/articles/662547-why-are-the-same-comments-showing-up-on-multiple-pages-
        throw "Please ensure that the `disqus-identifier` and `disqus-url` attributes are both set."  if typeof scope.disqus_identifier is "undefined" or typeof scope.disqus_url is "undefined"
        scope.$watch "readyToBind", (isReady) ->
          
          # If the directive has been called without the 'ready-to-bind' attribute, we
          # set the default to "true" so that Disqus will be loaded straight away.
          isReady = "true"  unless angular.isDefined(isReady)
          if scope.$eval(isReady)
            
            # put the config variables into separate global vars so that the Disqus script can see them
            $window.disqus_shortname = scope.disqus_shortname
            $window.disqus_identifier = scope.disqus_identifier
            $window.disqus_title = scope.disqus_title
            $window.disqus_url = scope.disqus_url
            $window.disqus_category_id = scope.disqus_category_id
            $window.disqus_disable_mobile = scope.disqus_disable_mobile
            
            # get the remote Disqus script and insert it into the DOM, but only if it not already loaded (as that will cause warnings)
            unless $window.DISQUS
              dsq = document.createElement("script")
              dsq.type = "text/javascript"
              dsq.async = true
              dsq.src = "//" + scope.disqus_shortname + ".disqus.com/embed.js"
              (document.getElementsByTagName("head")[0] or document.getElementsByTagName("body")[0]).appendChild dsq
            else
              $window.DISQUS.reset
                reload: true
                config: ->
                  @page.identifier = scope.disqus_identifier
                  @page.url = scope.disqus_url
                  @page.title = scope.disqus_title
    )