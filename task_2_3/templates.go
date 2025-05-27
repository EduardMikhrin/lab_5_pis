package main

import "html/template"

var tpl = template.Must(template.New("weather").Parse(`<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <title>Погода в Києві</title>
    <style>
      body { font-family: Arial, sans-serif; margin: 2rem; }
      .card { border: 1px solid #ccc; padding: 1rem; border-radius: 8px; max-width: 320px; }
      h1 { margin-bottom: 0.5rem; }
    </style>
</head>
<body>
    <h1>Поточна погода ({{.Date}})</h1>
    <div class="card">
        <p>Температура: <strong>{{.TempC}}°C</strong></p>
        <p>Відчувається як: <strong>{{.FeelsLikeC}}°C</strong></p>
        <p>Вітер: <strong>{{.WindspeedKmph}} км/год</strong></p>
        <p>Вологість: <strong>{{.Humidity}}%</strong></p>
    </div>
</body>
</html>`))
