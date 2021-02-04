// Code generated by rice embed-go; DO NOT EDIT.
package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1611306739, 0),

		Content: string("<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n    <title>Homepage</title>\r\n    <link rel=\"stylesheet\" href=\"style.css\" />\r\n</head>\r\n<body>\r\n    <h1>Hello, Sumeet.</h1>\r\n</body>\r\n</html>"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "style.css",
		FileModTime: time.Unix(1611306797, 0),

		Content: string("body {\r\n  background: #000000;\r\n  display: flex;\r\n  align-items: center;\r\n  justify-content: center;\r\n  min-height: 100vh;\r\n  margin: 0;\r\n  font-family: sans-serif;\r\n}\r\n\r\nh1 {\r\n  font-size: 75px;\r\n  font-family: \"Courier New\", Courier, monospace;\r\n  color: rgb(255, 255, 255);\r\n}\r\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1610107200, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"
			file3, // "style.css"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`Templates`, &embedded.EmbeddedBox{
		Name: `Templates`,
		Time: time.Unix(1610107200, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
			"style.css":  file3,
		},
	})
}
