package peerless

import (
	"encoding/xml"
)

type requestBody struct {
	ActivateSOA                 *activateSOA                 `xml:"ns:activateSOA,omitempty" json:"activateSOA,omitempty" yaml:"activateSOA,omitempty"`
	AddNotes                    *addNotes                    `xml:"ns:addNotes,omitempty" json:"addNotes,omitempty" yaml:"addNotes,omitempty"`
	CreateException             *createException             `xml:"ns:createException,omitempty" json:"createException,omitempty" yaml:"createException,omitempty"`
	DisconnectOrder             *disconnectOrder             `xml:"ns:disconnectOrder,omitempty" json:"disconnectOrder,omitempty" yaml:"disconnectOrder,omitempty"`
	Download                    *download                    `xml:"ns:download,omitempty" json:"download,omitempty" yaml:"download,omitempty"`
	GetHierarchicalView         *getHierarchicalView         `xml:"ns:getHierarchicalView,omitempty" json:"getHierarchicalView,omitempty" yaml:"getHierarchicalView,omitempty"`
	GetNewNumberSearchFilters   *getNewNumberSearchFilters   `xml:"ns:getNewNumberSearchFilters,omitempty" json:"getNewNumberSearchFilters,omitempty" yaml:"getNewNumberSearchFilters,omitempty"`
	GetOrderStatus              *getOrderStatus              `xml:"ns:getOrderStatus,omitempty" json:"getOrderStatus,omitempty" yaml:"getOrderStatus,omitempty"`
	GetOrdersByPONSearch        *getOrdersByPONSearch        `xml:"ns:getOrdersByPONSearch,omitempty" json:"getOrdersByPONSearch,omitempty" yaml:"getOrdersByPONSearch,omitempty"`
	GetPortInRelatedOrders      *getPortInRelatedOrders      `xml:"ns:getPortInRelatedOrders,omitempty" json:"getPortInRelatedOrders,omitempty" yaml:"getPortInRelatedOrders,omitempty"`
	GetStatusByNumberSearch     *getStatusByNumberSearch     `xml:"ns:getStatusByNumberSearch,omitempty" json:"getStatusByNumberSearch,omitempty" yaml:"getStatusByNumberSearch,omitempty"`
	GetTnInventoryReport        *getTnInventoryReport        `xml:"ns:getTnInventoryReport,omitempty" json:"getTnInventoryReport,omitempty" yaml:"getTnInventoryReport,omitempty"`
	PlaceOrder                  *placeOrder                  `xml:"ns:placeOrder,omitempty" json:"placeOrder,omitempty" yaml:"placeOrder,omitempty"`
	PlaceTFDisconnectOrder      *placeTFDisconnectOrder      `xml:"ns:placeTFDisconnectOrder,omitempty" json:"placeTFDisconnectOrder,omitempty" yaml:"placeTFDisconnectOrder,omitempty"`
	PlaceTFOrder                *placeTFOrder                `xml:"ns:placeTFOrder,omitempty" json:"placeTFOrder,omitempty" yaml:"placeTFOrder,omitempty"`
	PortabilityCheck            *portabilityCheck            `xml:"ns:portabilityCheck,omitempty" json:"portabilityCheck,omitempty" yaml:"portabilityCheck,omitempty"`
	SearchNumbers               *searchNumbers               `xml:"ns:searchNumbers,omitempty" json:"searchNumbers,omitempty" yaml:"searchNumbers,omitempty"`
	SearchOrderDetailsByOrderId *searchOrderDetailsByOrderId `xml:"ns:searchOrderDetailsByOrderId,omitempty" json:"searchOrderDetailsByOrderId,omitempty" yaml:"searchOrderDetailsByOrderId,omitempty"`
	SupplementOrder             *supplementOrder             `xml:"ns:supplementOrder,omitempty" json:"supplementOrder,omitempty" yaml:"supplementOrder,omitempty"`
	Upload                      *upload                      `xml:"ns:upload,omitempty" json:"upload,omitempty" yaml:"upload,omitempty"`
	ValidateE911Address         *validateE911Address         `xml:"ns:validateE911Address,omitempty" json:"validateE911Address,omitempty" yaml:"validateE911Address,omitempty"`
	ViewNumberDetails           *viewNumberDetails           `xml:"ns:viewNumberDetails,omitempty" json:"viewNumberDetails,omitempty" yaml:"viewNumberDetails,omitempty"`
}

