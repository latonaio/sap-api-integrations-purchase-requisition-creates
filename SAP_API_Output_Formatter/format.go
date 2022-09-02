package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-requisition-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionType:        data.PurchaseRequisitionType,
		SourceDetermination:            data.SourceDetermination,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		PurReqnReleaseStatus:           data.PurReqnReleaseStatus,
		PurchasingDocumentItemCategory: data.PurchasingDocumentItemCategory,
		PurchaseRequisitionItemText:    data.PurchaseRequisitionItemText,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		Material:                       data.Material,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingDocumentCategory:     data.PurchasingDocumentCategory,
		RequestedQuantity:              data.RequestedQuantity,
		BaseUnit:                       data.BaseUnit,
		PurchaseRequisitionPrice:       data.PurchaseRequisitionPrice,
		PurReqnPriceQuantity:           data.PurReqnPriceQuantity,
		MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
		ReleaseCode:                    data.ReleaseCode,
		PurchaseRequisitionReleaseDate: data.PurchaseRequisitionReleaseDate,
		PurchasingOrganization:         data.PurchasingOrganization,
		PurchasingGroup:                data.PurchasingGroup,
		Plant:                          data.Plant,
		CompanyCode:                    data.CompanyCode,
		SourceOfSupplyIsAssigned:       data.SourceOfSupplyIsAssigned,
		SupplyingPlant:                 data.SupplyingPlant,
		OrderedQuantity:                data.OrderedQuantity,
		DeliveryDate:                   data.DeliveryDate,
		CreationDate:                   data.CreationDate,
		ProcessingStatus:               data.ProcessingStatus,
		ExternalApprovalStatus:         data.ExternalApprovalStatus,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		Supplier:                       data.Supplier,
		FixedSupplier:                  data.FixedSupplier,
		RequisitionerName:              data.RequisitionerName,
		PurReqnItemCurrency:            data.PurReqnItemCurrency,
		MaterialPlannedDeliveryDurn:    data.MaterialPlannedDeliveryDurn,
		StorageLocation:                data.StorageLocation,
		PurReqnSourceOfSupplyType:      data.PurReqnSourceOfSupplyType,
		ConsumptionPosting:             data.ConsumptionPosting,
		PurReqnOrigin:                  data.PurReqnOrigin,
		IsPurReqnBlocked:               data.IsPurReqnBlocked,
		PurchaseRequisitionStatus:      data.PurchaseRequisitionStatus,
		Batch:                          data.Batch,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		RequirementTracking:            data.RequirementTracking,
		MRPController:                  data.MRPController,
		Reservation:                    data.Reservation,
		LastChangeDateTime:             data.LastChangeDateTime,
		IsDeleted:                      data.IsDeleted,
	}

	return item, nil
}
