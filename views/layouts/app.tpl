<!DOCTYPE html>
<html>
<head>
  <!-- Standard Meta -->
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">

  <!-- Site Properties -->
  <title>Open landings</title>
  <link rel="stylesheet" type="text/css" href="/static/css/components/reset.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/site.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/container.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/grid.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/header.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/image.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/menu.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/divider.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/dropdown.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/segment.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/button.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/list.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/icon.css">

  <link rel="stylesheet" type="text/css" href="/static/css/components/step.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/components/item.min.css">

  <link rel="stylesheet" href="/static/css/octicons.css">
  <link rel="stylesheet" href="/static/css/home.css">
  <link rel="stylesheet" href="/static/css/main.css">
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
        </div>
      </div>
    </div>
  </div>
  
  <div class="ui centered eleve wide column grid">
    {{ template "partials/steps.tpl" . }}

    {{ template "content" . }}

  </div> <!-- end grid -->

  <div class="ui inverted vertical footer segment">
    <div class="ui container">
      <div class="ui stackable inverted divided equal height stackable grid">
        <div class="three wide column">
          <div class="ui inverted link list">
            <a href="h" class="item"><span class="octicon octicon-mark-github"></span> Project Repository</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</body>
</html>
