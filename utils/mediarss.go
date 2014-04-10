package utils

// import (
// 	"encoding/xml"
// 	"github.com/gorilla/feeds"
// 	"time"
// )

// type RssItunesFeed struct {
// 	XMLName        xml.Name `xml:"channel"`
// 	Title          string   `xml:"title"` // required
// 	Link           *feeds.Link
// 	Description    string `xml:"description"`     // required
// 	ItunesSubtitle string `xml:"itunes:subtitle"` // required
// 	Language       string `xml:"language,omitempty"`
// 	Copyright      string `xml:"copyright,omitempty"`
// 	ManagingEditor string `xml:"managingEditor,omitempty"` // Author used
// 	WebMaster      string `xml:"webMaster,omitempty"`
// 	PubDate        string `xml:"pubDate,omitempty"`       // created or updated
// 	LastBuildDate  string `xml:"lastBuildDate,omitempty"` // updated used
// 	Category       string `xml:"category,omitempty"`
// 	Generator      string `xml:"generator,omitempty"`
// 	Docs           string `xml:"docs,omitempty"`
// 	Cloud          string `xml:"cloud,omitempty"`
// 	Ttl            int    `xml:"ttl,omitempty"`
// 	Rating         string `xml:"rating,omitempty"`
// 	SkipHours      string `xml:"skipHours,omitempty"`
// 	SkipDays       string `xml:"skipDays,omitempty"`
// 	Author         *feeds.Author
// 	Image          *Image
// 	Created        time.Time
// 	Items          []*RssMediaItem
// }

// type RssMediaItem struct {
// 	// Either the title or the description are required. The title of the item.
// 	Title string `xml:"title,omitempty"`

// 	// Optional. The URL of the item
// 	Link *feeds.Link

// 	// Either the title or the description are required. The item description.
// 	Description string `xml:"description,omitempty"`

// 	// Optional. The authors email address
// 	Author *feeds.Author

// 	// Optional. The items hierarchical categorizations.
// 	Categories []Category `xml:"category"`

// 	// Optional. The URL for the page containing the items comments.
// 	Comments string `xml:"comments,omitempty"`

// 	// Optional. A media object attached to the item.
// 	Enclosure *Enclosure `xml:"enclosure"`

// 	// Optional. A unique identifier for the item
// 	Guid *Guid `xml:"guid"`

// 	// Optional. Publication date of the item. See rssgo.ComposeRssDate and
// 	// rssgo.ParseRssDate
// 	PubDate string `xml:"pubDate,omitempty"`

// 	// Optional. The RSS channel the item came from.
// 	Source *Source `xml:"source"`

// 	// Not listed in the spec.
// 	Created time.Time
// 	// The content of the item.
// 	// Tagged as "content:encoded".
// 	Content string `xml:"encoded,omitempty"`

// 	// Alternate dates
// 	Date      string `xml:"date,omitempty"`
// 	Published string `xml:"published,omitempty"`

// 	Media *MediaContent `xml:"content"`
// }

// // A media object for an item
// type Enclosure struct {
// 	// Required. The enclosures URL.
// 	Url string `xml:"url,attr"`

// 	// Required. The enclosures size.
// 	Length int64 `xml:"length,attr,omitempty"`

// 	// Required. The enclosures MIME type.
// 	Type string `xml:"type,attr"`
// }

// // An RSS channel's image
// type Image struct {
// 	// Required. The URL to the GIF, JPEG, or PNG image
// 	Url string `xml:"url"`

// 	// Required. The image title (should probably match the channels title)
// 	Title string `xml:"title"`

// 	// Required. The image link (should probably match the channels link)
// 	Link string `xml:"link"`

// 	// Optional. The image width.
// 	// Note: If the element is missing from the XML this field will have a value
// 	// of 0. The field value should be treated as having a value of DefaultWidth
// 	Width int `xml:"width,omitempty"`

// 	// Optional. The image height.
// 	// Note: If the element is missing from the XML this field will have a value
// 	// of 0. The field value should be treated as having a value of DefaultHeight
// 	Height int `xml:"height,omitempty"`
// }

// type MediaContent struct {
// 	XMLBase string `xml:"http://search.yahoo.com/mrss/ content"`
// 	URL     string `xml:"url,attr"`
// 	Type    string `xml:"type,attr"`
// }

// // The RSS channel the item came from.
// type Source struct {
// 	// Required. The title of the channel where the item came from.
// 	Source string `xml:",chardata"`

// 	// Required. The URL of the channel where the item came from.
// 	Url string `xml:"url,attr"`
// }

// // A unique identifier for the item
// type Guid struct {

// 	// Required. The items GUID
// 	Guid string `xml:",chardata"`

// 	// Optional. If set to true the Guid must be a URL
// 	IsPermaLink bool `xml:"isPermaLink,attr,omitempty"`
// }

// // A hierarchical categorization type
// type Category struct {
// 	// Required. A hierarchical categorizations
// 	Category string `xml:",chardata"`

// 	// Optional. The domain URL
// 	Domain string `xml:"domain,attr,omitempty"`
// }

// // interface used by ToXML to get a object suitable for exporting XML.
// type XmlFeed interface {
// 	FeedXml() interface{}
// }

// type MediaRss struct {
// 	*XmlFeed
// }

// func ToXML(feed XmlFeed) (string, error) {
// 	x := feed.FeedXml()
// 	data, err := xml.MarshalIndent(x, "", "  ")
// 	if err != nil {
// 		return "", err
// 	}
// 	// strip empty line from default xml header
// 	s := xml.Header[:len(xml.Header)-1] + string(data)
// 	return s, nil
// }

// // creates an Rss representation of this feed
// func (f *Feed) ToRss() (string, error) {
// 	r := &MediaRss{f}
// 	return ToXML(r)
// }
