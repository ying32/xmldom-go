package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 2
 */

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-1950641247
type Node interface {
  NodeName() string;
}

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-745549614
type Element interface {
  Node; 
  TagName() string;
//  GetAttribute(name string) string;
//  SetAttribute(name string, value string);
}

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#i-Document
type Document interface {
  Node;
  DocumentElement() Element;
}

// internal structures that implement the above public interfaces

type elem struct {}
func (e *elem) NodeName() string { return "elem.NodeName() not implemented"; }
func (e *elem) TagName() string { return e.NodeName(); }

type doc struct {}
func (d *doc) NodeName() string { return ""; }
func (d *doc) DocumentElement() Element { return new(elem); }

func ParseString(s string) Document {
  var d = new(doc);
  return d;
}