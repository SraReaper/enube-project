CREATE TABLE fornecedores_dados (
    "id" serial PRIMARY KEY NOT NULL,
    "partner_id" varchar NOT NULL,
    "partner_name" varchar NOT NULL,
    "customer_id" varchar NOT NULL,
    "customer_name" varchar NOT NULL,
    "customer_domain_name" varchar NOT NULL,
    "customer_country" varchar NOT NULL,
    "mpn_id" int NOT NULL,
    "tier2_mpn_id" int NOT NULL,
    "invoice_number" varchar NOT NULL,
    "product_id" varchar NOT NULL,
    "sku_id" varchar NOT NULL,
    "availability_id" varchar NOT NULL,
    "sku_name" varchar NOT NULL,
    "product_name" varchar NOT NULL,
    "publisher_name" varchar NOT NULL,
    "publisher_id" varchar NOT NULL,
    "subscription_description" varchar NOT NULL,
    "subscription_id" varchar NOT NULL,
    "charge_start_date" timestamptz NOT NULL DEFAULT (now()),
    "charge_end_date" timestamptz NOT NULL DEFAULT (now()),
    "usage_date" timestamptz NOT NULL DEFAULT (now()),
    "meter_type" varchar NOT NULL,
    "meter_category" varchar NOT NULL,
    "meter_id" varchar NOT NULL,
    "meter_sub_category" varchar NOT NULL,
    "meter_name" varchar NOT NULL,
    "meter_region" varchar NOT NULL,
    "unit" varchar NOT NULL,
    "resource_location" varchar NOT NULL,
    "consumed_service" varchar NOT NULL,
    "resource_group" varchar NOT NULL,
    "resource_uri" varchar NOT NULL,
    "charge_type" varchar NOT NULL,
    "unit_price" NUMERIC NOT NULL,
    "quantity" NUMERIC NOT NULL,
    "unit_type" varchar NOT NULL,
    "billing_pre_tax_total" NUMERIC NOT NULL,
    "billing_currency" varchar NOT NULL,
    "pricing_pre_tax_total" NUMERIC NOT NULL,
    "pricing_currency" varchar NOT NULL,
    "service_info1" varchar NOT NULL,
    "service_info2" varchar NOT NULL,
    "tags" varchar NOT NULL,
    "additional_info" varchar NOT NULL,
    "effective_unit_price" NUMERIC NOT NULL,
    "pc_to_bc_exchange_rate" NUMERIC NOT NULL,
    "pc_to_bc_exchange_rate_date" timestamptz NOT NULL DEFAULT (now()),
    "entitlement_id" varchar NOT NULL,
    "entitlement_description" varchar NOT NULL,
    "partner_earned_credit_percentage" NUMERIC NOT NULL,
    "credit_percentage" NUMERIC NOT NULL,
    "credit_type" varchar NOT NULL,
    "benefit_order_id" varchar NOT NULL,
    "benefit_id" varchar NOT NULL,
    "benefit_type" varchar NOT NULL
);

