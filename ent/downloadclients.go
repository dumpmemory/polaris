// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"polaris/ent/downloadclients"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// DownloadClients is the model entity for the DownloadClients schema.
type DownloadClients struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Enable holds the value of the "enable" field.
	Enable bool `json:"enable,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Implementation holds the value of the "implementation" field.
	Implementation string `json:"implementation,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings string `json:"settings,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority string `json:"priority,omitempty"`
	// RemoveCompletedDownloads holds the value of the "remove_completed_downloads" field.
	RemoveCompletedDownloads bool `json:"remove_completed_downloads,omitempty"`
	// RemoveFailedDownloads holds the value of the "remove_failed_downloads" field.
	RemoveFailedDownloads bool `json:"remove_failed_downloads,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags         string `json:"tags,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DownloadClients) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case downloadclients.FieldEnable, downloadclients.FieldRemoveCompletedDownloads, downloadclients.FieldRemoveFailedDownloads:
			values[i] = new(sql.NullBool)
		case downloadclients.FieldID:
			values[i] = new(sql.NullInt64)
		case downloadclients.FieldName, downloadclients.FieldImplementation, downloadclients.FieldURL, downloadclients.FieldUser, downloadclients.FieldPassword, downloadclients.FieldSettings, downloadclients.FieldPriority, downloadclients.FieldTags:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DownloadClients fields.
func (dc *DownloadClients) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case downloadclients.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dc.ID = int(value.Int64)
		case downloadclients.FieldEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable", values[i])
			} else if value.Valid {
				dc.Enable = value.Bool
			}
		case downloadclients.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dc.Name = value.String
			}
		case downloadclients.FieldImplementation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field implementation", values[i])
			} else if value.Valid {
				dc.Implementation = value.String
			}
		case downloadclients.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				dc.URL = value.String
			}
		case downloadclients.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				dc.User = value.String
			}
		case downloadclients.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				dc.Password = value.String
			}
		case downloadclients.FieldSettings:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value.Valid {
				dc.Settings = value.String
			}
		case downloadclients.FieldPriority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				dc.Priority = value.String
			}
		case downloadclients.FieldRemoveCompletedDownloads:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field remove_completed_downloads", values[i])
			} else if value.Valid {
				dc.RemoveCompletedDownloads = value.Bool
			}
		case downloadclients.FieldRemoveFailedDownloads:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field remove_failed_downloads", values[i])
			} else if value.Valid {
				dc.RemoveFailedDownloads = value.Bool
			}
		case downloadclients.FieldTags:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value.Valid {
				dc.Tags = value.String
			}
		default:
			dc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DownloadClients.
// This includes values selected through modifiers, order, etc.
func (dc *DownloadClients) Value(name string) (ent.Value, error) {
	return dc.selectValues.Get(name)
}

// Update returns a builder for updating this DownloadClients.
// Note that you need to call DownloadClients.Unwrap() before calling this method if this DownloadClients
// was returned from a transaction, and the transaction was committed or rolled back.
func (dc *DownloadClients) Update() *DownloadClientsUpdateOne {
	return NewDownloadClientsClient(dc.config).UpdateOne(dc)
}

// Unwrap unwraps the DownloadClients entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dc *DownloadClients) Unwrap() *DownloadClients {
	_tx, ok := dc.config.driver.(*txDriver)
	if !ok {
		panic("ent: DownloadClients is not a transactional entity")
	}
	dc.config.driver = _tx.drv
	return dc
}

// String implements the fmt.Stringer.
func (dc *DownloadClients) String() string {
	var builder strings.Builder
	builder.WriteString("DownloadClients(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dc.ID))
	builder.WriteString("enable=")
	builder.WriteString(fmt.Sprintf("%v", dc.Enable))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(dc.Name)
	builder.WriteString(", ")
	builder.WriteString("implementation=")
	builder.WriteString(dc.Implementation)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(dc.URL)
	builder.WriteString(", ")
	builder.WriteString("user=")
	builder.WriteString(dc.User)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(dc.Password)
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(dc.Settings)
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(dc.Priority)
	builder.WriteString(", ")
	builder.WriteString("remove_completed_downloads=")
	builder.WriteString(fmt.Sprintf("%v", dc.RemoveCompletedDownloads))
	builder.WriteString(", ")
	builder.WriteString("remove_failed_downloads=")
	builder.WriteString(fmt.Sprintf("%v", dc.RemoveFailedDownloads))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(dc.Tags)
	builder.WriteByte(')')
	return builder.String()
}

// DownloadClientsSlice is a parsable slice of DownloadClients.
type DownloadClientsSlice []*DownloadClients
