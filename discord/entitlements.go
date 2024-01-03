package discord

import "time"

type Entitlement struct {
	ID             Snowflake       `json:"id"`
	SkuID          Snowflake       `json:"sku_id"`
	ApplicationID  Snowflake       `json:"application_id"`
	UserID         *Snowflake      `json:"user_id,omitempty"`
	Type           EntitlementType `json:"type"`
	Deleted        bool            `json:"deleted"`
	GiftCodeFlags  *GiftCodeFlags  `json:"gift_code_flags,omitempty"`
	StartsAt       *time.Time      `json:"starts_at,omitempty"`
	EndsAt         *time.Time      `json:"ends_at,omitempty"`
	GuildID        *Snowflake      `json:"guild_id,omitempty"`
	SubscriptionID *Snowflake      `json:"subscription_id,omitempty"`
}

// EntitlementParams represents the payload sent to discord.
type EntitlementParams struct {
	SkuID     Snowflake `json:"sku_id"`
	OwnerId   Snowflake `json:"owner_id"`
	OwnerType OwnerType `json:"owner_type"`
}

// EntitlementType represents the type of an entitlement.
type EntitlementType int8

const (
	EntitlementTypeApplicationSubscription EntitlementType = 8
)

// GiftCodeFlags is undocumented, but present in the API.
type GiftCodeFlags int8

// OwnerType represents who owns the entitlement.
type OwnerType int8

const (
	OwnerTypeGuild OwnerType = 1
	OwnerTypeUser  OwnerType = 2
)
