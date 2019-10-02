// Package peerless provides a SOAP HTTP client.
package peerless

import (
	"net/http"
	"reflect"
)

const (
	// Envelope is a SOAP envelope.
	Namespace          = "http://publicapi.api.s2.peerless.com/"
	// ProductionEndpoint endpoint used in production
	ProductionEndpoint = "https://animate.peerlessnetwork.com:8181/animateapi/axis/APIService"
	// StagingEndpoint endpoint used for development
	StagingEndpoint    = "https://aniweb02.peerlessnetwork.com:8181/animateapi/axis/APIService"
)

// APIService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type APIService interface {

	// ActivateSOA godoc
	ActivateSOA(ActivateSOA *ActivateSOA) (*ActivateSOAResponse, error)

	// AddNotes godoc
	AddNotes(AddNotes *AddNotes) (*AddNotesResponse, error)

	// CreateException godoc
	CreateException(CreateException *CreateException) (*CreateExceptionResponse, error)

	// DisconnectOrder godoc
	DisconnectOrder(DisconnectOrder *DisconnectOrder) (*DisconnectOrderResponse, error)

	// Download godoc
	Download(Download *Download) (*DownloadResponse, error)

	// GetHierarchicalView godoc
	GetHierarchicalView(GetHierarchicalView *GetHierarchicalView) (*GetHierarchicalViewResponse, error)

	// GetNewNumberSearchFilters godoc
	GetNewNumberSearchFilters(GetNewNumberSearchFilters *GetNewNumberSearchFilters) (*GetNewNumberSearchFiltersResponse, error)

	// GetOrderStatus godoc
	GetOrderStatus(orderID string) (*GetOrderStatusResponse, error)

	// GetOrdersByPONSearch godoc
	GetOrdersByPONSearch(GetOrdersByPONSearch *GetOrdersByPONSearch) (*GetOrdersByPONSearchResponse, error)

	// GetPortInRelatedOrders godoc
	GetPortInRelatedOrders(GetPortInRelatedOrders *GetPortInRelatedOrders) (*GetPortInRelatedOrdersResponse, error)

	// GetStatusByNumberSearch godoc
	GetStatusByNumberSearch(GetStatusByNumberSearch *GetStatusByNumberSearch) (*GetStatusByNumberSearchResponse, error)

	// GetTnInventoryReport godoc
	GetTnInventoryReport(GetTnInventoryReport *GetTnInventoryReport) (*GetTnInventoryReportResponse, error)

	// PlaceOrder was godoc
	PlaceOrder(PlaceOrder *PlaceOrder) (*PlaceOrderResponse, error)

	// PlaceTFDisconnectOrder godoc
	PlaceTFDisconnectOrder(PlaceTFDisconnectOrder *PlaceTFDisconnectOrder) (*PlaceTFDisconnectOrderResponse, error)

	// PlaceTFOrder godoc
	PlaceTFOrder(PlaceTFOrder *PlaceTFOrder) (*PlaceTFOrderResponse, error)

	// PortabilityCheck godoc
	PortabilityCheck(PortabilityCheck *PortabilityCheck) (*PortabilityCheckResponse, error)

	// SearchNumbers godoc
	SearchNumbers(SearchNumbers *SearchNumbers) (*SearchNumbersResponse, error)

	// SearchOrderDetailsByOrderId godoc
	SearchOrderDetailsByOrderId(SearchOrderDetailsByOrderId *SearchOrderDetailsByOrderId) (*SearchOrderDetailsByOrderIdResponse, error)

	// SupplementOrder godoc
	SupplementOrder(SupplementOrder *SupplementOrder) (*SupplementOrderResponse, error)

	// Upload godoc
	Upload(Upload *Upload) (*UploadResponse, error)

	// ValidateE911Address godoc
	ValidateE911Address(address BaseAddress) (*ValidateE911AddressResponse, error)

	// ViewNumberDetails godoc
	ViewNumberDetails(numbers []string) (*ViewNumberDetailsResponse, error)
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
func (s *service) ActivateSOA(ActivateSOA *ActivateSOA) (*ActivateSOAResponse, error) {
	req := Request{
		ActivateSOA: ActivateSOA,
	}
	res,err := s.call("activateSOA", req);
	if  err != nil {
		return nil, err
	}
	return res.ActivateSOA, nil
}

