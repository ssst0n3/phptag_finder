package phptag_finder

import (
	"fmt"
	"github.com/z7zmey/php-parser/php7"
)

func Find(src []byte, version string) (tags []Tag) {
	parser := php7.NewParser(src, version)
	parser.Parse()

	for _, e := range parser.GetErrors() {
		fmt.Println(e)
	}

	//dump := visitor.Dumper{Writer: os.Stdout}
	//rootNode := parser.GetRootNode()
	//rootNode.Walk(&dump)

	v := NewVisitor()
	rootNode := parser.GetRootNode()
	rootNode.Walk(v)
	return v.Tags
}
