package requests

type HeaderItem struct {
	Header
	ToItem `json:"to_PurchaseReqnItem"`
}

type ToItem struct {
	Results []Item `json:"results"`
}
