package dom

/*
 * Attr implementation
 *
 * Copyright (c) 2009, Jeff Schiller
 */

type _attr struct {
  *_node;
}
func (a *_attr) NodeName() string { return "Not implemented"; }
func (a *_attr) NodeType() int { return 2; }
func (a *_attr) AppendChild(n Node) Node { return n; }
func (a *_attr) RemoveChild(n Node) Node { return n; }
func (a *_attr) ParentNode() Node { return Node(nil); }
func (a *_attr) ChildNodes() NodeList { return NodeList(nil); }
func (a *_attr) Attributes() NamedNodeMap { return NamedNodeMap(nil); }

func newAttr() (*_attr) {
  return &_attr {
        new(_node)
        };
}