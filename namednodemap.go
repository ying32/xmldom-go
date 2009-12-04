package dom

/*
 * NamedNodeMap implementation
 *
 * Copyright (c) 2009, Jeff Schiller
 */

// used to return the live attributes of a node
type _attrnamednodemap struct {
}

func (nnm *_attrnamednodemap) Length() uint {
  return uint(0);
}
func (nnm *_attrnamednodemap) Item(index uint) Node {
  return Node(nil);
}
