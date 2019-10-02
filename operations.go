package peerless

type Request struct {
	ActivateSOA                 *ActivateSOA                 `xml:"ns:activateSOA,omitempty" json:"activateSOA,omitempty" yaml:"activateSOA,omitempty"`
	AddNotes                    *AddNotes                    `xml:"ns:addNotes,omitempty" json:"addNotes,omitempty" yaml:"addNotes,omitempty"`
	CreateException             *CreateException             `xml:"ns:createException,omitempty" json:"createException,omitempty" yaml:"createException,omitempty"`
	DisconnectOrder             *DisconnectOrder             `xml:"ns:disconnectOrder,omitempty" json:"disconnectOrder,omitempty" yaml:"disconnectOrder,omitempty"`
	Download                    *Download                    `xml:"ns:download,omitempty" json:"download,omitempty" yaml:"download,omitempty"`
	GetHierarchicalView         *GetHierarchicalView         `xml:"ns:getHierarchicalView,omitempty" json:"getHierarchicalView,omitempty" yaml:"getHierarchicalView,omitempty"`
	GetNewNumberSearchFilters   *GetNewNumberSearchFilters   `xml:"ns:getNewNumberSearchFilters,omitempty" json:"getNewNumberSearchFilters,omitempty" yaml:"getNewNumberSearchFilters,omitempty"`
	GetOrderStatus              *GetOrderStatus              `xml:"ns:getOrderStatus,omitempty" json:"getOrderStatus,omitempty" yaml:"getOrderStatus,omitempty"`
	GetOrdersByPONSearch        *GetOrdersByPONSearch        `xml:"ns:getOrdersByPONSearch,omitempty" json:"getOrdersByPONSearch,omitempty" yaml:"getOrdersByPONSearch,omitempty"`
	GetPortInRelatedOrders      *GetPortInRelatedOrders      `xml:"ns:getPortInRelatedOrders,omitempty" json:"getPortInRelatedOrders,omitempty" yaml:"getPortInRelatedOrders,omitempty"`
	GetStatusByNumberSearch     *GetStatusByNumberSearch     `xml:"ns:getStatusByNumberSearch,omitempty" json:"getStatusByNumberSearch,omitempty" yaml:"getStatusByNumberSearch,omitempty"`
	GetTnInventoryReport        *GetTnInventoryReport        `xml:"ns:getTnInventoryReport,omitempty" json:"getTnInventoryReport,omitempty" yaml:"getTnInventoryReport,omitempty"`
	PlaceOrder                  *PlaceOrder                  `xml:"ns:placeOrder,omitempty" json:"placeOrder,omitempty" yaml:"placeOrder,omitempty"`
	PlaceTFDisconnectOrder      *PlaceTFDisconnectOrder      `xml:"ns:placeTFDisconnectOrder,omitempty" json:"placeTFDisconnectOrder,omitempty" yaml:"placeTFDisconnectOrder,omitempty"`
	PlaceTFOrder                *PlaceTFOrder                `xml:"ns:placeTFOrder,omitempty" json:"placeTFOrder,omitempty" yaml:"placeTFOrder,omitempty"`
	PortabilityCheck            *PortabilityCheck            `xml:"ns:portabilityCheck,omitempty" json:"portabilityCheck,omitempty" yaml:"portabilityCheck,omitempty"`
	SearchNumbers               *SearchNumbers               `xml:"ns:searchNumbers,omitempty" json:"searchNumbers,omitempty" yaml:"searchNumbers,omitempty"`
	SearchOrderDetailsByOrderId *SearchOrderDetailsByOrderId `xml:"ns:searchOrderDetailsByOrderId,omitempty" json:"searchOrderDetailsByOrderId,omitempty" yaml:"searchOrderDetailsByOrderId,omitempty"`
	SupplementOrder             *SupplementOrder             `xml:"ns:supplementOrder,omitempty" json:"supplementOrder,omitempty" yaml:"supplementOrder,omitempty"`
	Upload                      *Upload                      `xml:"ns:upload,omitempty" json:"upload,omitempty" yaml:"upload,omitempty"`
	ValidateE911Address         *ValidateE911Address         `xml:"ns:validateE911Address,omitempty" json:"validateE911Address,omitempty" yaml:"validateE911Address,omitempty"`
	ViewNumberDetails           *ViewNumberDetails           `xml:"ns:viewNumberDetails,omitempty" json:"viewNumberDetails,omitempty" yaml:"viewNumberDetails,omitempty"`
}

