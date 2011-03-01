package dom

/*
 * Document interface implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

import (
	"xml"
)

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#i-Document
type Document struct {
	*_node;
}


func (d *Document) NodeValue() string {
	return ""
}

func (d *Document) AppendChild(c Node) Node { 
	return appendChild(d,c)
}

func (d *Document) RemoveChild(c Node) Node { 
	return removeChild(d,c)
}

func (d *Document) DocumentElement() Element { 
	return d.ChildNodes().Item(0).(Element); 
}

func (d *Document) OwnerDocument() *Document { 
	return ownerDocument(d)
}

func (d *Document) CreateElement(tag string) Element {
	return newElem(xml.StartElement { xml.Name { "", tag }, nil })
}

func (d *Document) setRoot(r Element) Element {
	// empty the children vector
	if d.ChildNodes().Length() > 0 {
		panic( "Document.setRoot used on document with non-empty list of child nodes" )
	}
 	appendChild(d,r);
	return r;
}

// DOM Level 2
func (d *Document) GetElementById(id string) Element {
	return getElementById( d.DocumentElement(), id)
}

func newDoc() (*Document) {
  n := newNode( DOCUMENT_NODE )
  n.self = Node( n )
  return &Document{ n }
}

