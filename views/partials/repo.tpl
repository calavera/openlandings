<div class="item">
  <div class="content">
    <a href="{{ .HTMLURL }}" class="header">{{ .FullName }}</a>
    <div class="description">
      {{ if .Description }}
      <p>{{ .Description }}</p>
      {{ end }}
    </div>
    <div class="extra">
      <div class="ui right floated secondary button">
        Configure <span class="octicon octicon-repo-pull"></span>
      </div>
    </div>
  </div>
</div>