// AddNotes was auto-generated from WSDL.
func (s *service) AddNotes(AddNotes *AddNotes) (*AddNotesResponse, error) {
	req := Request{
		AddNotes: AddNotes,
	}
	res,err := s.call("addNotes", req);
	if err != nil {
		return nil, err
	}
	return res.AddNotes, nil
}

// CreateException was auto-generated from WSDL.
func (s *service) CreateException(CreateException *CreateException) (*CreateExceptionResponse, error) {
	req := Request{
		CreateException: CreateException,
	}
	res,err := s.call("createException", req);
	if  err != nil {
		return nil, err
	}
	return res.CreateException, nil
}

// DisconnectOrder was auto-generated from WSDL.
func (s *service) DisconnectOrder(DisconnectOrder *DisconnectOrder) (*DisconnectOrderResponse, error) {
	req := Request{
		DisconnectOrder: DisconnectOrder,
	}
	res,err := s.call("disconnectOrder", req);
	if err != nil {
		return nil, err
	}
	return res.DisconnectOrder, nil
}

// Download was auto-generated from WSDL.
func (s *service) Download(Download *Download) (*DownloadResponse, error) {
	req := Request{
		Download: Download,
	}

	res,err := s.call("download", req);
	if  err != nil {
		return nil, err
	}
	return res.Download, nil
}

// GetHierarchicalView was auto-generated from WSDL.
func (s *service) GetHierarchicalView(GetHierarchicalView *GetHierarchicalView) (*GetHierarchicalViewResponse, error) {
	req := Request{
		GetHierarchicalView: GetHierarchicalView,
	}
	res,err := s.call("getHierarchicalView", req);
	if  err != nil {
		return nil, err
	}
	return res.GetHierarchicalView, nil
}

// GetNewNumberSearchFilters was auto-generated from WSDL.
func (s *service) GetNewNumberSearchFilters(GetNewNumberSearchFilters *GetNewNumberSearchFilters) (*GetNewNumberSearchFiltersResponse, error) {
	req := Request{
		GetNewNumberSearchFilters: GetNewNumberSearchFilters,
	}
	res,err := s.call("getNewNumberSearchFilters", req);
	if err != nil {
		return nil, err
	}
	return res.GetNewNumberSearchFilters, nil
}

// GetOrderStatus was auto-generated from WSDL.
func (s *service) GetOrderStatus(orderID string) (*GetOrderStatusResponse, error) {
	req := Request{
		GetOrderStatus: &GetOrderStatus{
			OrderId: orderID,
		},
	}
	res,err := s.call("getOrderStatus", req);
	if err != nil {
		return nil, err
	}
	return res.GetOrderStatus, nil
}

// GetOrdersByPONSearch was auto-generated from WSDL.
func (s *service) GetOrdersByPONSearch(GetOrdersByPONSearch *GetOrdersByPONSearch) (*GetOrdersByPONSearchResponse, error) {
	req := Request{
		GetOrdersByPONSearch: GetOrdersByPONSearch,
	}

	res,err := s.call("getOrdersByPONSearch", req);
	if  err != nil {
		return nil, err
	}
	return res.GetOrdersByPONSearch, nil
}

// GetPortInRelatedOrders was auto-generated from WSDL.
func (s *service) GetPortInRelatedOrders(GetPortInRelatedOrders *GetPortInRelatedOrders) (*GetPortInRelatedOrdersResponse, error) {
	req := Request{
		GetPortInRelatedOrders: GetPortInRelatedOrders,
	}
	res,err := s.call("getPortInRelatedOrders", req);
	if err != nil {
		return nil, err
	}
	return res.GetPortInRelatedOrders, nil
}

// GetStatusByNumberSearch was auto-generated from WSDL.
func (s *service) GetStatusByNumberSearch(GetStatusByNumberSearch *GetStatusByNumberSearch) (*GetStatusByNumberSearchResponse, error) {
	req := Request{
		GetStatusByNumberSearch: GetStatusByNumberSearch,
	}
	res,err := s.call("getStatusByNumberSearch", req);
	if err != nil {
		return nil, err
	}
	return res.GetStatusByNumberSearch, nil
}

// GetTnInventoryReport was auto-generated from WSDL.
func (s *service) GetTnInventoryReport(GetTnInventoryReport *GetTnInventoryReport) (*GetTnInventoryReportResponse, error) {
	req := Request{
		GetTnInventoryReport: GetTnInventoryReport,
	}
	res,err := s.call("getTnInventoryReport", req);
	if  err != nil {
		return nil, err
	}
	return res.GetTnInventoryReport, nil
}

