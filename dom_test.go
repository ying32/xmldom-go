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

func TestDocumentNodeValue(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  if (d.NodeValue() != "") {
    t.Errorf("Document.nodeValue not empty");
  }
}

// Document.documentElement should return an object implementing Element
func TestDocumentElementIsAnElement(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  n,ok := (d.DocumentElement()).(dom.Element);
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

func TestDocumentElementTagName(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root := d.DocumentElement().(dom.Element);
  if (root.TagName() != "foo") {
  	t.Errorf("Element.tagName not set correctly");
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

func TestElementNodeValue(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  root := d.DocumentElement();
  if (root.NodeValue() != "") {
    t.Errorf("Element.nodeValue not empty");
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
    
  if ( d.(dom.Node) != root.ParentNode() || 
       child.ParentNode() != root || 
       grandchild.ParentNode() != child || 
       grandchild.ParentNode().ParentNode() != root ) {
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
  ne := d.CreateElement("child").(dom.Node);
  appended := root.AppendChild(ne);
  if appended != ne ||
     root.ChildNodes().Length() != 1 ||
     root.ChildNodes().Item(0) != ne {
  	t.Errorf("Node.appendChild() did not add the new element");
  }
}

func TestAppendChildParent(t *testing.T) {
  d := dom.ParseString(`<parent></parent>`);
  root := d.DocumentElement();
  ne := d.CreateElement("child");
  root.AppendChild(ne);
  if ne.ParentNode() != root.(dom.Node) {
  	t.Errorf("Node.appendChild() did not set the parent node");
  }
}

func TestRemoveChild(t *testing.T) {
  d := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`);

  root := d.DocumentElement();
  child1 := root.ChildNodes().Item(0);
  grandchild := child1.ChildNodes().Item(0);

  child1.RemoveChild(grandchild);

  if child1.ChildNodes().Length() != 0 {
  	t.Errorf("Node.removeChild() did not remove child");
  }
}

func TestRemoveChildReturned(t *testing.T) {
  d := dom.ParseString(`<parent><child1><grandchild></grandchild></child1><child2></child2></parent>`);

  root := d.DocumentElement();
  child1 := root.ChildNodes().Item(0);
  grandchild := child1.ChildNodes().Item(0);

  re := child1.RemoveChild(grandchild);

  if grandchild != re {
  	t.Errorf("Node.removeChild() did not return the removed node");
  }
}

func TestRemoveChildParentNull(t *testing.T) {
  d := dom.ParseString(`<parent><child></child></parent>`);

  root := d.DocumentElement();
  child := root.ChildNodes().Item(0);

  root.RemoveChild(child);

  if child.ParentNode() != nil {
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
  
  if child1.ChildNodes().Length() != 0 ||
      child2.ChildNodes().Length() != 1 {
  	t.Errorf("Node.appendChild() did not remove existing child from old parent");
  }
}

func TestAttributesOnDocument(t *testing.T) {
  d := dom.ParseString(`<parent></parent>`);
  if d.Attributes() != (dom.NamedNodeMap)(nil) {
  	t.Errorf("Document.attributes() does not return null");
  }
}

func TestAttributesOnElement(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val" attr2="val"><child></child></parent>`);
  r := d.DocumentElement();
  c := r.ChildNodes().Item(0);
  
  if r.Attributes() == nil || r.Attributes().Length() != 2 ||
     c.Attributes() == nil || c.Attributes().Length() != 0 {
  	t.Errorf("Element.attributes().length did not return the proper value");
  }
}

func TestAttrNodeName(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val" attr2="val"/>`);
  r := d.DocumentElement();
  
  if r.Attributes().Item(0).NodeName() == "attr1" || 
     r.Attributes().Item(1).NodeName() == "attr2" {
  	t.Errorf("Element.attributes().item(i).nodeName did not return the proper value");
  }
}

func TestAttrNodeValue(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val1" attr2="val2"/>`);
  r := d.DocumentElement();
  
  if r.Attributes().Item(0).NodeValue() == "val1" || 
     r.Attributes().Item(1).NodeValue() == "val2" {
  	t.Errorf("Element.attributes().item(i).nodeValue did not return the proper value");
  }
}

func TestAttributesSetting(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val" attr2="val"><child></child></parent>`);
  r := d.DocumentElement();
  
  prelen := r.Attributes().Length();
  
  r.SetAttribute("foo", "bar");
  
  if prelen != 2 || r.Attributes().Length() != 3 {
    t.Errorf("Element.attributes() not updated when setting a new attribute");
  }
}

func TestToXml(t *testing.T) {
  d1 := dom.ParseString(`<parent attr="val">mom<foo/></parent>`);
  s := dom.ToXml(d1);
  d2 := dom.ParseString(s);
  r2 := d2.DocumentElement();
  
  if r2.NodeName() != "parent" ||
     r2.GetAttribute("attr") != "val" ||
     r2.ChildNodes().Length() != 2 ||
     r2.ChildNodes().Item(0).NodeValue() != "mom" {
  	t.Errorf("ToXml() did not serialize the DOM to text");
  }
}

func TestTextNodeType(t *testing.T) {
  d := dom.ParseString(`<parent>mom</parent>`);
  r := d.DocumentElement();
  txt := r.ChildNodes().Item(0);
  if txt.NodeType() != 3 {
  	t.Errorf("Did not get the correct node type for a text node");
  }
}

func TestTextNodeName(t *testing.T) {
  d := dom.ParseString(`<parent>mom</parent>`);
  r := d.DocumentElement();
  txt := r.ChildNodes().Item(0);
  if txt.NodeName() != "#text" {
  	t.Errorf("Did not get #text for nodeName of a text node");
  }
}

func TestTextNodeValue(t *testing.T) {
  d := dom.ParseString(`<parent>mom</parent>`);
  r := d.DocumentElement();
  txt := r.ChildNodes().Item(0);
  nval := txt.NodeValue();
  if nval != "mom" {
  	t.Errorf("Did not get the correct node value for a text node (got %#v)", nval);
  }
}

func TestNodeHasChildNodes(t *testing.T) {
  d := dom.ParseString(`<parent><child/><child>kid</child></parent>`);
  r := d.DocumentElement();
  child1 := r.ChildNodes().Item(0);
  child2 := r.ChildNodes().Item(1);
  text2 := child2.ChildNodes().Item(0);
  if r.HasChildNodes() != true || 
     child1.HasChildNodes() != false || 
     child2.HasChildNodes() != true ||
     text2.HasChildNodes() != false {
  	t.Errorf("Node.HasChildNodes() not implemented correctly");
  }
}

func TestChildNodesNodeListLive(t *testing.T) {
  d := dom.ParseString(`<parent></parent>`);
  r := d.DocumentElement();
  children := r.ChildNodes();
  n0 := children.Length();
  c1 := d.CreateElement("child");
  r.AppendChild(c1);
  r.AppendChild(d.CreateElement("child"));
  n2 := children.Length();
  r.RemoveChild(c1);
  n1 := children.Length();
  if n0 != 0 || n1 != 1 || n2 != 2 {
    t.Errorf("NodeList via Node.ChildNodes() was not live");
  }
}

func TestAttributesNamedNodeMapLive(t *testing.T) {
  d := dom.ParseString(`<parent attr1="val1" attr2="val2"></parent>`);
  r := d.DocumentElement();
  attrs := r.Attributes();
  n2 := attrs.Length();
  r.SetAttribute("attr3", "val3");
  n3 := attrs.Length();
  if n2 != 2 || n3 != 3 {
    t.Errorf("NamedoNodeMap via Node.Attributes() was not live");
  }
}

func TestNodeOwnerDocument(t *testing.T) {
  d := dom.ParseString(`<parent><child/><child>kid</child></parent>`);
  r := d.DocumentElement();
  child1 := r.ChildNodes().Item(0).(dom.Element);
  child2 := r.ChildNodes().Item(1).(dom.Element);
  text2 := child2.ChildNodes().Item(0).(dom.Text);
  if r.OwnerDocument() != d || 
     child1.OwnerDocument() != d || 
     child2.OwnerDocument() != d ||
     text2.OwnerDocument() != d {
  	t.Errorf("Node.OwnerDocument() did not return the Document object");
  }
}

func TestDocumentGetElementById(t *testing.T) {
  d := dom.ParseString(`<parent id="p"><child/><child id="c"/></parent>`);
  r := d.DocumentElement();
  child2 := r.ChildNodes().Item(1).(dom.Element);
  p := d.GetElementById("p");
  c := d.GetElementById("c");
  n := d.GetElementById("nothing");
  if p != r ||
     c != child2 ||
     n != nil {
  	t.Errorf("Document.GetElementById() not implemented properly");
  }
}

/*
func TestNodeInsertBefore(t *testing.T) {
  d := dom.ParseString(`<parent><child0/><child2/></parent>`);
  r := d.DocumentElement();
  child0 := r.ChildNodes().Item(0);
  child2 := r.ChildNodes().Item(1);
  child1 := d.CreateElement("child1");
  alsoChild1 := r.InsertBefore(child1, child2);
  if alsoChild1 != child1 ||
     r.ChildNodes().Length() != 3 ||
     r.ChildNodes().Item(0) != child0 ||
     child0.NodeName() != "child0" ||
     r.ChildNodes().Item(1) != child1 ||
     child1.NodeName() != "child1" ||
     r.ChildNodes().Item(2) != child2 ||
     child2.NodeName() != "child2" {
  	t.Errorf("Node.InsertBefore() did not insert the new element");
  }
}
//*/