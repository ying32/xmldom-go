package dom

/*
 * Document interface implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */

import (
  "xml";
  "os";
)

type _doc struct {
  *_node;
}

func (d *_doc) AppendChild(c Node) Node { return appendChild(d,c); }
func (d *_doc) RemoveChild(c Node) Node { return removeChild(d,c); }
func (d *_doc) DocumentElement() Element { return d.ChildNodes().Item(0).(Element); }
func (d *_doc) CreateElement(tag string) Element { 
  return newElem(xml.StartElement { xml.Name { "", tag }, nil }); 
}
func (d *_doc) setRoot(r Element) Element {
  // empty the children vector
  if d.ChildNodes().Length() > 0 {
    os.Exit(-1);
  }
  appendChild(d,r);
  return r;
}

func newDoc() (*_doc) {
  return &_doc{ newNode(9) };
}
