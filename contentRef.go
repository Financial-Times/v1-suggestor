package main

// ContentRef models the data as it comes from the metadata publishing event
type ContentRef struct {
	TagHolder Tags `xml:"tags"`
	PrimarySection Term `xml:"primarySection"`
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
	ID            string `xml:"id,attr"`
}

type TagScore struct {
	Confidence int `xml:"confidence,attr"`
	Relevance  int `xml:"relevance,attr"`
}
