// Package peerless provides a SOAP HTTP client.
package peerless

import (
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

	// ActivateSOA godoc
	ActivateSOA(orderID []string) (bool, error)

	// AddNotes godoc
	AddNotes(orderID, note string) (bool, error)

	// CreateException godoc
	CreateException(ExceptionNote *ExceptionNote) (bool, error)

	// DisconnectOrder godoc
	DisconnectOrder(DisconnectOrderRequest *DisconnectOrderRequest) (string, error)

	// Download godoc
	Download(orderID string) ([]Attachment, error)

	// GetHierarchicalView godoc
	GetHierarchicalView() (*HiearchicalView, error)

	// GetNewNumberSearchFilters godoc
	GetNewNumberSearchFilters(filters *NumberSearchParameters) (*NumberSearchParameters, error)

	// GetOrderStatus godoc
	GetOrderStatus(orderID string) (string, error)

	// GetOrdersByPONSearch godoc
	GetOrdersByPONSearch(pon string) (*PonOrders, error)

	// GetPortInRelatedOrders godoc
	GetPortInRelatedOrders(orderID int64) (*RelatedOrders, error)

	// GetStatusByNumberSearch godoc
	GetStatusByNumberSearch(telephoneNumber []string) (*NumberStatus, error)

	// GetTnInventoryReport godoc
	GetTnInventoryReport(searchParams *TnInventoryForApiSearchParams) ([]TnInventory, error)

	// PlaceOrder was godoc
	PlaceOrder(order *Order) (string, error)

	// PlaceTFDisconnectOrder godoc
	PlaceTFDisconnectOrder(disconnectOrderRequest *DisconnectOrderRequest) (string, error)

	// PlaceTFOrder godoc
	PlaceTFOrder(order *TollFreeOrder) (string, error)

	// PortabilityCheck godoc
	PortabilityCheck(portabilityCheckRequest *PortabilityCheckRequest) ([]string, error)

	// SearchNumbers godoc
	SearchNumbers(filters *NumberSearchParameters) (*NumberSearchParameters, error)

	// SearchOrderDetailsByOrderId godoc
	SearchOrderDetailsByOrderId(orderID int64, orderType string) (*OrderSearch, error)

	// SupplementOrder godoc
	SupplementOrder(supplementInfo *SupplementInfo, order *Order) (string, error)

	// Upload godoc
	Upload(orderID string, attachments []Attachment) (bool, error)

	// ValidateE911Address godoc
	ValidateE911Address(address BaseAddress) ([]BaseAddress, error)

	// ViewNumberDetails godoc
	ViewNumberDetails(numbers []string) ([]OrderNumber , error)
}

