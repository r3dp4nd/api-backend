package dtos

type RegisterCustomer struct {
	DNI       string `json:"dni" binding:"required,min=8,max=8"`
	Name      string `json:"name" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Telephone string `json:"telephone" binding:"required,max=9,min=9"`
	Email     string `json:"email" binding:"required,email"`
	BirthDate string `json:"birthDate" binding:"required"`
	City      string `json:"city" binding:"required"`
}

type GetCustomer struct {
	DNI string
}

type UpdateCustomer struct {
	DNI       string `json:"dni"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	BirthDate string `json:"birthDate"`
	City      string `json:"city"`
}

type DeleteCustomer struct {
	DNI string
}
