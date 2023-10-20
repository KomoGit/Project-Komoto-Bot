package main

type Category struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Company struct {
	Name     string `json:"name"`
	Industry string `json:"industry"`
}

type Job struct {
	ID          int      `json:"id"`
	CompanyID   int      `json:"employerId"`
	CategoryID  int      `json:"categoryId"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	ExpDate     string   `json:"expirationDate"`
	Cat         Category `json:"jobCategory"`
	Employer    Company  `json:"employer"`
}
