package export

import (
	"io"
	"testing"

	"github.com/aqyuki/zidx/internal/article"
	"github.com/google/go-cmp/cmp"
)

func TestConvertAndTrim(t *testing.T) {
	t.Parallel()

	metas := []article.Meta{
		{
			Username:    "username",
			Slug:        "slug",
			Title:       "title",
			Emoji:       "✨️",
			IsPublished: true,
		},
		{
			Username:    "username",
			Slug:        "slug",
			Title:       "title",
			Emoji:       "📖",
			IsPublished: false,
		},
	}

	got := ConvertAndTrim(metas)
	want := []Content{
		{
			URL:   "https://zenn.dev/username/articles/slug",
			Title: "title",
			Emoji: "✨️",
		},
	}

	if cmp.Equal(got, want) == false {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	metas := []Content{
		{
			URL:   "https://zenn.dev/username/articles/slug",
			Title: "title",
			Emoji: "✨️",
		},
		{
			URL:   "https://zenn.dev/username/articles/slug",
			Title: "title",
			Emoji: "📖",
		},
	}

	got, err := Generate(metas)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	b, err := io.ReadAll(got)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `# Table of Contents

- [✨️ title](https://zenn.dev/username/articles/slug)
- [📖 title](https://zenn.dev/username/articles/slug)
`
	if diff := cmp.Diff(string(b), want); diff != "" {
		// t.Errorf("diff (-got +want)\n%s", diff)
		t.Errorf("got\n%s\nwant\n%s\n", string(b), want)
	}
}
