package main

import (
	"fmt"
	"os"

	"github.com/antonlindstrom/imagechecker"
	"github.com/olekukonko/tablewriter"
)

func main() {
	client := imagechecker.New()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	doc, err := client.DocumentFromURL(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"URL", "Content-Type", "Response", "ETag"})

	for _, link := range doc.Links {
		table.Append([]string{link.URL, link.ContentType, fmt.Sprintf("%d", link.ResponseCode), link.ETag})
	}

	table.Render()
}
