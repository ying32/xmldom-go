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


type _comment struct {
	_node;
	content []byte
}

func (n *_comment) NodeType() uint { return COMMENT_NODE; }
func (n *_comment) NodeName() (s string) { return "#comment"; }
func (n *_comment) NodeValue() (s string) { return string(n.content); }
func (n *_comment) PreviousSibling() Node { return previousSibling( Node(n), n.p.ChildNodes() ) }
func (n *_comment) NextSibling() Node { return nextSibling( Node(n), n.p.ChildNodes() ) }
func (n *_comment) OwnerDocument() *Document { return ownerDocument(n) }

func newComment(token xml.Comment) (*_comment) {
	n := new( _comment )
	n.content = token.Copy()
	return n	
}

