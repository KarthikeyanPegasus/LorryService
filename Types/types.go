package Types

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
}

type Manager struct {
	Manager        User    `json:"manager"`
	ManagerDetails Details `json:"manager-details"`
}

type Driver struct {
	DriverName    string  `json:"driver_name"`
	Experience    int32   `json:"experience"`
	DriverDetails Details `json:"driver_details"`
}

type Details struct {
	Address   AddressInformation   `json:"address_details"`
	Contact   ContactInformation   `json:"contact_details"`
	Emergency EmergencyInformation `json:"emergency_details"`
	Personal  PersonalInformation  `json:"personal"`
}

type AddressInformation struct {
	DoorNo      string `json:"door_no"`
	AddressLine string `json:"address_line"`
	PinCode     string `json:"pin_code"`
}

type PersonalInformation struct {
	AadharNumber  string `json:"aadhar_number"`
	DriverLicence string `json:"driver_licence"` //only Driver will have the Driver Licence
}

type ContactInformation struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type EmergencyInformation struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Relation string `json:"relation"`
	Email    string `json:"email"`
}
