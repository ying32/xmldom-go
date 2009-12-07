package dom_test

import (
  "fmt";
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
  n,ok := (d.DocumentElement()).(dom.Node);
  if (!ok || n.NodeType() != 1) {
  	t.Errorf("Document.documentElement did not return an Element");
  }
}

func TestDocumentElementNodeName(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root := d.DocumentElement();
  if (root.NodeName() != "foo") {
  	t.Errorf("Element.nodeName not set correctly");
  }
}

// Element.nodeType should be 1
func TestElementNodeType(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root := d.DocumentElement();
  if (root.NodeType() != 1) {
    t.Errorf("Element.nodeType not equal to 1");
  }
}

func TestElementGetAttribute(t *testing.T) {
  var d = dom.ParseString("<foo bar='baz'></foo>");
  root := d.DocumentElement();
  if (root.GetAttribute("bar") != "baz") {
  	t.Errorf("Element.getAttribute() did not return the attribute value");
  }
}

func TestElementSetAttribute(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root := d.DocumentElement();
  root.SetAttribute("bar", "baz");
  if (root.GetAttribute("bar") != "baz") {
  	t.Errorf("Element.getAttribute() did not return the attribute value");
  }
}

func TestNodeListLength(t *testing.T) {
  var d = dom.ParseString(`<foo><bar></bar><baz></baz></foo>`);
  root := d.DocumentElement();
  children := root.ChildNodes();
  l := int(children.Length());
  if ( l != 2) {
  	t.Errorf("NodeList.length did not return the correct number of children ("+strconv.Itoa(l)+" instead of 2)");
  }
}

func TestNodeListItem(t *testing.T) {
  var d = dom.ParseString(`<foo><bar></bar><baz></baz></foo>`);
  root := d.DocumentElement();
  children := root.ChildNodes();
  if (children.Item(1).NodeName() != "baz" ||
      children.Item(0).NodeName() != "bar") {
  	t.Errorf("NodeList.item(i) did not return the correct child");
  }
}

func TestNodeListItemForNull(t *testing.T) {
  var d = dom.ParseString(`<foo><bar></bar><baz></baz></foo>`);
  root := d.DocumentElement();
  children := root.ChildNodes();
  if (children.Item(2) != nil ||
      children.Item(100000) != nil) {
  	t.Errorf("NodeList.item(i) did not return nil");
  }
}

func TestNodeParentNode(t *testing.T) {
  var d = dom.ParseString(`<foo><bar><baz></baz></bar></foo>`);
  
  root := d.DocumentElement().(dom.Node);
  child := root.ChildNodes().Item(0);
  grandchild := child.ChildNodes().Item(0);
  
  if (child.ParentNode() != root ||
      grandchild.ParentNode() != child ||
      grandchild.ParentNode().ParentNode() != root) {
  	t.Errorf("Node.ParentNode() did not return the correct parent");
  }
}

func TestNodeParentNodeOnRoot(t *testing.T) {
  var d = dom.ParseString(`<foo></foo>`);
  
  root := d.DocumentElement().(dom.Node);
  
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
  d := dom.ParseString(`<foo></foo>`);
  root := d.DocumentElement().(dom.Node);
  if (d.ChildNodes().Item(0) != root) {
  	t.Errorf("document.ChildNodes().Item(0) is not the documentElement");
  }
}

func TestDocumentCreateElement(t *testing.T) {
  d := dom.ParseString(`<foo></foo>`);
  ne := d.CreateElement("child");
  if (ne.NodeName() != "child") {
  	t.Errorf("document.CreateNode('child') did not create a <child> Element");
  }
}

func TestAppendChild(t *testing.T) {
  d := dom.ParseString(`<parent></parent>`);
  root := d.DocumentElement();
  ne := d.CreateElement("child");
  appended := root.AppendChild(ne).(dom.Node);
  if (appended != ne ||
      root.ChildNodes().Length() != 1 ||
      root.ChildNodes().Item(0) != ne.(dom.Node))
  {
  	t.Errorf("Node.appendChild() did not add the new element");
  }
}