// New creates an initializes a API service.
func New(url, customer, passcode, userid string) APIService {
	return &service{
		URL:       url,
		Namespace: Namespace,
		Authentication: &AuthInfo{
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
	Namespace              string       // SOAP Namespace
	URNamespace            string       // Uniform Resource Namespace
	ThisNamespace          string       // SOAP This-Namespace (tns)
	ExcludeActionNamespace bool         // Include Namespace to SOAP Action header
	Envelope               string       // Optional SOAP Envelope
	Header                 interface{}  // Optional SOAP Header
	ContentType            string       // Optional Content-Type (default text/xml)
	Config                 *http.Client // Optional HTTP client
	Authentication         *AuthInfo    // Authentication
}

// ActivateSOA was auto-generated from WSDL.
func (s *service) ActivateSOA(orderID []string) (bool, error) {
	req := request{
		ActivateSOA: &activateSOA{
			Authentication: s.Authentication,
			OrderID:        orderID,
		},
	}
	res, err := s.call("activateSOA", req)
	if err != nil {
		return false, err
	}
	return res.ActivateSOA.Return, nil
}

// AddNotes was auto-generated from WSDL.
func (s *service) AddNotes(orderID, note string) (bool, error) {
	req := request{
		AddNotes: &addNotes{
			Authentication: s.Authentication,
			OrderID:        orderID,
			Notes:          note,
		},
	}
	res, err := s.call("addNotes", req)
	if err != nil {
		return false, err
	}
	return res.AddNotes.Return, nil
}

// CreateException was auto-generated from WSDL.
func (s *service) CreateException(ExceptionNote *ExceptionNote) (bool, error) {
	req := request{
		CreateException: &createException{
			Authentication: s.Authentication,
			ExceptionNote:  ExceptionNote,
		},
	}
	res, err := s.call("createException", req)
	if err != nil {
		return false, err
	}
	return res.CreateException.Return, nil
}

// DisconnectOrder was auto-generated from WSDL.
func (s *service) DisconnectOrder(DisconnectOrderRequest *DisconnectOrderRequest) (string, error) {
	req := request{
		DisconnectOrder: &disconnectOrder{
			Authentication:         s.Authentication,
			DisconnectOrderRequest: DisconnectOrderRequest,
		},
	}
	res, err := s.call("disconnectOrder", req)
	if err != nil {
		return "", err
	}
	return res.DisconnectOrder.Return, nil
}

// Download was auto-generated from WSDL.
func (s *service) Download(orderID string) ([]Attachment, error) {
	req := request{
		Download: &download{
			Authentication: s.Authentication,
			OrderID:        orderID,
		},
	}
	res, err := s.call("download", req)
	if err != nil {
		return nil, err
	}
	return res.Download.Return, nil
}

// GetHierarchicalView was auto-generated from WSDL.
func (s *service) GetHierarchicalView() (*HiearchicalView, error) {
	req := request{
		GetHierarchicalView: &getHierarchicalView{
			Authentication: s.Authentication,
		},
	}
	res, err := s.call("getHierarchicalView", req)
	if err != nil {
		return nil, err
	}
	return res.GetHierarchicalView.Return, nil
}

// GetNewNumberSearchFilters was auto-generated from WSDL.
func (s *service) GetNewNumberSearchFilters(filters *NumberSearchParameters) (*NumberSearchParameters, error) {
	req := request{
		GetNewNumberSearchFilters: &getNewNumberSearchFilters{
			Authentication: s.Authentication,
			Filters:        filters,
		},
	}
	res, err := s.call("getNewNumberSearchFilters", req)
	if err != nil {
		return nil, err
	}
	return res.GetNewNumberSearchFilters.Return, nil
}

// GetOrderStatus was auto-generated from WSDL.
func (s *service) GetOrderStatus(orderID string) (string, error) {
	req := request{
		GetOrderStatus: &getOrderStatus{
			OrderID: orderID,
		},
	}
	res, err := s.call("getOrderStatus", req)
	if err != nil {
		return "", err
	}
	return res.GetOrderStatus.Return, nil
}

// GetOrdersByPONSearch was auto-generated from WSDL.
func (s *service) GetOrdersByPONSearch(pon string) (*PonOrders, error) {
	req := request{
		GetOrdersByPONSearch: &getOrdersByPONSearch{
			Authentication: s.Authentication,
			Pon:            pon,
		},
	}
	res, err := s.call("getOrdersByPONSearch", req)
	if err != nil {
		return nil, err
	}
	return res.GetOrdersByPONSearch.Return, nil
}

// GetPortInRelatedOrders was auto-generated from WSDL.
func (s *service) GetPortInRelatedOrders(orderID int64) (*RelatedOrders, error) {
	req := request{
		GetPortInRelatedOrders: &getPortInRelatedOrders{
			Authentication: s.Authentication,
			OrderID:        orderID,
		},
	}
	res, err := s.call("getPortInRelatedOrders", req)
	if err != nil {
		return nil, err
	}
	return res.GetPortInRelatedOrders.Return, nil
}

// GetStatusByNumberSearch was auto-generated from WSDL.
func (s *service) GetStatusByNumberSearch(telephoneNumber []string) (*NumberStatus, error) {
	req := request{
		GetStatusByNumberSearch: &getStatusByNumberSearch{
			Authentication:  s.Authentication,
			TelephoneNumber: telephoneNumber,
		},
	}
	res, err := s.call("getStatusByNumberSearch", req)
	if err != nil {
		return nil, err
	}
	return res.GetStatusByNumberSearch.Return, nil
}

// GetTnInventoryReport was auto-generated from WSDL.
func (s *service) GetTnInventoryReport(searchParams *TnInventoryForApiSearchParams) ([]TnInventory, error) {
	req := request{
		GetTnInventoryReport: &getTnInventoryReport{
			Authentication: s.Authentication,
			SearchParams:   searchParams,
		},
	}
	res, err := s.call("getTnInventoryReport", req)
	if err != nil {
		return nil, err
	}
	return res.GetTnInventoryReport.Return, nil
}

// PlaceOrder was auto-generated from WSDL.
func (s *service) PlaceOrder(order *Order) (string, error) {
	req := request{
		PlaceOrder: &placeOrder{
			Authentication: s.Authentication,
			Order:          order,
		},
	}
	res, err := s.call("placeOrder", req)
	if err != nil {
		return "", err
	}
	return res.PlaceOrder.Return, nil
}

// PlaceTFDisconnectOrder was auto-generated from WSDL.
func (s *service) PlaceTFDisconnectOrder(disconnectOrderRequest *DisconnectOrderRequest) (string, error) {
	req := request{
		PlaceTFDisconnectOrder: &placeTFDisconnectOrder{
			Authentication:         s.Authentication,
			DisconnectOrderRequest: disconnectOrderRequest,
		},
	}
	res, err := s.call("placeTFDisconnectOrder", req)
	if err != nil {
		return "", err
	}
	return res.PlaceTFDisconnectOrder.Return, nil
}

// PlaceTFOrder was auto-generated from WSDL.
func (s *service) PlaceTFOrder(order *TollFreeOrder) (string, error) {
	req := request{
		PlaceTFOrder: &placeTFOrder{
			Authentication: s.Authentication,
			TFNOrder:       order,
		},
	}
	res, err := s.call("placeTFOrder", req)
	if err != nil {
		return "", err
	}
	return res.PlaceTFOrder.Return, nil
}

// PortabilityCheck was auto-generated from WSDL.
func (s *service) PortabilityCheck(portabilityCheckRequest *PortabilityCheckRequest) ([]string, error) {
	req := request{
		PortabilityCheck: &portabilityCheck{
			Authentication:          s.Authentication,
			PortabilityCheckRequest: portabilityCheckRequest,
		},
	}
	res, err := s.call("portabilityCheck", req)
	if err != nil {
		return nil, err
	}
	return res.PortabilityCheck.Return.PortPsNumbers, nil
}

// SearchNumbers was auto-generated from WSDL.
func (s *service) SearchNumbers(filters *NumberSearchParameters) (*NumberSearchParameters, error) {
	req := request{
		SearchNumbers: &searchNumbers{
			Authentication: s.Authentication,
			Filters:        filters,
		},
	}
	res, err := s.call("searchNumbers", req)
	if err != nil {
		return nil, err
	}
	return res.SearchNumbers.Return, nil
}

// SearchOrderDetailsByOrderId was auto-generated from WSDL.
func (s *service) SearchOrderDetailsByOrderId(orderID int64, orderType string) (*OrderSearch, error) {
	req := request{
		SearchOrderDetailsByOrderId: &searchOrderDetailsByOrderId{
			Authentication: s.Authentication,
			OrderID:        orderID,
			OrderType:      orderType,
		},
	}
	res, err := s.call("searchOrderDetailsByOrderId", req)
	if err != nil {
		return nil, err
	}
	return res.SearchOrderDetailsByOrderId.Return, nil
}

// SupplementOrder was auto-generated from WSDL.
func (s *service) SupplementOrder(supplementInfo *SupplementInfo, order *Order) (string, error) {
	req := request{
		SupplementOrder: &supplementOrder{
			Authentication: s.Authentication,
			SupplementInfo: supplementInfo,
			Order:          order,
		},
	}
	res, err := s.call("supplementOrder", req)
	if err != nil {
		return "", err
	}
	return res.SupplementOrder.Return, nil
}

// Upload was auto-generated from WSDL.
func (s *service) Upload(orderID string, attachments []Attachment) (bool, error) {
	req := request{
		Upload: &upload{
			Authentication: s.Authentication,
			OrderID:        orderID,
			Attachments:    attachments,
		},
	}
	res, err := s.call("upload", req)
	if err != nil {
		return false, err
	}
	return res.Upload.Return, nil
}

// ValidateE911Address was auto-generated from WSDL.
func (s *service) ValidateE911Address(address BaseAddress) ([]BaseAddress, error) {
	req := request{
		ValidateE911Address: &validateE911Address{
			Authentication: s.Authentication,
			Address:        address,
		},
	}
	res, err := s.call("validateE911Address", req)
	if err != nil {
		return nil, err
	}
	return res.ValidateE911Address.Return, nil
}

// ViewNumberDetails was auto-generated from WSDL.
func (s *service) ViewNumberDetails(numbers []string) ([]OrderNumber , error) {
	req := request{
		ViewNumberDetails: &viewNumberDetails{
			Authentication:  s.Authentication,
			TelephoneNumber: numbers,
		},
	}
	res, err := s.call("viewNumberDetails", req)
	if err != nil {
		return nil, err
	}
	return res.ViewNumberDetails.Return, nil
}
