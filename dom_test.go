package dom_test

import (
  "testing";
  "xml/dom";
  "strconv";
)

// Document.nodeName should be #document
// see http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1841493061
func TestDocumentNodeName(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  if (d.NodeName() != "#document") {
    t.Errorf("Document.nodeName != #document");
  }
}

// Document.nodeType should be 9
func TestDocumentNodeType(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  if (d.NodeType() != 9) {
    t.Errorf("Document.nodeType not equal to 9");
  }
}

// Document.documentElement should return an object implementing Element
func TestDocumentElementIsAnElement(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  _,ok := (d.DocumentElement()).(dom.Element);
  if (!ok) {
  	t.Errorf("Document.documentElement did not return an Element");
  }
}

func TestDocumentElementNodeName(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root,_ := (d.DocumentElement()).(dom.Element);
  if (root.NodeName() != "foo") {
  	t.Errorf("Element.nodeName not set correctly");
  }
}

// Element.nodeType should be 1
func TestElementNodeType(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root,_ := (d.DocumentElement()).(dom.Element);
  if (root.NodeType() != 1) {
    t.Errorf("Element.nodeType not equal to 1");
  }
}

func TestElementGetAttribute(t *testing.T) {
  var d = dom.ParseString("<foo bar='baz'></foo>");
  root,_ := (d.DocumentElement()).(dom.Element);
  if (root.GetAttribute("bar") != "baz") {
  	t.Errorf("Element.getAttribute() did not return the attribute value");
  }
}

func TestElementSetAttribute(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root,_ := (d.DocumentElement()).(dom.Element);
  root.SetAttribute("bar", "baz");
  if (root.GetAttribute("bar") != "baz") {
  	t.Errorf("Element.getAttribute() did not return the attribute value");
  }
}

func TestNodeListLength(t *testing.T) {
  var d = dom.ParseString(
  	`<foo>
  		<bar></bar>
  		<baz></baz>
  	</foo>`);
  root,_ := (d.DocumentElement()).(dom.Element);
  children := root.ChildNodes();
  l := int(children.Length());
  if ( l != 2) {
  	t.Errorf("NodeList.length did not return the correct number of children ("+strconv.Itoa(l)+" instead of 2)");
  }
}

func TestNodeListItem(t *testing.T) {
  var d = dom.ParseString(
  	`<foo>
  		<bar></bar>
  		<baz></baz>
  	</foo>`);
  root,_ := (d.DocumentElement()).(dom.Element);
  children := root.ChildNodes();
  if (children.Item(1).NodeName() != "baz" ||
      children.Item(0).NodeName() != "bar") {
  	t.Errorf("NodeList.item(i) did not return the correct child");
  }
}

func TestNodeListItemForNull(t *testing.T) {
  var d = dom.ParseString(
  	`<foo>
  		<bar></bar>
  		<baz></baz>
  	</foo>`);
  root,_ := (d.DocumentElement()).(dom.Element);
  children := root.ChildNodes();
  if (children.Item(2) != nil ||
      children.Item(100000) != nil) {
  	t.Errorf("NodeList.item(i) did not return nil");
  }
}

func TestNodeParentNode(t *testing.T) {
  var d = dom.ParseString(
  	`<foo>
  		<bar>
          <baz></baz>  		
  		</bar>
  	</foo>`);
  
  root,_ := (d.DocumentElement()).(dom.Element);
  child := root.ChildNodes().Item(0);
  grandchild := child.ChildNodes().Item(0);
  
  if (child.ParentNode() != root.(dom.Node) ||
      grandchild.ParentNode() != child ||
      grandchild.ParentNode().ParentNode() != root.(dom.Node)) {
  	t.Errorf("Node.ParentNode() did not return the correct parent");
  }
}

func TestNodeParentNodeOnRoot(t *testing.T) {
  var d = dom.ParseString(`<foo></foo>`);
  
  root,_ := (d.DocumentElement()).(dom.Element);
  
  if (root.ParentNode() != d.(dom.Node)) {
  	t.Errorf("documentElement.ParentNode() did not return the document");
  }
}

func TestNodeParentNodeOnDocument(t *testing.T) {
  var d = dom.ParseString(`<foo></foo>`);
  if (d.ParentNode() != nil) {
  	t.Errorf("document.ParentNode() did not return nil");
  }
}

// the root node of the document is a child node
func TestNodeDocumentChildNodesLength(t *testing.T) {
  var d = dom.ParseString(`<foo></foo>`);
  if (d.ChildNodes().Length() != 1) {
  	t.Errorf("document.ChildNodes().Length() did not return the number of children");
  }
}

func TestNodeDocumentChildNodeIsRoot(t *testing.T) {
  var d = dom.ParseString(`<foo></foo>`);
  root,_ := (d.DocumentElement()).(dom.Element);
  if (d.ChildNodes().Item(0) != root.(dom.Node)) {
  	t.Errorf("document.ChildNodes().Item(0) is not the documentElement");
  }
}