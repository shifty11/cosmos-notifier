// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/contract"
	"github.com/shifty11/cosmos-notifier/ent/contractproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/schema"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
	"github.com/shifty11/cosmos-notifier/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chainMixin := schema.Chain{}.Mixin()
	chainMixinFields0 := chainMixin[0].Fields()
	_ = chainMixinFields0
	chainFields := schema.Chain{}.Fields()
	_ = chainFields
	// chainDescCreateTime is the schema descriptor for create_time field.
	chainDescCreateTime := chainMixinFields0[0].Descriptor()
	// chain.DefaultCreateTime holds the default value on creation for the create_time field.
	chain.DefaultCreateTime = chainDescCreateTime.Default.(func() time.Time)
	// chainDescUpdateTime is the schema descriptor for update_time field.
	chainDescUpdateTime := chainMixinFields0[1].Descriptor()
	// chain.DefaultUpdateTime holds the default value on creation for the update_time field.
	chain.DefaultUpdateTime = chainDescUpdateTime.Default.(func() time.Time)
	// chain.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	chain.UpdateDefaultUpdateTime = chainDescUpdateTime.UpdateDefault.(func() time.Time)
	// chainDescPath is the schema descriptor for path field.
	chainDescPath := chainFields[3].Descriptor()
	// chain.DefaultPath holds the default value on creation for the path field.
	chain.DefaultPath = chainDescPath.Default.(string)
	// chainDescDisplay is the schema descriptor for display field.
	chainDescDisplay := chainFields[4].Descriptor()
	// chain.DefaultDisplay holds the default value on creation for the display field.
	chain.DefaultDisplay = chainDescDisplay.Default.(string)
	// chainDescIsEnabled is the schema descriptor for is_enabled field.
	chainDescIsEnabled := chainFields[5].Descriptor()
	// chain.DefaultIsEnabled holds the default value on creation for the is_enabled field.
	chain.DefaultIsEnabled = chainDescIsEnabled.Default.(bool)
	// chainDescThumbnailURL is the schema descriptor for thumbnail_url field.
	chainDescThumbnailURL := chainFields[7].Descriptor()
	// chain.DefaultThumbnailURL holds the default value on creation for the thumbnail_url field.
	chain.DefaultThumbnailURL = chainDescThumbnailURL.Default.(string)
	chainproposalMixin := schema.ChainProposal{}.Mixin()
	chainproposalMixinFields0 := chainproposalMixin[0].Fields()
	_ = chainproposalMixinFields0
	chainproposalFields := schema.ChainProposal{}.Fields()
	_ = chainproposalFields
	// chainproposalDescCreateTime is the schema descriptor for create_time field.
	chainproposalDescCreateTime := chainproposalMixinFields0[0].Descriptor()
	// chainproposal.DefaultCreateTime holds the default value on creation for the create_time field.
	chainproposal.DefaultCreateTime = chainproposalDescCreateTime.Default.(func() time.Time)
	// chainproposalDescUpdateTime is the schema descriptor for update_time field.
	chainproposalDescUpdateTime := chainproposalMixinFields0[1].Descriptor()
	// chainproposal.DefaultUpdateTime holds the default value on creation for the update_time field.
	chainproposal.DefaultUpdateTime = chainproposalDescUpdateTime.Default.(func() time.Time)
	// chainproposal.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	chainproposal.UpdateDefaultUpdateTime = chainproposalDescUpdateTime.UpdateDefault.(func() time.Time)
	contractMixin := schema.Contract{}.Mixin()
	contractMixinFields0 := contractMixin[0].Fields()
	_ = contractMixinFields0
	contractFields := schema.Contract{}.Fields()
	_ = contractFields
	// contractDescCreateTime is the schema descriptor for create_time field.
	contractDescCreateTime := contractMixinFields0[0].Descriptor()
	// contract.DefaultCreateTime holds the default value on creation for the create_time field.
	contract.DefaultCreateTime = contractDescCreateTime.Default.(func() time.Time)
	// contractDescUpdateTime is the schema descriptor for update_time field.
	contractDescUpdateTime := contractMixinFields0[1].Descriptor()
	// contract.DefaultUpdateTime holds the default value on creation for the update_time field.
	contract.DefaultUpdateTime = contractDescUpdateTime.Default.(func() time.Time)
	// contract.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	contract.UpdateDefaultUpdateTime = contractDescUpdateTime.UpdateDefault.(func() time.Time)
	// contractDescThumbnailURL is the schema descriptor for thumbnail_url field.
	contractDescThumbnailURL := contractFields[4].Descriptor()
	// contract.DefaultThumbnailURL holds the default value on creation for the thumbnail_url field.
	contract.DefaultThumbnailURL = contractDescThumbnailURL.Default.(string)
	// contractDescRPCEndpoint is the schema descriptor for rpc_endpoint field.
	contractDescRPCEndpoint := contractFields[5].Descriptor()
	// contract.DefaultRPCEndpoint holds the default value on creation for the rpc_endpoint field.
	contract.DefaultRPCEndpoint = contractDescRPCEndpoint.Default.(string)
	// contractDescGetProposalsQuery is the schema descriptor for get_proposals_query field.
	contractDescGetProposalsQuery := contractFields[7].Descriptor()
	// contract.DefaultGetProposalsQuery holds the default value on creation for the get_proposals_query field.
	contract.DefaultGetProposalsQuery = contractDescGetProposalsQuery.Default.(string)
	contractproposalMixin := schema.ContractProposal{}.Mixin()
	contractproposalMixinFields0 := contractproposalMixin[0].Fields()
	_ = contractproposalMixinFields0
	contractproposalFields := schema.ContractProposal{}.Fields()
	_ = contractproposalFields
	// contractproposalDescCreateTime is the schema descriptor for create_time field.
	contractproposalDescCreateTime := contractproposalMixinFields0[0].Descriptor()
	// contractproposal.DefaultCreateTime holds the default value on creation for the create_time field.
	contractproposal.DefaultCreateTime = contractproposalDescCreateTime.Default.(func() time.Time)
	// contractproposalDescUpdateTime is the schema descriptor for update_time field.
	contractproposalDescUpdateTime := contractproposalMixinFields0[1].Descriptor()
	// contractproposal.DefaultUpdateTime holds the default value on creation for the update_time field.
	contractproposal.DefaultUpdateTime = contractproposalDescUpdateTime.Default.(func() time.Time)
	// contractproposal.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	contractproposal.UpdateDefaultUpdateTime = contractproposalDescUpdateTime.UpdateDefault.(func() time.Time)
	discordchannelMixin := schema.DiscordChannel{}.Mixin()
	discordchannelMixinFields0 := discordchannelMixin[0].Fields()
	_ = discordchannelMixinFields0
	discordchannelFields := schema.DiscordChannel{}.Fields()
	_ = discordchannelFields
	// discordchannelDescCreateTime is the schema descriptor for create_time field.
	discordchannelDescCreateTime := discordchannelMixinFields0[0].Descriptor()
	// discordchannel.DefaultCreateTime holds the default value on creation for the create_time field.
	discordchannel.DefaultCreateTime = discordchannelDescCreateTime.Default.(func() time.Time)
	// discordchannelDescUpdateTime is the schema descriptor for update_time field.
	discordchannelDescUpdateTime := discordchannelMixinFields0[1].Descriptor()
	// discordchannel.DefaultUpdateTime holds the default value on creation for the update_time field.
	discordchannel.DefaultUpdateTime = discordchannelDescUpdateTime.Default.(func() time.Time)
	// discordchannel.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	discordchannel.UpdateDefaultUpdateTime = discordchannelDescUpdateTime.UpdateDefault.(func() time.Time)
	telegramchatMixin := schema.TelegramChat{}.Mixin()
	telegramchatMixinFields0 := telegramchatMixin[0].Fields()
	_ = telegramchatMixinFields0
	telegramchatFields := schema.TelegramChat{}.Fields()
	_ = telegramchatFields
	// telegramchatDescCreateTime is the schema descriptor for create_time field.
	telegramchatDescCreateTime := telegramchatMixinFields0[0].Descriptor()
	// telegramchat.DefaultCreateTime holds the default value on creation for the create_time field.
	telegramchat.DefaultCreateTime = telegramchatDescCreateTime.Default.(func() time.Time)
	// telegramchatDescUpdateTime is the schema descriptor for update_time field.
	telegramchatDescUpdateTime := telegramchatMixinFields0[1].Descriptor()
	// telegramchat.DefaultUpdateTime holds the default value on creation for the update_time field.
	telegramchat.DefaultUpdateTime = telegramchatDescUpdateTime.Default.(func() time.Time)
	// telegramchat.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	telegramchat.UpdateDefaultUpdateTime = telegramchatDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
}
