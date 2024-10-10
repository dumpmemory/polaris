// Code generated by ent, DO NOT EDIT.

package blacklist

import (
	"polaris/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldLTE(FieldID, id))
}

// Notes applies equality check predicate on the "notes" field. It's identical to NotesEQ.
func Notes(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldEQ(FieldNotes, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNotIn(FieldType, vs...))
}

// NotesEQ applies the EQ predicate on the "notes" field.
func NotesEQ(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldEQ(FieldNotes, v))
}

// NotesNEQ applies the NEQ predicate on the "notes" field.
func NotesNEQ(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNEQ(FieldNotes, v))
}

// NotesIn applies the In predicate on the "notes" field.
func NotesIn(vs ...string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldIn(FieldNotes, vs...))
}

// NotesNotIn applies the NotIn predicate on the "notes" field.
func NotesNotIn(vs ...string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNotIn(FieldNotes, vs...))
}

// NotesGT applies the GT predicate on the "notes" field.
func NotesGT(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldGT(FieldNotes, v))
}

// NotesGTE applies the GTE predicate on the "notes" field.
func NotesGTE(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldGTE(FieldNotes, v))
}

// NotesLT applies the LT predicate on the "notes" field.
func NotesLT(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldLT(FieldNotes, v))
}

// NotesLTE applies the LTE predicate on the "notes" field.
func NotesLTE(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldLTE(FieldNotes, v))
}

// NotesContains applies the Contains predicate on the "notes" field.
func NotesContains(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldContains(FieldNotes, v))
}

// NotesHasPrefix applies the HasPrefix predicate on the "notes" field.
func NotesHasPrefix(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldHasPrefix(FieldNotes, v))
}

// NotesHasSuffix applies the HasSuffix predicate on the "notes" field.
func NotesHasSuffix(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldHasSuffix(FieldNotes, v))
}

// NotesIsNil applies the IsNil predicate on the "notes" field.
func NotesIsNil() predicate.Blacklist {
	return predicate.Blacklist(sql.FieldIsNull(FieldNotes))
}

// NotesNotNil applies the NotNil predicate on the "notes" field.
func NotesNotNil() predicate.Blacklist {
	return predicate.Blacklist(sql.FieldNotNull(FieldNotes))
}

// NotesEqualFold applies the EqualFold predicate on the "notes" field.
func NotesEqualFold(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldEqualFold(FieldNotes, v))
}

// NotesContainsFold applies the ContainsFold predicate on the "notes" field.
func NotesContainsFold(v string) predicate.Blacklist {
	return predicate.Blacklist(sql.FieldContainsFold(FieldNotes, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blacklist) predicate.Blacklist {
	return predicate.Blacklist(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blacklist) predicate.Blacklist {
	return predicate.Blacklist(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blacklist) predicate.Blacklist {
	return predicate.Blacklist(sql.NotPredicates(p))
}
