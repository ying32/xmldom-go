package dom_test

import(
	"testing"
	"xml/dom"
)

func TestConst(t *testing.T) {
	if dom.ELEMENT_NODE != 1 {
		t.Errorf( "Value of ELEMENT_NODE is incorrect." )
	}
	if dom.ATTRIBUTE_NODE != 2 {
		t.Errorf( "Value of ATTRIBUTE_NODE is incorrect." )
	}
	if dom.TEXT_NODE != 3 {
		t.Errorf( "Value of TEXT_NODE is incorrect." )
	}
	if dom.DOCUMENT_NODE != 9 {
		t.Errorf( "Value of DOCUMENT_NODE is incorrect." )
	}
}

