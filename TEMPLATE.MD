# Template Inheritance in Go

## Why `{{template "layout" .}}` is needed at the end of dashboard.html

### 1. Template Inheritance Structure
- The `dashboard.html` template defines three blocks: "title", "nav", and "content"
- These blocks are meant to be inserted into the main layout template
- Without the `{{template "layout" .}}` line, the template engine doesn't know how to combine these blocks with the layout

### 2. How Template Inheritance Works
- The layout template (`layout.html`) contains the base HTML structure
- It has placeholders where other templates can insert their content using `{{template "block_name" .}}`
- When we define blocks in `dashboard.html` using `{{define "block_name"}}`, we're creating content that should be inserted into those placeholders
- The `{{template "layout" .}}` line tells the template engine to use the layout template as the base and insert our defined blocks into it

### 3. The Flow
```
dashboard.html:
{{define "title"}}...{{end}}      →  layout.html:
{{define "nav"}}...{{end}}        →  <title>{{template "title" .}}</title>
{{define "content"}}...{{end}}    →  <nav>{{template "nav" .}}</nav>
{{template "layout" .}}           →  <main>{{template "content" .}}</main>
```

### 4. Why It's Necessary
- Without this line, the template engine would only see the block definitions but wouldn't know how to use them
- The blocks would be defined but never actually rendered because there's no connection to the layout template
- The `{{template "layout" .}}` line creates this connection, telling the engine to use the layout as the base and insert our blocks into it

This is a common pattern in Go's template system where child templates (like `dashboard.html`) define blocks that get inserted into a parent template (the layout), creating a complete HTML page. 