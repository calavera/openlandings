<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width initial-scale=1" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge">

  <title>{{ .site.Title }}</title>
  <meta name="description" content="{{ .site.Description }}">
  <meta name="author" content="{{ .site.Owner.Name }}">
  <meta name="HandheldFriendly" content="True">
  <meta name="MobileOptimized" content="320">
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="black-translucent">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

  <meta name="twitter:card" content="summary">
  <meta name="twitter:title" content="{{ .site.Title }}">
  <meta name="twitter:description" content="{{ .site.Description }}">

  <meta property="og:type" content="article">
  <meta property="og:title" content="{{ .site.Title }}">
  <meta property="og:description" content="{{ .site.Description }}">

  <meta name="msapplication-TileColor" content="#ffc40d">
  <meta name="theme-color" content="#ffffff">

  <link rel="stylesheet" href="{{ .site.BaseURL }}/css/main.css">
</head>

  <body>
<span class="mobile btn-mobile-menu">
  <i class="icon icon-list btn-mobile-menu__icon"></i>
  <i class="icon icon-x-circle btn-mobile-close__icon hidden"></i>
</span>
  
<header class="panel-cover" style="background-image: url({{ .site.BaseURL }}/images/cover.jpg)">
  <div class="panel-main">

    <div class="panel-main__inner panel-inverted">
    <div class="panel-main__content">
        <a href="{{ .site.BaseURL }}" title="link to home of {{ .site.Title }}">
          <img src="{{ .site.BaseURL }}/images/profile.jpg" class="user-image" alt="My Profile Photo">
          <h1 class="panel-cover__title panel-title">{{ .site.Title }}</h1>
        </a>
        <hr class="panel-cover__divider">
        <p class="panel-cover__description">{{ .site.Description }}</p>
        <hr class="panel-cover__divider panel-cover__divider--secondary">

        <div class="navigation-wrapper">

          <nav class="cover-navigation cover-navigation--primary">
            <ul class="navigation">
              <li class="navigation__item"><a href="{{ .site.BaseURL }}#readme" title="link to {{ .site.Title }} blog" class="blog-button">Readme</a></li>
            </ul>
          </nav>

          <nav class="cover-navigation navigation--social">
            <ul class="navigation">
          
            {{ if .site.Repo.URL }}
              <!-- GitHub -->
              <li class="navigation__item">
		      <a href="{{ .site.Repo.URL }}" target="_blank">
                  <i class="icon icon-social-github"></i>
                  <span class="label">GitHub</span>
                </a>
              </li>
            {{ end }}
          
            </ul>
          </nav>

        </div>

      </div>

    </div>

    <div class="panel-cover--overlay"></div>
  </div>
</header>

    <div class="content-wrapper">
      <div class="content-wrapper__inner">
        {{ .site.Content }}
      </div>

<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
<script type="text/javascript" src="{{ .site.BaseURL }}/js/main.js"></script>
{{ if .site.Analytics }}
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', '{{ site.Analytics }}', 'auto');
  ga('send', 'pageview');
</script>
{{ end }}
    </div>
  </body>
</html>
