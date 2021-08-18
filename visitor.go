package phptag_finder

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

type PhpTagFinder struct {
	Tags []Tag
}

func NewVisitor() *PhpTagFinder {
	return &PhpTagFinder{
		Tags: []Tag{},
	}
}

func (v *PhpTagFinder) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)
	switch typedNode := n.(type) {
	case *node.Root:
		log.Logger.Debug("root")
		if len(typedNode.Stmts) > 0 {
			firstStmt := typedNode.Stmts[0]
			switch firstStmt.(type) {
			case *stmt.InlineHtml:
				// next walk will access it
			default:
				fsPos := firstStmt.GetPosition()
				tag := NewTag(position.NewPosition(1, fsPos.StartLine, 0, fsPos.StartPos))
				v.Tags = append(v.Tags, tag)
			}

			lastStmt := typedNode.Stmts[len(typedNode.Stmts)-1]
			if _, is := lastStmt.(*stmt.InlineHtml); is {
				return true
			}
		}
	case *stmt.InlineHtml:
		log.Logger.Debug("html")
		pos := n.GetPosition()
		tag := NewTag(position.NewPosition(pos.EndLine, -1, pos.EndPos, -1))
		v.Tags = append(v.Tags, tag)
	default:
		log.Logger.Debug("default")
		if len(v.Tags) > 0 {
			last := v.Tags[len(v.Tags)-1]
			if last.Position.EndLine == -1 {
				pos := n.GetPosition()
				last.Position.EndLine = pos.StartLine
				last.Position.EndPos = pos.StartPos
			}
		}
	}
	return true
}

func (v *PhpTagFinder) LeaveNode(w walker.Walkable) {

}

func (v *PhpTagFinder) EnterChildNode(key string, w walker.Walkable) {
}
func (v *PhpTagFinder) LeaveChildNode(key string, w walker.Walkable) {
}

func (v *PhpTagFinder) EnterChildList(key string, w walker.Walkable) {
}

func (v *PhpTagFinder) LeaveChildList(key string, w walker.Walkable) {
}
