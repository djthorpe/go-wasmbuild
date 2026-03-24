package storybook

import (
	"fmt"
	"path"
	"strings"

	// Packages
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	yaml "gopkg.in/yaml.v3"
)

type componentDocFrontMatter struct {
	Description string `yaml:"description"`
}

// ComponentDoc returns an embedded Carbon component Markdown document as a view.
func ComponentDoc(filename string) mvc.View {
	frontMatter, body, err := readComponentDoc(filename)
	if err != nil {
		return carbon.Markdown(err.Error())
	}
	children := make([]any, 0, 2)
	if description := strings.TrimSpace(frontMatter.Description); description != "" {
		children = append(children, carbon.Markdown(
			description,
			carbon.WithMarkdownLinkResolver(componentDocLinkResolver(filename)),
		))
	}
	children = append(children, carbon.Markdown(
		body,
		carbon.WithMarkdownLinkResolver(componentDocLinkResolver(filename)),
	))
	return carbon.Page(children...)
}

func componentDocDescription(filename string) (string, error) {
	frontMatter, _, err := readComponentDoc(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(frontMatter.Description), nil
}

func componentDocLinkResolver(baseFilename string) carbon.MarkdownLinkResolver {
	return func(href string) string {
		if href == "" || strings.HasPrefix(href, "#") || strings.Contains(href, "://") || strings.HasPrefix(href, "mailto:") {
			return href
		}
		resolved := path.Clean(path.Join(path.Dir(baseFilename), href))
		name, anchor, _ := strings.Cut(resolved, "#")
		if strings.EqualFold(path.Ext(name), ".md") {
			route := "#" + strings.ToLower(strings.TrimSuffix(path.Base(name), path.Ext(name)))
			if anchor != "" {
				return route + "#" + anchor
			}
			return route
		}
		return resolved
	}
}

func parseComponentDoc(text string) (componentDocFrontMatter, string, error) {
	var frontMatter componentDocFrontMatter
	if !strings.HasPrefix(text, "---\n") {
		return frontMatter, text, nil
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return frontMatter, text, nil
	}
	frontMatterText := text[4 : end+4]
	body := strings.TrimLeft(text[end+9:], "\n")
	if err := yaml.Unmarshal([]byte(frontMatterText), &frontMatter); err != nil {
		return componentDocFrontMatter{}, "", err
	}
	return frontMatter, body, nil
}

func readComponentDoc(filename string) (componentDocFrontMatter, string, error) {
	data, err := carbon.DocsFS.ReadFile(filename)
	if err != nil {
		return componentDocFrontMatter{}, "", fmt.Errorf("Unable to load component documentation for `%s`: %w", filename, err)
	}
	frontMatter, body, err := parseComponentDoc(string(data))
	if err != nil {
		return componentDocFrontMatter{}, "", fmt.Errorf("Unable to parse component documentation for `%s`: %w", filename, err)
	}
	return frontMatter, body, nil
}
