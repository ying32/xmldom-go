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

type _node struct {
  p Node; // parent
  c vector.Vector; // children
}
func (n *_node) NodeName() string { return "Node.NodeName() not implemented"; }
func (n *_node) NodeType() int { return -1; }
func (n *_node) AppendChild(child Node) Node {
  // if the child is already in the tree somewhere,
  // remove it before reparenting
  if child.ParentNode() != nil {
    child.ParentNode().RemoveChild(child);
  }
  n.c.Push(child);
  return child;
}
func (n *_node) RemoveChild(child Node) Node {
  for i := n.c.Len()-1 ; i >= 0 ; i-- {
    if n.c.At(i).(Node) == child {
      n.c.Delete(i);
      break;
    }
  }
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
func (d *_doc) CreateElement( tag string ) Element {
  return newElem(
    xml.StartElement { xml.Name { "", tag }, nil });
}
  
  
func newDoc() (*_doc) {
  return &_doc {
        new(_node)
        };
}
// ====================================

type _cdata struct {
  *_node;
  content []byte;
}

type _text struct {
  *_cdata;
}

func newText(token xml.CharData) (*_text) {
  cd := &_cdata{
    new(_node),
    token
  };
  
  return &_text {
    cd      
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
      case xml.CharData:
        e.AppendChild(newText(token));
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

// called recursively
func toXml(n Node) string {
  s := "<" + n.NodeName();
  
  // TODO: iterate over attributes
  
  s += ">";
  
  // iterate over children
  for ch := uint(0); ch < n.ChildNodes().Length(); ch++ {
    s += toXml(n.ChildNodes().Item(ch));
  }
  
  s += "</" + n.NodeName() + ">";
  return s;
}

func ToXml(doc Document) string {
  return toXml(doc.DocumentElement());
}