type responseBody struct {
	XMLName                     xml.Name                             `xml:"Body"`
	ActivateSOA                 *activateSOAResponse                 `xml:"activateSOAResponse,omitempty" json:"activateSOAResponse,omitempty" yaml:"activateSOAResponse,omitempty"`
	AddNotes                    *addNotesResponse                    `xml:"addNotesResponse,omitempty" json:"addNotesResponse,omitempty" yaml:"addNotesResponse,omitempty"`
	CreateException             *createExceptionResponse             `xml:"createExceptionResponse,omitempty" json:"createExceptionResponse,omitempty" yaml:"createExceptionResponse,omitempty"`
	DisconnectOrder             *disconnectOrderResponse             `xml:"disconnectOrderResponse,omitempty" json:"disconnectOrderResponse,omitempty" yaml:"disconnectOrderResponse,omitempty"`
	Download                    *downloadResponse                    `xml:"downloadResponse,omitempty" json:"downloadResponse,omitempty" yaml:"downloadResponse,omitempty"`
	GetHierarchicalView         *getHierarchicalViewResponse         `xml:"getHierarchicalViewResponse,omitempty" json:"getHierarchicalViewResponse,omitempty" yaml:"getHierarchicalViewResponse,omitempty"`
	GetNewNumberSearchFilters   *getNewNumberSearchFiltersResponse   `xml:"getNewNumberSearchFiltersResponse,omitempty" json:"getNewNumberSearchFiltersResponse,omitempty" yaml:"getNewNumberSearchFiltersResponse,omitempty"`
	GetOrderStatus              *getOrderStatusResponse              `xml:"getOrderStatusResponse,omitempty" json:"getOrderStatusResponse,omitempty" yaml:"getOrderStatusResponse,omitempty"`
	GetOrdersByPONSearch        *getOrdersByPONSearchResponse        `xml:"getOrdersByPONSearchResponse,omitempty" json:"getOrdersByPONSearchResponse,omitempty" yaml:"getOrdersByPONSearchResponse,omitempty"`
	GetPortInRelatedOrders      *getPortInRelatedOrdersResponse      `xml:"getPortInRelatedOrdersResponse,omitempty" json:"getPortInRelatedOrdersResponse,omitempty" yaml:"getPortInRelatedOrdersResponse,omitempty"`
	GetStatusByNumberSearch     *getStatusByNumberSearchResponse     `xml:"getStatusByNumberSearchResponse,omitempty" json:"getStatusByNumberSearchResponse,omitempty" yaml:"getStatusByNumberSearchResponse,omitempty"`
	GetTnInventoryReport        *getTnInventoryReportResponse        `xml:"getTnInventoryReportResponse,omitempty" json:"getTnInventoryReportResponse,omitempty" yaml:"getTnInventoryReportResponse,omitempty"`
	PlaceOrder                  *placeOrderResponse                  `xml:"placeOrderResponse,omitempty" json:"placeOrderResponse,omitempty" yaml:"placeOrderResponse,omitempty"`
	PlaceTFDisconnectOrder      *placeTFDisconnectOrderResponse      `xml:"placeTFDisconnectOrderResponse,omitempty" json:"placeTFDisconnectOrderResponse,omitempty" yaml:"placeTFDisconnectOrderResponse,omitempty"`
	PlaceTFOrder                *placeTFOrderResponse                `xml:"placeTFOrderResponse,omitempty" json:"placeTFOrderResponse,omitempty" yaml:"placeTFOrderResponse,omitempty"`
	PortabilityCheck            *portabilityCheckResponse            `xml:"portabilityCheckResponse,omitempty" json:"portabilityCheckResponse,omitempty" yaml:"portabilityCheckResponse,omitempty"`
	SearchNumbers               *searchNumbersResponse               `xml:"searchNumbersResponse,omitempty" json:"searchNumbersResponse,omitempty" yaml:"searchNumbersResponse,omitempty"`
	SearchOrderDetailsByOrderId *searchOrderDetailsByOrderIdResponse `xml:"searchOrderDetailsByOrderIdResponse,omitempty" json:"searchOrderDetailsByOrderIdResponse,omitempty" yaml:"searchOrderDetailsByOrderIdResponse,omitempty"`
	SupplementOrder             *supplementOrderResponse             `xml:"supplementOrderResponse,omitempty" json:"supplementOrderResponse,omitempty" yaml:"supplementOrderResponse,omitempty"`
	Upload                      *uploadResponse                      `xml:"uploadResponse,omitempty" json:"uploadResponse,omitempty" yaml:"uploadResponse,omitempty"`
	ValidateE911Address         *validateE911AddressResponse         `xml:"validateE911AddressResponse,omitempty" json:"validateE911AddressResponse,omitempty" yaml:"validateE911AddressResponse,omitempty"`
	ViewNumberDetails           *viewNumberDetailsResponse           `xml:"viewNumberDetailsResponse,omitempty" json:"viewNumberDetailsResponse,omitempty" yaml:"viewNumberDetailsResponse,omitempty"`
	Fault                       *Fault                               `xml:",omitempty"`
}

type activateSOA struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        []string  `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
}

type activateSOAResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type addNotes struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        string    `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
	Notes          string    `xml:"notes,omitempty" json:"notes,omitempty" yaml:"notes,omitempty"`
}

type addNotesResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type createException struct {
	Authentication *authInfo      `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	ExceptionNote  *ExceptionNote `xml:"excpetionNote,omitempty" json:"excpetionNote,omitempty" yaml:"excpetionNote,omitempty"`
}

type createExceptionResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type disconnectOrder struct {
	Authentication         *authInfo               `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	DisconnectOrderRequest *DisconnectOrderRequest `xml:"disconnectOrderRequest,omitempty" json:"disconnectOrderRequest,omitempty" yaml:"disconnectOrderRequest,omitempty"`
}

type disconnectOrderResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type download struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        string    `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
}

type downloadResponse struct {
	Return []Attachment `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getHierarchicalView struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
}

type getHierarchicalViewResponse struct {
	Return *HiearchicalView `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getNewNumberSearchFilters struct {
	Authentication *authInfo               `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Filters        *NumberSearchParameters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

type getNewNumberSearchFiltersResponse struct {
	Return *NumberSearchParameters `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getOrderStatus struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        string    `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
}

type getOrderStatusResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getOrdersByPONSearch struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Pon            string    `xml:"pon,omitempty" json:"pon,omitempty" yaml:"pon,omitempty"`
}

type getOrdersByPONSearchResponse struct {
	Return struct {
		Result []ResultPONOrderDetails `xml:"result>ponOrders,omitempty" json:"result>ponOrders,omitempty" yaml:"result>ponOrders,omitempty"`
	} `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getPortInRelatedOrders struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        int64     `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
}

type getPortInRelatedOrdersResponse struct {
	Return struct {
		Result []ResultOrderDetails `xml:"result>relatedOrders,omitempty" json:"result>relatedOrders,omitempty" yaml:"result>relatedOrders,omitempty"`
	} `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getStatusByNumberSearch struct {
	Authentication  *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	TelephoneNumber []string  `xml:"telephoneNumber,omitempty" json:"telephoneNumber,omitempty" yaml:"telephoneNumber,omitempty"`
}

type getStatusByNumberSearchResponse struct {
	Return struct {
		Result []string `xml:"result>entry,omitempty" json:"result>entry,omitempty" yaml:"result>entry,omitempty"`
	} `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type getTnInventoryReport struct {
	Authentication *authInfo                      `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	SearchParams   *TnInventoryForApiSearchParams `xml:"searchParams,omitempty" json:"searchParams,omitempty" yaml:"searchParams,omitempty"`
}

type getTnInventoryReportResponse struct {
	Return []TnInventory `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type placeOrder struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Order          *Order    `xml:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty"`
}

type placeOrderResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type placeTFDisconnectOrder struct {
	Authentication         *authInfo               `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	DisconnectOrderRequest *DisconnectOrderRequest `xml:"disconnectOrderRequest,omitempty" json:"disconnectOrderRequest,omitempty" yaml:"disconnectOrderRequest,omitempty"`
}

type placeTFDisconnectOrderResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type placeTFOrder struct {
	Authentication *authInfo      `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	TFNOrder       *TollFreeOrder `xml:"TFNOrder,omitempty" json:"TFNOrder,omitempty" yaml:"TFNOrder,omitempty"`
}

type placeTFOrderResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type portabilityCheck struct {
	Authentication          *authInfo                `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	PortabilityCheckRequest *PortabilityCheckRequest `xml:"portabilityCheckRequest,omitempty" json:"portabilityCheckRequest,omitempty" yaml:"portabilityCheckRequest,omitempty"`
}

type portabilityCheckResponse struct {
	Return struct {
		PortPsNumbers []string `xml:"portPsNumbers,omitempty" json:"portPsNumbers,omitempty" yaml:"portPsNumbers,omitempty"`
	} `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type searchNumbers struct {
	Authentication *authInfo               `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Filters        *NumberSearchParameters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

type searchNumbersResponse struct {
	Return *NumberSearchParameters `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type searchOrderDetailsByOrderId struct {
	Authentication *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        int64     `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
	OrderType      string    `xml:"orderType,omitempty" json:"orderType,omitempty" yaml:"orderType,omitempty"`
}

type searchOrderDetailsByOrderIdResponse struct {
	Return *OrderSearch `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type supplementOrder struct {
	Authentication *authInfo       `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	SupplementInfo *SupplementInfo `xml:"supplementInfo,omitempty" json:"supplementInfo,omitempty" yaml:"supplementInfo,omitempty"`
	Order          *Order          `xml:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty"`
}

type supplementOrderResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type upload struct {
	Authentication *authInfo    `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	OrderID        string       `xml:"orderId,omitempty" json:"orderId,omitempty" yaml:"orderId,omitempty"`
	Attachments    []Attachment `xml:"attachments,omitempty" json:"attachments,omitempty" yaml:"attachments,omitempty"`
}

type uploadResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type validateE911Address struct {
	Authentication *authInfo   `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	Address        BaseAddress `xml:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
}

type validateE911AddressResponse struct {
	Return []BaseAddress `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

type viewNumberDetails struct {
	Authentication  *authInfo `xml:"authentication,omitempty" json:"authentication,omitempty" yaml:"authentication,omitempty"`
	TelephoneNumber []string  `xml:"telephoneNumber,omitempty" json:"telephoneNumber,omitempty" yaml:"telephoneNumber,omitempty"`
}

type viewNumberDetailsResponse struct {
	Return []OrderNumber `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}
