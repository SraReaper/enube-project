package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error connecting to the db:", err)
	}
	defer db.Close()

	file, err := excelize.OpenFile(os.Getenv("EXCEL_FILE"))
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	sheets := file.GetSheetList()
	fmt.Println("Abas disponíveis no arquivo:", sheets)

	rows, err := file.GetRows("Planilha1")
	if err != nil {
		log.Fatal("Error reading fornecedores file:", err)
	}

	headers := make(map[string]int)
	for i, col := range rows[0] {
		headers[col] = i
	}

	insertStmt := `
		INSERT INTO fornecedores_dados (
    partner_id,
    partner_name,
    customer_id,
    customer_name,
    customer_domain_name,
    customer_country,
    mpn_id,
    tier2_mpn_id,
    invoice_number,
    product_id,
    sku_id,
    availability_id,
    sku_name,
    product_name,
    publisher_name,
    publisher_id,
    subscription_description,
    subscription_id,
    charge_start_date,
    charge_end_date,
    usage_date,
    meter_type,
    meter_category,
    meter_id,
    meter_sub_category,
    meter_name,
    meter_region,
    unit,
    resource_location,
    consumed_service,
    resource_group,
    resource_uri,
    charge_type,
    unit_price,
    quantity,
    unit_type,
    billing_pre_tax_total,
    billing_currency,
    pricing_pre_tax_total,
    pricing_currency,
    service_info1,
    service_info2,
    tags,
    additional_info,
    effective_unit_price,
    pc_to_bc_exchange_rate,
    pc_to_bc_exchange_rate_date,
    entitlement_id,
    entitlement_description,
    partner_earned_credit_percentage,
    credit_percentage,
    credit_type,
    benefit_order_id,
    benefit_id,
    benefit_type,
	user_id
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
			$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
			$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,
			$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,
			$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,
			$51,$52,$53,$54,$55, $56
		)
	`

	for i, row := range rows[1:] {
		val := func(key string) string {
			idx, ok := headers[key]
			if !ok || idx >= len(row) {
				return ""
			}
			return row[idx]
		}

		getTime := func(key string) time.Time {
			t, err := time.Parse("2006-01-02", val(key))
			if err != nil {
				return time.Now()
			}
			return t
		}

		getFloat := func(key string) float64 {
			f, _ := strconv.ParseFloat(val(key), 64)
			return f
		}

		getInt := func(key string) int {
			n, _ := strconv.Atoi(val(key))
			return n
		}

		_, err := db.Exec(insertStmt,
			val("PartnerId"), val("PartnerName"), val("CustomerId"), val("CustomerName"), val("CustomerDomainName"),
			val("CustomerCountry"), getInt("MpnId"), getInt("Tier2MpnId"), val("InvoiceNumber"), val("ProductId"),
			val("SkuId"), val("AvailabilityId"), val("SkuName"), val("ProductName"), val("PublisherName"), val("PublisherId"),
			val("SubscriptionDescription"), val("SubscriptionId"), getTime("ChargeStartDate"), getTime("ChargeEndDate"),
			getTime("UsageDate"), val("MeterType"), val("MeterCategory"), val("MeterId"), val("MeterSubCategory"),
			val("MeterName"), val("MeterRegion"), val("Unit"), val("ResourceLocation"), val("ConsumedService"),
			val("ResourceGroup"), val("ResourceUri"), val("ChargeType"), getFloat("UnitPrice"), getFloat("Quantity"),
			val("UnitType"), getFloat("BillingPreTaxTotal"), val("BillingCurrency"), getFloat("PricingPreTaxTotal"),
			val("PricingCurrency"), val("ServiceInfo1"), val("ServiceInfo2"), val("Tags"), val("AdditionalInfo"),
			getFloat("EffectiveUnitPrice"), getFloat("PCToBCExchangeRate"), getTime("PCToBCExchangeRateDate"),
			val("EntitlementId"), val("EntitlementDescription"), getFloat("PartnerEarnedCreditPercentage"),
			getFloat("CreditPercentage"), val("CreditType"), val("BenefitOrderId"), val("BenefitId"), val("BenefitType"), 1,
		)

		if err != nil {
			log.Printf("Error in row %d: %v", i+2, err)
		}
	}

	log.Println("Import completed.")
}
