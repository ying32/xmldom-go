package dom

/*
 * Part of the xml/dom Go package
 *
 * Contains all the interfaces from DOM Core Level 3
 * http://www.w3.org/TR/DOM-Level-3-Core/
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */ 

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-1950641247
type Node interface {
	NodeName() string;
	NodeType() int;
  AppendChild(child *Node) (*Node);
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-745549614
type Element interface {
	Node;
	TagName() string;
	//GetAttribute(name string) string;
	//SetAttribute(name string, value string);
}

// DOM3: http://www.w3.org/TR/DOM-Level-3-Core/core.html#i-Document
type Document interface {
	Node;
	DocumentElement() Element;
}
