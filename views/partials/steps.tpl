<div class="row"> <!-- steps row -->
<div class="ui tablet stackable steps">
  <a href="/steps/browse" class="step {{ .steps.Browse.Status }}">
    <span class="mega-octicon octicon-search"></span>
    <div class="content">
      <div class="title">Browse</div>
      <div class="description">Find your repositories</div>
    </div>
  </a>

  <a  href="/steps/select" class="step {{ .steps.Select.Status }}">
    <span class="mega-octicon octicon-repo"></span>
    <div class="content">
      <div class="title">Select</div>
      <div class="description">Choose your repository</div>
    </div>
  </a>

  <a class="step {{ .steps.Configure.Status }}">
    <span class="mega-octicon octicon-gear"></span>
    <div class="content">
      <div class="title">Configure</div>
      <div class="description">Choose template and text</div>
    </div>
  </a>

  <a class="step {{ .steps.Publish.Status }}">
    <span class="mega-octicon octicon-squirrel"></span>
    <div class="content">
      <div class="title">Publish</div>
      <div class="description">Publish your new landing page</div>
    </div>
  </a>
</div>
</div> <!-- end steps row -->
