package dom

/*
 * Node implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */

// TODO: think about how to make this class a bit more generic to promote extensibility
//       (for instance, this class has to know about Attr, Element and Document types to
//        implement NodeName() among other things)

import (
  "container/vector";
  "xml";
)

type _node struct {
  T int; // node type
  p Node; // parent
  c vector.Vector; // children
  n xml.Name; // name
}

// internal methods used so that our workhorses can do the real work
func (n *_node) setParent(p Node) {
  n.p = p;
}
func (n *_node) insertChildAt(c Node, i uint) {
  n.c.Insert(int(i), c);
}
func (n *_node) removeChild(c Node) {
  for i := n.c.Len()-1 ; i >= 0 ; i-- {
    if n.c.At(i).(Node) == c {
      n.c.Delete(i);
      break;
    }
  }
}

func (n *_node) NodeName() string {
  switch n.T {
    case 1: return n.n.Local;
    case 2: return n.n.Local;
    case 9: return "#document";
  }
  return "Node.NodeName() not implemented";
}
func (n *_node) NodeValue() string { return "Node.NodeValue() not implemented"; }
func (n *_node) TagName() string { return n.NodeName(); }
func (n *_node) NodeType() int { return n.T; }
func (n *_node) AppendChild(c Node) Node { return appendChild(n,c); }
func (n *_node) RemoveChild(c Node) Node { return removeChild(n,c); }
func (n *_node) ChildNodes() NodeList { return newChildNodelist(n); }
func (n *_node) ParentNode() Node { return n.p; }
func (n *_node) Attributes() NamedNodeMap { return NamedNodeMap(nil); }
func (n *_node) HasChildNodes() (b bool) {
  b = false;
  if n.c.Len() > 0 {
    b = true;
  }
  return;
}

// has to be package-scoped because of 
func ownerDocument(n Node) (d Document) {
  d = nil;
  
  for n!=nil {
    if n.NodeType()==9 {
      return n.(Document);
    }
    n = n.ParentNode();
  }
  return Document(nil);
}

//func (n *_node) OwnerDocument(n Node) (d Document) {
  //d = nil;
  //p := n.p;
  //
  //for p!=nil {
  //  if p.NodeType()==9 {
  //    return (*_doc)(p);
  //  }
  //  p = n.p;
  //}
//  return Document(nil);
//}


func newNode(_t int) (n *_node) {
  n = new(_node);
  n.T = _t;
  return;
}
