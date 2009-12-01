package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 3
 * http://www.w3.org/TR/DOM-Level-3-Core/
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

func (n *_node) NodeName() string { return "Node.NodeName() not implemented"; }
func (n *_node) NodeType() int { return -1; }
func (n *_node) AppendChild(child Node) Node {
  n.c.Push(child);
  return child;
}

func newNode() (n *_node) {
  n = new(_node);
  return;
}

// implements the Element interface
type _elem struct {
  *_node;
  n xml.Name; // name
}
func (e *_elem) NodeName() string { return e.n.Local; }
func (e *_elem) NodeType() int { return 1; }
func (e *_elem) TagName() string { return e.NodeName(); }
func (e *_elem) GetAttribute(name string) string { return "Element.getAttribute() not implemented"; }

func newElem(token xml.StartElement) (e *_elem) {
  e = new(_elem);
  e.n = token.Name;
  return;
}

// implements the Document interface
type _doc struct {
  *_node;
  root Element;
}
func (d *_doc) NodeName() string { return "#document"; }
func (d *_doc) NodeType() int { return 9; }
func (d *_doc) DocumentElement() Element { return d.root; }
func (d *_doc) setRoot(r Element) Element {
  d.root = r;
  return r;
}

func ParseString(s string) Document {
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  t, err := p.Token();
  d := new(_doc);
  e := (Element)(nil); // e is the current parent
  for t != nil {
    switch token := t.(type) {
      case xml.StartElement:
        newElem := newElem(token);
        if e == nil {
          // set doc root
          e = d.setRoot(newElem);
        } else {
          // this element is a child of e, the last element we found
          e,_ = e.AppendChild(newElem).(Element);
        }
      case xml.EndElement:
      	// TODO: go up to parent
      default:
        fmt.Println("Unknown type");
    }
    // get the next token
    t, err = p.Token();
  }
  if err != os.EOF {
    fmt.Println(err.String());
  }
  return d;
}
