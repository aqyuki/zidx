package option

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("should return a default option", func(t *testing.T) {
		t.Parallel()

		got := New()
		want := &Option{
			ShowHelp:    false,
			ArticlePath: "./articles",
			Filename:    "toc.md",
			Username:    "",
		}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("differs: (-got +want)\n%s", diff)
		}
	})
}