func TestAppendChildParent(t *testing.T) {
  d := dom.ParseString(`<parent></parent>`);
  root := d.DocumentElement();
  ne := d.CreateElement("child");
  root.AppendChild(ne);
  if (ne.ParentNode() != root.(dom.Node))
  {
  	t.Errorf("Node.appendChild() did not set the parent node");
  }
}

func TestRemoveChild(t *testing.T) {
  d := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`);

  root := d.DocumentElement();
  child1 := root.ChildNodes().Item(0);
  grandchild := child1.ChildNodes().Item(0);

  child1.RemoveChild(grandchild);

  if (child1.ChildNodes().Length() != 0)
  {
  	t.Errorf("Node.removeChild() did not remove child");
  }
}

func TestRemoveChildReturned(t *testing.T) {
  d := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`);

  root := d.DocumentElement();
  child1 := root.ChildNodes().Item(0);
  grandchild := child1.ChildNodes().Item(0);

  re := child1.RemoveChild(grandchild);

  if (grandchild != re)
  {
  	t.Errorf("Node.removeChild() did not return the removed node");
  }
}

func TestRemoveChildParentNull(t *testing.T) {
  d := dom.ParseString(`<parent><child></child></parent>`);

  root := d.DocumentElement();
  child := root.ChildNodes().Item(0);

  root.RemoveChild(child);

  if (child.ParentNode() != nil)
  {
  	t.Errorf("Node.removeChild() did not null out the parentNode");
  }
}

// See http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-184E7107
// "If the newChild is already in the tree, it is first removed."
func TestAppendChildExisting(t *testing.T) {
  d := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`);

  root := d.DocumentElement();
  child1 := root.ChildNodes().Item(0);
  child2 := root.ChildNodes().Item(1);
  grandchild := child1.ChildNodes().Item(0);

  child2.AppendChild(grandchild);
  
  if (child1.ChildNodes().Length() != 0 ||
      child2.ChildNodes().Length() != 1)
  {
  	t.Errorf("Node.appendChild() did not remove existing child from old parent");
  }
}

func TestAttributesOnDocument(t *testing.T) {
  d := dom.ParseString(`<parent></parent>`);
  fmt.Println(d.Attributes());
  if (d.Attributes() != (dom.NamedNodeMap)(nil))
  {
  	t.Errorf("Document.attributes() does not return null");
  }
}

func TestAttributesOnElement(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val" attr2="val"><child></child></parent>`);
  r := d.DocumentElement();
  c := r.ChildNodes().Item(0);
  
  if (r.Attributes() == nil || r.Attributes().Length() != 2 ||
      c.Attributes() == nil || c.Attributes().Length() != 0)
  {
  	t.Errorf("Element.attributes().length did not return the proper value");
  }
}

func TestAttributesSetting(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val" attr2="val"><child></child></parent>`);
  r := d.DocumentElement();
  
  prelen := r.Attributes().Length();
  
  r.SetAttribute("foo", "bar");
  
  if (prelen != 2 || r.Attributes().Length() != 3) {
    t.Errorf("Element.attributes() not updated when setting a new attribute");
  }
}

func TestToXml(t *testing.T) {
  d1 := dom.ParseString(`<parent attr="val"><child><grandchild></grandchild></child></parent>`);
  s := dom.ToXml(d1);
  d2 := dom.ParseString(s);
  
  if (d1.DocumentElement().NodeName() != d2.DocumentElement().NodeName() ||
      d1.DocumentElement().ChildNodes().Length() != d2.DocumentElement().ChildNodes().Length() ||
      d1.DocumentElement().GetAttribute("attr") != d2.DocumentElement().GetAttribute("attr"))
  {
  	t.Errorf("ToXml() did not serialize the DOM to text");
  }
}