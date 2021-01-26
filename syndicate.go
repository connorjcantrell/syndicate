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
	ManufacturerID uuid.UUID `db:"manufacturer_id"`
	Name           string    `db:"name"`
	SeriesNumber   string    `db:"series_number"`
	Description    string    `db:"description"`
}

type SeriesStore interface {
	Series(id uuid.UUID) (Series, error)
	SeriesByManufacturer(manufacturerID uuid.UUID) ([]Series, error)
	CreateSeries(i *Series) error
	UpdateSeries(i *Series) error
	DeleteSeries(id uuid.UUID) error
}

type Model struct {
	ID             uuid.UUID `db:"id"`
	ManufacturerID uuid.UUID `db:"manufacturer_id"`
	SeriesID       uuid.UUID `db:"series_id"`
	ModelNumber    string    `db:"number"`
	Image          []byte    `db:"image"`
}

type ModelStore interface {
	Model(id uuid.UUID) (Model, error)
	ModelsByManufacturer(manufacturerID uuid.UUID) ([]Model, error)
	ModelsBySeries(seriesID uuid.UUID) ([]Model, error)
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
	CreateDistributor(d *Distributor) error
	UpdateDistributor(d *Distributor) error
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
	Branches() ([]Branch, error)
	BranchesByDistributor(id uuid.UUID) ([]Branch, error)
	CreateBranch(b *Branch) error
	UpdateBranch(b *Branch) error
	DeleteBranch(id uuid.UUID) error
}

type Store interface {
	ManufacturerStore
	SeriesStore
	ModelStore
	DistributorStore
	BranchStore
}

// type Contact struct {
// 	ID         uuid.UUID `db:"id"`
// 	Department string    `db:"department_id"`
// 	FirstName  string    `db:"first_name"`
// 	LastName   string    `db:"last_name"`
// 	Phone      string    `db:"phone"`
// 	Email      string    `db:"email"`
// 	Active     bool      `db:"active"`
// }

// type ContactStore interface {
// 	Contact(id uuid.UUID) (Contact, error)
// 	ContactsByBranch(id uuid.UUID) ([]Contact, error)
// 	ContactsByManufacturer(id uuid.UUID) ([]Contact, error)
// 	ContactsByDistributor(id uuid.UUID) ([]Contact, error)
// 	CreateContact(c *Contact) error
// 	UpdateContact(c *Contact) error
// 	DeleteContact(id uuid.UUID) error
// }

// type Address struct {
// 	ID       uuid.UUID `db:"id"`
// 	Address1 string    `db:"address_1"`
// 	Address2 string    `db:"address_2"`
// 	City     string    `db:"city"`
// 	State    string    `db:"state"`
// 	Zipcode  string    `db:"zipcode"`
// 	Geocode  string    `db:"geocode"`
// }

// type AddressStore interface {
// 	Address(id uuid.UUID) (Address, error)
// 	Addresses() ([]Address, error)
// 	CreateAddress(a *Address) error
// 	UpdateAddress(a *Address) error
// 	DeleteAddress(id uuid.UUID) error
// }
