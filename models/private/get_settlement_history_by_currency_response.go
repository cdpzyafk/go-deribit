package private

type GetSettlementHistoryByCurrencyResponse struct {
	Continuation string `json:"continuation" mapstructure:"continuation"`
	Settlements  []struct {
		IndexPrice        float64 `json:"index_price" mapstructure:"index_price"`
		InstrumentName    string  `json:"instrument_name" mapstructure:"instrument_name"`
		MarkPrice         float64 `json:"mark_price" mapstructure:"mark_price"`
		Position          int64   `json:"position" mapstructure:"position"`
		ProfitLoss        float64 `json:"profit_loss" mapstructure:"profit_loss"`
		SessionProfitLoss float64 `json:"session_profit_loss" mapstructure:"session_profit_loss"`
		Timestamp         int64   `json:"timestamp" mapstructure:"timestamp"`
		Type              string  `json:"type" mapstructure:"type"`
	} `json:"settlements" mapstructure:"settlements"`
}