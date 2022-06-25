// Code generated by entc, DO NOT EDIT.

package czkquote

import (
	"main/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Timestamp applies equality check predicate on the "Timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.CZKQuote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CZKQuote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.CZKQuote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CZKQuote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// TimestampEQ applies the EQ predicate on the "Timestamp" field.
func TimestampEQ(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "Timestamp" field.
func TimestampNEQ(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "Timestamp" field.
func TimestampIn(vs ...time.Time) predicate.CZKQuote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CZKQuote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "Timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.CZKQuote {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CZKQuote(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "Timestamp" field.
func TimestampGT(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "Timestamp" field.
func TimestampGTE(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "Timestamp" field.
func TimestampLT(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "Timestamp" field.
func TimestampLTE(v time.Time) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CZKQuote) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CZKQuote) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
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
func Not(p predicate.CZKQuote) predicate.CZKQuote {
	return predicate.CZKQuote(func(s *sql.Selector) {
		p(s.Not())
	})
}
