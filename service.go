// Package peerless provides a SOAP HTTP client.
package peerless

import (
	"context"
	"net/http"
)

const (
	// Envelope is a SOAP envelope.
	Namespace = "http://publicapi.api.s2.peerless.com/"
	// ProductionEndpoint endpoint used in production
	ProductionEndpoint = "https://animate.peerlessnetwork.com:8181/animateapi/axis/APIService"
	// StagingEndpoint endpoint used for development
	StagingEndpoint = "https://aniweb02.peerlessnetwork.com:8181/animateapi/axis/APIService"
)

// APIService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type APIService interface {
	// ActivateSOA The numbers for the given orderId(s) would be npac activated. Only the numbers which are peerless
	// and offnet numbers would have npac activation skipped.
	ActivateSOA(ctx context.Context, orderID []string) (bool, error)

	// AddNotes add notes to an Order
	AddNotes(ctx context.Context, orderID, note string) (bool, error)

	// CreateException An Exception Request should be used in the event that your search results do not return the desired
	// number availability. This request will generate an email notification to a designated Peerless
	// representative to further investigate. Please include as many details of your desired request in the
	// body of the email.
	CreateException(ctx context.Context, ExceptionNote *ExceptionNote) (bool, error)

	// DisconnectOrder If order is soft disconnect with or without features then available date will not be
	// displayed. And if it is hard disconnect then available date will be shown.
	DisconnectOrder(ctx context.Context, DisconnectOrderRequest *DisconnectOrderRequest) (string, error)

	// Download This method can be used to upload required attachments to an order
	Download(ctx context.Context, orderID string) ([]Attachment, error)

	// GetHierarchicalView Enables a user to identify the full coverage area in Peerless Network’s entire inventory by
	// State, LATA, Rate Center, NPAs and NXXs.
	GetHierarchicalView(ctx context.Context) (*HiearchicalView, error)

	// GetNewNumberSearchFilters With this operation users can request a list of all possible additional filter options
	// for a given filter.
	GetNewNumberSearchFilters(ctx context.Context, filters *NumberSearchParameters) (*NumberSearchParameters, error)

	// GetOrderStatus This operation returns Order Status
	GetOrderStatus(ctx context.Context, orderID string) (string, error)

	// GetOrdersByPONSearch  This is used to get the orders based on the PON entered in the request.
	GetOrdersByPONSearch(ctx context.Context, pon string) ([]ResultPONOrderDetails, error)

	// GetPortInRelatedOrders godoc
	GetPortInRelatedOrders(ctx context.Context, orderID int64) ([]ResultOrderDetails, error)

	// GetStatusByNumberSearch This is to get the status of the number given in the request. The number should belong
	// to the entity in question.
	GetStatusByNumberSearch(ctx context.Context, telephoneNumber []string) ([]string, error)

	// GetTnInventoryReport godoc
	GetTnInventoryReport(ctx context.Context, searchParams *TnInventoryForApiSearchParams) ([]TnInventory, error)

	// PlaceOrder was Placing your order will begin the activation process of your selected numbers. This request
	// includes features and destination instructions to provision your selected numbers.
	PlaceOrder(ctx context.Context, order *Order) (string, error)

	// PlaceTFDisconnectOrder godoc
	PlaceTFDisconnectOrder(ctx context.Context, disconnectOrderRequest *DisconnectOrderRequest) (string, error)

	// PlaceTFOrder godoc
	PlaceTFOrder(ctx context.Context, order *TollFreeOrder) (string, error)

	// PortabilityCheck godoc
	PortabilityCheck(ctx context.Context, portabilityCheckRequest *PortabilityCheckRequest) ([]string, error)

	// SearchNumbers This Operation will search the available inventory for numbers from Peerless.
	SearchNumbers(ctx context.Context, filters *NumberSearchParameters) (*NumberSearchParameters, error)

	// SearchOrderDetailsByOrderId godoc
	SearchOrderDetailsByOrderId(ctx context.Context, orderID int64, orderType string) (*OrderSearch, error)

	// SupplementOrder Supplement requests for PortIn order. Supported supplement Types:
	// <type> 1.cancel 2- change due date, 3- all other
	SupplementOrder(ctx context.Context, supplementInfo *SupplementInfo, order *Order) (string, error)

	// Upload This method can be used to upload required attachments to an order.
	Upload(ctx context.Context, orderID string, attachments []Attachment) (bool, error)

	// ValidateE911Address This operation will validate User Address for E911
	ValidateE911Address(ctx context.Context, address BaseAddress) ([]BaseAddress, error)

	// ViewNumberDetails This Operation will search all the features for the requested telephone numbers in the
	// inventory from Peerless.
	ViewNumberDetails(ctx context.Context, numbers []string) ([]OrderNumber, error)
}

