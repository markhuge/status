package main

import (
	"flag"

	"github.com/markhuge/status/config"
	"github.com/markhuge/status/site"
	"github.com/markhuge/status/views"
)

func main() {
	// Setup config
	configFile := flag.String("f", "config.yml", "path to config file")
	flag.Parse()

	conf := config.Read(*configFile)

	var checks []*site.Site

	for _, item := range conf.Sites {
		checks = append(checks, site.New(item.URL))
	}

	table := views.NewTable(checks)

	views.Init([]*views.Table{table})
	table.Render()
}
