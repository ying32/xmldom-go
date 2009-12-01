package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 2
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */

import (
  "strings";
  "xml";
)

// implements the Element interface
type elem struct {}
func (e *elem) NodeName() string { return "elem.NodeName() not implemented"; }
func (e *elem) NodeType() int { return 1; }
func (e *elem) TagName() string { return e.NodeName(); }

// implements the Document interface
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