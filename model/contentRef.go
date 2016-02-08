package model

type ContentRef struct {
	TagHolder Tags `xml:"tags"`
}

type Tags struct {
	Tags []Tag `xml:"tag"`
}

type Tag struct {
	Term     Term     `xml:"term"`
	TagScore TagScore `xml:"score"`
}

type Term struct {
	CanonicalName string `xml:"canonicalName"`
	Taxonomy      string `xml:"taxonomy,attr"`
	Id            string `xml:"id,attr"`
}

type TagScore struct {
	Confidence int `xml:"confidence,attr"`
	Relevance  int `xml:"relevance,attr"`
}
