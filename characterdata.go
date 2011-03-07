package dom

import (
	"strconv"
)

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


func (n *CharacterData) SubstringData( offset uint32, count uint32 ) string { 
	return string(n.content[offset:offset+count])
}

func (n *CharacterData) String() string {
	return string( n.content )
}

func (n *CharacterData) EscapedBytes() []byte {
	runes := []int( string( n.content ) )
	
	output := make( []byte, 0 )
	
	for _, r := range runes {
		switch {
		case r=='<':
			output = append( output, []byte( "&lt;" )... )
		case r=='>':
			output = append( output, []byte( "&gt;" )... )
		case r=='&':
			output = append( output, []byte( "&amp;" )... )
		case r<128:
			output = append( output, byte(r) )
		default:
			s := "&#" + strconv.Itoa( r ) + ";"
			output = append( output, []byte(s)... )
		}
	}
	
	return output
}
