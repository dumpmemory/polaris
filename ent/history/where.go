// Code generated by ent, DO NOT EDIT.

package history

import (
	"polaris/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.History {
	return predicate.History(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.History {
	return predicate.History(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldID, id))
}

// MediaID applies equality check predicate on the "media_id" field. It's identical to MediaIDEQ.
func MediaID(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldMediaID, v))
}

// EpisodeID applies equality check predicate on the "episode_id" field. It's identical to EpisodeIDEQ.
func EpisodeID(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldEpisodeID, v))
}

// SourceTitle applies equality check predicate on the "source_title" field. It's identical to SourceTitleEQ.
func SourceTitle(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldSourceTitle, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.History {
	return predicate.History(sql.FieldEQ(FieldDate, v))
}

// TargetDir applies equality check predicate on the "target_dir" field. It's identical to TargetDirEQ.
func TargetDir(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldTargetDir, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldSize, v))
}

// Saved applies equality check predicate on the "saved" field. It's identical to SavedEQ.
func Saved(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldSaved, v))
}

// MediaIDEQ applies the EQ predicate on the "media_id" field.
func MediaIDEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldMediaID, v))
}

// MediaIDNEQ applies the NEQ predicate on the "media_id" field.
func MediaIDNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldMediaID, v))
}

// MediaIDIn applies the In predicate on the "media_id" field.
func MediaIDIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldMediaID, vs...))
}

// MediaIDNotIn applies the NotIn predicate on the "media_id" field.
func MediaIDNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldMediaID, vs...))
}

// MediaIDGT applies the GT predicate on the "media_id" field.
func MediaIDGT(v int) predicate.History {
	return predicate.History(sql.FieldGT(FieldMediaID, v))
}

// MediaIDGTE applies the GTE predicate on the "media_id" field.
func MediaIDGTE(v int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldMediaID, v))
}

// MediaIDLT applies the LT predicate on the "media_id" field.
func MediaIDLT(v int) predicate.History {
	return predicate.History(sql.FieldLT(FieldMediaID, v))
}

// MediaIDLTE applies the LTE predicate on the "media_id" field.
func MediaIDLTE(v int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldMediaID, v))
}

// EpisodeIDEQ applies the EQ predicate on the "episode_id" field.
func EpisodeIDEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldEpisodeID, v))
}

// EpisodeIDNEQ applies the NEQ predicate on the "episode_id" field.
func EpisodeIDNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldEpisodeID, v))
}

// EpisodeIDIn applies the In predicate on the "episode_id" field.
func EpisodeIDIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldEpisodeID, vs...))
}

// EpisodeIDNotIn applies the NotIn predicate on the "episode_id" field.
func EpisodeIDNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldEpisodeID, vs...))
}

// EpisodeIDGT applies the GT predicate on the "episode_id" field.
func EpisodeIDGT(v int) predicate.History {
	return predicate.History(sql.FieldGT(FieldEpisodeID, v))
}

// EpisodeIDGTE applies the GTE predicate on the "episode_id" field.
func EpisodeIDGTE(v int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldEpisodeID, v))
}

// EpisodeIDLT applies the LT predicate on the "episode_id" field.
func EpisodeIDLT(v int) predicate.History {
	return predicate.History(sql.FieldLT(FieldEpisodeID, v))
}

// EpisodeIDLTE applies the LTE predicate on the "episode_id" field.
func EpisodeIDLTE(v int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldEpisodeID, v))
}

// SourceTitleEQ applies the EQ predicate on the "source_title" field.
func SourceTitleEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldSourceTitle, v))
}

// SourceTitleNEQ applies the NEQ predicate on the "source_title" field.
func SourceTitleNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldSourceTitle, v))
}

// SourceTitleIn applies the In predicate on the "source_title" field.
func SourceTitleIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldSourceTitle, vs...))
}

// SourceTitleNotIn applies the NotIn predicate on the "source_title" field.
func SourceTitleNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldSourceTitle, vs...))
}

// SourceTitleGT applies the GT predicate on the "source_title" field.
func SourceTitleGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldSourceTitle, v))
}

// SourceTitleGTE applies the GTE predicate on the "source_title" field.
func SourceTitleGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldSourceTitle, v))
}

// SourceTitleLT applies the LT predicate on the "source_title" field.
func SourceTitleLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldSourceTitle, v))
}

// SourceTitleLTE applies the LTE predicate on the "source_title" field.
func SourceTitleLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldSourceTitle, v))
}

// SourceTitleContains applies the Contains predicate on the "source_title" field.
func SourceTitleContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldSourceTitle, v))
}

// SourceTitleHasPrefix applies the HasPrefix predicate on the "source_title" field.
func SourceTitleHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldSourceTitle, v))
}

// SourceTitleHasSuffix applies the HasSuffix predicate on the "source_title" field.
func SourceTitleHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldSourceTitle, v))
}

// SourceTitleEqualFold applies the EqualFold predicate on the "source_title" field.
func SourceTitleEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldSourceTitle, v))
}

