package option

import (
	flag "github.com/spf13/pflag"
)

// Option is a struct that holds application options.
// Option is built from command line arguments.
type Option struct {
	// ShowHelp is a flag
	ShowHelp bool

	// ArticlePath is the path to the article directory.
	// default : `./articles`
	ArticlePath string

	// Filename sets which file the generated table of contents is output to.
	// default : `Index.md`
	Filename string

	// Username specifies the Zenn username.
	// If not specified, zidx will generate a table of contents with no links.
	Username string
}

var (
	showHelp    bool
	articlePath string
	filename    string
	username    string
)

func New() *Option {
	return &Option{
		ShowHelp:    showHelp,
		ArticlePath: articlePath,
		Filename:    filename,
		Username:    username,
	}
}

func init() {
	flag.BoolVarP(&showHelp, "help", "h", false, "show help")
	flag.StringVar(&articlePath, "article-dir", "./articles", "path to the article directory")
	flag.StringVarP(&filename, "filename", "f", "Index.md", "output filename")
	flag.StringVarP(&username, "username", "u", "", "Zenn username")
	flag.Parse()
}
