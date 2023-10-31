// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Company struct {
	CompanyID string `json:"companyId"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Jobs      []*Job `json:"jobs"`
}

type Job struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CompanyID int    `json:"companyId"`
}

type NewCompany struct {
	CompanyID string `json:"companyId"`
	Name      string `json:"name"`
	City      string `json:"city"`
}

type NewJob struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type NewUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
