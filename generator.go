package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

func generate(arg *argT) error {
	tpl, err := readAssets()
	if err != nil {
		return err
	}

	generator := &Generator{
		OutputDir:   filepath.Join(arg.Out, arg.Name),
		Name:        arg.Name,
		GoVersion:   arg.GoVersion,
		Prefix:      arg.Prefix,
		GoMod:       arg.GoMod,
		NodeVersion: arg.NodeVersion,
		Template:    tpl,
	}

	return generator.Generate()
}

func readAssets() (*template.Template, error) {
	tpl := template.New("")
	for path, f := range Assets.Files {
		if f.Data == nil {
			continue
		}
		t := tpl.New(path)
		t, err := t.Parse(string(f.Data))
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse template")
		}
	}
	return tpl, nil
}

type Generator struct {
	OutputDir   string
	Name        string
	GoVersion   string
	Prefix      string
	GoMod       bool
	NodeVersion string
	Template    *template.Template
}

func (g *Generator) Generate() error {
	if err := g.generateOutputDir(); err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to get current directory")
	}
	if filepath.IsAbs(g.OutputDir) {
		err = os.Chdir(g.OutputDir)
	} else {
		err = os.Chdir(filepath.Join(currentDir, g.OutputDir))
	}
	if err != nil {
		return errors.Wrap(err, "failed to change directory")
	}
	defer func() {
		_ = os.Chdir(currentDir)
	}()

	if err := g.generateTemplates(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) mkdir(dir string) error {
	if err := os.MkdirAll(dir, 0777); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to make directory `%s`", dir))
	}
	return nil
}

func (g *Generator) generateOutputDir() error {
	return g.mkdir(g.OutputDir)
}

func (g *Generator) execTemplate(name string, out string, data map[string]string) error {
	t := g.Template.Lookup(name)
	if t == nil {
		return errors.New(fmt.Sprintf("failed to find template `%s`", name))
	}

	f, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to open file `%s`", out))
	}

	err = t.Execute(f, data)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to execute template `%s`", out))
	}
	return nil
}

func (g *Generator) generateTemplates() error {
	p := g.params()

	for d := range Assets.Dirs {
		d = strings.TrimPrefix(d, "/assets")
		d = strings.TrimPrefix(d, "/")
		if d == "" {
			continue
		}
		if err := g.mkdir(d); err != nil {
			return err
		}
	}

	for fname, file := range Assets.Files {
		if file.IsDir() {
			continue
		}
		out := strings.TrimSuffix(fname, ".tmpl")
		out = strings.TrimPrefix(out, "/assets/")
		if err := g.execTemplate(fname, out, p); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) params() map[string]string {
	v := strings.Split(g.GoVersion, ".")
	return map[string]string{
		"name":           g.Name,
		"fullname":       path.Join(g.Prefix, g.Name),
		"goVersion":      g.GoVersion,
		"nodeVersion":    g.NodeVersion,
		"goMinorVersion": strings.Join(v[:2], "."),
	}
}
