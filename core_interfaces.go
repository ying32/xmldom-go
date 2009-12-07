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

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1950641247
type Node interface {
  SetParent(Node);
  AppendChild(Node) Node;
  RemoveChild(Node) Node;
  // attributes
  NodeName() string;
  NodeValue() string;
  NodeType() int;
  ParentNode() Node;
  ChildNodes() NodeList;
  Attributes() NamedNodeMap;
  TagName() string;
  GetAttribute(name string) string;
  SetAttribute(name string, value string);
  DocumentElement() Node;
  CreateElement(tagName string) Node;
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