package dom

/*
 * Text node implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

import (
  "xml";
)


type _cdata struct {
  *_node;
}

type _text struct {
  *_cdata;
  content []byte;
}

func (t *_text) NodeName() (s string) { return "#text"; }
func (t *_text) NodeValue() (s string) { return string(t.content); }

func (t *_text) OwnerDocument() *Document {
  return ownerDocument(t);
}

func newText(token xml.CharData) (*_text) {
  n := newNode(3);
  t := &_text{ &_cdata{n}, token.Copy() };
  n.self = Node(t)
  return t
}
