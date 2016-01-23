{{ template "layouts/app.tpl" . }}

{{ define "content" }}
<div class="row"> <!-- owners row -->
  <div class="eleven wide column">
    <div class="ui divided items browse">

      <div class="item"> <!-- user's own repositories -->
        <div class="image">
          <img src="{{ .currentUser.AvatarURL }}">
        </div>
        <div class="content">
          <a href="{{ .githubUser.HTMLURL }}" class="header">{{ .currentUser.Name }}</a>
          <div class="meta">
            <span class="login">{{ .currentUser.NickName }}</span>
          </div>
          <div class="extra">
            <div class="ui right floated primary button">
              See repositories <span class="octicon octicon-repo-pull"></span>
            </div>
          </div>
        </div>
      </div>

      <!-- organizations -->
      {{ range $idx, $org := .organizations }}
      <div class="item">
        <div class="image">
          <img src="{{ $org.AvatarURL }}">
        </div>
        <div class="content">
          <a href="{{ $org.HTMLURL }}" class="header">{{ $org.Name }}</a>
          <div class="meta">
            <span class="login">{{ $org.Login }}</span>
          </div>
          <div class="description">
	    {{ if $org.Blog }}
            <p><a href="{{ $org.Blog }}">{{ $org.Blog }}</a></p>
	    {{ else }}
            <p><a href="{{ $org.HTMLURL }}">{{ $org.HTMLURL }}</a></p>
	    {{ end }}
          </div>
          <div class="extra">
            <div class="ui right floated primary button">
              See repositories <span class="octicon octicon-repo-pull"></span>
            </div>
          </div>
        </div>
      </div>
      {{ end }}

    </div>
  </div>
</div> <!-- end owners row -->
{{ end }}