// New creates an initializes a API service.
func New(url, customer, passcode, userid string) APIService {
	return &service{
		URL: url,
		Authentication: &authInfo{
			Customer: customer,
			PassCode: passcode,
			UserId:   userid,
		},
	}
}

// service implements the APIService interface.
type service struct {
	URL                    string       // URL of the server
	UserAgent              string       // User-Agent header will be added to each request
	URNamespace            string       // Uniform Resource Namespace
	ThisNamespace          string       // SOAP This-Namespace (tns)
	ExcludeActionNamespace bool         // Include Namespace to SOAP Action header
	Envelope               string       // Optional SOAP Envelope
	Config                 *http.Client // Optional HTTP client
	Authentication         *authInfo    // Authentication
}

// ActivateSOA was auto-generated from WSDL.
func (s *service) ActivateSOA(ctx context.Context, orderID []string) (bool, error) {
	req := request{
		ActivateSOA: &activateSOA{
			Authentication: s.Authentication,
			OrderID:        orderID,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return false, err
	}
	return res.ActivateSOA.Return, nil
}

// AddNotes was auto-generated from WSDL.
func (s *service) AddNotes(ctx context.Context, orderID, note string) (bool, error) {
	req := request{
		AddNotes: &addNotes{
			Authentication: s.Authentication,
			OrderID:        orderID,
			Notes:          note,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return false, err
	}
	return res.AddNotes.Return, nil
}

// CreateException was auto-generated from WSDL.
func (s *service) CreateException(ctx context.Context, ExceptionNote *ExceptionNote) (bool, error) {
	req := request{
		CreateException: &createException{
			Authentication: s.Authentication,
			ExceptionNote:  ExceptionNote,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return false, err
	}
	return res.CreateException.Return, nil
}

// DisconnectOrder was auto-generated from WSDL.
func (s *service) DisconnectOrder(ctx context.Context, DisconnectOrderRequest *DisconnectOrderRequest) (string, error) {
	req := request{
		DisconnectOrder: &disconnectOrder{
			Authentication:         s.Authentication,
			DisconnectOrderRequest: DisconnectOrderRequest,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return "", err
	}
	return res.DisconnectOrder.Return, nil
}

// Download was auto-generated from WSDL.
func (s *service) Download(ctx context.Context, orderID string) ([]Attachment, error) {
	req := request{
		Download: &download{
			Authentication: s.Authentication,
			OrderID:        orderID,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Download.Return, nil
}

// GetHierarchicalView was auto-generated from WSDL.
func (s *service) GetHierarchicalView(ctx context.Context) (*HiearchicalView, error) {
	req := request{
		GetHierarchicalView: &getHierarchicalView{
			Authentication: s.Authentication,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetHierarchicalView.Return, nil
}

// GetNewNumberSearchFilters was auto-generated from WSDL.
func (s *service) GetNewNumberSearchFilters(ctx context.Context, filters *NumberSearchParameters) (*NumberSearchParameters, error) {
	req := request{
		GetNewNumberSearchFilters: &getNewNumberSearchFilters{
			Authentication: s.Authentication,
			Filters:        filters,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetNewNumberSearchFilters.Return, nil
}

// GetOrderStatus was auto-generated from WSDL.
func (s *service) GetOrderStatus(ctx context.Context, orderID string) (string, error) {
	req := request{
		GetOrderStatus: &getOrderStatus{
			OrderID: orderID,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetOrderStatus.Return, nil
}

// GetOrdersByPONSearch was auto-generated from WSDL.
func (s *service) GetOrdersByPONSearch(ctx context.Context, pon string) ([]ResultPONOrderDetails, error) {
	req := request{
		GetOrdersByPONSearch: &getOrdersByPONSearch{
			Authentication: s.Authentication,
			Pon:            pon,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetOrdersByPONSearch.Return.Result, nil
}

// GetPortInRelatedOrders was auto-generated from WSDL.
func (s *service) GetPortInRelatedOrders(ctx context.Context, orderID int64) ([]ResultOrderDetails, error) {
	req := request{
		GetPortInRelatedOrders: &getPortInRelatedOrders{
			Authentication: s.Authentication,
			OrderID:        orderID,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetPortInRelatedOrders.Return.Result, nil
}

// GetStatusByNumberSearch was auto-generated from WSDL.
func (s *service) GetStatusByNumberSearch(ctx context.Context, telephoneNumber []string) ([]string, error) {
	req := request{
		GetStatusByNumberSearch: &getStatusByNumberSearch{
			Authentication:  s.Authentication,
			TelephoneNumber: telephoneNumber,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetStatusByNumberSearch.Return.Result, nil
}

// GetTnInventoryReport was auto-generated from WSDL.
func (s *service) GetTnInventoryReport(ctx context.Context, searchParams *TnInventoryForApiSearchParams) ([]TnInventory, error) {
	req := request{
		GetTnInventoryReport: &getTnInventoryReport{
			Authentication: s.Authentication,
			SearchParams:   searchParams,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetTnInventoryReport.Return, nil
}

// PlaceOrder was auto-generated from WSDL.
func (s *service) PlaceOrder(ctx context.Context, order *Order) (string, error) {
	req := request{
		PlaceOrder: &placeOrder{
			Authentication: s.Authentication,
			Order:          order,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return "", err
	}
	return res.PlaceOrder.Return, nil
}

// PlaceTFDisconnectOrder was auto-generated from WSDL.
func (s *service) PlaceTFDisconnectOrder(ctx context.Context, disconnectOrderRequest *DisconnectOrderRequest) (string, error) {
	req := request{
		PlaceTFDisconnectOrder: &placeTFDisconnectOrder{
			Authentication:         s.Authentication,
			DisconnectOrderRequest: disconnectOrderRequest,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return "", err
	}
	return res.PlaceTFDisconnectOrder.Return, nil
}

// PlaceTFOrder was auto-generated from WSDL.
func (s *service) PlaceTFOrder(ctx context.Context, order *TollFreeOrder) (string, error) {
	req := request{
		PlaceTFOrder: &placeTFOrder{
			Authentication: s.Authentication,
			TFNOrder:       order,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return "", err
	}
	return res.PlaceTFOrder.Return, nil
}

// PortabilityCheck was auto-generated from WSDL.
func (s *service) PortabilityCheck(ctx context.Context, portabilityCheckRequest *PortabilityCheckRequest) ([]string, error) {
	req := request{
		PortabilityCheck: &portabilityCheck{
			Authentication:          s.Authentication,
			PortabilityCheckRequest: portabilityCheckRequest,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.PortabilityCheck.Return.PortPsNumbers, nil
}

// SearchNumbers was auto-generated from WSDL.
func (s *service) SearchNumbers(ctx context.Context, filters *NumberSearchParameters) (*NumberSearchParameters, error) {
	req := request{
		SearchNumbers: &searchNumbers{
			Authentication: s.Authentication,
			Filters:        filters,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.SearchNumbers.Return, nil
}

// SearchOrderDetailsByOrderId was auto-generated from WSDL.
func (s *service) SearchOrderDetailsByOrderId(ctx context.Context, orderID int64, orderType string) (*OrderSearch, error) {
	req := request{
		SearchOrderDetailsByOrderId: &searchOrderDetailsByOrderId{
			Authentication: s.Authentication,
			OrderID:        orderID,
			OrderType:      orderType,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.SearchOrderDetailsByOrderId.Return, nil
}

// SupplementOrder was auto-generated from WSDL.
func (s *service) SupplementOrder(ctx context.Context, supplementInfo *SupplementInfo, order *Order) (string, error) {
	req := request{
		SupplementOrder: &supplementOrder{
			Authentication: s.Authentication,
			SupplementInfo: supplementInfo,
			Order:          order,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return "", err
	}
	return res.SupplementOrder.Return, nil
}

// Upload was auto-generated from WSDL.
func (s *service) Upload(ctx context.Context, orderID string, attachments []Attachment) (bool, error) {
	req := request{
		Upload: &upload{
			Authentication: s.Authentication,
			OrderID:        orderID,
			Attachments:    attachments,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return false, err
	}
	return res.Upload.Return, nil
}

// ValidateE911Address was auto-generated from WSDL.
func (s *service) ValidateE911Address(ctx context.Context, address BaseAddress) ([]BaseAddress, error) {
	req := request{
		ValidateE911Address: &validateE911Address{
			Authentication: s.Authentication,
			Address:        address,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.ValidateE911Address.Return, nil
}

// ViewNumberDetails was auto-generated from WSDL.
func (s *service) ViewNumberDetails(ctx context.Context, numbers []string) ([]OrderNumber, error) {
	req := request{
		ViewNumberDetails: &viewNumberDetails{
			Authentication:  s.Authentication,
			TelephoneNumber: numbers,
		},
	}
	res, err := s.call(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.ViewNumberDetails.Return, nil
}
