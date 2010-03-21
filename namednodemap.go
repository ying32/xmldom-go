package dom

/*
 * NamedNodeMap implementation
 *
 * Copyright (c) 2010, Jeff Schiller
 */

// used to return the live attributes of a node
type _attrnamednodemap struct {
  e *_elem;
}

func (m *_attrnamednodemap) Length() uint {
  return uint(len(m.e.attribs));
}
func (m *_attrnamednodemap) Item(index uint) Node {
  if (index >= 0 && index < m.Length()) {
    i := uint(0);
    for name,val := range(m.e.attribs) {
      if i == index {
        return newAttr(name,val);
      }
      i += 1;
    }
  }
  return Node(nil);
}

func newAttrNamedNodeMap(e *_elem) (*_attrnamednodemap) {
  nm := new(_attrnamednodemap);
  nm.e = e;
  return nm;
}
