package dom

type Document struct {
  
}

func ParseString(s string) (d *Document){
  d = new(Document);
  return d;
}