{{ template "layouts/app.tpl" . }}

{{ define "content" }}
<div class="row"> <!-- owners row -->
  <div class="eleven wide column">
    <div class="ui divided items browse">
      {{ template "partials/card.tpl" .owner }}

      {{ range $idx, $site := .sites.All }}
	{{ template "partials/repo.tpl" $site }}
      {{ end }}

      <div class="ui item buttons">
	<a href="/steps/select?owner={{ .owner.Login }}&page={{ .sites.PrevPage }}"
	   class="{{ if eq .sites.PrevPage 0 }}disabled{{end}}">
           <button class="ui button {{ if eq .sites.PrevPage 0 }}disabled{{end}}">Previous</button>
	</a>
        <div class="or" data-text=""></div>
	<a href="/steps/select?owner={{ .owner.Login }}&page={{.sites.NextPage}}"
	   class="{{ if eq .sites.NextPage 0 }}disabled{{end}}">
           <button class="ui button {{ if eq .sites.NextPage 0 }}disabled{{end}}">Next</button>
	</a>
      </div>
    </div>
  </div>
</div> <!-- end owners row -->
{{ end }}
