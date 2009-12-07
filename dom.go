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
  T int; // node type
  p Node; // parent
  c vector.Vector; // children
  n xml.Name; // name
  attribs map[string] string; // attributes of the element
  content []byte;
}
func (n *_node) SetParent(_p Node) {
  n.p = _p;
}
func (n *_node) NodeName() string {
  switch n.T {
    case 1: return n.n.Local;
    case 9: return "#document";
  }
  return "Node.NodeName() not implemented";
}
func (n *_node) TagName() string { return n.NodeName(); }
func (n *_node) NodeType() int { return n.T; }

func (n *_node) DocumentElement() Node {
  return n.c.At(0).(Node);
}

func (n *_node) setRoot(r Node) Node {
  // empty the children vector
  if n.c.Len() > 0 {
    os.Exit(-1);
  }
  n.AppendChild(r);
  return r;
}

func (n *_node) CreateElement(tag string) Node {
  return Node(newElem(xml.StartElement { xml.Name { "", tag }, nil }));
}

func (n *_node) AppendChild(child Node) Node {
  // if the child is already in the tree somewhere,
  // remove it before reparenting
  if child.ParentNode() != nil {
    child.ParentNode().RemoveChild(child);
  }
  n.c.Push(child);
  child.SetParent(n);
  return child;
}
func (n *_node) RemoveChild(child Node) Node {
  for i := n.c.Len()-1 ; i >= 0 ; i-- {
    if n.c.At(i).(Node) == child {
      n.c.Delete(i);
      break;
    }
  }
  child.SetParent(nil);
  return child;
}

func (n *_node) ChildNodes() NodeList {
  return newChildNodelist(n);
}
func (n *_node) ParentNode() Node {
  return n.p;
}
func (n *_node) GetAttribute(name string) string {
  val, ok := n.attribs[name];
  if (!ok) {
    val = "";
  }
  return val;
}
func (n *_node) SetAttribute(attrname string, attrval string) {
  n.attribs[attrname]=attrval;
}

func (n *_node) Attributes() NamedNodeMap {
  if (n.NodeType() == 1) {
    return newAttrNamedNodeMap(n);
  }
  return NamedNodeMap(nil);
}

func newNode(_t int) (n *_node) {
  n = new(_node);
  n.T = _t;
  return;
}

func newElem(token xml.StartElement) (*_node) {
  n := newNode(1);
  n.n = token.Name;
  n.attribs = make(map[string] string);
  return n;
}
  
func newDoc() (*_node) {
  return newNode(9);
}

/*
type _cdata struct {
  *_node;
}

type _text struct {
  *_cdata;
}
*/

func newText(token xml.CharData) (*_node) {
  text := newNode(3);
  text.content = token;
  return text;
}
// ====================================


func ParseString(s string) Node {
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  t, err := p.Token();
  d := newDoc();
  e := (Node)(nil); // e is the current parent
  for t != nil {
    switch token := t.(type) {
      case xml.StartElement:
        el := newElem(token);
//        fmt.Println("Starting ", el.NodeName());
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
          e = e.AppendChild(el);
        }
      case xml.CharData:
        e.AppendChild(newText(token));
      case xml.EndElement:
        e = e.ParentNode();
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
  s := "";
  switch n.NodeType() {
    case 1: // Element Nodes
      s += "<" + n.NodeName();
  
      // TODO: iterate over attributes
  
      s += ">";
  
      // iterate over children
      for ch := uint(0); ch < n.ChildNodes().Length(); ch++ {
        s += toXml(n.ChildNodes().Item(ch));
      }
  
      s += "</" + n.NodeName() + ">";
      
    case 3: // Text Nodes
      break;
  }
  return s;
}

func ToXml(doc Node) string {
  return toXml(doc.DocumentElement());
}
