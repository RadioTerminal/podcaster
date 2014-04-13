'use strict'

angular.module('podcasterApp')
  .directive('waveformplayer', ->
    restrict: "E"
    templateUrl: "/views/player.html"
    link: (scope, element, attr) ->
      scope.audio = new Audio()  
      scope.audio.src = attr.source
      patchCanvasForIE = (canvas) ->
         if typeof window.G_vmlCanvasManager != "undefined"
          canvas = window.G_vmlCanvasManager.initElement(canvas)
          oldGetContext = canvas.getContext
          canvas.getContext = (a) ->
            ctx = oldGetContext.apply(canvas, arguments)
            canvas.getContext = oldGetContext
            ctx

      createCanvas= (container, width, height) ->
         canvas = document.createElement("canvas")
         container[0].appendChild(canvas)
         canvas.width  = width || 370
         canvas.height = height || 80
         canvas

      linearInterpolate= (before, after, atPoint) ->
          before + (after - before) * atPoint

      interpolateArray = (data, fitCount) ->
        newData = new Array()
        springFactor = new Number((data.length - 1) / (fitCount - 1))
        for i,x in data
          data[x] = Number(i)
        newData[0] = Number(data[0])
        i = 1

        while i < fitCount - 1
          tmp = i * springFactor
          before = new Number(Math.floor(tmp)).toFixed()
          after = new Number(Math.ceil(tmp)).toFixed()
          atPoint = tmp - before
          newData[i] = linearInterpolate(data[before], data[after], atPoint)
          i++
        newData[fitCount - 1] = data[data.length - 1]
        return newData

      redraw= () =>
          clear()
          if typeof(innerColor) == "function"
            context.fillStyle = innerColor()
          else
            context.fillStyle = innerColor
          middle = height / 2
          i = 0
          for d in data
            t = width / data.length
            context.fillStyle = innerColor(i/width, d) if typeof(innerColor) == "function"
            context.clearRect t*i, middle - middle * d, t, (middle * d * 2)
            context.fillRect t*i, middle - middle * d, t, middle * d * 2
            i++

      clear= ->
        context.fillStyle = outerColor
        context.clearRect(0, 0, width, height)
        context.fillRect(0, 0,  width, height)

      # tell audio element to play/pause, you can also use $scope.audio.play() or $scope.audio.pause();
      scope.playpause = ->
        a = (if scope.audio.paused then scope.audio.play() else scope.audio.pause())
        return
      canvas = createCanvas(element, element.parent().clientWidth, element.clientHeight)
      patchCanvasForIE(canvas)
      context = canvas.getContext("2d")
      width  = parseInt context.canvas.width, 10
      height = parseInt context.canvas.height, 10
      outerColor = "transparent"
      innerColor = (x, y) ->
        if x < scope.audio.currentTime / scope.audio.duration
          return "rgba(255,  102, 0, 0.8)"
        else
          return "rgba(0, 0, 0, 0.4)"
      data = interpolateArray(attr.waveform.split(","), width)
      redraw()

      getPosition = (event) ->
        oncanvas = ((100/width) *event.offsetX)
        scope.audio.currentTime = (scope.audio.duration/100) * oncanvas

      canvas.addEventListener "mousedown", getPosition, false
      scope.audio.addEventListener "timeupdate", redraw, false
      scope.audio.addEventListener "ended", ()->
        scope.$apply ->
          scope.audio.paused = true
          scope.audio.currentTime = 0
          redraw()
      , false

      scope.$on '$destroy', ()->
        
        scope.audio.pause()
        scope.audio.src=''
        scope.audio.removeAttribute("src")
        scope.audio = null
 )