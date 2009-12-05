package dom

/*
 * NamedNodeMap implementation
 *
 * Copyright (c) 2009, Jeff Schiller
 */

// used to return the live attributes of a node
type _attrnamednodemap struct {
  e *_elem;
}

func (m *_attrnamednodemap) Length() uint {
  return uint(len(m.e.attribs));
}
func (m *_attrnamednodemap) Item(index uint) Node {
  return Node(nil);
}

func newAttrNamedNodeMap(e *_elem) (*_attrnamednodemap) {
  nm := new(_attrnamednodemap);
  nm.e = e;
  return nm;
}
