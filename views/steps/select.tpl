{{ template "layouts/app.tpl" . }}

{{ define "content" }}
<div class="row"> <!-- owners row -->
  <div class="eleven wide column">
    <div class="ui divided items browse">
      {{ template "partials/card.tpl" .owner }}

      {{ range $idx, $repo := .repositories.List }}
	{{ template "partials/repo.tpl" $repo }}
      {{ end }}

      <div class="ui buttons">
	<a href="/steps/select?owner={{ .owner.Login }}&page={{ .repositories.PrevPage }}"
	   class="{{ if eq .repositories.PrevPage 0 }}disabled{{end}}">
           <button class="ui button {{ if eq .repositories.PrevPage 0 }}disabled{{end}}">Previous</button>
	</a>
        <div class="or" data-text=""></div>
	<a href="/steps/select?owner={{ .owner.Login }}&page={{.repositories.NextPage}}"
	   class="{{ if eq .repositories.NextPage 0 }}disabled{{end}}">
           <button class="ui button {{ if eq .repositories.NextPage 0 }}disabled{{end}}">Next</button>
	</a>
      </div>
    </div>
  </div>
</div> <!-- end owners row -->
{{ end }}
