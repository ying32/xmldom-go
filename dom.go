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
}

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-745549614
type Element interface {
	Node;
	TagName() string;
	GetAttribute(name string) string;
	SetAttribute(name string, value string);
}

// DOM2: http://www.w3.org/TR/DOM-Level-2-Core/core.html#i-Document
type Document interface {
	Node;
	DocumentElement() Element;
}

func ParseString(s string) (d *Document){
  d = new(Document);
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  _,_ = p.Token(); // todo: get the data token by token
  return d;
}

// doc is our internal implementation of the Document interface
type doc struct {
  
}

func (*doc) DocumentElement() (e Element) {
  return;
}
