<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en-us">
<head>
  <meta charset="utf-8">
  <title>{{ .site.Title }}</title>

  <meta name="author" content="{{ .site.Owner.Name }}" />
  <meta name="description" content="{{ .site.Description }}" />
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

  <link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet" />
  <link rel="stylesheet" href="{{ .site.BaseURL }}/css/base.css" type="text/css" media="screen, projection" />
  <link rel="stylesheet" href="{{ .site.BaseURL }}/css/pygments.css" type="text/css" />
  <link media="only screen and (max-device-width: 480px)" href="{{ .site.BaseURL }}/css/mobile.css" type="text/css" rel="stylesheet" />
  <link media="only screen and (device-width: 768px)" href="{{ .site.BaseURL }}/css/mobile.css" type="text/css" rel="stylesheet" />
  <link href='http://fonts.googleapis.com/css?family=Yanone+Kaffeesatz' rel='stylesheet' type='text/css'>
  <link rel="apple-touch-icon" href="{{ .site.BaseURL }}/apple-touch-icon.png" />
  <script type="text/javascript" src="{{ .site.BaseURL }}/js/application.js"></script>
</head>
<body>
  <section class="sidebar">
    <a href="/">
      <img src="{{ site.Owner.AvatarURL }}?s=150" height="75" width="75" class="avatar" />
    </a>
  
    <section class="name">
      <a href="/">
        <span id="fname">{{ .site.Owner.Name }}</span>
        <span id="lname">{{ .site.Repo.Login }}</span>
      </a>
    </section>
  
    <section class="meta">
      <a href="{{ .site.Repo.URL }}" target="_blank"><i class="fa fa-github"></i></a>
    </section>
  </section>

  <section class="content">
  {{ .site.Content }}
  </section>
</body>
</html>
