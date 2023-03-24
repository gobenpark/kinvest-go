package main

type CurrentPrice struct {
	CurrentPrice struct {
		IscdStatClsCode      string `json:"iscd_stat_cls_code"`
		MargRate             string `json:"marg_rate"`
		RprsMrktKorName      string `json:"rprs_mrkt_kor_name"`
		BstpKorIsnm          string `json:"bstp_kor_isnm"`
		TempStopYn           string `json:"temp_stop_yn"`
		OprcRangContYn       string `json:"oprc_rang_cont_yn"`
		ClprRangContYn       string `json:"clpr_rang_cont_yn"`
		CrdtAbleYn           string `json:"crdt_able_yn"`
		GrmnRateClsCode      string `json:"grmn_rate_cls_code"`
		ElwPblcYn            string `json:"elw_pblc_yn"`
		StckPrpr             string `json:"stck_prpr"`
		PrdyVrss             string `json:"prdy_vrss"`
		PrdyVrssSign         string `json:"prdy_vrss_sign"`
		PrdyCtrt             string `json:"prdy_ctrt"`
		AcmlTrPbmn           string `json:"acml_tr_pbmn"`
		AcmlVol              string `json:"acml_vol"`
		PrdyVrssVolRate      string `json:"prdy_vrss_vol_rate"`
		StckOprc             string `json:"stck_oprc"`
		StckHgpr             string `json:"stck_hgpr"`
		StckLwpr             string `json:"stck_lwpr"`
		StckMxpr             string `json:"stck_mxpr"`
		StckLlam             string `json:"stck_llam"`
		StckSdpr             string `json:"stck_sdpr"`
		WghnAvrgStckPrc      string `json:"wghn_avrg_stck_prc"`
		HtsFrgnEhrt          string `json:"hts_frgn_ehrt"`
		FrgnNtbyQty          string `json:"frgn_ntby_qty"`
		PgtrNtbyQty          string `json:"pgtr_ntby_qty"`
		PvtScndDmrsPrc       string `json:"pvt_scnd_dmrs_prc"`
		PvtFrstDmrsPrc       string `json:"pvt_frst_dmrs_prc"`
		PvtPontVal           string `json:"pvt_pont_val"`
		PvtFrstDmspPrc       string `json:"pvt_frst_dmsp_prc"`
		PvtScndDmspPrc       string `json:"pvt_scnd_dmsp_prc"`
		DmrsVal              string `json:"dmrs_val"`
		DmspVal              string `json:"dmsp_val"`
		Cpfn                 string `json:"cpfn"`
		RstcWdthPrc          string `json:"rstc_wdth_prc"`
		StckFcam             string `json:"stck_fcam"`
		StckSspr             string `json:"stck_sspr"`
		AsprUnit             string `json:"aspr_unit"`
		HtsDealQtyUnitVal    string `json:"hts_deal_qty_unit_val"`
		LstnStcn             string `json:"lstn_stcn"`
		HtsAvls              string `json:"hts_avls"`
		Per                  string `json:"per"`
		Pbr                  string `json:"pbr"`
		StacMonth            string `json:"stac_month"`
		VolTnrt              string `json:"vol_tnrt"`
		Eps                  string `json:"eps"`
		Bps                  string `json:"bps"`
		D250Hgpr             string `json:"d250_hgpr"`
		D250HgprDate         string `json:"d250_hgpr_date"`
		D250HgprVrssPrprRate string `json:"d250_hgpr_vrss_prpr_rate"`
		D250Lwpr             string `json:"d250_lwpr"`
		D250LwprDate         string `json:"d250_lwpr_date"`
		D250LwprVrssPrprRate string `json:"d250_lwpr_vrss_prpr_rate"`
		StckDryyHgpr         string `json:"stck_dryy_hgpr"`
		DryyHgprVrssPrprRate string `json:"dryy_hgpr_vrss_prpr_rate"`
		DryyHgprDate         string `json:"dryy_hgpr_date"`
		StckDryyLwpr         string `json:"stck_dryy_lwpr"`
		DryyLwprVrssPrprRate string `json:"dryy_lwpr_vrss_prpr_rate"`
		DryyLwprDate         string `json:"dryy_lwpr_date"`
		W52Hgpr              string `json:"w52_hgpr"`
		W52HgprVrssPrprCtrt  string `json:"w52_hgpr_vrss_prpr_ctrt"`
		W52HgprDate          string `json:"w52_hgpr_date"`
		W52Lwpr              string `json:"w52_lwpr"`
		W52LwprVrssPrprCtrt  string `json:"w52_lwpr_vrss_prpr_ctrt"`
		W52LwprDate          string `json:"w52_lwpr_date"`
		WholLoanRmndRate     string `json:"whol_loan_rmnd_rate"`
		SstsYn               string `json:"ssts_yn"`
		StckShrnIscd         string `json:"stck_shrn_iscd"`
		FcamCnnm             string `json:"fcam_cnnm"`
		CpfnCnnm             string `json:"cpfn_cnnm"`
		FrgnHldnQty          string `json:"frgn_hldn_qty"`
		ViClsCode            string `json:"vi_cls_code"`
		OvtmViClsCode        string `json:"ovtm_vi_cls_code"`
		LastSstsCntgQty      string `json:"last_ssts_cntg_qty"`
		InvtCafulYn          string `json:"invt_caful_yn"`
		MrktWarnClsCode      string `json:"mrkt_warn_cls_code"`
		ShortOverYn          string `json:"short_over_yn"`
		SltrYn               string `json:"sltr_yn"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

type CurrentConclusion struct {
	Output []struct {
		StckCntgHour string `json:"stck_cntg_hour"`
		StckPrpr     string `json:"stck_prpr"`
		PrdyVrss     string `json:"prdy_vrss"`
		PrdyVrssSign string `json:"prdy_vrss_sign"`
		CntgVol      string `json:"cntg_vol"`
		TdayRltv     string `json:"tday_rltv"`
		PrdyCtrt     string `json:"prdy_ctrt"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

type DailyPrice struct {
	Output []struct {
		StckBsopDate    string `json:"stck_bsop_date"`
		StckOprc        string `json:"stck_oprc"`
		StckHgpr        string `json:"stck_hgpr"`
		StckLwpr        string `json:"stck_lwpr"`
		StckClpr        string `json:"stck_clpr"`
		AcmlVol         string `json:"acml_vol"`
		PrdyVrssVolRate string `json:"prdy_vrss_vol_rate"`
		PrdyVrss        string `json:"prdy_vrss"`
		PrdyVrssSign    string `json:"prdy_vrss_sign"`
		PrdyCtrt        string `json:"prdy_ctrt"`
		HtsFrgnEhrt     string `json:"hts_frgn_ehrt"`
		FrgnNtbyQty     string `json:"frgn_ntby_qty"`
		FlngClsCode     string `json:"flng_cls_code"`
		AcmlPrttRate    string `json:"acml_prtt_rate"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

type ExpectPrice struct {
	Output1 struct {
		AsprAcptHour      string `json:"aspr_acpt_hour"`
		Askp1             string `json:"askp1"`
		Askp2             string `json:"askp2"`
		Askp3             string `json:"askp3"`
		Askp4             string `json:"askp4"`
		Askp5             string `json:"askp5"`
		Askp6             string `json:"askp6"`
		Askp7             string `json:"askp7"`
		Askp8             string `json:"askp8"`
		Askp9             string `json:"askp9"`
		Askp10            string `json:"askp10"`
		Bidp1             string `json:"bidp1"`
		Bidp2             string `json:"bidp2"`
		Bidp3             string `json:"bidp3"`
		Bidp4             string `json:"bidp4"`
		Bidp5             string `json:"bidp5"`
		Bidp6             string `json:"bidp6"`
		Bidp7             string `json:"bidp7"`
		Bidp8             string `json:"bidp8"`
		Bidp9             string `json:"bidp9"`
		Bidp10            string `json:"bidp10"`
		AskpRsqn1         string `json:"askp_rsqn1"`
		AskpRsqn2         string `json:"askp_rsqn2"`
		AskpRsqn3         string `json:"askp_rsqn3"`
		AskpRsqn4         string `json:"askp_rsqn4"`
		AskpRsqn5         string `json:"askp_rsqn5"`
		AskpRsqn6         string `json:"askp_rsqn6"`
		AskpRsqn7         string `json:"askp_rsqn7"`
		AskpRsqn8         string `json:"askp_rsqn8"`
		AskpRsqn9         string `json:"askp_rsqn9"`
		AskpRsqn10        string `json:"askp_rsqn10"`
		BidpRsqn1         string `json:"bidp_rsqn1"`
		BidpRsqn2         string `json:"bidp_rsqn2"`
		BidpRsqn3         string `json:"bidp_rsqn3"`
		BidpRsqn4         string `json:"bidp_rsqn4"`
		BidpRsqn5         string `json:"bidp_rsqn5"`
		BidpRsqn6         string `json:"bidp_rsqn6"`
		BidpRsqn7         string `json:"bidp_rsqn7"`
		BidpRsqn8         string `json:"bidp_rsqn8"`
		BidpRsqn9         string `json:"bidp_rsqn9"`
		BidpRsqn10        string `json:"bidp_rsqn10"`
		AskpRsqnIcdc1     string `json:"askp_rsqn_icdc1"`
		AskpRsqnIcdc2     string `json:"askp_rsqn_icdc2"`
		AskpRsqnIcdc3     string `json:"askp_rsqn_icdc3"`
		AskpRsqnIcdc4     string `json:"askp_rsqn_icdc4"`
		AskpRsqnIcdc5     string `json:"askp_rsqn_icdc5"`
		AskpRsqnIcdc6     string `json:"askp_rsqn_icdc6"`
		AskpRsqnIcdc7     string `json:"askp_rsqn_icdc7"`
		AskpRsqnIcdc8     string `json:"askp_rsqn_icdc8"`
		AskpRsqnIcdc9     string `json:"askp_rsqn_icdc9"`
		AskpRsqnIcdc10    string `json:"askp_rsqn_icdc10"`
		BidpRsqnIcdc1     string `json:"bidp_rsqn_icdc1"`
		BidpRsqnIcdc2     string `json:"bidp_rsqn_icdc2"`
		BidpRsqnIcdc3     string `json:"bidp_rsqn_icdc3"`
		BidpRsqnIcdc4     string `json:"bidp_rsqn_icdc4"`
		BidpRsqnIcdc5     string `json:"bidp_rsqn_icdc5"`
		BidpRsqnIcdc6     string `json:"bidp_rsqn_icdc6"`
		BidpRsqnIcdc7     string `json:"bidp_rsqn_icdc7"`
		BidpRsqnIcdc8     string `json:"bidp_rsqn_icdc8"`
		BidpRsqnIcdc9     string `json:"bidp_rsqn_icdc9"`
		BidpRsqnIcdc10    string `json:"bidp_rsqn_icdc10"`
		TotalAskpRsqn     string `json:"total_askp_rsqn"`
		TotalBidpRsqn     string `json:"total_bidp_rsqn"`
		TotalAskpRsqnIcdc string `json:"total_askp_rsqn_icdc"`
		TotalBidpRsqnIcdc string `json:"total_bidp_rsqn_icdc"`
		OvtmTotalAskpIcdc string `json:"ovtm_total_askp_icdc"`
		OvtmTotalBidpIcdc string `json:"ovtm_total_bidp_icdc"`
		OvtmTotalAskpRsqn string `json:"ovtm_total_askp_rsqn"`
		OvtmTotalBidpRsqn string `json:"ovtm_total_bidp_rsqn"`
		NtbyAsprRsqn      string `json:"ntby_aspr_rsqn"`
		NewMkopClsCode    string `json:"new_mkop_cls_code"`
	} `json:"output1"`
	Output2 struct {
		AntcMkopClsCode  string `json:"antc_mkop_cls_code"`
		StckPrpr         string `json:"stck_prpr"`
		StckOprc         string `json:"stck_oprc"`
		StckHgpr         string `json:"stck_hgpr"`
		StckLwpr         string `json:"stck_lwpr"`
		StckSdpr         string `json:"stck_sdpr"`
		AntcCnpr         string `json:"antc_cnpr"`
		AntcCntgVrssSign string `json:"antc_cntg_vrss_sign"`
		AntcCntgVrss     string `json:"antc_cntg_vrss"`
		AntcCntgPrdyCtrt string `json:"antc_cntg_prdy_ctrt"`
		AntcVol          string `json:"antc_vol"`
		StckShrnIscd     string `json:"stck_shrn_iscd"`
		ViClsCode        string `json:"vi_cls_code"`
	} `json:"output2"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

type Investor struct {
	Output []struct {
		StckBsopDate   string `json:"stck_bsop_date"`
		StckClpr       string `json:"stck_clpr"`
		PrdyVrss       string `json:"prdy_vrss"`
		PrdyVrssSign   string `json:"prdy_vrss_sign"`
		PrsnNtbyQty    string `json:"prsn_ntby_qty"`
		FrgnNtbyQty    string `json:"frgn_ntby_qty"`
		OrgnNtbyQty    string `json:"orgn_ntby_qty"`
		PrsnNtbyTrPbmn string `json:"prsn_ntby_tr_pbmn"`
		FrgnNtbyTrPbmn string `json:"frgn_ntby_tr_pbmn"`
		OrgnNtbyTrPbmn string `json:"orgn_ntby_tr_pbmn"`
		PrsnShnuVol    string `json:"prsn_shnu_vol"`
		FrgnShnuVol    string `json:"frgn_shnu_vol"`
		OrgnShnuVol    string `json:"orgn_shnu_vol"`
		PrsnShnuTrPbmn string `json:"prsn_shnu_tr_pbmn"`
		FrgnShnuTrPbmn string `json:"frgn_shnu_tr_pbmn"`
		OrgnShnuTrPbmn string `json:"orgn_shnu_tr_pbmn"`
		PrsnSelnVol    string `json:"prsn_seln_vol"`
		FrgnSelnVol    string `json:"frgn_seln_vol"`
		OrgnSelnVol    string `json:"orgn_seln_vol"`
		PrsnSelnTrPbmn string `json:"prsn_seln_tr_pbmn"`
		FrgnSelnTrPbmn string `json:"frgn_seln_tr_pbmn"`
		OrgnSelnTrPbmn string `json:"orgn_seln_tr_pbmn"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

type Output[T any, P *T] struct {
	Output P      `json:"output"`
	RtCd   string `json:"rt_cd"`
	MsgCd  string `json:"msg_cd"`
	Msg1   string `json:"msg1"`
}

type Member struct {
	Output struct {
		SelnMbcrNo1          string `json:"seln_mbcr_no1"`
		SelnMbcrNo2          string `json:"seln_mbcr_no2"`
		SelnMbcrNo3          string `json:"seln_mbcr_no3"`
		SelnMbcrNo4          string `json:"seln_mbcr_no4"`
		SelnMbcrNo5          string `json:"seln_mbcr_no5"`
		SelnMbcrName1        string `json:"seln_mbcr_name1"`
		SelnMbcrName2        string `json:"seln_mbcr_name2"`
		SelnMbcrName3        string `json:"seln_mbcr_name3"`
		SelnMbcrName4        string `json:"seln_mbcr_name4"`
		SelnMbcrName5        string `json:"seln_mbcr_name5"`
		TotalSelnQty1        string `json:"total_seln_qty1"`
		TotalSelnQty2        string `json:"total_seln_qty2"`
		TotalSelnQty3        string `json:"total_seln_qty3"`
		TotalSelnQty4        string `json:"total_seln_qty4"`
		TotalSelnQty5        string `json:"total_seln_qty5"`
		SelnMbcrRlim1        string `json:"seln_mbcr_rlim1"`
		SelnMbcrRlim2        string `json:"seln_mbcr_rlim2"`
		SelnMbcrRlim3        string `json:"seln_mbcr_rlim3"`
		SelnMbcrRlim4        string `json:"seln_mbcr_rlim4"`
		SelnMbcrRlim5        string `json:"seln_mbcr_rlim5"`
		SelnQtyIcdc1         string `json:"seln_qty_icdc1"`
		SelnQtyIcdc2         string `json:"seln_qty_icdc2"`
		SelnQtyIcdc3         string `json:"seln_qty_icdc3"`
		SelnQtyIcdc4         string `json:"seln_qty_icdc4"`
		SelnQtyIcdc5         string `json:"seln_qty_icdc5"`
		ShnuMbcrNo1          string `json:"shnu_mbcr_no1"`
		ShnuMbcrNo2          string `json:"shnu_mbcr_no2"`
		ShnuMbcrNo3          string `json:"shnu_mbcr_no3"`
		ShnuMbcrNo4          string `json:"shnu_mbcr_no4"`
		ShnuMbcrNo5          string `json:"shnu_mbcr_no5"`
		ShnuMbcrName1        string `json:"shnu_mbcr_name1"`
		ShnuMbcrName2        string `json:"shnu_mbcr_name2"`
		ShnuMbcrName3        string `json:"shnu_mbcr_name3"`
		ShnuMbcrName4        string `json:"shnu_mbcr_name4"`
		ShnuMbcrName5        string `json:"shnu_mbcr_name5"`
		TotalShnuQty1        string `json:"total_shnu_qty1"`
		TotalShnuQty2        string `json:"total_shnu_qty2"`
		TotalShnuQty3        string `json:"total_shnu_qty3"`
		TotalShnuQty4        string `json:"total_shnu_qty4"`
		TotalShnuQty5        string `json:"total_shnu_qty5"`
		ShnuMbcrRlim1        string `json:"shnu_mbcr_rlim1"`
		ShnuMbcrRlim2        string `json:"shnu_mbcr_rlim2"`
		ShnuMbcrRlim3        string `json:"shnu_mbcr_rlim3"`
		ShnuMbcrRlim4        string `json:"shnu_mbcr_rlim4"`
		ShnuMbcrRlim5        string `json:"shnu_mbcr_rlim5"`
		ShnuQtyIcdc1         string `json:"shnu_qty_icdc1"`
		ShnuQtyIcdc2         string `json:"shnu_qty_icdc2"`
		ShnuQtyIcdc3         string `json:"shnu_qty_icdc3"`
		ShnuQtyIcdc4         string `json:"shnu_qty_icdc4"`
		ShnuQtyIcdc5         string `json:"shnu_qty_icdc5"`
		GlobTotalSelnQty     string `json:"glob_total_seln_qty"`
		GlobSelnRlim         string `json:"glob_seln_rlim"`
		GlobNtbyQty          string `json:"glob_ntby_qty"`
		GlobTotalShnuQty     string `json:"glob_total_shnu_qty"`
		GlobShnuRlim         string `json:"glob_shnu_rlim"`
		SelnMbcrGlobYn1      string `json:"seln_mbcr_glob_yn_1"`
		SelnMbcrGlobYn2      string `json:"seln_mbcr_glob_yn_2"`
		SelnMbcrGlobYn3      string `json:"seln_mbcr_glob_yn_3"`
		SelnMbcrGlobYn4      string `json:"seln_mbcr_glob_yn_4"`
		SelnMbcrGlobYn5      string `json:"seln_mbcr_glob_yn_5"`
		ShnuMbcrGlobYn1      string `json:"shnu_mbcr_glob_yn_1"`
		ShnuMbcrGlobYn2      string `json:"shnu_mbcr_glob_yn_2"`
		ShnuMbcrGlobYn3      string `json:"shnu_mbcr_glob_yn_3"`
		ShnuMbcrGlobYn4      string `json:"shnu_mbcr_glob_yn_4"`
		ShnuMbcrGlobYn5      string `json:"shnu_mbcr_glob_yn_5"`
		GlobTotalSelnQtyIcdc string `json:"glob_total_seln_qty_icdc"`
		GlobTotalShnuQtyIcdc string `json:"glob_total_shnu_qty_icdc"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

type ELW struct {
	Output struct {
		ElwPrpr          string `json:"elw_prpr"`
		PrdyVrss         string `json:"prdy_vrss"`
		PrdyCtrt         string `json:"prdy_ctrt"`
		AcmlVol          string `json:"acml_vol"`
		PrdyVrssVolRate  string `json:"prdy_vrss_vol_rate"`
		UnasIsnm         string `json:"unas_isnm"`
		UnasPrpr         string `json:"unas_prpr"`
		UnasPrdyVrss     string `json:"unas_prdy_vrss"`
		UnasPrdyVrssSign string `json:"unas_prdy_vrss_sign"`
		UnasPrdyCtrt     string `json:"unas_prdy_ctrt"`
		Bidp             string `json:"bidp"`
		Askp             string `json:"askp"`
		AcmlTrPbmn       string `json:"acml_tr_pbmn"`
		VolTnrt          string `json:"vol_tnrt"`
		ElwOprc          string `json:"elw_oprc"`
		ElwHgpr          string `json:"elw_hgpr"`
		ElwLwpr          string `json:"elw_lwpr"`
		StckPrdyClpr     string `json:"stck_prdy_clpr"`
		HtsThpr          string `json:"hts_thpr"`
		Dprt             string `json:"dprt"`
		AtmClsName       string `json:"atm_cls_name"`
		HtsIntsVltl      string `json:"hts_ints_vltl"`
		Acpr             string `json:"acpr"`
		PvtScndDmrsPrc   string `json:"pvt_scnd_dmrs_prc"`
		PvtFrstDmrsPrc   string `json:"pvt_frst_dmrs_prc"`
		PvtPontVal       string `json:"pvt_pont_val"`
		PvtFrstDmspPrc   string `json:"pvt_frst_dmsp_prc"`
		PvtScndDmspPrc   string `json:"pvt_scnd_dmsp_prc"`
		DmspVal          string `json:"dmsp_val"`
		DmrsVal          string `json:"dmrs_val"`
		ElwSdpr          string `json:"elw_sdpr"`
		ApprchRate       string `json:"apprch_rate"`
		TickConvPrc      string `json:"tick_conv_prc"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}
