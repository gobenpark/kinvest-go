package kv

type Balance struct {
	CtxAreaFk100 string        `json:"ctx_area_fk100"`
	CtxAreaNk100 string        `json:"ctx_area_nk100"`
	Output1      []interface{} `json:"output1"`
	Output2      []struct {
		DncaTotAmt         string `json:"dnca_tot_amt"`
		NxdyExccAmt        string `json:"nxdy_excc_amt"`
		PrvsRcdlExccAmt    string `json:"prvs_rcdl_excc_amt"`
		CmaEvluAmt         string `json:"cma_evlu_amt"`
		BfdyBuyAmt         string `json:"bfdy_buy_amt"`
		ThdtBuyAmt         string `json:"thdt_buy_amt"`
		NxdyAutoRdptAmt    string `json:"nxdy_auto_rdpt_amt"`
		BfdySllAmt         string `json:"bfdy_sll_amt"`
		ThdtSllAmt         string `json:"thdt_sll_amt"`
		D2AutoRdptAmt      string `json:"d2_auto_rdpt_amt"`
		BfdyTlexAmt        string `json:"bfdy_tlex_amt"`
		ThdtTlexAmt        string `json:"thdt_tlex_amt"`
		TotLoanAmt         string `json:"tot_loan_amt"`
		SctsEvluAmt        string `json:"scts_evlu_amt"`
		TotEvluAmt         string `json:"tot_evlu_amt"`
		NassAmt            string `json:"nass_amt"`
		FncgGldAutoRdptYn  string `json:"fncg_gld_auto_rdpt_yn"`
		PchsAmtSmtlAmt     string `json:"pchs_amt_smtl_amt"`
		EvluAmtSmtlAmt     string `json:"evlu_amt_smtl_amt"`
		EvluPflsSmtlAmt    string `json:"evlu_pfls_smtl_amt"`
		TotStlnSlngChgs    string `json:"tot_stln_slng_chgs"`
		BfdyTotAsstEvluAmt string `json:"bfdy_tot_asst_evlu_amt"`
		AsstIcdcAmt        string `json:"asst_icdc_amt"`
		AsstIcdcErngRt     string `json:"asst_icdc_erng_rt"`
	} `json:"output2"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}
