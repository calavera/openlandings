<!DOCTYPE html>
<html>
<head>
  <!-- Standard Meta -->
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">

  <!-- Site Properties -->
  <title>Open landings</title>
  <link rel="stylesheet" type="text/css" href="/static/css/components/reset.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/site.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/container.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/grid.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/header.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/image.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/menu.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/divider.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/dropdown.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/segment.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/button.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/list.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/icon.min.css">

  <link rel="stylesheet" type="text/css" href="/static/css/components/step.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/item.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/form.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/input.min.css">

  <link rel="stylesheet" href="/static/css/tab.css">
  <link rel="stylesheet" href="/static/css/octicons.css">
  <link rel="stylesheet" href="/static/css/home.css">
  <link rel="stylesheet" href="/static/css/main.css">

  <script src="/static/js/jquery-2.2.0.min.js"></script>
</head>
<body>

<!-- Page Contents -->
<div class="pusher">
  <div class="ui inverted vertical center aligned segment">
    <div class="ui container">
      <div class="ui large secondary inverted menu">
	<a class="item"><span class="mega-octicon octicon-home"></span></a>
        <div class="right item">
	  <img class="ui mini circular image" src="{{ .currentUser.AvatarURL }}">
    	  <div class="content avatar">{{ .currentUser.NickName }}</div>
	  <span class="divider">|</span>
	  <div class="content logout"><a href="/logout">log out</a></div>
        </div>
      </div>
    </div>
  </div>
  
  <div class="ui centered eleven wide column grid">
    {{ template "partials/steps.tpl" . }}

    {{ template "content" . }}

  </div> <!-- end grid -->

  <div class="ui inverted vertical footer segment">
    <div class="ui container">
      <div class="ui stackable inverted divided equal height stackable grid">
        <div class="three wide column">
          <div class="ui inverted link list">
            <a href="https://github.com/calavera/openlandings" class="item"><span class="octicon octicon-mark-github"></span> Project Repository</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<script src="/static/js/main.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/markdown-it/5.1.0/markdown-it.min.js"></script>
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-73416646-1', 'auto');
  ga('send', 'pageview');
</script>
</body>
</html>
