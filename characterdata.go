package dom

/*
 * Text node implementation
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2010, Jeff Schiller
 */

type CharacterData struct {
	_node;
	content []byte;
}

func (n *CharacterData) NodeType() uint { return CDATA_SECTION_NODE; }
func (n *CharacterData) NodeName() (s string) { return "#cdata-section"; }
func (n *CharacterData) NodeValue() (s string) { return string(n.content); }
func (n *CharacterData) PreviousSibling() Node { return previousSibling( Node(n), n.p.ChildNodes() ) }
func (n *CharacterData) NextSibling() Node { return nextSibling( Node(n), n.p.ChildNodes() ) }
func (n *CharacterData) OwnerDocument() *Document { return ownerDocument(n); }


func (n *CharacterData) SubstringData( offset uint32, count uint32 ) string { return string(n.content[offset:offset+count]); }

