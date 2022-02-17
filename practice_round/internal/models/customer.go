package models

type Customer struct {
	Likes    []string
	Dislikes []string
}

type CustomerDataset struct {
	Customers []Customer
	Count     int
}