// SourceTitleContainsFold applies the ContainsFold predicate on the "source_title" field.
func SourceTitleContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldSourceTitle, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.History {
	return predicate.History(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.History {
	return predicate.History(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.History {
	return predicate.History(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.History {
	return predicate.History(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.History {
	return predicate.History(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.History {
	return predicate.History(sql.FieldLTE(FieldDate, v))
}

// TargetDirEQ applies the EQ predicate on the "target_dir" field.
func TargetDirEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldTargetDir, v))
}

// TargetDirNEQ applies the NEQ predicate on the "target_dir" field.
func TargetDirNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldTargetDir, v))
}

// TargetDirIn applies the In predicate on the "target_dir" field.
func TargetDirIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldTargetDir, vs...))
}

// TargetDirNotIn applies the NotIn predicate on the "target_dir" field.
func TargetDirNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldTargetDir, vs...))
}

// TargetDirGT applies the GT predicate on the "target_dir" field.
func TargetDirGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldTargetDir, v))
}

// TargetDirGTE applies the GTE predicate on the "target_dir" field.
func TargetDirGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldTargetDir, v))
}

// TargetDirLT applies the LT predicate on the "target_dir" field.
func TargetDirLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldTargetDir, v))
}

// TargetDirLTE applies the LTE predicate on the "target_dir" field.
func TargetDirLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldTargetDir, v))
}

// TargetDirContains applies the Contains predicate on the "target_dir" field.
func TargetDirContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldTargetDir, v))
}

// TargetDirHasPrefix applies the HasPrefix predicate on the "target_dir" field.
func TargetDirHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldTargetDir, v))
}

// TargetDirHasSuffix applies the HasSuffix predicate on the "target_dir" field.
func TargetDirHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldTargetDir, v))
}

// TargetDirEqualFold applies the EqualFold predicate on the "target_dir" field.
func TargetDirEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldTargetDir, v))
}

// TargetDirContainsFold applies the ContainsFold predicate on the "target_dir" field.
func TargetDirContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldTargetDir, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.History {
	return predicate.History(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.History {
	return predicate.History(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.History {
	return predicate.History(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.History {
	return predicate.History(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.History {
	return predicate.History(sql.FieldLTE(FieldSize, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.History {
	return predicate.History(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.History {
	return predicate.History(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldStatus, vs...))
}

// SavedEQ applies the EQ predicate on the "saved" field.
func SavedEQ(v string) predicate.History {
	return predicate.History(sql.FieldEQ(FieldSaved, v))
}

// SavedNEQ applies the NEQ predicate on the "saved" field.
func SavedNEQ(v string) predicate.History {
	return predicate.History(sql.FieldNEQ(FieldSaved, v))
}

// SavedIn applies the In predicate on the "saved" field.
func SavedIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldIn(FieldSaved, vs...))
}

// SavedNotIn applies the NotIn predicate on the "saved" field.
func SavedNotIn(vs ...string) predicate.History {
	return predicate.History(sql.FieldNotIn(FieldSaved, vs...))
}

// SavedGT applies the GT predicate on the "saved" field.
func SavedGT(v string) predicate.History {
	return predicate.History(sql.FieldGT(FieldSaved, v))
}

// SavedGTE applies the GTE predicate on the "saved" field.
func SavedGTE(v string) predicate.History {
	return predicate.History(sql.FieldGTE(FieldSaved, v))
}

// SavedLT applies the LT predicate on the "saved" field.
func SavedLT(v string) predicate.History {
	return predicate.History(sql.FieldLT(FieldSaved, v))
}

// SavedLTE applies the LTE predicate on the "saved" field.
func SavedLTE(v string) predicate.History {
	return predicate.History(sql.FieldLTE(FieldSaved, v))
}

// SavedContains applies the Contains predicate on the "saved" field.
func SavedContains(v string) predicate.History {
	return predicate.History(sql.FieldContains(FieldSaved, v))
}

// SavedHasPrefix applies the HasPrefix predicate on the "saved" field.
func SavedHasPrefix(v string) predicate.History {
	return predicate.History(sql.FieldHasPrefix(FieldSaved, v))
}

// SavedHasSuffix applies the HasSuffix predicate on the "saved" field.
func SavedHasSuffix(v string) predicate.History {
	return predicate.History(sql.FieldHasSuffix(FieldSaved, v))
}

// SavedIsNil applies the IsNil predicate on the "saved" field.
func SavedIsNil() predicate.History {
	return predicate.History(sql.FieldIsNull(FieldSaved))
}

// SavedNotNil applies the NotNil predicate on the "saved" field.
func SavedNotNil() predicate.History {
	return predicate.History(sql.FieldNotNull(FieldSaved))
}

// SavedEqualFold applies the EqualFold predicate on the "saved" field.
func SavedEqualFold(v string) predicate.History {
	return predicate.History(sql.FieldEqualFold(FieldSaved, v))
}

// SavedContainsFold applies the ContainsFold predicate on the "saved" field.
func SavedContainsFold(v string) predicate.History {
	return predicate.History(sql.FieldContainsFold(FieldSaved, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.History) predicate.History {
	return predicate.History(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.History) predicate.History {
	return predicate.History(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.History) predicate.History {
	return predicate.History(sql.NotPredicates(p))
}
