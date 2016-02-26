{{ template "layouts/app.tpl" . }}

{{ define "content" }}
    <div class="row">
	<div class="twelve wide column ui large header">
    	<span class="mega-octicon octicon-browser"></span>
	{{ .repository.Owner.Login}} / {{ .repository.Name }}
	</div>
    </div>

   <div class="row tabs">
     <div class="twelve wide column">
       <div class="ui tabular menu">
         <a class="item content active">
          <span class="octicon octicon-file-text"></span>
           Content
         </a>
         <a class="item template">
           <span class="octicon octicon-file-media"></span>
           Template
         </a>
       </div>
     </div>
   </div>

    <div class="row md ui tab content active">
	<div class="twelve wide column">
	  <form class="ui form" method="post" action="/steps/configure-site">
	    <input type="hidden" name="template" id="template" value="">
	    <input type="hidden" name="landing" id="landing" value="">
	    <input type="hidden" name="nwo" value="{{ .repository.FullName }}">

		<div class="ui fluid">
		  <div class="ui action input">
		    <input type="text" name="file" placeholder="README.md">
		    <div type="submit" class="ui button">
    			<span class="octicon octicon-search"></span>
		    </div>
		  </div>
		  <div class="ui right floated">
			<span class="mega-octicon octicon-markdown"></span>
	          </div>
		</div>

      	  <div class="ui fluid content">
      		<textarea rows="100" id="markdown-content" name="content">{{ .repository.Content }}</textarea>
      	  </div>
	  </form>
	</div>
     </div>

    <div class="ui tab row template" data-tab="tempate">
     <div class="twelve wide column">
        <div class="ui divided items browse">
          {{ range $idx, $tmpl := .templates }}
    	    {{ template "partials/template.tpl" $tmpl }}
          {{ end }}
        </div>
     </div>
    </div>
{{ end }}

<script>
$(document).ready(function() {

$('.ui.menu .item.template')
  .on('click', function() {
	$(".ui.tab.content").removeClass("active")
	$(".ui.tab.template").addClass("active")
	$(this).addClass("active")
	$(".ui.menu .item.content").removeClass("active")
  })
;

$('.ui.menu .item.content')
  .on('click', function() {
	$(".ui.tab.template").removeClass("active")
	$(".ui.tab.content").addClass("active")
	$(this).addClass("active")
	$(".ui.menu .item.template").removeClass("active")
  })
;

$('.post.configure-site')
  .on('click', function() {
	configureSite($(this).data("template"))
  })
;

})
</script>