// PlaceOrder was auto-generated from WSDL.
func (s *service) PlaceOrder(PlaceOrder *PlaceOrder) (*PlaceOrderResponse, error) {
	req := Request{
		PlaceOrder: PlaceOrder,
	}
	res,err := s.call("placeOrder", req);
	if  err != nil {
		return nil, err
	}
	return res.PlaceOrder, nil
}

// PlaceTFDisconnectOrder was auto-generated from WSDL.
func (s *service) PlaceTFDisconnectOrder(PlaceTFDisconnectOrder *PlaceTFDisconnectOrder) (*PlaceTFDisconnectOrderResponse, error) {
	req := Request{
		PlaceTFDisconnectOrder: PlaceTFDisconnectOrder,
	}
	res,err := s.call("placeTFDisconnectOrder", req);
	if  err != nil {
		return nil, err
	}
	return res.PlaceTFDisconnectOrder, nil
}

// PlaceTFOrder was auto-generated from WSDL.
func (s *service) PlaceTFOrder(PlaceTFOrder *PlaceTFOrder) (*PlaceTFOrderResponse, error) {
	req := Request{
		PlaceTFOrder: PlaceTFOrder,
	}
	res,err := s.call("placeTFOrder", req);
	if err != nil {
		return nil, err
	}
	return res.PlaceTFOrder, nil
}

// PortabilityCheck was auto-generated from WSDL.
func (s *service) PortabilityCheck(PortabilityCheck *PortabilityCheck) (*PortabilityCheckResponse, error) {
	req := Request{
		PortabilityCheck: PortabilityCheck,
	}
	res,err := s.call("portabilityCheck", req);
	if  err != nil {
		return nil, err
	}
	return res.PortabilityCheck, nil
}

// SearchNumbers was auto-generated from WSDL.
func (s *service) SearchNumbers(SearchNumbers *SearchNumbers) (*SearchNumbersResponse, error) {
	req := Request{
		SearchNumbers: SearchNumbers,
	}
	res,err := s.call("searchNumbers", req);
	if err != nil {
		return nil, err
	}
	return res.SearchNumbers, nil
}

// SearchOrderDetailsByOrderId was auto-generated from WSDL.
func (s *service) SearchOrderDetailsByOrderId(SearchOrderDetailsByOrderId *SearchOrderDetailsByOrderId) (*SearchOrderDetailsByOrderIdResponse, error) {
	req := Request{
		SearchOrderDetailsByOrderId: SearchOrderDetailsByOrderId,
	}
	res,err := s.call("searchOrderDetailsByOrderId", req);
	if err != nil {
		return nil, err
	}
	return res.SearchOrderDetailsByOrderId, nil
}

// SupplementOrder was auto-generated from WSDL.
func (s *service) SupplementOrder(SupplementOrder *SupplementOrder) (*SupplementOrderResponse, error) {
	req := Request{
		SupplementOrder: SupplementOrder,
	}
	res,err := s.call("supplementOrder", req);
	if err != nil {
		return nil, err
	}
	return res.SupplementOrder, nil
}

// Upload was auto-generated from WSDL.
func (s *service) Upload(Upload *Upload) (*UploadResponse, error) {
	req := Request{
		Upload: Upload,
	}
	res,err := s.call("upload", req);
	if err != nil {
		return nil, err
	}
	return res.Upload, nil
}

// ValidateE911Address was auto-generated from WSDL.
func (s *service) ValidateE911Address(address BaseAddress) (*ValidateE911AddressResponse, error) {
	req := Request{
		ValidateE911Address: &ValidateE911Address{
			Authentication: s.Authentication,
			Address:        address,
		},
	}
	res,err := s.call("validateE911Address", req);
	if  err != nil {
		return nil, err
	}
	return res.ValidateE911Address, nil
}

// ViewNumberDetails was auto-generated from WSDL.
func (s *service) ViewNumberDetails(numbers []string) (*ViewNumberDetailsResponse, error) {
	req := Request{
		ViewNumberDetails: &ViewNumberDetails{
			Authentication:  s.Authentication,
			TelephoneNumber: numbers,
		},
	}
	res,err := s.call("viewNumberDetails", req);
	if err != nil {
		return nil, err
	}
	return res.ViewNumberDetails, nil
}
