package hmrcinterface

import (
	"fmt"
	"testing"
)

func TestVATReturn(T *testing.T) {
	v := &VATReturn{
		Vrn:                          123456789,
		PeriodKey:                    "#001",
		VatDueSales:                  100.00,
		VatDueAcquisitions:           100.00,
		TotalVatDue:                  200.00,
		VatReclaimedCurrPeriod:       100.00,
		NetVatDue:                    100.00,
		TotalValueSalesExVAT:         500,
		TotalValuePurchasesExVAT:     500,
		TotalValueGoodsSuppliedExVAT: 500,
		TotalAcquisitionsExVAT:       500,
		Finalised:                    true,
	}

	authToken := Authenticate("write:vat")

	r := v.Post(authToken)

	fmt.Println(r)
}