type Response struct {
	ActivateSOA                 *ActivateSOAResponse                 `xml:"activateSOAResponse,omitempty" json:"activateSOAResponse,omitempty" yaml:"activateSOAResponse,omitempty"`
	AddNotes                    *AddNotesResponse                    `xml:"addNotesResponse,omitempty" json:"addNotesResponse,omitempty" yaml:"addNotesResponse,omitempty"`
	CreateException             *CreateExceptionResponse             `xml:"createExceptionResponse,omitempty" json:"createExceptionResponse,omitempty" yaml:"createExceptionResponse,omitempty"`
	DisconnectOrder             *DisconnectOrderResponse             `xml:"disconnectOrderResponse,omitempty" json:"disconnectOrderResponse,omitempty" yaml:"disconnectOrderResponse,omitempty"`
	Download                    *DownloadResponse                    `xml:"downloadResponse,omitempty" json:"downloadResponse,omitempty" yaml:"downloadResponse,omitempty"`
	GetHierarchicalView         *GetHierarchicalViewResponse         `xml:"getHierarchicalViewResponse,omitempty" json:"getHierarchicalViewResponse,omitempty" yaml:"getHierarchicalViewResponse,omitempty"`
	GetNewNumberSearchFilters   *GetNewNumberSearchFiltersResponse   `xml:"getNewNumberSearchFiltersResponse,omitempty" json:"getNewNumberSearchFiltersResponse,omitempty" yaml:"getNewNumberSearchFiltersResponse,omitempty"`
	GetOrderStatus              *GetOrderStatusResponse              `xml:"getOrderStatusResponse,omitempty" json:"getOrderStatusResponse,omitempty" yaml:"getOrderStatusResponse,omitempty"`
	GetOrdersByPONSearch        *GetOrdersByPONSearchResponse        `xml:"getOrdersByPONSearchResponse,omitempty" json:"getOrdersByPONSearchResponse,omitempty" yaml:"getOrdersByPONSearchResponse,omitempty"`
	GetPortInRelatedOrders      *GetPortInRelatedOrdersResponse      `xml:"getPortInRelatedOrdersResponse,omitempty" json:"getPortInRelatedOrdersResponse,omitempty" yaml:"getPortInRelatedOrdersResponse,omitempty"`
	GetStatusByNumberSearch     *GetStatusByNumberSearchResponse     `xml:"getStatusByNumberSearchResponse,omitempty" json:"getStatusByNumberSearchResponse,omitempty" yaml:"getStatusByNumberSearchResponse,omitempty"`
	GetTnInventoryReport        *GetTnInventoryReportResponse        `xml:"getTnInventoryReportResponse,omitempty" json:"getTnInventoryReportResponse,omitempty" yaml:"getTnInventoryReportResponse,omitempty"`
	PlaceOrder                  *PlaceOrderResponse                  `xml:"placeOrderResponse,omitempty" json:"placeOrderResponse,omitempty" yaml:"placeOrderResponse,omitempty"`
	PlaceTFDisconnectOrder      *PlaceTFDisconnectOrderResponse      `xml:"placeTFDisconnectOrderResponse,omitempty" json:"placeTFDisconnectOrderResponse,omitempty" yaml:"placeTFDisconnectOrderResponse,omitempty"`
	PlaceTFOrder                *PlaceTFOrderResponse                `xml:"placeTFOrderResponse,omitempty" json:"placeTFOrderResponse,omitempty" yaml:"placeTFOrderResponse,omitempty"`
	PortabilityCheck            *PortabilityCheckResponse            `xml:"portabilityCheckResponse,omitempty" json:"portabilityCheckResponse,omitempty" yaml:"portabilityCheckResponse,omitempty"`
	SearchNumbers               *SearchNumbersResponse               `xml:"searchNumbersResponse,omitempty" json:"searchNumbersResponse,omitempty" yaml:"searchNumbersResponse,omitempty"`
	SearchOrderDetailsByOrderId *SearchOrderDetailsByOrderIdResponse `xml:"searchOrderDetailsByOrderIdResponse,omitempty" json:"searchOrderDetailsByOrderIdResponse,omitempty" yaml:"searchOrderDetailsByOrderIdResponse,omitempty"`
	SupplementOrder             *SupplementOrderResponse             `xml:"supplementOrderResponse,omitempty" json:"supplementOrderResponse,omitempty" yaml:"supplementOrderResponse,omitempty"`
	Upload                      *UploadResponse                      `xml:"uploadResponse,omitempty" json:"uploadResponse,omitempty" yaml:"uploadResponse,omitempty"`
	ValidateE911Address         *ValidateE911AddressResponse         `xml:"validateE911AddressResponse,omitempty" json:"validateE911AddressResponse,omitempty" yaml:"validateE911AddressResponse,omitempty"`
	ViewNumberDetails           *ViewNumberDetailsResponse           `xml:"viewNumberDetailsResponse,omitempty" json:"viewNumberDetailsResponse,omitempty" yaml:"viewNumberDetailsResponse,omitempty"`
}
