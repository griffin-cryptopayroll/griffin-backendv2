// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CryptoCurrencyColumns holds the columns for the "crypto_currency" table.
	CryptoCurrencyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "ticker", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "source", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "crypto_prc_source_price_of", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// CryptoCurrencyTable holds the schema information for the "crypto_currency" table.
	CryptoCurrencyTable = &schema.Table{
		Name:       "crypto_currency",
		Columns:    CryptoCurrencyColumns,
		PrimaryKey: []*schema.Column{CryptoCurrencyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "crypto_currency_crypto_prc_source_price_of",
				Columns:    []*schema.Column{CryptoCurrencyColumns[3]},
				RefColumns: []*schema.Column{CryptoPrcSourceColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CryptoPrcSourceColumns holds the columns for the "crypto_prc_source" table.
	CryptoPrcSourceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
	}
	// CryptoPrcSourceTable holds the schema information for the "crypto_prc_source" table.
	CryptoPrcSourceTable = &schema.Table{
		Name:       "crypto_prc_source",
		Columns:    CryptoPrcSourceColumns,
		PrimaryKey: []*schema.Column{CryptoPrcSourceColumns[0]},
	}
	// EmployeeColumns holds the columns for the "employee" table.
	EmployeeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "gid", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "employer_gid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "position", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "wallet", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "payroll", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "FLOAT"}},
		{Name: "currency", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "payday", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "employ", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "email", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "work_start", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "work_ends", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "updated_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "crypto_currency_employee_paid", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "employer_user_info_work_under", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "employ_type_employee_type_to", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// EmployeeTable holds the schema information for the "employee" table.
	EmployeeTable = &schema.Table{
		Name:       "employee",
		Columns:    EmployeeColumns,
		PrimaryKey: []*schema.Column{EmployeeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "employee_crypto_currency_employee_paid",
				Columns:    []*schema.Column{EmployeeColumns[17]},
				RefColumns: []*schema.Column{CryptoCurrencyColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "employee_employer_user_info_work_under",
				Columns:    []*schema.Column{EmployeeColumns[18]},
				RefColumns: []*schema.Column{EmployerUserInfoColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "employee_employ_type_employee_type_to",
				Columns:    []*schema.Column{EmployeeColumns[19]},
				RefColumns: []*schema.Column{EmployTypeColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmployerUserInfoColumns holds the columns for the "employer_user_info" table.
	EmployerUserInfoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "username", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "gid", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "corp_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "corp_email", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "wallet", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "updated_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
	}
	// EmployerUserInfoTable holds the schema information for the "employer_user_info" table.
	EmployerUserInfoTable = &schema.Table{
		Name:       "employer_user_info",
		Columns:    EmployerUserInfoColumns,
		PrimaryKey: []*schema.Column{EmployerUserInfoColumns[0]},
	}
	// EmployTypeColumns holds the columns for the "employ_type" table.
	EmployTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "is_permanent", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "pay_freq", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
	}
	// EmployTypeTable holds the schema information for the "employ_type" table.
	EmployTypeTable = &schema.Table{
		Name:       "employ_type",
		Columns:    EmployTypeColumns,
		PrimaryKey: []*schema.Column{EmployTypeColumns[0]},
	}
	// PaymentHistoryColumns holds the columns for the "payment_history" table.
	PaymentHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "employee_gid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(200)"}},
		{Name: "employee_payment_history", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// PaymentHistoryTable holds the schema information for the "payment_history" table.
	PaymentHistoryTable = &schema.Table{
		Name:       "payment_history",
		Columns:    PaymentHistoryColumns,
		PrimaryKey: []*schema.Column{PaymentHistoryColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_history_employee_payment_history",
				Columns:    []*schema.Column{PaymentHistoryColumns[4]},
				RefColumns: []*schema.Column{EmployeeColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CryptoCurrencyTable,
		CryptoPrcSourceTable,
		EmployeeTable,
		EmployerUserInfoTable,
		EmployTypeTable,
		PaymentHistoryTable,
	}
)

func init() {
	CryptoCurrencyTable.ForeignKeys[0].RefTable = CryptoPrcSourceTable
	CryptoCurrencyTable.Annotation = &entsql.Annotation{
		Table: "crypto_currency",
	}
	CryptoPrcSourceTable.Annotation = &entsql.Annotation{
		Table: "crypto_prc_source",
	}
	EmployeeTable.ForeignKeys[0].RefTable = CryptoCurrencyTable
	EmployeeTable.ForeignKeys[1].RefTable = EmployerUserInfoTable
	EmployeeTable.ForeignKeys[2].RefTable = EmployTypeTable
	EmployeeTable.Annotation = &entsql.Annotation{
		Table: "employee",
	}
	EmployerUserInfoTable.Annotation = &entsql.Annotation{
		Table: "employer_user_info",
	}
	EmployTypeTable.Annotation = &entsql.Annotation{
		Table: "employ_type",
	}
	PaymentHistoryTable.ForeignKeys[0].RefTable = EmployeeTable
	PaymentHistoryTable.Annotation = &entsql.Annotation{
		Table: "payment_history",
	}
}
