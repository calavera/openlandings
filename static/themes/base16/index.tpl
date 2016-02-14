{{ $colors := "fedcba98765432" }}
<!DOCTYPE html>
<html lang="en-US">
<head>
<meta charset="utf-8">
<meta name="description" content="{{ .site.Description }}">
<meta name="author" content="{{ .site.Owner.Name }}">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="{{ .site.BaseURL }}/css/style.min.css" type="text/css">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Code+Pro:400,700" type="text/css">
<title>{{ .site.Title }}</title>
</head>
<body>

<header>
  <div class="container">
    <a class="path" href="{{ .site.Repo.URL }}">[{{ .site.Title }}]</a>
    <span class="caret"># _</span>
  </div>
</header>


<div class="container">

<main role="main" class="homepage">
  <h1 class="site-title">
    <span class="base05">[</span>{{ range $idx, $char := runes .site.Repo.Login }}{{ $i := mod $idx 13 }}{{ $b := charAt $colors $i }}<span class="base0{{ $b }}">{{ printf "%c" $char }}</span>{{ end }}<span class="base05">]</span>
    <span class="base05"># _</span>
  </h1>
  
  <div class="hero-logo">
    <img src="/images/base16-eighties.svg">
  </div>

  <div class="article">
	<article class="single" itemscope itemtype="http://schema.org/BlogPosting">
	  <section class="body" itemprop="articleBody">
	    {{ .site.Content }}
	  </section>
	</article>
  </div>
</main>

</div>

<footer>
  <div class="container">
    <span class="copyright">&copy; 2016 {{ .site.Title }} - <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">CC BY 4.0</a></span>
  </div>
</footer>

</body>
</html>
