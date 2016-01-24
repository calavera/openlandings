<div class="item"> <!-- user's own repositories -->
  <div class="image screenshot">
    <img src="/static/themes/{{ .BasePath }}/screenshot.png">
  </div>
  <div class="content">
    <a href="{{ .Home }}" class="header">{{ .Name }}</a>
    <div class="description">
      <p>{{ .Description }}</p>
    </div>
    <div class="extra">
      <div class="ui right floated button">
        Preview <span class="octicon octicon-eye"></span>
      </div>
      <a class="post configure-site" data-template="{{ .BasePath }}">
      <div class="ui right floated primary button">
        Select <span class="octicon octicon-milestone"></span>
      </div>
      </a>
    </div>
  </div>
</div>

<script>
$(document).ready(function() {
$('.post.configure-site')
  .on('click', function() {
	configureSite($(this).data("template"))
  })
})
</script>
