package newsfeed

type Adder interface {
	Add(item *Item)
}

type Getter interface {
	GetAll() Items
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Items []*Item

type Repo struct {
	Items Items
}

func New() *Repo {
	return &Repo{
		Items: Items{},
	}
}

func (r *Repo) Add(item *Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() Items {
	return r.Items
}
