// Code generated by ent, DO NOT EDIT.

package contractproposal

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldUpdateTime, v))
}

// ProposalID applies equality check predicate on the "proposal_id" field. It's identical to ProposalIDEQ.
func ProposalID(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldProposalID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldDescription, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldExpiresAt, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldUpdateTime, v))
}

// ProposalIDEQ applies the EQ predicate on the "proposal_id" field.
func ProposalIDEQ(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldProposalID, v))
}

// ProposalIDNEQ applies the NEQ predicate on the "proposal_id" field.
func ProposalIDNEQ(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldProposalID, v))
}

// ProposalIDIn applies the In predicate on the "proposal_id" field.
func ProposalIDIn(vs ...int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldProposalID, vs...))
}

// ProposalIDNotIn applies the NotIn predicate on the "proposal_id" field.
func ProposalIDNotIn(vs ...int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldProposalID, vs...))
}

// ProposalIDGT applies the GT predicate on the "proposal_id" field.
func ProposalIDGT(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldProposalID, v))
}

// ProposalIDGTE applies the GTE predicate on the "proposal_id" field.
func ProposalIDGTE(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldProposalID, v))
}

// ProposalIDLT applies the LT predicate on the "proposal_id" field.
func ProposalIDLT(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldProposalID, v))
}

// ProposalIDLTE applies the LTE predicate on the "proposal_id" field.
func ProposalIDLTE(v int) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldProposalID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldContainsFold(FieldDescription, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldLTE(FieldExpiresAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ContractProposal {
	return predicate.ContractProposal(sql.FieldNotIn(FieldStatus, vs...))
}

// HasContract applies the HasEdge predicate on the "contract" edge.
func HasContract() predicate.ContractProposal {
	return predicate.ContractProposal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContractTable, ContractColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContractWith applies the HasEdge predicate on the "contract" edge with a given conditions (other predicates).
func HasContractWith(preds ...predicate.Contract) predicate.ContractProposal {
	return predicate.ContractProposal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ContractInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContractTable, ContractColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ContractProposal) predicate.ContractProposal {
	return predicate.ContractProposal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ContractProposal) predicate.ContractProposal {
	return predicate.ContractProposal(func(s *sql.Selector) {
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
func Not(p predicate.ContractProposal) predicate.ContractProposal {
	return predicate.ContractProposal(func(s *sql.Selector) {
		p(s.Not())
	})
}
