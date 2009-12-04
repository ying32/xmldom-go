package dom

/*
 * NodeList implementations
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */

// A _childNodelist only stores a reference to its parent node.
// This way the list can be live, each time Length() or Item is
// called, fresh results are returned.
type _childNodelist struct {
  p *_node;
}

func (nl *_childNodelist) Length() uint {
  return uint(nl.p.c.Len());
}
func (nl *_childNodelist) Item(index uint) Node {
  n := Node(nil);
  if index < uint(nl.p.c.Len()) {
    // TODO: what if index == nl.p.c.Len() -1 and a node is deleted right now?
    n = nl.p.c.At(int(index)).(Node);
  }
  return n;
}
func newChildNodelist(p *_node) (*_childNodelist) {
  nl := new(_childNodelist);
  nl.p = p;
  return nl;
}
