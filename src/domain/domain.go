package domain

type Item struct {
	title string
	post  string
}

func NewItem(title string, post string) *Item {
	return &Item{
		title: title,
		post:  post,
	}
}

func (i *Item) Title() string {
	return i.title
}

func (i *Item) Post() string {
	return i.post
}
