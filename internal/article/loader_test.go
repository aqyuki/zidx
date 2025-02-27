package article

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_collectArticles(t *testing.T) {
	t.Parallel()

	t.Run("article found", func(t *testing.T) {
		got := collectArticles("testdata/exist")
		want := []string{"article1.md", "article2.md"}

		if cmp.Equal(got, want) == false {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})

	t.Run("article not found", func(t *testing.T) {
		got := collectArticles("testdata/empty")
		want := []string{}

		if cmp.Equal(got, want) == false {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})
}

const parsableFrontmatter = `---
title: "title"
emoji: "✨️"
published: true
---`

// This frontmatter is invalid because the value of the `published` field is a string.
const unparsableFrontmatter = `---
title: "title"
emoji: "✨️"
published: "true"
---
`

func Test_loadFrontmatter(t *testing.T) {
	t.Parallel()

	t.Run("success to parse frontmatter", func(t *testing.T) {
		t.Parallel()

		got, err := loadFrontmatter(strings.NewReader(parsableFrontmatter))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		want := Frontmatter{
			Title:       "title",
			Emoji:       "✨️",
			IsPublished: true,
		}

		if cmp.Equal(got, want) == false {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})

	t.Run("fail to parse frontmatter", func(t *testing.T) {
		t.Parallel()

		_, err := loadFrontmatter(strings.NewReader(unparsableFrontmatter))
		if err == nil {
			t.Error("expected an error but got none")
		}
	})
}

func Test_Load(t *testing.T) {
	t.Parallel()

	t.Run("success to load article metadata", func(t *testing.T) {
		t.Parallel()

		got, err := Load("testdata/exist", "username")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		want := []Meta{
			{
				Slug:        "article1",
				Username:    "username",
				Title:       "article1",
				Emoji:       "✨️",
				IsPublished: true,
			},
			{
				Slug:        "article2",
				Username:    "username",
				Title:       "article2",
				Emoji:       "📖",
				IsPublished: false,
			},
		}

		if cmp.Equal(got, want) == false {
			t.Errorf("got = %v, want = %v", got, want)
		}
	})

	t.Run("article not found", func(t *testing.T) {
		t.Parallel()

		_, err := Load("testdata/empty", "username")
		if err != ErrArticleNotFound {
			t.Errorf("got = %v, want = %v", err, ErrArticleNotFound)
		}
	})
}
