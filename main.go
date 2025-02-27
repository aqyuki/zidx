package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aqyuki/zidx/internal/article"
	"github.com/aqyuki/zidx/internal/export"
	"github.com/aqyuki/zidx/internal/option"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	ops := option.New()

	if ops.ShowHelp {
		usage()
		return nil
	}
	return exec(ops)
}

func usage() {
	fmt.Printf(`zidx generates a table of contents for Zenn articles.

Usage: zidx [options]

Options:
  -h, --help       Show help
  --article-dir    Path to the article directory (default: ./articles)
  -f, --filename   Output filename (default: toc.md)
  -u, --username   Zenn username (required)
`)
}

func exec(ops *option.Option) error {
	if ops.Username == "" {
		return fmt.Errorf("username is required")
	}

	metas, err := article.Load(ops.ArticlePath, ops.Username)
	if err != nil {
		return fmt.Errorf("failed to load articles: %w", err)
	}

	contents := export.ConvertAndTrim(metas)
	reader, err := export.Generate(contents)
	if err != nil {
		return fmt.Errorf("failed to generate export: %w", err)
	}

	f, err := os.Create(ops.Filename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, reader); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}
	return nil
}
