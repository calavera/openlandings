{{ template "layouts/app.tpl" . }}

{{ define "content" }}
    <div class="row">
	<div class="twelve wide column segment">
		Congratulations, you're done!
		Go see your site at {{ .domainURL }}
	</div>
    </div>
{{ end }}
