package entities

// list of entities brands
type Brands struct {
	Brand []Brand
}

// entity brand
type Brand struct {
	Id      string
	UrlCode string
	Title   string
	Image   string
	Logo    string
	Visible string
}
