package viewmodel

type IndexViewModel struct {
	Title        string
	ErrorMessage string
}

func NewIndexViewModel(title string, err error) (vm IndexViewModel) {
	vm = IndexViewModel{
		Title: title,
	}
	if err != nil {
		vm.ErrorMessage = "Echec de chargment des données. Merci de reessayer"
	}
	return vm
}
