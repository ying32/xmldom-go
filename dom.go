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

// ====================================
// NodeList implementation
type _nodelist struct {
}
func (nl *_nodelist) Length() uint {
  return 0;
}
func (nl *_nodelist) Item(index uint) Node {
  return new(_node);
}

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

// ====================================

// ====================================
type _node struct {
  p Node; // parent
  c vector.Vector; // children
}
func (n *_node) NodeName() string { return "Node.NodeName() not implemented"; }
func (n *_node) NodeType() int { return -1; }
func (n *_node) AppendChild(child Node) Node {
  n.c.Push(child);
  return child;
}

func (n *_node) ChildNodes() NodeList {
  return newChildNodelist(n);
}

func (n *_node) ParentNode() Node {
  return n.p;
}

// TODO: never called now?
func newNode() (n *_node) {
  n = new(_node);
  return;
}
// ====================================

// ====================================
// implements the Element interface
type _elem struct {
  *_node;
  n xml.Name; // name
  attribs map[string] string; // attributes of the element
}
func (e *_elem) NodeName() string { return e.n.Local; }
func (e *_elem) NodeType() int { return 1; }
func (e *_elem) TagName() string { return e.NodeName(); }
func (e *_elem) GetAttribute(name string) string {
  val, ok := e.attribs[name];
  if (!ok) {
    val = "";
  }
  return val;
}
func (e *_elem) SetAttribute(attrname string, attrval string) {
  e.attribs[attrname]=attrval;
}

// this is our _elem constructor, it takes care to initialize
// the unnamed *_node field
func newElem(token xml.StartElement) (*_elem) {
  return &_elem {
        new(_node), 
        token.Name, 
        make(map[string] string)
      };
}
// ====================================

// ====================================
// implements the Document interface
type _doc struct {
  *_node;
}
func (d *_doc) NodeName() string { return "#document"; }
func (d *_doc) NodeType() int { return 9; }
func (d *_doc) DocumentElement() Element {
  return d.c.At(0).(Element);
}
func (d *_doc) setRoot(r Element) Element {
  // empty the children vector
  if d.c.Len() > 0 {
    os.Exit(-1);
  }
  d.AppendChild(r);
  return r;
}
func (d *_doc) CreateElement(tagName string) Element {
  return &_elem{};
}
func newDoc() (*_doc) {
  return &_doc {
        new(_node)
        };
}
// ====================================

func ParseString(s string) Document {
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  t, err := p.Token();
  d := newDoc();
  e := (Element)(nil); // e is the current parent
  for t != nil {
    switch token := t.(type) {
      case xml.StartElement:
        el := newElem(token);
        for ar := range(token.Attr) {
          el.SetAttribute(token.Attr[ar].Name.Local, token.Attr[ar].Value);
        }
        if e == nil {
          // set doc root
          el.p = d;
          e = d.setRoot(el);
        } else {
          // this element is a child of e, the last element we found
          el.p = e;
          e,_ = e.AppendChild(el).(Element);
        }
      case xml.EndElement:
        // up the tree
        switch q := e.ParentNode().(type) {
          case Document:
            e = nil;
          case Element:
            e = q;
        }
      default:
      	// TODO: add handling for other types (text nodes, etc)
//        fmt.Println("Unknown type");
    }
    // get the next token
    t, err = p.Token();
  }
  if err != os.EOF {
    fmt.Println(err.String());
  }
  return d;
}