CREATE INDEX idx_fornecedores_dados_partner_id ON fornecedores_dados (partner_id);
CREATE INDEX idx_fornecedores_dados_customer_id ON fornecedores_dados (customer_id);
CREATE INDEX idx_fornecedores_dados_subscription_id ON fornecedores_dados (subscription_id);
CREATE INDEX idx_fornecedores_dados_invoice_number ON fornecedores_dados (invoice_number);
CREATE INDEX idx_fornecedores_dados_product_id ON fornecedores_dados (product_id);
CREATE INDEX idx_fornecedores_dados_sku_id ON fornecedores_dados (sku_id);
CREATE INDEX idx_fornecedores_dados_availability_id ON fornecedores_dados (availability_id);
CREATE INDEX idx_fornecedores_dados_mpn_id ON fornecedores_dados (mpn_id);
CREATE INDEX idx_fornecedores_dados_tier2_mpn_id ON fornecedores_dados (tier2_mpn_id);
CREATE INDEX idx_fornecedores_dados_usage_date ON fornecedores_dados (usage_date);
CREATE INDEX idx_fornecedores_dados_charge_start_date ON fornecedores_dados (charge_start_date);
CREATE INDEX idx_fornecedores_dados_charge_end_date ON fornecedores_dados (charge_end_date);
CREATE INDEX idx_fornecedores_dados_meter_id ON fornecedores_dados (meter_id);
CREATE INDEX idx_fornecedores_dados_meter_type ON fornecedores_dados (meter_type);
CREATE INDEX idx_fornecedores_dados_meter_category ON fornecedores_dados (meter_category);
CREATE INDEX idx_fornecedores_dados_meter_sub_category ON fornecedores_dados (meter_sub_category);
CREATE INDEX idx_fornecedores_dados_meter_name ON fornecedores_dados (meter_name);
CREATE INDEX idx_fornecedores_dados_meter_region ON fornecedores_dados (meter_region);
CREATE INDEX idx_fornecedores_dados_resource_location ON fornecedores_dados (resource_location);
CREATE INDEX idx_fornecedores_dados_consumed_service ON fornecedores_dados (consumed_service);
CREATE INDEX idx_fornecedores_dados_resource_group ON fornecedores_dados (resource_group);
CREATE INDEX idx_fornecedores_dados_resource_uri ON fornecedores_dados (resource_uri);
CREATE INDEX idx_fornecedores_dados_charge_type ON fornecedores_dados (charge_type);
CREATE INDEX idx_fornecedores_dados_unit ON fornecedores_dados (unit);
CREATE INDEX idx_fornecedores_dados_billing_currency ON fornecedores_dados (billing_currency);
CREATE INDEX idx_fornecedores_dados_pricing_currency ON fornecedores_dados (pricing_currency);
CREATE INDEX idx_fornecedores_dados_service_info1 ON fornecedores_dados (service_info1);
CREATE INDEX idx_fornecedores_dados_service_info2 ON fornecedores_dados (service_info2);
CREATE INDEX idx_fornecedores_dados_tags ON fornecedores_dados (tags);
CREATE INDEX idx_fornecedores_dados_additional_info ON fornecedores_dados (additional_info);
CREATE INDEX idx_fornecedores_dados_effective_unit_price ON fornecedores_dados (effective_unit_price);
CREATE INDEX idx_fornecedores_dados_pc_to_bc_exchange_rate ON fornecedores_dados (pc_to_bc_exchange_rate);
CREATE INDEX idx_fornecedores_dados_pc_to_bc_exchange_rate_date ON fornecedores_dados (pc_to_bc_exchange_rate_date);
CREATE INDEX idx_fornecedores_dados_entitlement_id ON fornecedores_dados (entitlement_id);
CREATE INDEX idx_fornecedores_dados_entitlement_description ON fornecedores_dados (entitlement_description);
CREATE INDEX idx_fornecedores_dados_partner_earned_credit_percentage ON fornecedores_dados (partner_earned_credit_percentage);
CREATE INDEX idx_fornecedores_dados_credit_percentage ON fornecedores_dados (credit_percentage);
CREATE INDEX idx_fornecedores_dados_credit_type ON fornecedores_dados (credit_type);
CREATE INDEX idx_fornecedores_dados_benefit_order_id ON fornecedores_dados (benefit_order_id);
CREATE INDEX idx_fornecedores_dados_benefit_id ON fornecedores_dados (benefit_id);
CREATE INDEX idx_fornecedores_dados_benefit_type ON fornecedores_dados (benefit_type);

CREATE TABLE "users" (
    "id" serial PRIMARY KEY NOT NULL,
    "username" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "fornecedores_dados" ADD COLUMN user_id int NOT NULL;
ALTER TABLE "fornecedores_dados" 
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES "users" (id);
