package test

import (
	"encoding/json"
	"ms-tagpeak/internal/models"
	"ms-tagpeak/internal/service"
	"strconv"
	"testing"
)

func storeJsonToModel() models.Store {
	storeJson := []byte(`{
  "uuid": "c79d3b85-7094-48ce-97a9-0215deeaac4e",
  "name": "Sport Zone",
  "logo": "https://app.tagpeak.com/images/images/4cdf9dbd-8224-46cf-85b7-e8364ac58ec0/logo.webp",
  "banner": "https://app.tagpeak.com/images/images/cc775ed6-d797-4036-a9c5-7dded22faf2d/resized.webp",
  "shortDescription": "Sportzone is a sports retailer that specializes in offering a wide range of athletic apparel, footwear, and equipment for various sports and outdoor activities.",
  "description": "<p><b>Important:</b> always accept cookies</p>\n",
  "urlSlug": "sport-zone",
  "averageRewardActivationTime": "35 days",
  "state": "ACTIVE",
  "affiliateLink": "https://www.awin1.com/cread.php?awinmid=27902&awinaffid=1256165",
  "termsAndConditions": "<p><strong> Very important steps to get rewards? </strong></p><p></p><p>- Click through the &quot;Go to store&quot; button to visit the partner store</p><p><span class=\"ql-font-IBMPlexSans\">- Always accept cookies before purchasing</span></p><p><span class=\"ql-font-IBMPlexSans\">- Remove/disable any blockers that might affect cookies acceptance</span></p><p><span class=\"ql-font-IBMPlexSans\">- Complete your purchase in the same browser tab or page opened, within a reasonable timeframe</span></p><p></p><p><strong> In what products or categories are cash rewards available? </strong></p><p></p><p>- All products eligible.</p><p></p><p><strong> Other important information: </strong></p><p></p><p>- This merchant calculates cashback amount excluding VAT, delivery and any other charges.</p><p></p><p>- Tagpeak cash rewards are not cumulative with the use of coupon codes, other cashback platforms, and/or other affiliate marketing solutions. In order to be eligible for our cash reward, do not use any of these other benefits in the same purchase.</p><p></p><p>- Purchases must be completed immediately and fully online to be eligible for cash rewards.</p><p></p><p>- When you click through to the partner website, please ensure you complete your booking in one smooth transaction. If you move away from your purchase on to other pages, or other websites either before or during your transaction, cash rewards will not successfully track and will not be awarded.</p>",
  "cashbackType": "FIXED",
  "cashbackValue": 6,
  "percentageCashout": 2,
  "country": [
    "PT"
  ],
  "category": [
    "sports"
  ],
  "partnerIdentity": "",
  "affiliatePartnerCode": "awin_uk"
}`)
	var store models.Store
	_ = json.Unmarshal(storeJson, &store)

	return store
}

// TestCalculateCashback_PercentualCashbackType tests the cashback calculation with percentual cashback type
// This test validates the calculation when store has a percentual cashback value
func TestCalculateCashback_PercentualCashbackType(t *testing.T) {
	t.Run("Should calculate the cashback for a transaction", func(t *testing.T) {
		store := storeJsonToModel()
		fee := 2.0
		amount := 487.74
		var commission *float64 = nil
		cashback := 0.0

		conv, _ := strconv.ParseFloat("34.14", 64)
		commission = &conv
		if *store.CashbackType == "FIXED" && commission != nil {
			t.Logf("Commission: %f", *commission)
			cashback = service.CalculateTransactionCashback(
				*commission,
				"FIXED",
				fee,
				amount,
			)

			t.Logf("storeCashbackValue: %f, storeCashbackType: %s, transactionFee: %f, transactionAmount: %f - Result: %f", *commission, "FIXED", fee, amount, cashback)
			expected := float64(30.73)
			if cashback != expected {
				t.Errorf("Expected %f, got %f", expected, cashback)
			}
			return
		}

		cashback = service.CalculateTransactionCashback(
			*store.CashbackValue,
			*store.CashbackType,
			fee,
			amount,
		)
		t.Logf("storeCashbackValue: %f, storeCashbackType: %s, transactionFee: %f, transactionAmount: %f - Result: %f", *store.CashbackValue, *store.CashbackType, fee, amount, cashback)

		expected := float64(19.51)
		if cashback != expected {
			t.Errorf("Expected %f, got %f", expected, cashback)
		}
	})
}
