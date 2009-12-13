package dom

/*
 * Implements a very small, very non-compliant subset of the DOM Core Level 3
 * http://www.w3.org/TR/DOM-Level-3-Core/
 *
 * Copyright (c) 2009, Rob Russell
 * Copyright (c) 2009, Jeff Schiller
 */
 
// FIXME: we use the empty string "" to denote a 'null' value when the data type
// according to the DOM API is expected to be a string. Perhaps return a pointer to a string?

import (
  "strings";
  "xml";
  "fmt";
  "os";
)

const (
  DEBUG = true;
)

func dbg(x ...) {
  if DEBUG {
    fmt.Println(x);
  }
}

// ====================================

// these are the package-level functions that are the real workhorses
// they only use interface types

func appendChild(p Node, c Node) Node {
  // if the child is already in the tree somewhere,
  // remove it before reparenting
  if c.ParentNode() != nil {
    removeChild(c.ParentNode(), c);
  }
  i := p.ChildNodes().Length();
  p.insertChildAt(c, i);
  c.setParent(p);
  return c;
}

func removeChild(p Node, c Node) Node {
  p.removeChild(c);
  c.setParent(nil);
  return c;
}

func getElementById(e Element, id string) Element {
  if e.NodeType() == 1 {
    // check for an id
    av := e.GetAttribute("id");
    if av != "" {
      if av==id {
        return e;
      }
    }
    // if not found, check the children
    cnodes := e.ChildNodes()
    var ix uint
    clen := cnodes.Length();
    for ix = 0 ; ix < clen ; ix++ {
    //for c := range e.c {
      // return the first one found
      //ce := cnodes.Item(ix).(Element).GetElementById(id);
      cnode := cnodes.Item(ix)
      // can't cast safely unless it's an Element for reals
      if cnode.NodeType() == 1 { 
        ce := getElementById(cnode.(Element),id);
        if ce != nil {
          return ce;
        }
      }
    }
  }
  return nil;
}

func ParseString(s string) Document {
  r := strings.NewReader(s);
  p := xml.NewParser(r);
  t, err := p.Token();
  d := newDoc();
  e := (Node)(nil); // e is the current parent
  for t != nil {
    switch token := t.(type) {
      case xml.StartElement:
        el := newElem(token);
        for ar := range(token.Attr) {
          el.SetAttribute(token.Attr[ar].Name.Local, token.Attr[ar].Value);
        }
        if e == nil {
          // set doc root
          // this element is a child of e, the last element we found
          e = d.setRoot(el);
        } else {
          // this element is a child of e, the last element we found
          e = e.AppendChild(el);
        }
      case xml.CharData:
        e.AppendChild(newText(token));
      case xml.EndElement:
        e = e.ParentNode();
      default:
      	// TODO: add handling for other types (text nodes, etc)
    }
    // get the next token
    t, err = p.Token();
  }
  if err != os.EOF {
    fmt.Println(err.String());
  }
  return d;
}

// called recursively
func toXml(n Node) string {
  s := "";
  switch n.NodeType() {
    case 1: // Element Nodes
      s += "<" + n.NodeName();
  
      // iterate over attributes
      for i := uint(0); i < n.Attributes().Length(); i++ {
        a := n.Attributes().Item(i);
        s += " " + a.NodeName() + "=\"" + a.NodeValue() + "\"";
      }
  
      s += ">";
  
      // iterate over children
      for ch := uint(0); ch < n.ChildNodes().Length(); ch++ {
        s += toXml(n.ChildNodes().Item(ch));
      }
  
      s += "</" + n.NodeName() + ">";
      
    case 3: // Text Nodes
      s += n.NodeValue();
      break;
  }
  return s;
}

func ToXml(doc Document) string {
  return toXml(doc.DocumentElement());
}
