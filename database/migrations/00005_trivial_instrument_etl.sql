-- +goose Up

CREATE TABLE etl.etl_bats_instrument (
    company_name                                    VARCHAR(128),
    bats_name                                       VARCHAR(12),
    isin					                        VARCHAR(12),
    currency				                        VARCHAR(3),
    mic						                        VARCHAR(4),
    reuters_exchange_code	                        VARCHAR(4),
    lis_local				                        DECIMAL(21,9),
    live					                        VARCHAR(1),
    tick_type				                        VARCHAR(12),
    reference_price			                        DECIMAL(21,9),
    bats_prev_close			                        DECIMAL(21,9),
    live_date				                        DATE,
    bloomberg_primary		                        VARCHAR(32),
    bloomberg_bats			                        VARCHAR(32),
    mifid_share				                        VARCHAR(1),
    asset_class				                        VARCHAR(8),
    matching_unit			                        INTEGER,
    euroccp_enabled			                        BOOLEAN,
    xclr_enabled			                        BOOLEAN,
    lchl_enabled			                        BOOLEAN,
    reuters_ric_primary		                        VARCHAR(99),
    reuters_ric_bats		                        VARCHAR(99),
    reference_adt_eur		                        DECIMAL(21,9),
    csd						                        VARCHAR(16),
    corporate_action_status	                        VARCHAR(99),
    supported_services		                        VARCHAR(8),
    trading_segment			                        VARCHAR(4),
    printed_name						            VARCHAR(8),
    periodic_auction_max_duration			        INTEGER,
    periodic_auction_min_order_entry_size	        INTEGER,
    periodic_auction_min_order_entry_notional		INTEGER,
    max_otr_count						            INTEGER,
    max_otr_volume						            INTEGER,
    capped						                    INTEGER,
    venue_cap_percentage						    DECIMAL(21,9),
    venue_uncap_date                                DATE
);

-- +goose Down
DROP TABLE etl_bats_instrument;
