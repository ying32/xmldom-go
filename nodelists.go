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

type _tagNodeList struct {
  e *_elem;
  tag string
}

func (nl *_tagNodeList) Length() uint {
  return 0;
}

func (nl *_tagNodeList) Item(index uint) Node {
  var count uint = 0
  e := nl.e
  if e.NodeType() == 1 {
    // check for an id
    if e.TagName() == nl.tag {
      if index == count {
        return e
      }
      count++
    }
    // if not found, check the children
    cnodes := e.ChildNodes()
    var ix uint
    clen := cnodes.Length();
    for ix = 0 ; ix < clen ; ix++ {
      cnode := cnodes.Item(ix)
      // can't cast safely unless it's an Element for reals
      if cnode.NodeType() == 1 {
        result := nl.Item(index - count)
        if result != nil {
          return result
        }
      }
    }
  }
  return Node(nil);
}

func newTagNodeList(p *_elem, t string) (*_tagNodeList) {
  nl := new(_tagNodeList);
  nl.e = p;
  return nl;
}
