<div class="item">
  <div class="content">
    <a href="{{ .Repository.HTMLURL }}" class="header">{{ .Repository.FullName }}</a>
    <div class="description">
      {{ if .Repository.Description }}
      <p>{{ .Repository.Description }}</p>
      {{ end }}
    </div>
    <div class="extra">
      {{ if .Site }}
	  <form method="post" action="/sites/{{ .Site.ID }}">
	    <input type="hidden" name="_method" value="DELETE" />
            <button class="ui right floated red basic button delete" type="submit">Delete site <span class="octicon octicon-trashcan"></span></button>
	  </form>
      {{ else }}
        <a href="/steps/configure?nwo={{ .Repository.FullName }}" class="ui right floated secondary button">
            Configure <span class="octicon octicon-repo-pull"></span>
        </a>
      {{ end }}
    </div>
  </div>
</div>
