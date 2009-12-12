package dom

/*
 * Part of the xml/dom Go package
 *
 * Declares the interfaces from DOM Core Level 3
 * http://www.w3.org/TR/DOM-Level-3-Core/
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */ 

// TODO: split this out into separate interfaces again eventually


// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1950641247
type Node interface {
  AppendChild(Node) Node;
  RemoveChild(Node) Node;
  // attributes
  NodeName() string;
  NodeValue() string;
  NodeType() int;
  ParentNode() Node;
  ChildNodes() NodeList;
  Attributes() NamedNodeMap;
  HasChildNodes() bool;

  // internal interface methods needed for implementations (not part of the DOM)
  setParent(Node);
  insertChildAt(Node,uint);
  removeChild(Node);
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-745549614
type Element interface {
  Node;
  TagName() string;
  GetAttribute(name string) string;
  SetAttribute(name string, value string);
  OwnerDocument() Document;
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#i-Document
type Document interface {
  Node;
  DocumentElement() Element;
  CreateElement(tagName string) Element;
  OwnerDocument() Document;
}

// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-FF21A306
type CharacterData interface {
  Node;
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1312295772
type Text interface {
  CharacterData;
  OwnerDocument() Document;
}

// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-637646024
type Attr interface {
  Node;
  OwnerDocument() Document;
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-536297177
type NodeList interface {
  Length() uint;
  Item(index uint) Node;
}

// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1780488922
type NamedNodeMap interface {
  Length() uint;
  Item(index uint) Node;
}
