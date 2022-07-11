package types

// Enumeration is a struct which
// describes an enumeration type for database.
type Enumeration struct {
	ID        int                // ID is a primary key of enumeration.
	Workspace int                // Workspace is an ID of workspace where it was initialized.
	Title     string             // Title is a title of enumeration.
	Values    []EnumerationValue // Values is an array of enumeration values.
}

// EnumerationModel is a struct which
// describes an enumeration type while getting it via request.
type EnumerationModel struct {
	ID            int    `json:"id" form:"id"`                       // ID is a primary key of enumeration in database.
	SelectedValue string `json:"selectedValue" form:"selectedValue"` // SelectedValue is a value of enumeration which is selected.
}

// EnumerationValue is a struct which
// describes an enumeration value.
type EnumerationValue struct {
	Value string `json:"value" form:"value"` // Value is a value of enumeration.
	Color string `json:"color" form:"color"` // Color is a color of enumeration value.
}

// AddValue adds a value to enumeration.
func (enum *Enumeration) AddValue(value EnumerationValue) {
	enum.Values = append(enum.Values, value)
}

// Equals is a method which that compares 2 enumeration values.
func (val *EnumerationValue) Equals(value *EnumerationValue) bool {
	return val.Value == value.Value && val.Color == value.Color
}

// RemoveValue removes a value from enumeration.
func (enum *Enumeration) RemoveValue(value *EnumerationValue) {
	for i := 0; i < len(enum.Values); i++ {
		if enum.Values[i].Equals(value) {
			enum.Values = append(enum.Values[:i], enum.Values[i+1:]...)
			break
		}
	}
}
