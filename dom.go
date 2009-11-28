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

// internal structures that implements the above public interfaces
type elementImpl struct {}
func (e *elementImpl) NodeName() string { return ""; }
func (e *elementImpl) TagName() string { return e.NodeName(); }

type documentImpl struct {}
func (d *documentImpl) NodeName() string { return ""; }
func (d *documentImpl) DocumentElement() Element { return new(elementImpl); }

func ParseString(s string) *documentImpl {
  var d = new(documentImpl);
  return d;
}