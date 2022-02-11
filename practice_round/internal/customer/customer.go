package customer

type Customer struct {
	Likes    []string
	Dislikes []string
}

type CustomerDataset struct {
	Customers []Customer
	Count     int
}

type PerfectPizzaDataset struct {
	Ingredients []string
}
