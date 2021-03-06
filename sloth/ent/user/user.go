// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID       = "id"       // FieldUsername holds the string denoting the username vertex property in the database.
	FieldUsername = "username" // FieldPassword holds the string denoting the password vertex property in the database.
	FieldPassword = "password"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the user in the database.
	Table = "users"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "activities"
	// OwnerInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	OwnerInverseTable = "activities"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_owner"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
