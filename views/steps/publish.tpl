{{ template "layouts/app.tpl" . }}

{{ define "content" }}
    <div class="row">
	<div class="twelve wide column ui large header">
    	<span class="mega-octicon octicon-browser"></span>
	{{ .repository.Owner.Login}} / {{ .repository.Name }}
	</div>
    </div>

    <div class="row md ui tab content active">
	<div class="twelve wide column">
	  <form class="ui form" method="post" action="/steps/publish">
	    <input type="hidden" name="landing" value="{{ .preview.Landing }}">
	    <input type="hidden" name="template" value="{{ .preview.Template }}">
	    <input type="hidden" name="nwo" value="{{ .preview.Nwo }}">

	    <h4 class="ui dividing header">Domain information</h4>
            <div class="field">
              <label>Default name</label>
              <input type="text" name="default" placeholder="{{ .repository.Owner.Login }}-{{ .repository.Name }}.netlify.com" readonly>
            </div>
            <div class="field">
              <label>Custom name</label>
              <input type="text" name="custom" placeholder="awesome-project.com">
            </div>
            <button class="ui secondary button" type="submit">Publish <span class="octicon octicon-cloud-upload"></span></button>
	  </form>
	</div>
    </div>
{{ end }}
