{{ block "index" . }}
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>title</title>
    <script
      src="https://unpkg.com/htmx.org@1.9.12"
      integrity="sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2"
      crossorigin="anonymous"
    ></script>
  </head>
  <body>
    <header>   
      Home page
    </header>
    <section class="contacts">
      {{ template "contacts" . }}
    </section> 
  </body>
</html>
{{ end }}
