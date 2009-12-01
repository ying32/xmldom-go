package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 2
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */

import (
  "container/vector";
  "strings";
  "xml";
  "fmt";
  "os";
)

type _node struct {
  p *_node; // parent
  c vector.Vector; // children
}

func (n *_node) NodeName() string {
  return "";
}

func (n *_node) NodeType() int {
  return -1;
}

func newNode() (n *_node) {
  n = new(_node);
  return;
}

func (n *_node) AppendChild(child *Node) (*Node) {
  n.c.Push(child);
  return child;
}

// implements the Element interface
type _elem struct {
  *_node;
}
func (e *_elem) NodeName() string { return "elem.NodeName() not implemented"; }
func (e *_elem) NodeType() int { return 1; }
func (e *_elem) TagName() string { return e.NodeName(); }

// implements the Document interface
type _doc struct {
  *_node;
  root *_node;
}
func (d *_doc) NodeName() string { return "#document"; }
func (d *_doc) NodeType() int { return 9; }
func (d *_doc) DocumentElement() Element { return new(_elem); }
func (d *_doc) setRoot(n *_node) *_node {
  d.root = n;
  return n;
}

func ParseString(s string) Document {
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  t, err := p.Token();
  d := new(_doc);
  e := (*_node)(nil); // e is the current parent
  for t != nil {
    t, err = p.Token();
    switch t1 := t.(type) {
      case xml.StartElement:
        n := newNode();
        
        if e == nil {
          // set doc root
          e = d.setRoot(n);
        } else {
          // this element is a child of e, the last element we found
          e = e.AppendChild(n);
        }
      case xml.EndElement:
    }
  }
  if err != os.EOF {
    fmt.Println(err.String());
  }
  return d;
}