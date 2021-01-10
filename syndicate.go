package syndicate

import "github.com/google/uuid"

type Manufacturer struct {
	ID uuid.UUID `db:"id"`
	name string	`db:"name"`
	active bool `db:"active"`
	address_id uuid.UUID `db:"address_id"`
}

type ManufacturerStore interface {
	Manufacturer(id uuid.UUID) (Manufacturer, error)
	Manufacturers() ([]Manufacturer, error)
	CreateManufacturer(m *Manufacturer) error
	UpdateManufacturer(m *Manufacturer) error
	DeleteManufacturer(id uuid.UUID) error
}

type Series struct {
	ID uuid.UUID `db:"id"`
	manufacturer_id uuid.UUID `db:"id"`
	name string `db:"name"`
	number string `db:"number"`
	description string `db:"description"`
}

type SeriesStore interface {
	Series(id uuid.UUID) (Series, error)
	Seriess() ([]Series, error)
	CreateSeries(m *Series) error
	UpdateSeries(m *Series) error
	DeleteSeries(id uuid.UUID) error
}

type Model struct {
	ID uuid.UUID `db:"id"`
	series_id uuid.UUID `db:"id"`
	number string `db:"number"`
	image []byte `db:"image"`
}

type ModelStore interface {
	Model(id uuid.UUID) (Model, error)
	Models() ([]Model, error)
	CreateModel(m *Model) error
	UpdateModel(m *Model) error
	DeleteModel(id uuid.UUID) error
}

type Distributor struct {
	ID uuid.UUID `db:"id"`
	name string	`db:"name"`
	active bool `db:"active"`
	address_id uuid.UUID `db:"address_id"`
}

type DistributorStore interface {
	Distributor(id uuid.UUID) (Distributor, error)
	Distributors() ([]Distributor, error)
	CreateDistributor(m *Distributor) error
	UpdateDistributor(m *Distributor) error
	DeleteDistributor(id uuid.UUID) error
}

type Branch struct {
	ID uuid.UUID `db:"id"`
	distibutor_id uuid.UUID `db:"distributor_id"`
	active bool `db:"active"`
	address_id uuid.UUID `db:"address_id"`
}

type BranchStore interface {
	Branch(id uuid.UUID) (Branch, error)
	Branchs() ([]Branch, error)
	CreateBranch(m *Branch) error
	UpdateBranch(m *Branch) error
	DeleteBranch(id uuid.UUID) error
}

type Contact struct {
	ID uuid.UUID `db:"id"`
	department string `db:"department_id"`
	first_name string `db:"name"`
	last_name string `db:"name"`
	phone string `db:"phone"`
	email string `db:"email"`
	active bool `db:"active"`
}

type ContactStore interface {
	Contact(id uuid.UUID) (Contact, error)
	Contacts() ([]Contact, error)
	CreateContact(m *Contact) error
	UpdateContact(m *Contact) error
	DeleteContact(id uuid.UUID) error
}

type Address struct {
	ID uuid.UUID `db:"id"`
	address_1 string `db:"address_1"`
	address_2 string `db:"address_2"`
	city string `db:"city"`
	state string `db:"state"`
	zipcode string `db:"zipcode"`
	geocode string `db:"geocode"`
}

type AddressStore interface {
	Address(id uuid.UUID) (Address, error)
	Addresses() ([]Address, error)
	CreateAddress(m *Address) error
	UpdateAddress(m *Address) error
	DeleteAddress(id uuid.UUID) error
}