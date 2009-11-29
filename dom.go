package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 2
 */

import (
  "strings";
  "xml";
)

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-1950641247
type Node interface {
	NodeName() string;
	NodeType() int;
}

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-745549614
type Element interface {
	Node;
	TagName() string;
	//GetAttribute(name string) string;
	//SetAttribute(name string, value string);
}

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#i-Document
type Document interface {
	Node;
	DocumentElement() Element;
}


type elem struct {}
func (e *elem) NodeName() string { return "elem.NodeName() not implemented"; }
func (e *elem) NodeType() int { return 1; }
func (e *elem) TagName() string { return e.NodeName(); }

// doc is our internal implementation of the Document interface
type doc struct {}
func (d *doc) NodeName() string { return "#document"; }
func (d *doc) NodeType() int { return 9; }
func (d *doc) DocumentElement() Element { return new(elem); }

func ParseString(s string) Document {
  var d = new(doc);
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  _,_ = p.Token(); // todo: get the data token by token
  return d;
}