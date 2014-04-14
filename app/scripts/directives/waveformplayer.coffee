'use strict'

angular.module('podcasterApp')
  .directive('waveformplayer', ->
    restrict: "E"
    scope: { item: '=' }
    templateUrl: "/views/player.html"
    link: (scope, element, attr) ->
      scope.audio = new Audio()
      scope.preload = "metadata"
      scope.now = 0
      scope.duration = 0
      scope.cursor = 0
      scope.cursor_over = false

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

      toHHMMSS = (times)->
        sec_num = parseInt(times, 10) # don't forget the second param
        hours = Math.floor(sec_num / 3600)
        minutes = Math.floor((sec_num - (hours * 3600)) / 60)
        seconds = sec_num - (hours * 3600) - (minutes * 60)
        hours = "0" + hours  if hours < 10
        minutes = "0" + minutes  if minutes < 10
        seconds = "0" + seconds  if seconds < 10
        time = hours + ":" + minutes + ":" + seconds
        time

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
        if scope.audio.paused
          if scope.audio.src == "/api/media/head/#{scope.item.id}"
            scope.audio.src = "/api/media/play/#{scope.item.id}"
          scope.audio.play()
        else
          scope.audio.pause()

      scope.audio.addEventListener "ended", ()->
        scope.$apply ->
          scope.audio.paused = true
          scope.audio.currentTime = 0
          redraw()
      , false

      canvas = createCanvas(element, element.parent().clientWidth, element.clientHeight)
      patchCanvasForIE(canvas)
      context = canvas.getContext("2d")
      width  = parseInt context.canvas.width, 10
      height = parseInt context.canvas.height, 10
      outerColor = "transparent"
      gradient = context.createLinearGradient(0, 0, 0, height);
      gradient.addColorStop(0.0, "#f60");
      gradient.addColorStop(1.0, "#ff1b00");
      innerColor = (x, y) ->
        try
          buffered = scope.audio.buffered.end(0)
        catch e
          buffered = 0
        if (x < scope.cursor / width) && scope.cursor_over
          return "#f60"
        else if x < scope.audio.currentTime / scope.audio.duration
          return gradient
        else if x < buffered / scope.audio.duration
          return "rgba(0, 0, 0, 0.7)"
        else  
          return "rgba(0, 0, 0, 0.5)"
      data = []
      canvas.addEventListener "mousedown", (event) ->
        oncanvas = ((100/width) *event.offsetX)
        if scope.audio.paused
          scope.playpause()
        scope.audio.currentTime = (scope.audio.duration/100) * oncanvas
      , false

      canvas.addEventListener "mouseover", (event) ->
        scope.cursor_over = true
        scope.cursor = event.offsetX 
        redraw()
      , false
      canvas.addEventListener "mouseout", (event) ->
        scope.cursor_over = false
        scope.cursor = event.offsetX 
        redraw()
      , false
      canvas.addEventListener "mousemove", (event) ->
        scope.cursor = event.offsetX 
        redraw()
      , false

      scope.audio.addEventListener "timeupdate", ()->
        redraw() 
        scope.$apply ->
          scope.duration = toHHMMSS(scope.audio.duration)
          scope.now = toHHMMSS(scope.audio.currentTime)
      , false

      scope.$watch 'item', (item)->
        data = interpolateArray(item.wave.split(","), width)
        scope.item = item
        scope.audio.src = "/api/media/head/#{item.id}"
        redraw()

      scope.$on '$destroy', ()->
        
        scope.audio.pause()
        scope.audio.src=''
        scope.audio.removeAttribute("src")
        scope.audio = null
 )