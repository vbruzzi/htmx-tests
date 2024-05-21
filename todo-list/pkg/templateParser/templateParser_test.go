package templateparser_test

import (
	"path/filepath"
	"testing"
	templateparser "vbruzzi/todo-list/pkg/templateParser"

	"github.com/stretchr/testify/assert"
)

func Walk(root string, fn filepath.WalkFunc) error {
	files := []string{"foo.html", "foo.css", "bar/", "bar", "foo/bar.html", "", "foo/bar.css"}
	for _, file := range files {
		fn(file, nil, nil)
	}

	return nil
}

func TestParser(t *testing.T) {
	templates := templateparser.ParseTemplates(Walk)
	assert := assert.New(t)
	assert.NotNil(t, templates)
}
