package version1

type GuidePageV1 struct {
	Title   map[string]string `json:"title"`
	Content map[string]string `json:"content"`
	MoreUrl string            `json:"more_url"`
	Color   string            `json:"color"`
	PicId   string            `json:"pic_id"`
	PicUri  string            `json:"pic_uri"`
}

func NewGuidePageV1(title, content map[string]string, moreUrl, color, picId, picUri string) *GuidePageV1 {
	c := &GuidePageV1{}

	c.Title = title
	c.Content = content
	c.MoreUrl = moreUrl
	c.Color = color
	c.PicId = picId
	c.PicUri = picUri

	return c
}
