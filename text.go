package dom

/*
 * Node implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
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

func (t *_text) NodeValue() (s string) {
  return string(t.content);
}

func newText(token xml.CharData) (*_text) {
  return &_text{ &_cdata{newNode(3)}, token.Copy() };
}
