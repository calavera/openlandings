{{ template "layouts/app.tpl" . }}

{{ define "content" }}
    <div class="row">
      <div class="twelve wide column finish ui raised contaner segment">
      <h2 class="ui center aligned icon header">
      	<span class=" circular mega-octicon octicon-check"></span>
      	<p>Congratulations, you're done!</p>
      </h2>
      	<h3 class="ui cneter aligned icon header">Go see your site at <a target="_blank" href="{{ .domainURL }}">{{ .domainURL }}</a></h3>
	</div>
    </div>
{{ end }}
