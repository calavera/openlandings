## Community templates

This folder includes all the templates contributed by the community.

Each template has the same structure:

1. An `info.json` file with information about the template.
2. An `index.tmpl` file as the landing page template.
3. A `static` directory with other static files. These static files must be structured into `css`, `js`, and `fonts` directories.

### How to contribute a new template

1. Create an `info.json` file with this information:

```json
{
	"name": "name of the template",
	"description": "description of the template",
	"home": "URL to the home page of this template"
}
```

2. Add the static files necessaries for the template.

### What are the variables that I can use in my index.tmpl?

We use Go's template engine to parse the `index.tmpl` provided. This is all the information the template has access to:


`{{ .site.Owner.Name }}`
`{{ .site.Owner.URL }}`
`{{ .site.Owner.Avatar }}`
`{{ .site.Repo.Name }}`
`{{ .site.Repo.Description }}`
`{{ .site.Repo.URL }}`
`{{ .site.Content }}`
`{{ .site.Analytics }}`
`{{ .site.BaseURL }}`
`{{ .site.Title }}`
`{{ .site.Description }}`
