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
