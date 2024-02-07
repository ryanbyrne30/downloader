package main

type scanResponse struct {
	File string 	`json:"file"`
	Ok bool 		`json:"ok"`
	Status string	`json:"status"`
}
