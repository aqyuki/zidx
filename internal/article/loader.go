package article

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/adrg/frontmatter"
)

var ErrArticleNotFound = errors.New("no articles found")

type Frontmatter struct {
	Title       string `yaml:"title"`
	Emoji       string `yaml:"emoji"`
	IsPublished bool   `yaml:"published"`
}

func Load(dir string, username string) ([]Meta, error) {
	articleDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path of article directory: %w", err)
	}

	filenames := collectArticles(articleDir)
	if len(filenames) == 0 {
		return nil, ErrArticleNotFound
	}

	metas := make([]Meta, 0, len(filenames))
	errs := make([]error, 0, len(filenames))
	for _, filename := range filenames {
		path := filepath.Join(articleDir, filename)
		f, err := os.Open(path)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to open file: %w", err))
			continue
		}
		defer f.Close()

		fm, err := loadFrontmatter(f)
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to load frontmatter: %w", err))
			continue
		}

		metas = append(metas, Meta{
			Slug:        filepath.Base(filename[:len(filename)-len(filepath.Ext(filename))]),
			Username:    username,
			Title:       fm.Title,
			Emoji:       fm.Emoji,
			IsPublished: fm.IsPublished,
		})
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("failed to load article metadata: %v", errors.Join(errs...))
	}
	return metas, nil
}

func collectArticles(target string) []string {
	entries, err := os.ReadDir(target)
	if err != nil {
		return nil
	}

	articles := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".md" {
			continue
		}
		articles = append(articles, entry.Name())
	}
	return articles
}

func loadFrontmatter(r io.Reader) (Frontmatter, error) {
	var fm Frontmatter
	if _, err := frontmatter.Parse(r, &fm); err != nil {
		return Frontmatter{}, fmt.Errorf("failed to parse frontmatter: %w", err)
	}
	return fm, nil
}
