// Code generated by typeshare 1.11.0. DO NOT EDIT.
package onepassword

import "encoding/json"

type ItemCategory string

const (
	ItemCategoryLogin                ItemCategory = "Login"
	ItemCategorySecureNote           ItemCategory = "SecureNote"
	ItemCategoryCreditCard           ItemCategory = "CreditCard"
	ItemCategoryCryptoWallet         ItemCategory = "CryptoWallet"
	ItemCategoryIdentity             ItemCategory = "Identity"
	ItemCategoryPassword             ItemCategory = "Password"
	ItemCategoryDocument             ItemCategory = "Document"
	ItemCategoryAPICredentials       ItemCategory = "ApiCredentials"
	ItemCategoryBankAccount          ItemCategory = "BankAccount"
	ItemCategoryDatabase             ItemCategory = "Database"
	ItemCategoryDriverLicense        ItemCategory = "DriverLicense"
	ItemCategoryEmail                ItemCategory = "Email"
	ItemCategoryMedicalRecord        ItemCategory = "MedicalRecord"
	ItemCategoryMembership           ItemCategory = "Membership"
	ItemCategoryOutdoorLicense       ItemCategory = "OutdoorLicense"
	ItemCategoryPassport             ItemCategory = "Passport"
	ItemCategoryRewards              ItemCategory = "Rewards"
	ItemCategoryRouter               ItemCategory = "Router"
	ItemCategoryServer               ItemCategory = "Server"
	ItemCategorySSHKey               ItemCategory = "SshKey"
	ItemCategorySocialSecurityNumber ItemCategory = "SocialSecurityNumber"
	ItemCategorySoftwareLicense      ItemCategory = "SoftwareLicense"
	ItemCategoryPerson               ItemCategory = "Person"
	ItemCategoryUnsupported          ItemCategory = "Unsupported"
)

type ItemFieldType string

const (
	ItemFieldTypeText           ItemFieldType = "Text"
	ItemFieldTypeConcealed      ItemFieldType = "Concealed"
	ItemFieldTypeCreditCardType ItemFieldType = "CreditCardType"
	ItemFieldTypePhone          ItemFieldType = "Phone"
	ItemFieldTypeURL            ItemFieldType = "Url"
	ItemFieldTypeTOTP           ItemFieldType = "Totp"
	ItemFieldTypeUnsupported    ItemFieldType = "Unsupported"
)

// Field type-specific attributes.
type ItemFieldDetailsTypes string

const (
	// The computed OTP code and other details
	ItemFieldDetailsTypeVariantOTP ItemFieldDetailsTypes = "Otp"
)

type ItemFieldDetails struct {
	Type    ItemFieldDetailsTypes `json:"type"`
	content interface{}
}

func (i *ItemFieldDetails) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag     ItemFieldDetailsTypes `json:"type"`
		Content json.RawMessage       `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	i.Type = enum.Tag
	switch i.Type {
	case ItemFieldDetailsTypeVariantOTP:
		var res OTPFieldDetails
		i.content = &res

	}
	if err := json.Unmarshal(enum.Content, &i.content); err != nil {
		return err
	}

	return nil
}

func (i ItemFieldDetails) MarshalJSON() ([]byte, error) {
	var enum struct {
		Tag     ItemFieldDetailsTypes `json:"type"`
		Content interface{}           `json:"content,omitempty"`
	}
	enum.Tag = i.Type
	enum.Content = i.content
	return json.Marshal(enum)
}

func (i ItemFieldDetails) OTP() *OTPFieldDetails {
	res, _ := i.content.(*OTPFieldDetails)
	return res
}

func NewItemFieldDetailsTypeVariantOTP(content *OTPFieldDetails) ItemFieldDetails {
	return ItemFieldDetails{
		Type:    ItemFieldDetailsTypeVariantOTP,
		content: content,
	}
}

// Represents a field within an item.
type ItemField struct {
	// The field's ID
	ID string `json:"id"`
	// The field's title
	Title string `json:"title"`
	// The ID of the section containing the field. Built-in fields such as usernames and passwords don't require a section.
	SectionID *string `json:"sectionId,omitempty"`
	// The field's type
	FieldType ItemFieldType `json:"fieldType"`
	// The string representation of the field's value
	Value string `json:"value"`
	// Field type-specific attributes.
	Details *ItemFieldDetails `json:"details,omitempty"`
}

// A section groups together multiple fields in an item.
type ItemSection struct {
	// The section's unique ID
	ID string `json:"id"`
	// The section's title
	Title string `json:"title"`
}

// Controls the auto-fill behavior of a website.
//
// For more information, visit https://support.1password.com/autofill-behavior/
type AutofillBehavior string

const (
	// Auto-fill any page that’s part of the website, including subdomains
	AutofillBehaviorAnywhereOnWebsite AutofillBehavior = "AnywhereOnWebsite"
	// Auto-fill only if the domain (hostname and port) is an exact match.
	AutofillBehaviorExactDomain AutofillBehavior = "ExactDomain"
	// Never auto-fill on this website
	AutofillBehaviorNever AutofillBehavior = "Never"
)

type Website struct {
	// The website URL
	URL string `json:"url"`
	// The label of the website, e.g. 'website', 'sign-in address'
	Label string `json:"label"`
	// The auto-fill behavior of the website
	//
	// For more information, visit https://support.1password.com/autofill-behavior/
	AutofillBehavior AutofillBehavior `json:"autofill_behavior"`
}

// Represents a 1Password item.
type Item struct {
	// The item's ID
	ID string `json:"id"`
	// The item's title
	Title string `json:"title"`
	// The item's category
	Category ItemCategory `json:"category"`
	// The ID of the vault where the item is saved
	VaultID string `json:"vaultId"`
	// The item's fields
	Fields []ItemField `json:"fields"`
	// The item's sections
	Sections []ItemSection `json:"sections"`
	// The item's tags
	Tags []string `json:"tags"`
	// The websites used for autofilling for items of the Login and Password categories.
	Websites []Website `json:"websites"`
	// The item's version
	Version uint32 `json:"version"`
}
type ItemCreateParams struct {
	// The item's category
	Category ItemCategory `json:"category"`
	// The ID of the vault where the item is saved
	VaultID string `json:"vaultId"`
	// The item's title
	Title string `json:"title"`
	// The item's fields
	Fields []ItemField `json:"fields,omitempty"`
	// The item's sections
	Sections []ItemSection `json:"sections,omitempty"`
	// The item's tags
	Tags []string `json:"tags,omitempty"`
	// The websites used for autofilling for items of the Login and Password categories.
	Websites []Website `json:"websites,omitempty"`
}

// Represents a decrypted 1Password item.
type ItemOverview struct {
	// The item's ID
	ID string `json:"id"`
	// The item's title
	Title string `json:"title"`
	// The item's category
	Category ItemCategory `json:"category"`
	// The ID of the vault where the item is saved
	VaultID string `json:"vaultId"`
	// The websites used for autofilling for items of the Login and Password categories.
	Websites []Website `json:"websites"`
}

// Additional attributes for OTP fields.
type OTPFieldDetails struct {
	// The OTP code, if successfully computed
	Code *string `json:"code,omitempty"`
	// The error message, if the OTP code could not be computed
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Represents a decrypted 1Password vault.
type VaultOverview struct {
	// The vault's ID
	ID string `json:"id"`
	// The vault's title
	Title string `json:"title"`
}
