package main

// ContentRef models the data as it comes from the metadata publishing event
type ContentRef struct {
	TagHolder      tags `xml:"tags"`
	PrimarySection term `xml:"primarySection"`
}

type tags struct {
	Tags []tag `xml:"tag"`
}

type tag struct {
	Term     term     `xml:"term"`
	TagScore tagScore `xml:"score"`
}

type term struct {
	CanonicalName string `xml:"canonicalName"`
	Taxonomy      string `xml:"taxonomy,attr"`
	ID            string `xml:"id,attr"`
}

type tagScore struct {
	Confidence int `xml:"confidence,attr"`
	Relevance  int `xml:"relevance,attr"`
}
