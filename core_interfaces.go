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
  NodeName() string;
  NodeType() int;
  AppendChild(Node) Node;
  RemoveChild(Node);
  ChildNodes() NodeList;
  // attributes
  ParentNode() Node;
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-745549614
type Element interface {
  Node;
  TagName() string;
  GetAttribute(name string) string;
  SetAttribute(name string, value string);
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#i-Document
type Document interface {
  Node;
  DocumentElement() Element;
  CreateElement(tagName string) Element;
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-536297177
type NodeList interface {
  Length() uint;
  Item(index uint) Node;
}

type CharacterData interface {
  Node;
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1312295772
type Text interface {
  CharacterData;
}