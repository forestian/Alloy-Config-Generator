package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"alloy-config-generator/internal/config"
	embeddedtemplates "alloy-config-generator/internal/templates"
)

type templateData struct {
	config.PipelineConfig
	ConfigAlloy string
}

func renderTemplate(name string, data templateData) (string, error) {
	tmpl, err := template.New(name).Funcs(template.FuncMap{
		"indent": indent,
	}).Option("missingkey=error").ParseFS(embeddedtemplates.FS, name)
	if err != nil {
		return "", fmt.Errorf("parse template %s: %w", name, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("render template %s: %w", name, err)
	}

	return buf.String(), nil
}

func indent(spaces int, text string) string {
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(strings.TrimRight(text, "\n"), "\n")
	for i, line := range lines {
		if line == "" {
			lines[i] = pad
			continue
		}
		lines[i] = pad + line
	}
	return strings.Join(lines, "\n")
}
