package viewmodel

// Home provides the data used when rendering the
// home content view (template)
type Home struct {
	PageTitle string
}

// NewHome creates a new viewmodel for home page used to render the home view
func NewHome() Home {
	return Home{PageTitle: "Earthquake portal - main page"}
}
