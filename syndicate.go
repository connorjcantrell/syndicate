package syndicate

import "github.com/google/uuid"

type Manufacturer struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	AddressID uuid.UUID `db:"address_id"`
}

type ManufacturerStore interface {
	Manufacturer(id uuid.UUID) (Manufacturer, error)
	Manufacturers() ([]Manufacturer, error)
	CreateManufacturer(m *Manufacturer) error
	UpdateManufacturer(m *Manufacturer) error
	DeleteManufacturer(id uuid.UUID) error
}

type Series struct {
	ID             uuid.UUID `db:"id"`
	ManufacturerID uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	SeriesNumber   string    `db:"number"`
	Description    string    `db:"description"`
}

type SeriesStore interface {
	Series(id uuid.UUID) (Series, error)
	Seriess() ([]Series, error)
	CreateSeries(m *Series) error
	UpdateSeries(m *Series) error
	DeleteSeries(id uuid.UUID) error
}

type Model struct {
	ID          uuid.UUID `db:"id"`
	SeriesID    uuid.UUID `db:"id"`
	ModelNumber string    `db:"number"`
	Image       []byte    `db:"image"`
}

type ModelStore interface {
	Model(id uuid.UUID) (Model, error)
	Models() ([]Model, error)
	CreateModel(m *Model) error
	UpdateModel(m *Model) error
	DeleteModel(id uuid.UUID) error
}

type Distributor struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Active    bool      `db:"active"`
	AddressID uuid.UUID `db:"address_id"`
}

type DistributorStore interface {
	Distributor(id uuid.UUID) (Distributor, error)
	Distributors() ([]Distributor, error)
	CreateDistributor(m *Distributor) error
	UpdateDistributor(m *Distributor) error
	DeleteDistributor(id uuid.UUID) error
}

type Branch struct {
	ID            uuid.UUID `db:"id"`
	DistributorID uuid.UUID `db:"distributor_id"`
	Active        bool      `db:"active"`
	AddressID     uuid.UUID `db:"address_id"`
}

type BranchStore interface {
	Branch(id uuid.UUID) (Branch, error)
	Branchs() ([]Branch, error)
	CreateBranch(m *Branch) error
	UpdateBranch(m *Branch) error
	DeleteBranch(id uuid.UUID) error
}

type Contact struct {
	ID         uuid.UUID `db:"id"`
	Department string    `db:"department_id"`
	FirstName  string    `db:"name"`
	LastName   string    `db:"name"`
	Phone      string    `db:"phone"`
	Email      string    `db:"email"`
	Active     bool      `db:"active"`
}

type ContactStore interface {
	Contact(id uuid.UUID) (Contact, error)
	Contacts() ([]Contact, error)
	CreateContact(m *Contact) error
	UpdateContact(m *Contact) error
	DeleteContact(id uuid.UUID) error
}

type Address struct {
	ID       uuid.UUID `db:"id"`
	Address1 string    `db:"address_1"`
	Address2 string    `db:"address_2"`
	City     string    `db:"city"`
	State    string    `db:"state"`
	Zipcode  string    `db:"zipcode"`
	Geocode  string    `db:"geocode"`
}

type AddressStore interface {
	Address(id uuid.UUID) (Address, error)
	Addresses() ([]Address, error)
	CreateAddress(m *Address) error
	UpdateAddress(m *Address) error
	DeleteAddress(id uuid.UUID) error
}
