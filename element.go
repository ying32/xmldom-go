package dom

/*
 * Element implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

import (
	"xml"
)

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-745549614
type Element struct {
  _node
  attribs map[string] string; // attributes of the element
}

func (e *Element) NodeType() uint { return ELEMENT_NODE; }
func (n *Element) NodeName() string { return n.n.Local; }
func (n *Element) NodeValue() string { return ""; }
func (n *Element) PreviousSibling() Node { return previousSibling( Node(n), n.p.ChildNodes() ) }
func (n *Element) NextSibling() Node { return nextSibling( Node(n), n.p.ChildNodes() ) }
func (n *Element) AppendChild(c Node) Node { return appendChild(n,c); }
func (n *Element) RemoveChild(c Node) Node { return removeChild(n,c); }
func (n *Element) OwnerDocument() *Document { return ownerDocument(n); }
func (n *Element) TagName() string { return n.NodeName(); }
func (n *Element) Attributes() NamedNodeMap { return newAttrNamedNodeMap(n); }
func (n *Element) GetElementById(id string) *Element {
  return getElementById(n,id)
}
func (n *Element) GetAttribute(name string) string {
  val, ok := n.attribs[name];
  if (!ok) {
    val = "";
  }
  return val;
}
func (n *Element) SetAttribute(attrname string, attrval string) {
  n.attribs[attrname]=attrval;
}
// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-6D6AC0F9
func (n *Element) RemoveAttribute(name string) {
  n.attribs[name] = "",false;
}
// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-ElHasAttr
func (n *Element) HasAttribute(name string) bool {
  _,has := n.attribs[name];
  return has;
}

func (n *Element) GetElementsByTagName(name string) NodeList {
  return newTagNodeList(n, name);
}

func newElem(token xml.StartElement) (*Element) {
	n := new(Element)
	n.n = token.Name
	n.attribs = make(map[string] string)
	return n;
}

// Custom routines solely for golang
func (n *Element) ToXml() string {
	return toXml( Node(n) )
}

