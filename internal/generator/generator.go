package generator

import (
	"os"

	"alloy-config-generator/internal/config"
	"alloy-config-generator/internal/validate"
)

func Generate(cfg config.PipelineConfig) error {
	files, err := RenderFiles(&cfg)
	if err != nil {
		return err
	}
	return writeFiles(cfg.OutputDir, cfg.Force, files)
}

func RenderFiles(cfg *config.PipelineConfig) ([]GeneratedFile, error) {
	if err := validate.NormalizeAndValidate(cfg); err != nil {
		return nil, err
	}

	data := templateData{PipelineConfig: *cfg}
	configAlloy, err := renderTemplate("config.alloy.tmpl", data)
	if err != nil {
		return nil, err
	}
	data.ConfigAlloy = configAlloy

	files := []GeneratedFile{}

	readme, err := renderTemplate("README.md.tmpl", data)
	if err != nil {
		return nil, err
	}
	files = append(files, GeneratedFile{Path: "README.md", Content: readme, Mode: 0o644})

	if cfg.IncludesConfig() {
		files = append(files, GeneratedFile{Path: "config/config.alloy", Content: configAlloy, Mode: 0o644})
	}

	if cfg.IncludesHelm() {
		values, err := renderTemplate("values.yaml.tmpl", data)
		if err != nil {
			return nil, err
		}
		files = append(files, GeneratedFile{Path: "helm/values.yaml", Content: values, Mode: 0o644})
	}

	if cfg.IncludesExamples() {
		install, err := renderTemplate("install.sh.tmpl", data)
		if err != nil {
			return nil, err
		}
		uninstall, err := renderTemplate("uninstall.sh.tmpl", data)
		if err != nil {
			return nil, err
		}
		files = append(files, GeneratedFile{Path: "examples/install.sh", Content: install, Mode: os.FileMode(0o755)})
		files = append(files, GeneratedFile{Path: "examples/uninstall.sh", Content: uninstall, Mode: os.FileMode(0o755)})
	}

	return files, nil
}
