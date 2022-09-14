// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CryptoCurrencYsColumns holds the columns for the "crypto_currenc_ys" table.
	CryptoCurrencYsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "ticker", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(10)"}},
		{Name: "source", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "crypto_prc_source_price_of", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// CryptoCurrencYsTable holds the schema information for the "crypto_currenc_ys" table.
	CryptoCurrencYsTable = &schema.Table{
		Name:       "crypto_currenc_ys",
		Columns:    CryptoCurrencYsColumns,
		PrimaryKey: []*schema.Column{CryptoCurrencYsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "crypto_currenc_ys_crypto_prc_sourc_es_price_of",
				Columns:    []*schema.Column{CryptoCurrencYsColumns[3]},
				RefColumns: []*schema.Column{CryptoPrcSourcEsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CryptoPrcSourcEsColumns holds the columns for the "crypto_prc_sourc_es" table.
	CryptoPrcSourcEsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
	}
	// CryptoPrcSourcEsTable holds the schema information for the "crypto_prc_sourc_es" table.
	CryptoPrcSourcEsTable = &schema.Table{
		Name:       "crypto_prc_sourc_es",
		Columns:    CryptoPrcSourcEsColumns,
		PrimaryKey: []*schema.Column{CryptoPrcSourcEsColumns[0]},
	}
	// EmployeEsColumns holds the columns for the "employe_es" table.
	EmployeEsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "gid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(6)"}},
		{Name: "employer_gid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(6)"}},
		{Name: "last_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "first_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "position", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "wallet", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "payroll", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "FLOAT"}},
		{Name: "currency", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "payday", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "employ", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "updated_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "crypto_currency_employee_paid", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "employer_user_info_work_under", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "employ_type_employee_type_to", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// EmployeEsTable holds the schema information for the "employe_es" table.
	EmployeEsTable = &schema.Table{
		Name:       "employe_es",
		Columns:    EmployeEsColumns,
		PrimaryKey: []*schema.Column{EmployeEsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "employe_es_crypto_currenc_ys_employee_paid",
				Columns:    []*schema.Column{EmployeEsColumns[15]},
				RefColumns: []*schema.Column{CryptoCurrencYsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "employe_es_employer_user_inf_os_work_under",
				Columns:    []*schema.Column{EmployeEsColumns[16]},
				RefColumns: []*schema.Column{EmployerUserInfOsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "employe_es_employ_typ_es_employee_type_to",
				Columns:    []*schema.Column{EmployeEsColumns[17]},
				RefColumns: []*schema.Column{EmployTypEsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmployerUserInfOsColumns holds the columns for the "employer_user_inf_os" table.
	EmployerUserInfOsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "username", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(20)"}},
		{Name: "password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(20)"}},
		{Name: "gid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(6)"}},
		{Name: "corp_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "corp_email", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "wallet", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "updated_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
	}
	// EmployerUserInfOsTable holds the schema information for the "employer_user_inf_os" table.
	EmployerUserInfOsTable = &schema.Table{
		Name:       "employer_user_inf_os",
		Columns:    EmployerUserInfOsColumns,
		PrimaryKey: []*schema.Column{EmployerUserInfOsColumns[0]},
	}
	// EmployTypEsColumns holds the columns for the "employ_typ_es" table.
	EmployTypEsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "is_permanent", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(5)"}},
		{Name: "contract_start", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "contract_period", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// EmployTypEsTable holds the schema information for the "employ_typ_es" table.
	EmployTypEsTable = &schema.Table{
		Name:       "employ_typ_es",
		Columns:    EmployTypEsColumns,
		PrimaryKey: []*schema.Column{EmployTypEsColumns[0]},
	}
	// PaymentHistorYsColumns holds the columns for the "payment_histor_ys" table.
	PaymentHistorYsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "employee_gid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(6)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "DATETIME"}},
		{Name: "created_by", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(45)"}},
		{Name: "employee_payment_history", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// PaymentHistorYsTable holds the schema information for the "payment_histor_ys" table.
	PaymentHistorYsTable = &schema.Table{
		Name:       "payment_histor_ys",
		Columns:    PaymentHistorYsColumns,
		PrimaryKey: []*schema.Column{PaymentHistorYsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_histor_ys_employe_es_payment_history",
				Columns:    []*schema.Column{PaymentHistorYsColumns[4]},
				RefColumns: []*schema.Column{EmployeEsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CryptoCurrencYsTable,
		CryptoPrcSourcEsTable,
		EmployeEsTable,
		EmployerUserInfOsTable,
		EmployTypEsTable,
		PaymentHistorYsTable,
	}
)

func init() {
	CryptoCurrencYsTable.ForeignKeys[0].RefTable = CryptoPrcSourcEsTable
	EmployeEsTable.ForeignKeys[0].RefTable = CryptoCurrencYsTable
	EmployeEsTable.ForeignKeys[1].RefTable = EmployerUserInfOsTable
	EmployeEsTable.ForeignKeys[2].RefTable = EmployTypEsTable
	PaymentHistorYsTable.ForeignKeys[0].RefTable = EmployeEsTable
}
