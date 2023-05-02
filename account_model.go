package kv

type Balance struct {
	CtxAreaFk100 string        `json:"ctx_area_fk100"`
	CtxAreaNk100 string        `json:"ctx_area_nk100"`
	Output1      []interface{} `json:"output1"`
	Output2      []struct {
		//예수금 총금액
		DncaTotAmt string `json:"dnca_tot_amt"`
		//익일정산금액
		NxdyExccAmt string `json:"nxdy_excc_amt"`
		//가수도정산금액
		PrvsRcdlExccAmt string `json:"prvs_rcdl_excc_amt"`
		//CMA평가금액
		CmaEvluAmt string `json:"cma_evlu_amt"`
		//전일매수금액
		BfdyBuyAmt string `json:"bfdy_buy_amt"`
		//금일매수금액
		ThdtBuyAmt string `json:"thdt_buy_amt"`
		//익일자동상환금액
		NxdyAutoRdptAmt string `json:"nxdy_auto_rdpt_amt"`
		//전일매도금액
		BfdySllAmt string `json:"bfdy_sll_amt"`
		//금일매도금액
		ThdtSllAmt string `json:"thdt_sll_amt"`
		//D+2자동상환금액
		D2AutoRdptAmt string `json:"d2_auto_rdpt_amt"`
		//전일제비용금액
		BfdyTlexAmt string `json:"bfdy_tlex_amt"`
		//금일제비용금액
		ThdtTlexAmt string `json:"thdt_tlex_amt"`
		//총대출금액
		TotLoanAmt string `json:"tot_loan_amt"`
		//유가평가금액
		SctsEvluAmt string `json:"scts_evlu_amt"`
		//총평가금액
		TotEvluAmt string `json:"tot_evlu_amt"`
		//순자산금액
		NassAmt string `json:"nass_amt"`
		//융자금자동상환여부
		FncgGldAutoRdptYn string `json:"fncg_gld_auto_rdpt_yn"`
		//매입금액합계금액
		PchsAmtSmtlAmt string `json:"pchs_amt_smtl_amt"`
		//평가금액합계금액
		EvluAmtSmtlAmt string `json:"evlu_amt_smtl_amt"`
		//평가손익합계금액
		EvluPflsSmtlAmt string `json:"evlu_pfls_smtl_amt"`
		//총대주매각대금
		TotStlnSlngChgs string `json:"tot_stln_slng_chgs"`
		//전일총자산평가금액
		BfdyTotAsstEvluAmt string `json:"bfdy_tot_asst_evlu_amt"`
		//자산증감액
		AsstIcdcAmt string `json:"asst_icdc_amt"`
		//자산증감수익율
		AsstIcdcErngRt string `json:"asst_icdc_erng_rt"`
	} `json:"output2"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}
