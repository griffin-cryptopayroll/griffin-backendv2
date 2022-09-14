// Code generated by ent, DO NOT EDIT.

package employ_type

import (
	"griffin-dao/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IsPermanent applies equality check predicate on the "is_permanent" field. It's identical to IsPermanentEQ.
func IsPermanent(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPermanent), v))
	})
}

// ContractStart applies equality check predicate on the "contract_start" field. It's identical to ContractStartEQ.
func ContractStart(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractStart), v))
	})
}

// ContractPeriod applies equality check predicate on the "contract_period" field. It's identical to ContractPeriodEQ.
func ContractPeriod(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractPeriod), v))
	})
}

// IsPermanentEQ applies the EQ predicate on the "is_permanent" field.
func IsPermanentEQ(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentNEQ applies the NEQ predicate on the "is_permanent" field.
func IsPermanentNEQ(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentIn applies the In predicate on the "is_permanent" field.
func IsPermanentIn(vs ...string) predicate.EMPLOY_TYPE {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIsPermanent), v...))
	})
}

// IsPermanentNotIn applies the NotIn predicate on the "is_permanent" field.
func IsPermanentNotIn(vs ...string) predicate.EMPLOY_TYPE {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIsPermanent), v...))
	})
}

// IsPermanentGT applies the GT predicate on the "is_permanent" field.
func IsPermanentGT(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentGTE applies the GTE predicate on the "is_permanent" field.
func IsPermanentGTE(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentLT applies the LT predicate on the "is_permanent" field.
func IsPermanentLT(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentLTE applies the LTE predicate on the "is_permanent" field.
func IsPermanentLTE(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentContains applies the Contains predicate on the "is_permanent" field.
func IsPermanentContains(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentHasPrefix applies the HasPrefix predicate on the "is_permanent" field.
func IsPermanentHasPrefix(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentHasSuffix applies the HasSuffix predicate on the "is_permanent" field.
func IsPermanentHasSuffix(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentEqualFold applies the EqualFold predicate on the "is_permanent" field.
func IsPermanentEqualFold(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIsPermanent), v))
	})
}

// IsPermanentContainsFold applies the ContainsFold predicate on the "is_permanent" field.
func IsPermanentContainsFold(v string) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIsPermanent), v))
	})
}

// ContractStartEQ applies the EQ predicate on the "contract_start" field.
func ContractStartEQ(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractStart), v))
	})
}

// ContractStartNEQ applies the NEQ predicate on the "contract_start" field.
func ContractStartNEQ(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContractStart), v))
	})
}

// ContractStartIn applies the In predicate on the "contract_start" field.
func ContractStartIn(vs ...time.Time) predicate.EMPLOY_TYPE {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContractStart), v...))
	})
}

// ContractStartNotIn applies the NotIn predicate on the "contract_start" field.
func ContractStartNotIn(vs ...time.Time) predicate.EMPLOY_TYPE {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContractStart), v...))
	})
}

// ContractStartGT applies the GT predicate on the "contract_start" field.
func ContractStartGT(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContractStart), v))
	})
}

// ContractStartGTE applies the GTE predicate on the "contract_start" field.
func ContractStartGTE(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContractStart), v))
	})
}

// ContractStartLT applies the LT predicate on the "contract_start" field.
func ContractStartLT(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContractStart), v))
	})
}

// ContractStartLTE applies the LTE predicate on the "contract_start" field.
func ContractStartLTE(v time.Time) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContractStart), v))
	})
}

// ContractPeriodEQ applies the EQ predicate on the "contract_period" field.
func ContractPeriodEQ(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractPeriod), v))
	})
}

// ContractPeriodNEQ applies the NEQ predicate on the "contract_period" field.
func ContractPeriodNEQ(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContractPeriod), v))
	})
}

// ContractPeriodIn applies the In predicate on the "contract_period" field.
func ContractPeriodIn(vs ...int) predicate.EMPLOY_TYPE {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContractPeriod), v...))
	})
}

// ContractPeriodNotIn applies the NotIn predicate on the "contract_period" field.
func ContractPeriodNotIn(vs ...int) predicate.EMPLOY_TYPE {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContractPeriod), v...))
	})
}

// ContractPeriodGT applies the GT predicate on the "contract_period" field.
func ContractPeriodGT(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContractPeriod), v))
	})
}

// ContractPeriodGTE applies the GTE predicate on the "contract_period" field.
func ContractPeriodGTE(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContractPeriod), v))
	})
}

// ContractPeriodLT applies the LT predicate on the "contract_period" field.
func ContractPeriodLT(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContractPeriod), v))
	})
}

// ContractPeriodLTE applies the LTE predicate on the "contract_period" field.
func ContractPeriodLTE(v int) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContractPeriod), v))
	})
}

// HasEmployeeTypeTo applies the HasEdge predicate on the "employee_type_to" edge.
func HasEmployeeTypeTo() predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTypeToTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmployeeTypeToTable, EmployeeTypeToColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeTypeToWith applies the HasEdge predicate on the "employee_type_to" edge with a given conditions (other predicates).
func HasEmployeeTypeToWith(preds ...predicate.EMPLOYEE) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeTypeToInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmployeeTypeToTable, EmployeeTypeToColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EMPLOY_TYPE) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EMPLOY_TYPE) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EMPLOY_TYPE) predicate.EMPLOY_TYPE {
	return predicate.EMPLOY_TYPE(func(s *sql.Selector) {
		p(s.Not())
	})
}
