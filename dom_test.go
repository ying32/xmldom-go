package dom_test

import (
  "testing";
  "xml/dom";
)

func TestDocumentNodeName(t *testing.T) {
  var d = dom.ParseString("<foo></foo>");
  if (d.NodeName() != "#document") {
    t.Errorf("Document.nodeName != #document");
  }
}

