{{ block "contacts" . }}
    {{ range .Data }}
        {{ template "contact" .}}
    {{ end }}
    <div hx-swap="outerHTML" hx-get="/contacts?start={{ .Next }}&take=10" hx-trigger="revealed">
        .
    </div>
{{end}}