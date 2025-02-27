package option

import flag "github.com/spf13/pflag"

// Option is a struct that holds application options.
// Option is built from command line arguments.
type Option struct {
	// ShowHelp is a flag
	ShowHelp bool

	// Filename sets which file the generated table of contents is output to.
	// default : `Index.md`
	Filename string

	// Username specifies the Zenn username.
	// If not specified, zidx will generate a table of contents with no links.
	Username string
}

func New() *Option {
	var opts Option

	flag.BoolVarP(&opts.ShowHelp, "help", "h", false, "show help")
	flag.StringVarP(&opts.Filename, "filename", "f", "Index.md", "output filename")
	flag.StringVarP(&opts.Username, "username", "u", "", "Zenn username")

	return &opts
}

func init() {
	flag.Parse()
}
