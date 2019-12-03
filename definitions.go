// Package peerless provides a SOAP HTTP client.
package peerless

import (
	"reflect"
)

// DateTime in WSDL format.
type DateTime string

// Action was auto-generated from WSDL.
type Action string

// Validate validates Action.
func (v Action) Validate() bool {
	for _, vv := range []string{
		"CREATE",
		"CONFIRM",
		"APPROVE",
		"REQUEST_CSR",
		"SUBMIT_LSR",
		"DECOMPOSE",
		"SINGLE_ACCOUNT",
		"MULTIPLE_ACCOUNT",
		"CSR_NOT_REQUIRED",
		"CLOSE",
		"REJECT",
		"ACTIVATE_NPAC",
		"SET_FOC",
		"REQUEST_CLARIFICATION",
		"ASSIGN",
		"ADD_GENERAL_NOTE",
		"ADD_INTERNAL_NOTE",
		"SUBSCRIBE_NPAC",
		"CHANGE_DUE_DATE",
		"CANCEL",
		"BUILT_SUBSCRIPTION_DOWNLOAD",
		"ACKNOWLEDGE_FAILURE",
		"UPDATE_NPAC",
		"REPROCESS",
		"UPDATE",
		"SUPPLEMENT",
		"PROVISIONING_IN_PROCESS",
		"SWITCH_PROVISIONED",
		"SOA_ACTIVATE",
		"BYPASS_CNAM",
		"LOA_SUBMITTED",
		"RESP_ORG_RELEASED",
		"ACTIVATION_DATE_SET",
		"IXC_ROUTING",
		"INTRALATA_DECOMPOSITION",
		"INTRALATA_CANCELED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Exception was auto-generated from WSDL.
type Exception struct {
	Message string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// Apie911Address was auto-generated from WSDL.
type Apie911Address struct {
	Number       string `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	Name         string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	StreetNumber string `xml:"streetNumber,omitempty" json:"streetNumber,omitempty" yaml:"streetNumber,omitempty"`
	Direction    string `xml:"direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	StreetName   string `xml:"streetName,omitempty" json:"streetName,omitempty" yaml:"streetName,omitempty"`
	City         string `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	State        string `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Zip          string `xml:"zip,omitempty" json:"zip,omitempty" yaml:"zip,omitempty"`
	Suite        string `xml:"suite,omitempty" json:"suite,omitempty" yaml:"suite,omitempty"`
}

// Attachment was auto-generated from WSDL.
type Attachment struct {
	FileContent []byte `xml:"fileContent,omitempty" json:"fileContent,omitempty" yaml:"fileContent,omitempty"`
	FileName    string `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
}

// BaseAddress was auto-generated from WSDL.
type BaseAddress struct {
	StreetNumber string `xml:"streetNumber,omitempty" json:"streetNumber,omitempty" yaml:"streetNumber,omitempty"`
	Direction    string `xml:"direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	StreetName   string `xml:"streetName,omitempty" json:"streetName,omitempty" yaml:"streetName,omitempty"`
	City         string `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	State        string `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
}

// DisconnectOrderRequest was auto-generated from WSDL.
type DisconnectOrderRequest struct {
	Pon string          `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	Tns []DisconnectTns `xml:"tns,omitempty" json:"tns,omitempty" yaml:"tns,omitempty"`
}

// DisconnectTns was auto-generated from WSDL.
type DisconnectTns struct {
	Tn        string `xml:"tn,omitempty" json:"tn,omitempty" yaml:"tn,omitempty"`
	DiscoType int    `xml:"discoType,omitempty" json:"discoType,omitempty" yaml:"discoType,omitempty"`
}

// ExceptionNote was auto-generated from WSDL.
type ExceptionNote struct {
	EmailID       string `xml:"emailId,omitempty" json:"emailId,omitempty" yaml:"emailId,omitempty"`
	ExceptionNote string `xml:"exceptionNote,omitempty" json:"exceptionNote,omitempty" yaml:"exceptionNote,omitempty"`
	Subject       string `xml:"subject,omitempty" json:"subject,omitempty" yaml:"subject,omitempty"`
}

// HiearchicalView was auto-generated from WSDL.
type HiearchicalView struct {
	States []StateVo `xml:"states>state,omitempty" json:"states>state,omitempty" yaml:"states>state,omitempty"`
}

// LataVo was auto-generated from WSDL.
type LataVo struct {
	Name        string         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	RateCenters []RateCenterVo `xml:"rateCenters,omitempty" json:"rateCenters,omitempty" yaml:"rateCenters,omitempty"`
}

// Note was auto-generated from WSDL.
type Note struct {
	When     string  `xml:"when,omitempty" json:"when,omitempty" yaml:"when,omitempty"`
	Notes    string  `xml:"notes,omitempty" json:"notes,omitempty" yaml:"notes,omitempty"`
	Type     string  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Action   *Action `xml:"action,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
	UserName string  `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	UserID   int     `xml:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty"`
}

// NpaVo was auto-generated from WSDL.
type NpaVo struct {
	Name  string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	NxxVo *NxxVo `xml:"nxxVo,omitempty" json:"nxxVo,omitempty" yaml:"nxxVo,omitempty"`
}

// NumberSearchParameters was auto-generated from WSDL.
type NumberSearchParameters struct {
	Pon          string     `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	States       []string   `xml:"states>state,omitempty" json:"states>state,omitempty" yaml:"states>state,omitempty"`
	Latas        []string   `xml:"latas>lata,omitempty" json:"latas>lata,omitempty" yaml:"latas>lata,omitempty"`
	Locations    []string   `xml:"locations>location,omitempty" json:"locations>location,omitempty" yaml:"locations>location,omitempty"`
	RateCenters  []string   `xml:"rateCenters>rateCenter,omitempty" json:"rateCenters>rateCenter,omitempty" yaml:"rateCenters>rateCenter,omitempty"`
	Npas         []string   `xml:"npas>npa,omitempty" json:"npas>npa,omitempty" yaml:"npas>npa,omitempty"`
	Nxxs         []string   `xml:"nxxs>nxx,omitempty" json:"nxxs>nxx,omitempty" yaml:"nxxs>nxx,omitempty"`
	Quantity     int64      `xml:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	Consecutive  int64      `xml:"consecutive,omitempty" json:"consecutive,omitempty" yaml:"consecutive,omitempty"`
	VanityDigits string     `xml:"vanityDigits,omitempty" json:"vanityDigits,omitempty" yaml:"vanityDigits,omitempty"`
	Categories   []string   `xml:"categories>category,omitempty" json:"categories>category,omitempty" yaml:"categories>category,omitempty"`
	Result       []ResultTN `xml:"result>tn,omitempty" json:"result>tn,omitempty" yaml:"result>tn,omitempty"`
	SelectedTN   []int64    `xml:"selectedTN>tn,omitempty" json:"selectedTN>tn,omitempty" yaml:"selectedTN>tn,omitempty"`
	TnCount      int64      `xml:"tnCount,omitempty" json:"tnCount,omitempty" yaml:"tnCount,omitempty"`
}

// NxxVo was auto-generated from WSDL.
type NxxVo struct {
	Nxxs []string `xml:"nxxs,omitempty" json:"nxxs,omitempty" yaml:"nxxs,omitempty"`
}

// Order was auto-generated from WSDL.
type Order struct {
	Pon          string        `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	OrderType    string        `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	OrderNumbers []OrderNumber `xml:"orderNumbers>number,omitempty" json:"orderNumbers>number,omitempty" yaml:"orderNumbers>number,omitempty"`
	OrderDetails *OrderDetails `xml:"orderDetails,omitempty" json:"orderDetails,omitempty" yaml:"orderDetails,omitempty"`
}

// OrderDetails was auto-generated from WSDL.
type OrderDetails struct {
	AccountNumber         string    `xml:"accountNumber,omitempty" json:"accountNumber,omitempty" yaml:"accountNumber,omitempty"`
	Act                   string    `xml:"act,omitempty" json:"act,omitempty" yaml:"act,omitempty"`
	AssignedTo            string    `xml:"assignedTo,omitempty" json:"assignedTo,omitempty" yaml:"assignedTo,omitempty"`
	Atn                   string    `xml:"atn,omitempty" json:"atn,omitempty" yaml:"atn,omitempty"`
	AuthDate              string    `xml:"authDate,omitempty" json:"authDate,omitempty" yaml:"authDate,omitempty"`
	Authnm                string    `xml:"authnm,omitempty" json:"authnm,omitempty" yaml:"authnm,omitempty"`
	Bldg                  string    `xml:"bldg,omitempty" json:"bldg,omitempty" yaml:"bldg,omitempty"`
	CarrierName           string    `xml:"carrierName,omitempty" json:"carrierName,omitempty" yaml:"carrierName,omitempty"`
	Cc                    string    `xml:"cc,omitempty" json:"cc,omitempty" yaml:"cc,omitempty"`
	Ccna                  string    `xml:"ccna,omitempty" json:"ccna,omitempty" yaml:"ccna,omitempty"`
	Chc                   string    `xml:"chc,omitempty" json:"chc,omitempty" yaml:"chc,omitempty"`
	City                  string    `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	Confirmed             bool      `xml:"confirmed,omitempty" json:"confirmed,omitempty" yaml:"confirmed,omitempty"`
	CreatedDate           *DateTime `xml:"createdDate,omitempty" json:"createdDate,omitempty" yaml:"createdDate,omitempty"`
	CsrAccountNumber      string    `xml:"csrAccountNumber,omitempty" json:"csrAccountNumber,omitempty" yaml:"csrAccountNumber,omitempty"`
	CsrRequested          bool      `xml:"csrRequested,omitempty" json:"csrRequested,omitempty" yaml:"csrRequested,omitempty"`
	DesiredDueDate        string    `xml:"desiredDueDate,omitempty" json:"desiredDueDate,omitempty" yaml:"desiredDueDate,omitempty"`
	DueDate               string    `xml:"dueDate,omitempty" json:"dueDate,omitempty" yaml:"dueDate,omitempty"`
	EarliestPossible      bool      `xml:"earliestPossible,omitempty" json:"earliestPossible,omitempty" yaml:"earliestPossible,omitempty"`
	Email                 string    `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	EndCustomerName       string    `xml:"endCustomerName,omitempty" json:"endCustomerName,omitempty" yaml:"endCustomerName,omitempty"`
	EndUserVerification   bool      `xml:"endUserVerification,omitempty" json:"endUserVerification,omitempty" yaml:"endUserVerification,omitempty"`
	FaxNumber             string    `xml:"faxNumber,omitempty" json:"faxNumber,omitempty" yaml:"faxNumber,omitempty"`
	Floor                 string    `xml:"floor,omitempty" json:"floor,omitempty" yaml:"floor,omitempty"`
	FocDate               *DateTime `xml:"focDate,omitempty" json:"focDate,omitempty" yaml:"focDate,omitempty"`
	FocDateSync           bool      `xml:"focDateSync,omitempty" json:"focDateSync,omitempty" yaml:"focDateSync,omitempty"`
	GeneralNote           string    `xml:"generalNote,omitempty" json:"generalNote,omitempty" yaml:"generalNote,omitempty"`
	Init                  string    `xml:"init,omitempty" json:"init,omitempty" yaml:"init,omitempty"`
	InternalNote          string    `xml:"internalNote,omitempty" json:"internalNote,omitempty" yaml:"internalNote,omitempty"`
	Mi                    string    `xml:"mi,omitempty" json:"mi,omitempty" yaml:"mi,omitempty"`
	ModifiedDate          *DateTime `xml:"modifiedDate,omitempty" json:"modifiedDate,omitempty" yaml:"modifiedDate,omitempty"`
	MouThresholdExceeded  bool      `xml:"mouThresholdExceeded,omitempty" json:"mouThresholdExceeded,omitempty" yaml:"mouThresholdExceeded,omitempty"`
	NewAtn                string    `xml:"newAtn,omitempty" json:"newAtn,omitempty" yaml:"newAtn,omitempty"`
	Nnsp                  string    `xml:"nnsp,omitempty" json:"nnsp,omitempty" yaml:"nnsp,omitempty"`
	NpacActivated         bool      `xml:"npacActivated,omitempty" json:"npacActivated,omitempty" yaml:"npacActivated,omitempty"`
	NpacActivatedDate     *DateTime `xml:"npacActivatedDate,omitempty" json:"npacActivatedDate,omitempty" yaml:"npacActivatedDate,omitempty"`
	Npdi                  string    `xml:"npdi,omitempty" json:"npdi,omitempty" yaml:"npdi,omitempty"`
	NumberOfTns           int64     `xml:"numberOfTns,omitempty" json:"numberOfTns,omitempty" yaml:"numberOfTns,omitempty"`
	PortOutAssignedEntity string    `xml:"portOutAssignedEntity,omitempty" json:"portOutAssignedEntity,omitempty" yaml:"portOutAssignedEntity,omitempty"`
	PurchaseOrderNumber   string    `xml:"purchaseOrderNumber,omitempty" json:"purchaseOrderNumber,omitempty" yaml:"purchaseOrderNumber,omitempty"`
	Remarks               string    `xml:"remarks,omitempty" json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Room                  string    `xml:"room,omitempty" json:"room,omitempty" yaml:"room,omitempty"`
	SadloAdditionalInfo   string    `xml:"sadloAdditionalInfo,omitempty" json:"sadloAdditionalInfo,omitempty" yaml:"sadloAdditionalInfo,omitempty"`
	SanoHouseNumber       string    `xml:"sanoHouseNumber,omitempty" json:"sanoHouseNumber,omitempty" yaml:"sanoHouseNumber,omitempty"`
	SaprHousePrefix       string    `xml:"saprHousePrefix,omitempty" json:"saprHousePrefix,omitempty" yaml:"saprHousePrefix,omitempty"`
	SasdStrDir            string    `xml:"sasdStrDir,omitempty" json:"sasdStrDir,omitempty" yaml:"sasdStrDir,omitempty"`
	SasfHouseNumberSuffix string    `xml:"sasfHouseNumberSuffix,omitempty" json:"sasfHouseNumberSuffix,omitempty" yaml:"sasfHouseNumberSuffix,omitempty"`
	SasnStrName           string    `xml:"sasnStrName,omitempty" json:"sasnStrName,omitempty" yaml:"sasnStrName,omitempty"`
	SassStrSuffix         string    `xml:"sassStrSuffix,omitempty" json:"sassStrSuffix,omitempty" yaml:"sassStrSuffix,omitempty"`
	Sath                  string    `xml:"sath,omitempty" json:"sath,omitempty" yaml:"sath,omitempty"`
	State                 string    `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	SubscriptionBuilt     bool      `xml:"subscriptionBuilt,omitempty" json:"subscriptionBuilt,omitempty" yaml:"subscriptionBuilt,omitempty"`
	Supplement            bool      `xml:"supplement,omitempty" json:"supplement,omitempty" yaml:"supplement,omitempty"`
	TelephoneNumber       string    `xml:"telephoneNumber,omitempty" json:"telephoneNumber,omitempty" yaml:"telephoneNumber,omitempty"`
	Version               int64     `xml:"version,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
	WirelessPin           string    `xml:"wirelessPin,omitempty" json:"wirelessPin,omitempty" yaml:"wirelessPin,omitempty"`
	ZipCode               string    `xml:"zipCode,omitempty" json:"zipCode,omitempty" yaml:"zipCode,omitempty"`
}

// OrderNumber was auto-generated from WSDL.
type OrderNumber struct {
	Tn                      string          `xml:"tn,omitempty" json:"tn,omitempty" yaml:"tn,omitempty"`
	Mou                     string          `xml:"mou,omitempty" json:"mou,omitempty" yaml:"mou,omitempty"`
	RouteLabel              string          `xml:"routeLabel,omitempty" json:"routeLabel,omitempty" yaml:"routeLabel,omitempty"`
	CustomName              string          `xml:"customName,omitempty" json:"customName,omitempty" yaml:"customName,omitempty"`
	CnamDelivery            bool            `xml:"cnamDelivery,omitempty" json:"cnamDelivery,omitempty" yaml:"cnamDelivery,omitempty"`
	CnamStorage             bool            `xml:"cnamStorage,omitempty" json:"cnamStorage,omitempty" yaml:"cnamStorage,omitempty"`
	CnamStorageName         string          `xml:"cnamStorageName,omitempty" json:"cnamStorageName,omitempty" yaml:"cnamStorageName,omitempty"`
	PeerlessMsgProvisioning string          `xml:"peerlessMsgProvisioning,omitempty" json:"peerlessMsgProvisioning,omitempty" yaml:"peerlessMsgProvisioning,omitempty"`
	Sms                     bool            `xml:"sms,omitempty" json:"sms,omitempty" yaml:"sms,omitempty"`
	E911                    bool            `xml:"e911,omitempty" json:"e911,omitempty" yaml:"e911,omitempty"`
	Address                 *Apie911Address `xml:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
}

// OrderSearch was auto-generated from WSDL.
type OrderSearch struct {
	OrderSearch *OrderSearchDetails `xml:"orderSearch,omitempty" json:"orderSearch,omitempty" yaml:"orderSearch,omitempty"`
}

// OrderSearchDetails was auto-generated from WSDL.
type OrderSearchDetails struct {
	OrderNumber      string        `xml:"orderNumber,omitempty" json:"orderNumber,omitempty" yaml:"orderNumber,omitempty"`
	Pon              string        `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	Entity           string        `xml:"entity,omitempty" json:"entity,omitempty" yaml:"entity,omitempty"`
	Numbers          []OrderNumber `xml:"numbers,omitempty" json:"numbers,omitempty" yaml:"numbers,omitempty"`
	Destination      string        `xml:"destination,omitempty" json:"destination,omitempty" yaml:"destination,omitempty"`
	NumberOfTNs      int64         `xml:"numberOfTNs,omitempty" json:"numberOfTNs,omitempty" yaml:"numberOfTNs,omitempty"`
	RequestedDueDate string        `xml:"requestedDueDate,omitempty" json:"requestedDueDate,omitempty" yaml:"requestedDueDate,omitempty"`
	FOCDate          string        `xml:"FOCDate,omitempty" json:"FOCDate,omitempty" yaml:"FOCDate,omitempty"`
	Status           string        `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Attachment       []string      `xml:"attachment,omitempty" json:"attachment,omitempty" yaml:"attachment,omitempty"`
	Notes            []Note        `xml:"notes,omitempty" json:"notes,omitempty" yaml:"notes,omitempty"`
}

// PortPSNumber was auto-generated from WSDL.
type PortPSNumber struct {
	Number      string `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	RateCenter  string `xml:"rateCenter,omitempty" json:"rateCenter,omitempty" yaml:"rateCenter,omitempty"`
	Lata        int64  `xml:"lata,omitempty" json:"lata,omitempty" yaml:"lata,omitempty"`
	State       string `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Location    string `xml:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty"`
	CarrierName string `xml:"carrierName,omitempty" json:"carrierName,omitempty" yaml:"carrierName,omitempty"`
	Ocn         string `xml:"ocn,omitempty" json:"ocn,omitempty" yaml:"ocn,omitempty"`
	OcnName     string `xml:"ocnName,omitempty" json:"ocnName,omitempty" yaml:"ocnName,omitempty"`
	Provider    string `xml:"provider,omitempty" json:"provider,omitempty" yaml:"provider,omitempty"`
	OffNet      string `xml:"offNet,omitempty" json:"offNet,omitempty" yaml:"offNet,omitempty"`
	Portable    string `xml:"portable,omitempty" json:"portable,omitempty" yaml:"portable,omitempty"`
}

// PortabilityCheckRequest was auto-generated from WSDL.
type PortabilityCheckRequest struct {
	Tns []string `xml:"tns>tns,omitempty" json:"tns>tns,omitempty" yaml:"tns>tns,omitempty"`
}

// RateCenterVo was auto-generated from WSDL.
type RateCenterVo struct {
	Name   string  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	NpasVo []NpaVo `xml:"npasVo,omitempty" json:"npasVo,omitempty" yaml:"npasVo,omitempty"`
}

// ResultOrderDetails was auto-generated from WSDL.
type ResultOrderDetails struct {
	OrderNumber string `xml:"orderNumber,omitempty" json:"orderNumber,omitempty" yaml:"orderNumber,omitempty"`
	Pon         string `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	NumberOfTNs int64  `xml:"numberOfTNs,omitempty" json:"numberOfTNs,omitempty" yaml:"numberOfTNs,omitempty"`
	Carrier     string `xml:"carrier,omitempty" json:"carrier,omitempty" yaml:"carrier,omitempty"`
	CSRAccount  string `xml:"CSRAccount,omitempty" json:"CSRAccount,omitempty" yaml:"CSRAccount,omitempty"`
	Status      string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	FOCDateSync string `xml:"FOCDateSync,omitempty" json:"FOCDateSync,omitempty" yaml:"FOCDateSync,omitempty"`
}

// ResultPONOrderDetails was auto-generated from WSDL.
type ResultPONOrderDetails struct {
	OrderNumber  *int64  `xml:"orderNumber,omitempty" json:"orderNumber,omitempty" yaml:"orderNumber,omitempty"`
	OrderType    *string `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	Pon          *string `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	Status       *string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	DateReceived *string `xml:"dateReceived,omitempty" json:"dateReceived,omitempty" yaml:"dateReceived,omitempty"`
}

// ResultTN was auto-generated from WSDL.
type ResultTN struct {
	Number   int64  `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	Category string `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
}

// StateVo was auto-generated from WSDL.
type StateVo struct {
	Name  string   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Latas []LataVo `xml:"latas>lata,omitempty" json:"latas>lata,omitempty" yaml:"latas>lata,omitempty"`
}

// SupplementInfo was auto-generated from WSDL.
type SupplementInfo struct {
	OrderID int64     `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
	Type    int64     `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	DueDate *DateTime `xml:"dueDate,omitempty" json:"dueDate,omitempty" yaml:"dueDate,omitempty"`
}

// TfOrderNumber was auto-generated from WSDL.
type TfOrderNumber struct {
	Tn               string `xml:"tn,omitempty" json:"tn,omitempty" yaml:"tn,omitempty"`
	Quantity         string `xml:"quantity,omitempty" json:"quantity,omitempty" yaml:"quantity,omitempty"`
	Aos              string `xml:"aos,omitempty" json:"aos,omitempty" yaml:"aos,omitempty"`
	Destination      string `xml:"destination,omitempty" json:"destination,omitempty" yaml:"destination,omitempty"`
	CnamDelivery     string `xml:"cnamDelivery,omitempty" json:"cnamDelivery,omitempty" yaml:"cnamDelivery,omitempty"`
	PayPhoneBlocking string `xml:"payPhoneBlocking,omitempty" json:"payPhoneBlocking,omitempty" yaml:"payPhoneBlocking,omitempty"`
}

// TnInventory was auto-generated from WSDL.
type TnInventory struct {
	Address                 *Apie911Address `xml:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Category                string          `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	CnamDeliveryProvisioned bool            `xml:"cnamDeliveryProvisioned,omitempty" json:"cnamDeliveryProvisioned,omitempty" yaml:"cnamDeliveryProvisioned,omitempty"`
	CnamStorageName         string          `xml:"cnamStorageName,omitempty" json:"cnamStorageName,omitempty" yaml:"cnamStorageName,omitempty"`
	Coverage                string          `xml:"coverage,omitempty" json:"coverage,omitempty" yaml:"coverage,omitempty"`
	CustomName              string          `xml:"customName,omitempty" json:"customName,omitempty" yaml:"customName,omitempty"`
	Destination             string          `xml:"destination,omitempty" json:"destination,omitempty" yaml:"destination,omitempty"`
	E911Provisioned         bool            `xml:"e911Provisioned,omitempty" json:"e911Provisioned,omitempty" yaml:"e911Provisioned,omitempty"`
	Entity                  string          `xml:"entity,omitempty" json:"entity,omitempty" yaml:"entity,omitempty"`
	Lata                    int64           `xml:"lata,omitempty" json:"lata,omitempty" yaml:"lata,omitempty"`
	Location                string          `xml:"location,omitempty" json:"location,omitempty" yaml:"location,omitempty"`
	ModifiedDate            *DateTime       `xml:"modifiedDate,omitempty" json:"modifiedDate,omitempty" yaml:"modifiedDate,omitempty"`
	Number                  string          `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	NumberType              string          `xml:"numberType,omitempty" json:"numberType,omitempty" yaml:"numberType,omitempty"`
	RcAbbrev                string          `xml:"rcAbbrev,omitempty" json:"rcAbbrev,omitempty" yaml:"rcAbbrev,omitempty"`
	Residency               string          `xml:"residency,omitempty" json:"residency,omitempty" yaml:"residency,omitempty"`
	SmsProvisioned          bool            `xml:"smsProvisioned,omitempty" json:"smsProvisioned,omitempty" yaml:"smsProvisioned,omitempty"`
	State                   string          `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Status                  string          `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
}

// TnInventoryForApiSearchParams was auto-generated from WSDL.
type TnInventoryForApiSearchParams struct {
	Categories       []string `xml:"categories>category,omitempty" json:"categories>category,omitempty" yaml:"categories>category,omitempty"`
	CnamDelivery     int64    `xml:"cnamDelivery,omitempty" json:"cnamDelivery,omitempty" yaml:"cnamDelivery,omitempty"`
	CnamStorage      int64    `xml:"cnamStorage,omitempty" json:"cnamStorage,omitempty" yaml:"cnamStorage,omitempty"`
	CnamStorageName  string   `xml:"cnamStorageName,omitempty" json:"cnamStorageName,omitempty" yaml:"cnamStorageName,omitempty"`
	Coverages        []string `xml:"coverages>coverage,omitempty" json:"coverages>coverage,omitempty" yaml:"coverages>coverage,omitempty"`
	Destinations     []string `xml:"destinations>destination,omitempty" json:"destinations>destination,omitempty" yaml:"destinations>destination,omitempty"`
	E911             int64    `xml:"e911,omitempty" json:"e911,omitempty" yaml:"e911,omitempty"`
	Entities         []string `xml:"entities>entity,omitempty" json:"entities>entity,omitempty" yaml:"entities>entity,omitempty"`
	LastmodifiedFrom string   `xml:"lastmodifiedFrom,omitempty" json:"lastmodifiedFrom,omitempty" yaml:"lastmodifiedFrom,omitempty"`
	LastmodifiedTo   string   `xml:"lastmodifiedTo,omitempty" json:"lastmodifiedTo,omitempty" yaml:"lastmodifiedTo,omitempty"`
	Latas            []int64  `xml:"latas>lata,omitempty" json:"latas>lata,omitempty" yaml:"latas>lata,omitempty"`
	ListOfTns        []string `xml:"listOfTns>tn,omitempty" json:"listOfTns>tn,omitempty" yaml:"listOfTns>tn,omitempty"`
	PeerlessMsg      int      `xml:"peerlessMsg,omitempty" json:"peerlessMsg,omitempty" yaml:"peerlessMsg,omitempty"`
	Ratecenters      []string `xml:"ratecenters>ratecenter,omitempty" json:"ratecenters>ratecenter,omitempty" yaml:"ratecenters>ratecenter,omitempty"`
	Residencies      []string `xml:"residencies>residency,omitempty" json:"residencies>residency,omitempty" yaml:"residencies>residency,omitempty"`
	Sms              int      `xml:"sms,omitempty" json:"sms,omitempty" yaml:"sms,omitempty"`
	States           []string `xml:"states>state,omitempty" json:"states>state,omitempty" yaml:"states>state,omitempty"`
	TnStatuses       []string `xml:"tnStatuses>tnStatus,omitempty" json:"tnStatuses>tnStatus,omitempty" yaml:"tnStatuses>tnStatus,omitempty"`
	Types            []string `xml:"types>type,omitempty" json:"types>type,omitempty" yaml:"types>type,omitempty"`
}

// TollFreeOrder was auto-generated from WSDL.
type TollFreeOrder struct {
	Pon                     string          `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
	OrderType               string          `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
	RequestedDueDate        string          `xml:"requestedDueDate,omitempty" json:"requestedDueDate,omitempty" yaml:"requestedDueDate,omitempty"`
	ActivateReservation     string          `xml:"activateReservation,omitempty" json:"activateReservation,omitempty" yaml:"activateReservation,omitempty"`
	EndUserName             string          `xml:"endUserName,omitempty" json:"endUserName,omitempty" yaml:"endUserName,omitempty"`
	RespOrgID               string          `xml:"respOrgId,omitempty" json:"respOrgId,omitempty" yaml:"respOrgId,omitempty"`
	DefaultAos              string          `xml:"defaultAos,omitempty" json:"defaultAos,omitempty" yaml:"defaultAos,omitempty"`
	DefaultDestination      string          `xml:"defaultDestination,omitempty" json:"defaultDestination,omitempty" yaml:"defaultDestination,omitempty"`
	DefaultPayPhoneBlocking string          `xml:"defaultPayPhoneBlocking,omitempty" json:"defaultPayPhoneBlocking,omitempty" yaml:"defaultPayPhoneBlocking,omitempty"`
	TollFreeOrderNumbers    []TfOrderNumber `xml:"tollFreeOrderNumbers>tollFreeOrderNumbers,omitempty" json:"tollFreeOrderNumbers>tollFreeOrderNumbers,omitempty" yaml:"tollFreeOrderNumbers>tollFreeOrderNumbers,omitempty"`
}

type authInfo struct {
	Customer string `xml:"customer,omitempty" json:"customer,omitempty" yaml:"customer,omitempty"`
	PassCode string `xml:"passCode,omitempty" json:"passCode,omitempty" yaml:"passCode,omitempty"`
	UserId   string `xml:"userId,omitempty" json:"userId,omitempty" yaml:"userId,omitempty"`
}
