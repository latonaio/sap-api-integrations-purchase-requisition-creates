package responses

type Header struct {
	D struct {
		PurchaseRequisition     string `json:"PurchaseRequisition"`
		PurchaseRequisitionType string `json:"PurchaseRequisitionType"`
		SourceDetermination     bool   `json:"SourceDetermination"`
	} `json:"d"`
}
