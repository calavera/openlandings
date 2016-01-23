{{ template "layouts/app.tpl" . }}

{{ define "content" }}
<div class="row"> <!-- owners row -->
  <div class="eleven wide column">
    <div class="ui divided items browse">

      {{ template "partials/owner.tpl" .githubUser }}

      <!-- organizations -->
      {{ range $idx, $org := .organizations }}
      	{{ template "partials/owner.tpl" $org }}
      {{ end }}

    </div>
  </div>
</div> <!-- end owners row -->
{{ end }}
