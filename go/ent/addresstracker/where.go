// Code generated by ent, DO NOT EDIT.

package addresstracker

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldUpdateTime, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldAddress, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLTE(FieldUpdateTime, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.AddressTracker {
	return predicate.AddressTracker(sql.FieldContainsFold(FieldAddress, v))
}

// HasChain applies the HasEdge predicate on the "chain" edge.
func HasChain() predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChainTable, ChainColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChainWith applies the HasEdge predicate on the "chain" edge with a given conditions (other predicates).
func HasChainWith(preds ...predicate.Chain) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChainInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChainTable, ChainColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDiscordChannel applies the HasEdge predicate on the "discord_channel" edge.
func HasDiscordChannel() predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DiscordChannelTable, DiscordChannelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiscordChannelWith applies the HasEdge predicate on the "discord_channel" edge with a given conditions (other predicates).
func HasDiscordChannelWith(preds ...predicate.DiscordChannel) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiscordChannelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DiscordChannelTable, DiscordChannelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTelegramChat applies the HasEdge predicate on the "telegram_chat" edge.
func HasTelegramChat() predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TelegramChatTable, TelegramChatColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTelegramChatWith applies the HasEdge predicate on the "telegram_chat" edge with a given conditions (other predicates).
func HasTelegramChatWith(preds ...predicate.TelegramChat) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TelegramChatInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TelegramChatTable, TelegramChatColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChainProposals applies the HasEdge predicate on the "chain_proposals" edge.
func HasChainProposals() predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChainProposalsTable, ChainProposalsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChainProposalsWith applies the HasEdge predicate on the "chain_proposals" edge with a given conditions (other predicates).
func HasChainProposalsWith(preds ...predicate.ChainProposal) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChainProposalsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChainProposalsTable, ChainProposalsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AddressTracker) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AddressTracker) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
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
func Not(p predicate.AddressTracker) predicate.AddressTracker {
	return predicate.AddressTracker(func(s *sql.Selector) {
		p(s.Not())
	})
}
