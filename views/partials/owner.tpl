<div class="item"> <!-- user's own repositories -->
  <div class="image">
    <img src="{{ .AvatarURL }}">
  </div>
  <div class="content">
    <a href="{{ .HTMLURL }}" class="header">{{ .Name }}</a>
    <div class="meta">
      <span class="login">{{ .Login }}</span>
    </div>
    <div class="description">
      {{ if .Blog }}
      <p><a href="{{ .Blog }}">{{ .Blog }}</a></p>
      {{ else }}
      <p><a href="{{ .HTMLURL }}">{{ .HTMLURL }}</a></p>
      {{ end }}
    </div>
    <div class="extra">
      <a href="/steps/select?owner={{ .Login }}">
      <div class="ui right floated primary button">
        See repositories <span class="octicon octicon-repo-pull"></span>
      </div>
      </a>
    </div>
  </div>
</div>
