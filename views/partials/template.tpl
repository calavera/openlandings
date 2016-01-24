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
      <div class="ui right floated primary button">
        Select <span class="octicon octicon-milestone"></span>
      </div>
    </div>
  </div>
</div>
