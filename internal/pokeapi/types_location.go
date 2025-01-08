package pokeapi

type LocationAreaResponse struct {
    Count int
    Next *string
    Previous *string
    Results []LocationArea
}

type LocationArea struct {
    Name string
    Url  string
}

