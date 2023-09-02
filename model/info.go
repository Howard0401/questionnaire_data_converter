package model

var ColNames = []string{
	"result_sn",
	"N_Reg_WST_C",
	"N_Reg_WET_C",
	"N_Reg_WLT_C",
	"N_Reg_AT_C",
	"N_Reg_PeakT_C",
	"N_Reg_WST_H",
	"N_Reg_WLT_H",
	"N_ELD_WST_C",
	"N_ELD_WLT_C",
	"N_ELD_AT_C",
	"N_ELD_PeakT_C",
	"N_ELD_WST_H",
	"N_ELD_WLT_H",
	"N_Reg_WD_C",
	"N_Reg_WTH_C",
	"N_Reg_CGSTNfreqFwy_C",
	"N_Reg_CGSTNfreqFwyDay_C",
	"N_Reg_CGSTNfreqUrb_C",
	"N_Reg_CGSTNfreqUrbDay_C",
	"N_Reg_CGSTNfreq_C",
	"N_Reg_WD_H",
	"N_Reg_WTH_H",
	"N_ELD_WD_C",
	"N_ELD_WTH_C",
	"N_ELD_CGSTNfreqFwy_C",
	"N_ELD_CGSTNfreqFwyDay_C",
	"N_ELD_CGSTNfreqUrb_C",
	"N_ELD_CGSTNfreqUrbDay_C",
	"N_ELD_CGSTNfreq_C",
	"N_ELD_WD_H",
	"N_ELD_WTH_H",
	"N_Reg_WOT_C",
	"N_Reg_PeakT_C1",
	"N_Reg_PeakT_C2",
	"N_ELD_WOT_C",
	"N_ELD_PeakT_C1",
	"N_ELD_PeakT_C2",
	"N_ELW_WTH_C",
	"N_ELW_WTH_H",
	"N_Reg_WET_H",
	"N_ELD_WET_C",
	"N_ELD_WET_H",
	"N_Reg_DT_C",
	"N_ELD_DT_C",
	"C_Reg_WST_C",
	"C_Reg_WET_C",
	"C_Reg_WLT_C",
	"C_Reg_AT_C",
	"C_Reg_WST_H",
	"C_Reg_WET_H",
	"C_Reg_WLT_H",
	"C_ELD_WST_C",
	"C_ELD_WET_C",
	"C_ELD_AT_C",
	"C_ELD_WST_H",
	"C_ELD_WET_H",
	"C_ELD_WLT_H",
	"C_ELD_WLT_C",
	"C_Reg_WD_C",
	"C_Reg_WTH_C",
	"C_Reg_CGSTNfreqFwy_C",
	"C_Reg_CGSTNfreqFwyDay_C",
	"C_Reg_CGSTNfreqUrb_C",
	"C_Reg_CGSTNfreqUrbDay_C",
	"C_Reg_CGSTNfreq_C",
	"C_Reg_WD_H",
	"C_Reg_WTH_H",
	"C_ELD_WD_C",
	"C_ELD_CGSTNfreqFwy_C",
	"C_ELD_CGSTNfreqFwyDay_C",
	"C_ELD_CGSTNfreqUrb_C",
	"C_ELD_CGSTNfreqUrbDay_C",
	"C_ELD_CGSTNfreq_C",
	"C_ELD_WD_H",
	"C_Reg_WOT_C",
	"C_Reg_PeakT_C1",
	"C_Reg_PeakT_C2",
	"C_ELD_WOT_C",
	"C_ELD_PeakT_C1",
	"C_ELD_PeakT_C2",
	"C_ELD_WTH_C",
	"C_ELW_WTH_C",
	"C_ELD_WTH_H",
	"C_ELW_WTH_H",
	"C_Reg_DT_C",
	"C_ELD_DT_C",
	"late_rule",
	"late_times",
	"late_reason",
	"late_penalty",
	"penalty_feel",
	"Free_PAKG",
	"PAKG",
	"late_rulersn",
	"late_reason_value",
	"late_penalty_value",
	"late_min",
	"CMUTE_yy",
	"CMUTE_mm",
	"CMUTE_dd",
	"CMUTE_w",
	"TTDT_ORIG",
	"TTDT_EACPT",
	"TTDT_LACPT",
	"TTARR_EACPT",
	"TTARR_LACPT",
	"zipcode_home",
	"road_home",
	"zipcode_work",
	"road_work",
	"city_home",
	"Urban_home",
	"city_work",
	"Urban_work",
	"consider_COVID",
	"habitual_DT",
	"Odtime_thesame",
	"work_type",
	"CarSeat_No",
	"Stop_No",
	"Stop_purp_rsn",
	"otherValue4",
	"TDwalk",
	"OVTTwalk",
	"TDbike",
	"OVTTbike",
	"TCbike",
	"TD_Urbn",
	"V_URRD",
	"IVTT_URRD",
	"TD_HWY",
	"V_HWY",
	"IVTT_HWY",
	"TD_FWY",
	"V_FWY",
	"IVTT_FWY",
	"TC_fee",
	"TC_fuel",
	"TC_PKG",
	"IVTT_PRKG",
	"Total_seat",
	"CarSeat",
	"Stop",
	"IDveh",
	"IVTT_STOP",
	"PKG_home_1",
	"PKG_home_2",
	"route_sim",
	"Choice_SN1",
	"Choice_rsn",
	"EtCO2_A1",
	"EtCO2_A2",
	"EtCO2_B",
	"EtCO2_C",
	"EyCO2_A1",
	"EyCO2_A2",
	"EyCO2_B",
	"EyCO2_C",
	"ST_WRK",
	"OVTT",
	"TTDT_EACPT_Time",
	"TTDT_LACPT_Time",
	"TDFwy_map",
	"TDHwy_map",
	"TDUrban_map",
	"TT_map",
	"TTFwy_map",
	"TTHwy_map",
	"TTUrban_map",
	"IVTT_ORIG",
	"TTT_ORIG",
	"EPT_ORIG",
	"IVTTS_ORIG",
	"IVTTL_ORIG",
	"IVTTA_ORIG",
	"COE_TT_ORIG",
	"TTCGSTN_ORIG",
	"TTFwy_ORIG",
	"TTHwy_ORIG",
	"TTUrban_ORIG",
	"EtCO2_ORIG",
	"EyCO2_ORIG",
	"ECtrip_ORIG",
	"ECyr_ORIG",
	"ECtrip_A",
	"IVTT_EA",
	"TTT_EA",
	"TTDT_EA",
	"EPT_EA",
	"IVTTS_EA",
	"IVTTL_EA",
	"IVTTA_EA",
	"COE_TT_EA",
	"TTCGSTN_EA",
	"TTFwy_EA",
	"TTHwy_EA",
	"TTUrban_EA",
	"EtCO2_EA",
	"EyCO2_EA",
	"ECtrip_EA",
	"ECyr_EA",
	"DEtCO2_EA",
	"DEyCO2_EA",
	"DEtCost_EA",
	"DEyCost_EA",
	"IVTT_LA",
	"TTT_LA",
	"TTDT_LA",
	"EPT_LA",
	"IVTTS_LA",
	"IVTTL_LA",
	"IVTTA_LA",
	"COE_TT_LA",
	"TTCGSTN_LA",
	"TTFwy_LA",
	"TTHwy_LA",
	"TTUrban_LA",
	"EtCO2_LA",
	"EyCO2_LA",
	"ECtrip_LA",
	"ECyr_LA",
	"DEtCO2_LA",
	"DEyCO2_LA",
	"DEtCost_LA",
	"DEyCost_LA",
	"WTPChoice_NP_A",
	"WTP_NP_A",
	"WTPChoice_NP_B",
	"WTP_NP_B",
	"WTPChoice_NP_C",
	"WTP_NP_C",
	"WTPChoice_P_A",
	"WTP_P_A",
	"WTPChoice_P_B",
	"WTP_P_B",
	"WTPChoice_P_C",
	"WTP_P_C",
	"WTP_zerorsn",
	"WTP_zerorsn_other",
	"WTP_usage",
	"WTP_usage1type",
	"WTP_usage_rsn",
	"WTA_DT_Bnp_choice",
	"WTA_DT_Bnp",
	"WTA_DT_Cnp_choice",
	"WTA_DT_Cnp",
	"WTA_DT_Bp_choice",
	"WTA_DT_Bp1",
	"WTA_DT_Bp2",
	"WTA_DT_Cp_choice",
	"WTA_DT_Cp1",
	"WTA_DT_Cp2",
	"WTA_DT_zero",
	"WTA_DT_zerorsn",
	"WTA_DT_RMN",
	"WTA_DT_RMN_rsn",
	"COVID_DT_Choice",
	"carpoolD_Freq_before",
	"carpoolP_Freq_before",
	"Uber_Freq_before",
	"CScar_Freq_before",
	"CSoto_Freq_before",
	"CSebike_Freq_before",
	"CSbike_Freq_before",
	"carpoolD_Freq_now",
	"carpoolP_Freq_now",
	"Uber_Freq_now",
	"CScar_Freq_now",
	"CSmoto_Freq_now",
	"CSebike_Freq_now",
	"CSbike_Freq_now",
	"CarShare_PREF1",
	"CarShare_PREF2",
	"CarShare_PREF3",
	"CarShare_PREF4",
	"CarShare_PREF5",
	"CarShare_PREF6",
	"CarShare_PREF7",
	"CarShare_PREF8",
	"CPL_D_PAX",
	"IVTT_PU_D",
	"CPL_P_PAX",
	"IVTT_PU_P",
	"OVTT_WT_PAX",
	"Uber_PAX",
	"IVTT_PU_Uber",
	"OVTT_WT_Uber",
	"CSCAR_PAX",
	"IVTTRT_C",
	"OVTTwalk_CSCAR",
	"CSMK_PAX",
	"IVTTRT_MB",
	"OVTTwalk_CSMB",
	"IVTTRT_EB",
	"OVTTwalk_CSEB",
	"IVTTRT_B",
	"OVTTwalk_CSB",
	"CarShare_PREF",
	"Choice_SN3",
	"IVTT_A",
	"IVTT_A_A",
	"IVTT_pkup_A",
	"IVTT_PRKG_A",
	"OVTT_A",
	"OVTTwalk_A",
	"OVTTbike_A",
	"OVTTwait_A",
	"TTT_A",
	"EPT_A",
	"TC_A",
	"TC_E_A",
	"TC_fee_A",
	"TC_PKG_A",
	"PAX_A",
	"EtCO2_A",
	"EyCO2_A",
	"ECyr_A",
	"IVTT_B",
	"IVTT_A_B",
	"IVTT_pkup_B",
	"IVTT_PRKG_B",
	"OVTT_B",
	"OVTTwalk_B",
	"OVTTbike_B",
	"OVTTwait_B",
	"TTT_B",
	"EPT_B",
	"TC_B",
	"TC_E_B",
	"TC_fee_B",
	"TC_PKG_B",
	"PAX_B",
	"ECtrip_B",
	"ECyr_B",
	"IVTT_C",
	"IVTT_A_C",
	"IVTT_pkup_C",
	"IVTT_PRKG_C",
	"OVTT_C",
	"OVTTwalk_C",
	"OVTTbike_C",
	"OVTTwait_C",
	"TTT_C",
	"EPT_C",
	"TC_C",
	"TC_E_C",
	"TC_fee_C",
	"TC_PKG_C",
	"PAX_C",
	"ECtrip_C",
	"ECyr_C",
	"IVTT_D",
	"IVTT_A_D",
	"IVTT_pkup_D",
	"IVTT_PRKG_D",
	"OVTT_D",
	"OVTTwalk_D",
	"OVTTbike_D",
	"OVTTwait_D",
	"TTT_D",
	"EPT_D",
	"TC_E_D",
	"TC_fee_D",
	"TC_PKG_D",
	"PAX_D",
	"EtCO2_D",
	"EyCO2_D",
	"ECtrip_D",
	"ECyr_D",
	"IVTT_E",
	"IVTT_A_E",
	"IVTT_pkup_E",
	"IVTT_PRKG_E",
	"OVTT_E",
	"OVTTwalk_E",
	"OVTTbike_E",
	"OVTTwait_E",
	"TTT_E",
	"EPT_E",
	"TC_E",
	"TC_E_E",
	"TC_fee_E",
	"TC_PKG_E",
	"PAX_E",
	"EtCO2_E",
	"EyCO2_E",
	"ECtrip_E",
	"ECyr_E",
	"IVTT_F",
	"IVTT_A_F",
	"IVTT_pkup_F",
	"IVTT_PRKG_F",
	"OVTT_F",
	"OVTTwalk_F",
	"OVTTbike_F",
	"OVTTwait_F",
	"TTT_F",
	"EPT_F",
	"TC_F",
	"TC_E_F",
	"TC_fee_F",
	"TC_PKG_F",
	"PAX_F",
	"EtCO2_F",
	"EyCO2_F",
	"ECtrip_F",
	"ECyr_F",
	"IVTT_G",
	"IVTT_A_G",
	"IVTT_pkup_G",
	"IVTT_PRKG_G",
	"OVTT_G",
	"OVTTwalk_G",
	"OVTTbike_G",
	"OVTTwait_G",
	"TTT_G",
	"EPT_G",
	"TC_G",
	"TC_E_G",
	"TC_fee_G",
	"TC_PKG_G",
	"PAX_G",
	"EtCO2_G",
	"EyCO2_G",
	"ECtrip_G",
	"ECyr_G",
	"IVTT_H",
	"IVTT_A_H",
	"IVTT_pkup_H",
	"IVTT_PRKG_H",
	"OVTT_H",
	"OVTTwalk_H",
	"OVTTbike_H",
	"OVTTwait_H",
	"TTT_H",
	"EPT_H",
	"TC_H",
	"TC_E_H",
	"TC_fee_H",
	"TC_PKG_H",
	"PAX_H",
	"EtCO2_H",
	"EyCO2_H",
	"ECtrip_H",
	"ECyr_H",
	"WTA_Mode_Bnp_choice",
	"WTA_Mode_Bnp",
	"WTA_Mode_Cnp_choice",
	"WTA_Mode_Cnp",
	"WTA_Mode_Dnp_choice",
	"WTA_Mode_Dnp",
	"WTA_Mode_Enp_choice",
	"WTA_Mode_Enp",
	"WTA_Mode_Fnp_choice",
	"WTA_Mode_Fnp",
	"WTA_Mode_Gnp_choice",
	"WTA_Mode_Gnp",
	"WTA_Mode_Hnp_choice",
	"WTA_Mode_Hnp",
	"WTA_Mode_zero",
	"WTA_Mode_zerorsn",
	"WTA_Mode_RMN",
	"WTA_Mode_RMN_rsn",
	"COVID_Mode_Choice",
	"WTA_Attract",
	"WTA_Attract_rsn",
	"gender",
	"age",
	"education",
	"income_pers",
	"work_period",
	"work_period_2",
	"resident_status",
	"resident_HH",
	"income_pers_value",
	"income_HH",
	"resident_kid",
	"income_HH_value",
	"driving_EXPC",
	"Scooter_HH",
	"Scooter_HH_use",
	"Car_HH",
	"Car_HH_use",
	"Vehicle_INCR",
	"Vehicle_DECR",
	"Vehicle_nn_change",
	"WFHday",
	"SR04",
	"SP05",
	"SP09",
	"SR10",
	"SR12",
	"SR14",
	"SP15",
	"SR18",
	"SP19",
	"SP21",
	"SR24",
	"SP31",
	"SP35",
	"SP37",
	"SR38",
	"SP39",
	"SR44",
	"SR46",
	"SP47",
	"SP48",
	"FilterCity",
	"FilterLicience",
	"FilterMode",
}

type InfoDetail struct {
	ResultSN             int    `json:"result_sn,omitempty" db:"result_sn" csv:"result_sn"`
	NRegWSTC             string `json:"N_Reg_WST_C,omitempty" db:"N_Reg_WST_C" csv:"N_Reg_WST_C"`
	NRegWETC             string `json:"N_Reg_WET_C,omitempty" db:"N_Reg_WET_C" csv:"N_Reg_WET_C"`
	NRegWLTC             string `json:"N_Reg_WLT_C,omitempty" db:"N_Reg_WLT_C" csv:"N_Reg_WLT_C"`
	NRegATC              string `json:"N_Reg_AT_C,omitempty" db:"N_Reg_AT_C" csv:"N_Reg_AT_C"`
	NRegPeakTC           string `json:"N_Reg_PeakT_C,omitempty" db:"N_Reg_PeakT_C" csv:"N_Reg_PeakT_C"`
	NRegWSTH             string `json:"N_Reg_WST_H,omitempty" db:"N_Reg_WST_H" csv:"N_Reg_WST_H"`
	NRegWLTH             string `json:"N_Reg_WLT_H,omitempty" db:"N_Reg_WLT_H" csv:"N_Reg_WLT_H"`
	NEldWstC             string `json:"N_ELD_WST_C,omitempty" db:"N_ELD_WST_C" csv:"N_ELD_WST_C"`
	NEldWltC             string `json:"N_ELD_WLT_C,omitempty" db:"N_ELD_WLT_C" csv:"N_ELD_WLT_C"`
	NEldAtC              string `json:"N_ELD_AT_C,omitempty" db:"N_ELD_AT_C" csv:"N_ELD_AT_C"`
	NELDPeakTC           string `json:"N_ELD_PeakT_C,omitempty" db:"N_ELD_PeakT_C" csv:"N_ELD_PeakT_C"`
	NEldWstH             string `json:"N_ELD_WST_H,omitempty" db:"N_ELD_WST_H" csv:"N_ELD_WST_H"`
	NEldWltH             string `json:"N_ELD_WLT_H,omitempty" db:"N_ELD_WLT_H" csv:"N_ELD_WLT_H"`
	NRegWDC              string `json:"N_Reg_WD_C,omitempty" db:"N_Reg_WD_C" csv:"N_Reg_WD_C"`
	NRegWTHC             string `json:"N_Reg_WTH_C,omitempty" db:"N_Reg_WTH_C" csv:"N_Reg_WTH_C"`
	NRegCGSTNfreqFwyC    string `json:"N_Reg_CGSTNfreqFwy_C,omitempty" db:"N_Reg_CGSTNfreqFwy_C" csv:"N_Reg_CGSTNfreqFwy_C"`
	NRegCGSTNfreqFwyDayC string `json:"N_Reg_CGSTNfreqFwyDay_C,omitempty" db:"N_Reg_CGSTNfreqFwyDay_C" csv:"N_Reg_CGSTNfreqFwyDay_C"`
	NRegCGSTNfreqUrbC    string `json:"N_Reg_CGSTNfreqUrb_C,omitempty" db:"N_Reg_CGSTNfreqUrb_C" csv:"N_Reg_CGSTNfreqUrb_C"`
	NRegCGSTNfreqUrbDayC string `json:"N_Reg_CGSTNfreqUrbDay_C,omitempty" db:"N_Reg_CGSTNfreqUrbDay_C" csv:"N_Reg_CGSTNfreqUrbDay_C"`
	NRegCGSTNfreqC       string `json:"N_Reg_CGSTNfreq_C,omitempty" db:"N_Reg_CGSTNfreq_C" csv:"N_Reg_CGSTNfreq_C"`
	NRegWDH              string `json:"N_Reg_WD_H,omitempty" db:"N_Reg_WD_H" csv:"N_Reg_WD_H"`
	NRegWTHH             string `json:"N_Reg_WTH_H,omitempty" db:"N_Reg_WTH_H" csv:"N_Reg_WTH_H"`
	NEldWdC              string `json:"N_ELD_WD_C,omitempty" db:"N_ELD_WD_C" csv:"N_ELD_WD_C"`
	NEldWthC             string `json:"N_ELD_WTH_C,omitempty" db:"N_ELD_WTH_C" csv:"N_ELD_WTH_C"`
	NELDCGSTNfreqFwyC    string `json:"N_ELD_CGSTNfreqFwy_C,omitempty" db:"N_ELD_CGSTNfreqFwy_C" csv:"N_ELD_CGSTNfreqFwy_C"`
	NELDCGSTNfreqFwyDayC string `json:"N_ELD_CGSTNfreqFwyDay_C,omitempty" db:"N_ELD_CGSTNfreqFwyDay_C" csv:"N_ELD_CGSTNfreqFwyDay_C"`
	NELDCGSTNfreqUrbC    string `json:"N_ELD_CGSTNfreqUrb_C,omitempty" db:"N_ELD_CGSTNfreqUrb_C" csv:"N_ELD_CGSTNfreqUrb_C"`
	NELDCGSTNfreqUrbDayC string `json:"N_ELD_CGSTNfreqUrbDay_C,omitempty" db:"N_ELD_CGSTNfreqUrbDay_C" csv:"N_ELD_CGSTNfreqUrbDay_C"`
	NELDCGSTNfreqC       string `json:"N_ELD_CGSTNfreq_C,omitempty" db:"N_ELD_CGSTNfreq_C" csv:"N_ELD_CGSTNfreq_C"`
	NEldWdH              string `json:"N_ELD_WD_H,omitempty" db:"N_ELD_WD_H" csv:"N_ELD_WD_H"`
	NEldWthH             string `json:"N_ELD_WTH_H,omitempty" db:"N_ELD_WTH_H" csv:"N_ELD_WTH_H"`
	NRegWOTC             string `json:"N_Reg_WOT_C,omitempty" db:"N_Reg_WOT_C" csv:"N_Reg_WOT_C"`
	NRegPeakTC1          string `json:"N_Reg_PeakT_C1,omitempty" db:"N_Reg_PeakT_C1" csv:"N_Reg_PeakT_C1"`
	NRegPeakTC2          string `json:"N_Reg_PeakT_C2,omitempty" db:"N_Reg_PeakT_C2" csv:"N_Reg_PeakT_C2"`
	NEldWotC             string `json:"N_ELD_WOT_C,omitempty" db:"N_ELD_WOT_C" csv:"N_ELD_WOT_C"`
	NELDPeakTC1          string `json:"N_ELD_PeakT_C1,omitempty" db:"N_ELD_PeakT_C1" csv:"N_ELD_PeakT_C1"`
	NELDPeakTC2          string `json:"N_ELD_PeakT_C2,omitempty" db:"N_ELD_PeakT_C2" csv:"N_ELD_PeakT_C2"`
	NElwWthC             string `json:"N_ELW_WTH_C,omitempty" db:"N_ELW_WTH_C" csv:"N_ELW_WTH_C"`
	NElwWthH             string `json:"N_ELW_WTH_H,omitempty" db:"N_ELW_WTH_H" csv:"N_ELW_WTH_H"`
	NRegWETH             string `json:"N_Reg_WET_H,omitempty" db:"N_Reg_WET_H" csv:"N_Reg_WET_H"`
	NEldWetC             string `json:"N_ELD_WET_C,omitempty" db:"N_ELD_WET_C" csv:"N_ELD_WET_C"`
	NEldWetH             string `json:"N_ELD_WET_H,omitempty" db:"N_ELD_WET_H" csv:"N_ELD_WET_H"`
	NRegDTC              string `json:"N_Reg_DT_C,omitempty" db:"N_Reg_DT_C" csv:"N_Reg_DT_C"`
	NEldDtC              string `json:"N_ELD_DT_C,omitempty" db:"N_ELD_DT_C" csv:"N_ELD_DT_C"`
	CRegWSTC             string `json:"C_Reg_WST_C,omitempty" db:"C_Reg_WST_C" csv:"C_Reg_WST_C"`
	CRegWETC             string `json:"C_Reg_WET_C,omitempty" db:"C_Reg_WET_C" csv:"C_Reg_WET_C"`
	CRegWLTC             string `json:"C_Reg_WLT_C,omitempty" db:"C_Reg_WLT_C" csv:"C_Reg_WLT_C"`
	CRegATC              string `json:"C_Reg_AT_C,omitempty" db:"C_Reg_AT_C" csv:"C_Reg_AT_C"`
	CRegWSTH             string `json:"C_Reg_WST_H,omitempty" db:"C_Reg_WST_H" csv:"C_Reg_WST_H"`
	CRegWETH             string `json:"C_Reg_WET_H,omitempty" db:"C_Reg_WET_H" csv:"C_Reg_WET_H"`
	CRegWLTH             string `json:"C_Reg_WLT_H,omitempty" db:"C_Reg_WLT_H" csv:"C_Reg_WLT_H"`
	CEldWstC             string `json:"C_ELD_WST_C,omitempty" db:"C_ELD_WST_C" csv:"C_ELD_WST_C"`
	CEldWetC             string `json:"C_ELD_WET_C,omitempty" db:"C_ELD_WET_C" csv:"C_ELD_WET_C"`
	CEldAtC              string `json:"C_ELD_AT_C,omitempty" db:"C_ELD_AT_C" csv:"C_ELD_AT_C"`
	CEldWstH             string `json:"C_ELD_WST_H,omitempty" db:"C_ELD_WST_H" csv:"C_ELD_WST_H"`
	CEldWetH             string `json:"C_ELD_WET_H,omitempty" db:"C_ELD_WET_H" csv:"C_ELD_WET_H"`
	CEldWltH             string `json:"C_ELD_WLT_H,omitempty" db:"C_ELD_WLT_H" csv:"C_ELD_WLT_H"`
	CEldWltC             string `json:"C_ELD_WLT_C,omitempty" db:"C_ELD_WLT_C" csv:"C_ELD_WLT_C"`
	CRegWDC              string `json:"C_Reg_WD_C,omitempty" db:"C_Reg_WD_C" csv:"C_Reg_WD_C"`
	CRegWTHC             string `json:"C_Reg_WTH_C,omitempty" db:"C_Reg_WTH_C" csv:"C_Reg_WTH_C"`
	CRegCGSTNfreqFwyC    string `json:"C_Reg_CGSTNfreqFwy_C,omitempty" db:"C_Reg_CGSTNfreqFwy_C" csv:"C_Reg_CGSTNfreqFwy_C"`
	CRegCGSTNfreqFwyDayC string `json:"C_Reg_CGSTNfreqFwyDay_C,omitempty" db:"C_Reg_CGSTNfreqFwyDay_C" csv:"C_Reg_CGSTNfreqFwyDay_C"`
	CRegCGSTNfreqUrbC    string `json:"C_Reg_CGSTNfreqUrb_C,omitempty" db:"C_Reg_CGSTNfreqUrb_C" csv:"C_Reg_CGSTNfreqUrb_C"`
	CRegCGSTNfreqUrbDayC string `json:"C_Reg_CGSTNfreqUrbDay_C,omitempty" db:"C_Reg_CGSTNfreqUrbDay_C" csv:"C_Reg_CGSTNfreqUrbDay_C"`
	CRegCGSTNfreqC       string `json:"C_Reg_CGSTNfreq_C,omitempty" db:"C_Reg_CGSTNfreq_C" csv:"C_Reg_CGSTNfreq_C"`
	CRegWDH              string `json:"C_Reg_WD_H,omitempty" db:"C_Reg_WD_H" csv:"C_Reg_WD_H"`
	CRegWTHH             string `json:"C_Reg_WTH_H,omitempty" db:"C_Reg_WTH_H" csv:"C_Reg_WTH_H"`
	CEldWdC              string `json:"C_ELD_WD_C,omitempty" db:"C_ELD_WD_C" csv:"C_ELD_WD_C"`
	CELDCGSTNfreqFwyC    string `json:"C_ELD_CGSTNfreqFwy_C,omitempty" db:"C_ELD_CGSTNfreqFwy_C" csv:"C_ELD_CGSTNfreqFwy_C"`
	CELDCGSTNfreqFwyDayC string `json:"C_ELD_CGSTNfreqFwyDay_C,omitempty" db:"C_ELD_CGSTNfreqFwyDay_C" csv:"C_ELD_CGSTNfreqFwyDay_C"`
	CELDCGSTNfreqUrbC    string `json:"C_ELD_CGSTNfreqUrb_C,omitempty" db:"C_ELD_CGSTNfreqUrb_C" csv:"C_ELD_CGSTNfreqUrb_C"`
	CELDCGSTNfreqUrbDayC string `json:"C_ELD_CGSTNfreqUrbDay_C,omitempty" db:"C_ELD_CGSTNfreqUrbDay_C" csv:"C_ELD_CGSTNfreqUrbDay_C"`
	CELDCGSTNfreqC       string `json:"C_ELD_CGSTNfreq_C,omitempty" db:"C_ELD_CGSTNfreq_C" csv:"C_ELD_CGSTNfreq_C"`
	CEldWdH              string `json:"C_ELD_WD_H,omitempty" db:"C_ELD_WD_H" csv:"C_ELD_WD_H"`
	CRegWOTC             string `json:"C_Reg_WOT_C,omitempty" db:"C_Reg_WOT_C" csv:"C_Reg_WOT_C"`
	CRegPeakTC1          string `json:"C_Reg_PeakT_C1,omitempty" db:"C_Reg_PeakT_C1" csv:"C_Reg_PeakT_C1"`
	CRegPeakTC2          string `json:"C_Reg_PeakT_C2,omitempty" db:"C_Reg_PeakT_C2" csv:"C_Reg_PeakT_C2"`
	CEldWotC             string `json:"C_ELD_WOT_C,omitempty" db:"C_ELD_WOT_C" csv:"C_ELD_WOT_C"`
	CELDPeakTC1          string `json:"C_ELD_PeakT_C1,omitempty" db:"C_ELD_PeakT_C1" csv:"C_ELD_PeakT_C1"`
	CELDPeakTC2          string `json:"C_ELD_PeakT_C2,omitempty" db:"C_ELD_PeakT_C2" csv:"C_ELD_PeakT_C2"`
	CEldWthC             string `json:"C_ELD_WTH_C,omitempty" db:"C_ELD_WTH_C" csv:"C_ELD_WTH_C"`
	CElwWthC             string `json:"C_ELW_WTH_C,omitempty" db:"C_ELW_WTH_C" csv:"C_ELW_WTH_C"`
	CEldWthH             string `json:"C_ELD_WTH_H,omitempty" db:"C_ELD_WTH_H" csv:"C_ELD_WTH_H"`
	CElwWthH             string `json:"C_ELW_WTH_H,omitempty" db:"C_ELW_WTH_H" csv:"C_ELW_WTH_H"`
	CRegDTC              string `json:"C_Reg_DT_C,omitempty" db:"C_Reg_DT_C" csv:"C_Reg_DT_C"`
	CEldDtC              string `json:"C_ELD_DT_C,omitempty" db:"C_ELD_DT_C" csv:"C_ELD_DT_C"`
	LateRule             string `json:"late_rule,omitempty" db:"late_rule" csv:"late_rule"`
	LateTimes            string `json:"late_times,omitempty" db:"late_times" csv:"late_times"`
	LateReason           string `json:"late_reason,omitempty" db:"late_reason" csv:"late_reason"`
	LatePenalty          string `json:"late_penalty,omitempty" db:"late_penalty" csv:"late_penalty"`
	PenaltyFeel          string `json:"penalty_feel,omitempty" db:"penalty_feel" csv:"penalty_feel"`
	FreePAKG             string `json:"Free_PAKG,omitempty" db:"Free_PAKG" csv:"Free_PAKG"`
	Pakg                 string `json:"PAKG,omitempty" db:"PAKG" csv:"PAKG"`
	LateRulersn          string `json:"late_rulersn,omitempty" db:"late_rulersn" csv:"late_rulersn"`
	LateReasonValue      string `json:"late_reason_value,omitempty" db:"late_reason_value" csv:"late_reason_value"`
	LatePenaltyValue     string `json:"late_penalty_value,omitempty" db:"late_penalty_value" csv:"late_penalty_value"`
	LateMin              string `json:"late_min,omitempty" db:"late_min" csv:"late_min"`
	CMUTEYy              string `json:"CMUTE_yy,omitempty" db:"CMUTE_yy" csv:"CMUTE_yy"`
	CMUTEMm              string `json:"CMUTE_mm,omitempty" db:"CMUTE_mm" csv:"CMUTE_mm"`
	CMUTEDd              string `json:"CMUTE_dd,omitempty" db:"CMUTE_dd" csv:"CMUTE_dd"`
	CMUTEW               string `json:"CMUTE_w,omitempty" db:"CMUTE_w" csv:"CMUTE_w"`
	TtdtOrig             string `json:"TTDT_ORIG,omitempty" db:"TTDT_ORIG" csv:"TTDT_ORIG"`
	TtdtEacpt            string `json:"TTDT_EACPT,omitempty" db:"TTDT_EACPT" csv:"TTDT_EACPT"`
	TtdtLacpt            string `json:"TTDT_LACPT,omitempty" db:"TTDT_LACPT" csv:"TTDT_LACPT"`
	TtarrEacpt           string `json:"TTARR_EACPT,omitempty" db:"TTARR_EACPT" csv:"TTARR_EACPT"`
	TtarrLacpt           string `json:"TTARR_LACPT,omitempty" db:"TTARR_LACPT" csv:"TTARR_LACPT"`
	ZipcodeHome          string `json:"zipcode_home,omitempty" db:"zipcode_home" csv:"zipcode_home"`
	RoadHome             string `json:"road_home,omitempty" db:"road_home" csv:"road_home"`
	ZipcodeWork          string `json:"zipcode_work,omitempty" db:"zipcode_work" csv:"zipcode_work"`
	RoadWork             string `json:"road_work,omitempty" db:"road_work" csv:"road_work"`
	CityHome             string `json:"city_home,omitempty" db:"city_home" csv:"city_home"`
	UrbanHome            string `json:"Urban_home,omitempty" db:"Urban_home" csv:"Urban_home"`
	CityWork             string `json:"city_work,omitempty" db:"city_work" csv:"city_work"`
	UrbanWork            string `json:"Urban_work,omitempty" db:"Urban_work" csv:"Urban_work"`
	ConsiderCOVID        string `json:"consider_COVID,omitempty" db:"consider_COVID" csv:"consider_COVID"`
	HabitualDT           string `json:"habitual_DT,omitempty" db:"habitual_DT" csv:"habitual_DT"`
	OdtimeThesame        string `json:"Odtime_thesame,omitempty" db:"Odtime_thesame" csv:"Odtime_thesame"`
	WorkType             string `json:"work_type,omitempty" db:"work_type" csv:"work_type"`
	CarSeatNo            string `json:"CarSeat_No,omitempty" db:"CarSeat_No" csv:"CarSeat_No"`
	StopNo               string `json:"Stop_No,omitempty" db:"Stop_No" csv:"Stop_No"`
	StopPurpRsn          string `json:"Stop_purp_rsn,omitempty" db:"Stop_purp_rsn" csv:"Stop_purp_rsn"`
	OtherValue4          string `json:"otherValue4,omitempty" db:"otherValue4" csv:"otherValue4"`
	TDwalk               string `json:"TDwalk,omitempty" db:"TDwalk" csv:"TDwalk"`
	OVTTwalk             string `json:"OVTTwalk,omitempty" db:"OVTTwalk" csv:"OVTTwalk"`
	TDbike               string `json:"TDbike,omitempty" db:"TDbike" csv:"TDbike"`
	OVTTbike             string `json:"OVTTbike,omitempty" db:"OVTTbike" csv:"OVTTbike"`
	TCbike               string `json:"TCbike,omitempty" db:"TCbike" csv:"TCbike"`
	TDUrbn               string `json:"TD_Urbn,omitempty" db:"TD_Urbn" csv:"TD_Urbn"`
	VUrrd                string `json:"V_URRD,omitempty" db:"V_URRD" csv:"V_URRD"`
	IvttUrrd             string `json:"IVTT_URRD,omitempty" db:"IVTT_URRD" csv:"IVTT_URRD"`
	TdHwy                string `json:"TD_HWY,omitempty" db:"TD_HWY" csv:"TD_HWY"`
	VHwy                 string `json:"V_HWY,omitempty" db:"V_HWY" csv:"V_HWY"`
	IvttHwy              string `json:"IVTT_HWY,omitempty" db:"IVTT_HWY" csv:"IVTT_HWY"`
	TdFwy                string `json:"TD_FWY,omitempty" db:"TD_FWY" csv:"TD_FWY"`
	VFwy                 string `json:"V_FWY,omitempty" db:"V_FWY" csv:"V_FWY"`
	IvttFwy              string `json:"IVTT_FWY,omitempty" db:"IVTT_FWY" csv:"IVTT_FWY"`
	TCFee                string `json:"TC_fee,omitempty" db:"TC_fee" csv:"TC_fee"`
	TCFuel               string `json:"TC_fuel,omitempty" db:"TC_fuel" csv:"TC_fuel"`
	TcPkg                string `json:"TC_PKG,omitempty" db:"TC_PKG" csv:"TC_PKG"`
	IvttPrkg             string `json:"IVTT_PRKG,omitempty" db:"IVTT_PRKG" csv:"IVTT_PRKG"`
	TotalSeat            string `json:"Total_seat,omitempty" db:"Total_seat" csv:"Total_seat"`
	CarSeat              string `json:"CarSeat,omitempty" db:"CarSeat" csv:"CarSeat"`
	Stop                 string `json:"Stop,omitempty" db:"Stop" csv:"Stop"`
	IDveh                string `json:"IDveh,omitempty" db:"IDveh" csv:"IDveh"`
	IvttStop             string `json:"IVTT_STOP,omitempty" db:"IVTT_STOP" csv:"IVTT_STOP"`
	PKGHome1             string `json:"PKG_home_1,omitempty" db:"PKG_home_1" csv:"PKG_home_1"`
	PKGHome2             string `json:"PKG_home_2,omitempty" db:"PKG_home_2" csv:"PKG_home_2"`
	RouteSim             string `json:"route_sim,omitempty" db:"route_sim" csv:"route_sim"`
	ChoiceSN1            string `json:"Choice_SN1,omitempty" db:"Choice_SN1" csv:"Choice_SN1"`
	ChoiceRsn            string `json:"Choice_rsn,omitempty" db:"Choice_rsn" csv:"Choice_rsn"`
	EtCO2A1              string `json:"EtCO2_A1,omitempty" db:"EtCO2_A1" csv:"EtCO2_A1"`
	EtCO2A2              string `json:"EtCO2_A2,omitempty" db:"EtCO2_A2" csv:"EtCO2_A2"`
	EtCO2B               string `json:"EtCO2_B,omitempty" db:"EtCO2_B" csv:"EtCO2_B"`
	EtCO2C               string `json:"EtCO2_C,omitempty" db:"EtCO2_C" csv:"EtCO2_C"`
	EyCO2A1              string `json:"EyCO2_A1,omitempty" db:"EyCO2_A1" csv:"EyCO2_A1"`
	EyCO2A2              string `json:"EyCO2_A2,omitempty" db:"EyCO2_A2" csv:"EyCO2_A2"`
	EyCO2B               string `json:"EyCO2_B,omitempty" db:"EyCO2_B" csv:"EyCO2_B"`
	EyCO2C               string `json:"EyCO2_C,omitempty" db:"EyCO2_C" csv:"EyCO2_C"`
	StWrk                string `json:"ST_WRK,omitempty" db:"ST_WRK" csv:"ST_WRK"`
	Ovtt                 string `json:"OVTT,omitempty" db:"OVTT" csv:"OVTT"`
	TTDTEACPTTime        string `json:"TTDT_EACPT_Time,omitempty" db:"TTDT_EACPT_Time" csv:"TTDT_EACPT_Time"`
	TTDTLACPTTime        string `json:"TTDT_LACPT_Time,omitempty" db:"TTDT_LACPT_Time" csv:"TTDT_LACPT_Time"`
	TDFwyMap             string `json:"TDFwy_map,omitempty" db:"TDFwy_map" csv:"TDFwy_map"`
	TDHwyMap             string `json:"TDHwy_map,omitempty" db:"TDHwy_map" csv:"TDHwy_map"`
	TDUrbanMap           string `json:"TDUrban_map,omitempty" db:"TDUrban_map" csv:"TDUrban_map"`
	TTMap                string `json:"TT_map,omitempty" db:"TT_map" csv:"TT_map"`
	TTFwyMap             string `json:"TTFwy_map,omitempty" db:"TTFwy_map" csv:"TTFwy_map"`
	TTHwyMap             string `json:"TTHwy_map,omitempty" db:"TTHwy_map" csv:"TTHwy_map"`
	TTUrbanMap           string `json:"TTUrban_map,omitempty" db:"TTUrban_map" csv:"TTUrban_map"`
	IvttOrig             string `json:"IVTT_ORIG,omitempty" db:"IVTT_ORIG" csv:"IVTT_ORIG"`
	TttOrig              string `json:"TTT_ORIG,omitempty" db:"TTT_ORIG" csv:"TTT_ORIG"`
	EptOrig              string `json:"EPT_ORIG,omitempty" db:"EPT_ORIG" csv:"EPT_ORIG"`
	IvttsOrig            string `json:"IVTTS_ORIG,omitempty" db:"IVTTS_ORIG" csv:"IVTTS_ORIG"`
	IvttlOrig            string `json:"IVTTL_ORIG,omitempty" db:"IVTTL_ORIG" csv:"IVTTL_ORIG"`
	IvttaOrig            string `json:"IVTTA_ORIG,omitempty" db:"IVTTA_ORIG" csv:"IVTTA_ORIG"`
	COETTORIG            string `json:"COE_TT_ORIG,omitempty" db:"COE_TT_ORIG" csv:"COE_TT_ORIG"`
	TtcgstnOrig          string `json:"TTCGSTN_ORIG,omitempty" db:"TTCGSTN_ORIG" csv:"TTCGSTN_ORIG"`
	TTFwyORIG            string `json:"TTFwy_ORIG,omitempty" db:"TTFwy_ORIG" csv:"TTFwy_ORIG"`
	TTHwyORIG            string `json:"TTHwy_ORIG,omitempty" db:"TTHwy_ORIG" csv:"TTHwy_ORIG"`
	TTUrbanORIG          string `json:"TTUrban_ORIG,omitempty" db:"TTUrban_ORIG" csv:"TTUrban_ORIG"`
	EtCO2ORIG            string `json:"EtCO2_ORIG,omitempty" db:"EtCO2_ORIG" csv:"EtCO2_ORIG"`
	EyCO2ORIG            string `json:"EyCO2_ORIG,omitempty" db:"EyCO2_ORIG" csv:"EyCO2_ORIG"`
	ECtripORIG           string `json:"ECtrip_ORIG,omitempty" db:"ECtrip_ORIG" csv:"ECtrip_ORIG"`
	ECyrORIG             string `json:"ECyr_ORIG,omitempty" db:"ECyr_ORIG" csv:"ECyr_ORIG"`
	ECtripA              string `json:"ECtrip_A,omitempty" db:"ECtrip_A" csv:"ECtrip_A"`
	IvttEa               string `json:"IVTT_EA,omitempty" db:"IVTT_EA" csv:"IVTT_EA"`
	TttEa                string `json:"TTT_EA,omitempty" db:"TTT_EA" csv:"TTT_EA"`
	TtdtEa               string `json:"TTDT_EA,omitempty" db:"TTDT_EA" csv:"TTDT_EA"`
	EptEa                string `json:"EPT_EA,omitempty" db:"EPT_EA" csv:"EPT_EA"`
	IvttsEa              string `json:"IVTTS_EA,omitempty" db:"IVTTS_EA" csv:"IVTTS_EA"`
	IvttlEa              string `json:"IVTTL_EA,omitempty" db:"IVTTL_EA" csv:"IVTTL_EA"`
	IvttaEa              string `json:"IVTTA_EA,omitempty" db:"IVTTA_EA" csv:"IVTTA_EA"`
	COETTEA              string `json:"COE_TT_EA,omitempty" db:"COE_TT_EA" csv:"COE_TT_EA"`
	TtcgstnEa            string `json:"TTCGSTN_EA,omitempty" db:"TTCGSTN_EA" csv:"TTCGSTN_EA"`
	TTFwyEA              string `json:"TTFwy_EA,omitempty" db:"TTFwy_EA" csv:"TTFwy_EA"`
	TTHwyEA              string `json:"TTHwy_EA,omitempty" db:"TTHwy_EA" csv:"TTHwy_EA"`
	TTUrbanEA            string `json:"TTUrban_EA,omitempty" db:"TTUrban_EA" csv:"TTUrban_EA"`
	EtCO2EA              string `json:"EtCO2_EA,omitempty" db:"EtCO2_EA" csv:"EtCO2_EA"`
	EyCO2EA              string `json:"EyCO2_EA,omitempty" db:"EyCO2_EA" csv:"EyCO2_EA"`
	ECtripEA             string `json:"ECtrip_EA,omitempty" db:"ECtrip_EA" csv:"ECtrip_EA"`
	ECyrEA               string `json:"ECyr_EA,omitempty" db:"ECyr_EA" csv:"ECyr_EA"`
	DEtCO2EA             string `json:"DEtCO2_EA,omitempty" db:"DEtCO2_EA" csv:"DEtCO2_EA"`
	DEyCO2EA             string `json:"DEyCO2_EA,omitempty" db:"DEyCO2_EA" csv:"DEyCO2_EA"`
	DEtCostEA            string `json:"DEtCost_EA,omitempty" db:"DEtCost_EA" csv:"DEtCost_EA"`
	DEyCostEA            string `json:"DEyCost_EA,omitempty" db:"DEyCost_EA" csv:"DEyCost_EA"`
	IvttLa               string `json:"IVTT_LA,omitempty" db:"IVTT_LA" csv:"IVTT_LA"`
	TttLa                string `json:"TTT_LA,omitempty" db:"TTT_LA" csv:"TTT_LA"`
	TtdtLa               string `json:"TTDT_LA,omitempty" db:"TTDT_LA" csv:"TTDT_LA"`
	EptLa                string `json:"EPT_LA,omitempty" db:"EPT_LA" csv:"EPT_LA"`
	IvttsLa              string `json:"IVTTS_LA,omitempty" db:"IVTTS_LA" csv:"IVTTS_LA"`
	IvttlLa              string `json:"IVTTL_LA,omitempty" db:"IVTTL_LA" csv:"IVTTL_LA"`
	IvttaLa              string `json:"IVTTA_LA,omitempty" db:"IVTTA_LA" csv:"IVTTA_LA"`
	COETTLA              string `json:"COE_TT_LA,omitempty" db:"COE_TT_LA" csv:"COE_TT_LA"`
	TtcgstnLa            string `json:"TTCGSTN_LA,omitempty" db:"TTCGSTN_LA" csv:"TTCGSTN_LA"`
	TTFwyLA              string `json:"TTFwy_LA,omitempty" db:"TTFwy_LA" csv:"TTFwy_LA"`
	TTHwyLA              string `json:"TTHwy_LA,omitempty" db:"TTHwy_LA" csv:"TTHwy_LA"`
	TTUrbanLA            string `json:"TTUrban_LA,omitempty" db:"TTUrban_LA" csv:"TTUrban_LA"`
	EtCO2LA              string `json:"EtCO2_LA,omitempty" db:"EtCO2_LA" csv:"EtCO2_LA"`
	EyCO2LA              string `json:"EyCO2_LA,omitempty" db:"EyCO2_LA" csv:"EyCO2_LA"`
	ECtripLA             string `json:"ECtrip_LA,omitempty" db:"ECtrip_LA" csv:"ECtrip_LA"`
	ECyrLA               string `json:"ECyr_LA,omitempty" db:"ECyr_LA" csv:"ECyr_LA"`
	DEtCO2LA             string `json:"DEtCO2_LA,omitempty" db:"DEtCO2_LA" csv:"DEtCO2_LA"`
	DEyCO2LA             string `json:"DEyCO2_LA,omitempty" db:"DEyCO2_LA" csv:"DEyCO2_LA"`
	DEtCostLA            string `json:"DEtCost_LA,omitempty" db:"DEtCost_LA" csv:"DEtCost_LA"`
	DEyCostLA            string `json:"DEyCost_LA,omitempty" db:"DEyCost_LA" csv:"DEyCost_LA"`
	WTPChoiceNPA         string `json:"WTPChoice_NP_A,omitempty" db:"WTPChoice_NP_A" csv:"WTPChoice_NP_A"`
	WtpNpA               string `json:"WTP_NP_A,omitempty" db:"WTP_NP_A" csv:"WTP_NP_A"`
	WTPChoiceNPB         string `json:"WTPChoice_NP_B,omitempty" db:"WTPChoice_NP_B" csv:"WTPChoice_NP_B"`
	WtpNpB               string `json:"WTP_NP_B,omitempty" db:"WTP_NP_B" csv:"WTP_NP_B"`
	WTPChoiceNPC         string `json:"WTPChoice_NP_C,omitempty" db:"WTPChoice_NP_C" csv:"WTPChoice_NP_C"`
	WtpNpC               string `json:"WTP_NP_C,omitempty" db:"WTP_NP_C" csv:"WTP_NP_C"`
	WTPChoicePA          string `json:"WTPChoice_P_A,omitempty" db:"WTPChoice_P_A" csv:"WTPChoice_P_A"`
	WtpPA                string `json:"WTP_P_A,omitempty" db:"WTP_P_A" csv:"WTP_P_A"`
	WTPChoicePB          string `json:"WTPChoice_P_B,omitempty" db:"WTPChoice_P_B" csv:"WTPChoice_P_B"`
	WtpPB                string `json:"WTP_P_B,omitempty" db:"WTP_P_B" csv:"WTP_P_B"`
	WTPChoicePC          string `json:"WTPChoice_P_C,omitempty" db:"WTPChoice_P_C" csv:"WTPChoice_P_C"`
	WtpPC                string `json:"WTP_P_C,omitempty" db:"WTP_P_C" csv:"WTP_P_C"`
	WTPZerorsn           string `json:"WTP_zerorsn,omitempty" db:"WTP_zerorsn" csv:"WTP_zerorsn"`
	WTPZerorsnOther      string `json:"WTP_zerorsn_other,omitempty" db:"WTP_zerorsn_other" csv:"WTP_zerorsn_other"`
	WTPUsage             string `json:"WTP_usage,omitempty" db:"WTP_usage" csv:"WTP_usage"`
	WTPUsage1Type        string `json:"WTP_usage1type,omitempty" db:"WTP_usage1type" csv:"WTP_usage1type"`
	WTPUsageRsn          string `json:"WTP_usage_rsn,omitempty" db:"WTP_usage_rsn" csv:"WTP_usage_rsn"`
	WTADTBnpChoice       string `json:"WTA_DT_Bnp_choice,omitempty" db:"WTA_DT_Bnp_choice" csv:"WTA_DT_Bnp_choice"`
	WTADTBnp             string `json:"WTA_DT_Bnp,omitempty" db:"WTA_DT_Bnp" csv:"WTA_DT_Bnp"`
	WTADTCnpChoice       string `json:"WTA_DT_Cnp_choice,omitempty" db:"WTA_DT_Cnp_choice" csv:"WTA_DT_Cnp_choice"`
	WTADTCnp             string `json:"WTA_DT_Cnp,omitempty" db:"WTA_DT_Cnp" csv:"WTA_DT_Cnp"`
	WTADTBpChoice        string `json:"WTA_DT_Bp_choice,omitempty" db:"WTA_DT_Bp_choice" csv:"WTA_DT_Bp_choice"`
	WTADTBp1             string `json:"WTA_DT_Bp1,omitempty" db:"WTA_DT_Bp1" csv:"WTA_DT_Bp1"`
	WTADTBp2             string `json:"WTA_DT_Bp2,omitempty" db:"WTA_DT_Bp2" csv:"WTA_DT_Bp2"`
	WTADTCpChoice        string `json:"WTA_DT_Cp_choice,omitempty" db:"WTA_DT_Cp_choice" csv:"WTA_DT_Cp_choice"`
	WTADTCp1             string `json:"WTA_DT_Cp1,omitempty" db:"WTA_DT_Cp1" csv:"WTA_DT_Cp1"`
	WTADTCp2             string `json:"WTA_DT_Cp2,omitempty" db:"WTA_DT_Cp2" csv:"WTA_DT_Cp2"`
	WTADTZero            string `json:"WTA_DT_zero,omitempty" db:"WTA_DT_zero" csv:"WTA_DT_zero"`
	WTADTZerorsn         string `json:"WTA_DT_zerorsn,omitempty" db:"WTA_DT_zerorsn" csv:"WTA_DT_zerorsn"`
	WtaDtRmn             string `json:"WTA_DT_RMN,omitempty" db:"WTA_DT_RMN" csv:"WTA_DT_RMN"`
	WTADTRMNRsn          string `json:"WTA_DT_RMN_rsn,omitempty" db:"WTA_DT_RMN_rsn" csv:"WTA_DT_RMN_rsn"`
	COVIDDTChoice        string `json:"COVID_DT_Choice,omitempty" db:"COVID_DT_Choice" csv:"COVID_DT_Choice"`
	CarpoolDFreqBefore   string `json:"carpoolD_Freq_before,omitempty" db:"carpoolD_Freq_before" csv:"carpoolD_Freq_before"`
	CarpoolPFreqBefore   string `json:"carpoolP_Freq_before,omitempty" db:"carpoolP_Freq_before" csv:"carpoolP_Freq_before"`
	UberFreqBefore       string `json:"Uber_Freq_before,omitempty" db:"Uber_Freq_before" csv:"Uber_Freq_before"`
	CScarFreqBefore      string `json:"CScar_Freq_before,omitempty" db:"CScar_Freq_before" csv:"CScar_Freq_before"`
	CSotoFreqBefore      string `json:"CSoto_Freq_before,omitempty" db:"CSoto_Freq_before" csv:"CSoto_Freq_before"`
	CSebikeFreqBefore    string `json:"CSebike_Freq_before,omitempty" db:"CSebike_Freq_before" csv:"CSebike_Freq_before"`
	CSbikeFreqBefore     string `json:"CSbike_Freq_before,omitempty" db:"CSbike_Freq_before" csv:"CSbike_Freq_before"`
	CarpoolDFreqNow      string `json:"carpoolD_Freq_now,omitempty" db:"carpoolD_Freq_now" csv:"carpoolD_Freq_now"`
	CarpoolPFreqNow      string `json:"carpoolP_Freq_now,omitempty" db:"carpoolP_Freq_now" csv:"carpoolP_Freq_now"`
	UberFreqNow          string `json:"Uber_Freq_now,omitempty" db:"Uber_Freq_now" csv:"Uber_Freq_now"`
	CScarFreqNow         string `json:"CScar_Freq_now,omitempty" db:"CScar_Freq_now" csv:"CScar_Freq_now"`
	CSmotoFreqNow        string `json:"CSmoto_Freq_now,omitempty" db:"CSmoto_Freq_now" csv:"CSmoto_Freq_now"`
	CSebikeFreqNow       string `json:"CSebike_Freq_now,omitempty" db:"CSebike_Freq_now" csv:"CSebike_Freq_now"`
	CSbikeFreqNow        string `json:"CSbike_Freq_now,omitempty" db:"CSbike_Freq_now" csv:"CSbike_Freq_now"`
	CarSharePREF1        string `json:"CarShare_PREF1,omitempty" db:"CarShare_PREF1" csv:"CarShare_PREF1"`
	CarSharePREF2        string `json:"CarShare_PREF2,omitempty" db:"CarShare_PREF2" csv:"CarShare_PREF2"`
	CarSharePREF3        string `json:"CarShare_PREF3,omitempty" db:"CarShare_PREF3" csv:"CarShare_PREF3"`
	CarSharePREF4        string `json:"CarShare_PREF4,omitempty" db:"CarShare_PREF4" csv:"CarShare_PREF4"`
	CarSharePREF5        string `json:"CarShare_PREF5,omitempty" db:"CarShare_PREF5" csv:"CarShare_PREF5"`
	CarSharePREF6        string `json:"CarShare_PREF6,omitempty" db:"CarShare_PREF6" csv:"CarShare_PREF6"`
	CarSharePREF7        string `json:"CarShare_PREF7,omitempty" db:"CarShare_PREF7" csv:"CarShare_PREF7"`
	CarSharePREF8        string `json:"CarShare_PREF8,omitempty" db:"CarShare_PREF8" csv:"CarShare_PREF8"`
	CplDPax              string `json:"CPL_D_PAX,omitempty" db:"CPL_D_PAX" csv:"CPL_D_PAX"`
	IvttPuD              string `json:"IVTT_PU_D,omitempty" db:"IVTT_PU_D" csv:"IVTT_PU_D"`
	CplPPax              string `json:"CPL_P_PAX,omitempty" db:"CPL_P_PAX" csv:"CPL_P_PAX"`
	IvttPuP              string `json:"IVTT_PU_P,omitempty" db:"IVTT_PU_P" csv:"IVTT_PU_P"`
	OvttWtPax            string `json:"OVTT_WT_PAX,omitempty" db:"OVTT_WT_PAX" csv:"OVTT_WT_PAX"`
	UberPAX              string `json:"Uber_PAX,omitempty" db:"Uber_PAX" csv:"Uber_PAX"`
	IVTTPUUber           string `json:"IVTT_PU_Uber,omitempty" db:"IVTT_PU_Uber" csv:"IVTT_PU_Uber"`
	OVTTWTUber           string `json:"OVTT_WT_Uber,omitempty" db:"OVTT_WT_Uber" csv:"OVTT_WT_Uber"`
	CscarPax             string `json:"CSCAR_PAX,omitempty" db:"CSCAR_PAX" csv:"CSCAR_PAX"`
	IvttrtC              string `json:"IVTTRT_C,omitempty" db:"IVTTRT_C" csv:"IVTTRT_C"`
	OVTTwalkCSCAR        string `json:"OVTTwalk_CSCAR,omitempty" db:"OVTTwalk_CSCAR" csv:"OVTTwalk_CSCAR"`
	CsmkPax              string `json:"CSMK_PAX,omitempty" db:"CSMK_PAX" csv:"CSMK_PAX"`
	IvttrtMb             string `json:"IVTTRT_MB,omitempty" db:"IVTTRT_MB" csv:"IVTTRT_MB"`
	OVTTwalkCSMB         string `json:"OVTTwalk_CSMB,omitempty" db:"OVTTwalk_CSMB" csv:"OVTTwalk_CSMB"`
	IvttrtEb             string `json:"IVTTRT_EB,omitempty" db:"IVTTRT_EB" csv:"IVTTRT_EB"`
	OVTTwalkCSEB         string `json:"OVTTwalk_CSEB,omitempty" db:"OVTTwalk_CSEB" csv:"OVTTwalk_CSEB"`
	IvttrtB              string `json:"IVTTRT_B,omitempty" db:"IVTTRT_B" csv:"IVTTRT_B"`
	OVTTwalkCSB          string `json:"OVTTwalk_CSB,omitempty" db:"OVTTwalk_CSB" csv:"OVTTwalk_CSB"`
	CarSharePREF         string `json:"CarShare_PREF,omitempty" db:"CarShare_PREF" csv:"CarShare_PREF"`
	ChoiceSN3            string `json:"Choice_SN3,omitempty" db:"Choice_SN3" csv:"Choice_SN3"`
	IvttA                string `json:"IVTT_A,omitempty" db:"IVTT_A" csv:"IVTT_A"`
	IvttAA               string `json:"IVTT_A_A,omitempty" db:"IVTT_A_A" csv:"IVTT_A_A"`
	IVTTPkupA            string `json:"IVTT_pkup_A,omitempty" db:"IVTT_pkup_A" csv:"IVTT_pkup_A"`
	IvttPrkgA            string `json:"IVTT_PRKG_A,omitempty" db:"IVTT_PRKG_A" csv:"IVTT_PRKG_A"`
	OvttA                string `json:"OVTT_A,omitempty" db:"OVTT_A" csv:"OVTT_A"`
	OVTTwalkA            string `json:"OVTTwalk_A,omitempty" db:"OVTTwalk_A" csv:"OVTTwalk_A"`
	OVTTbikeA            string `json:"OVTTbike_A,omitempty" db:"OVTTbike_A" csv:"OVTTbike_A"`
	OVTTwaitA            string `json:"OVTTwait_A,omitempty" db:"OVTTwait_A" csv:"OVTTwait_A"`
	TttA                 string `json:"TTT_A,omitempty" db:"TTT_A" csv:"TTT_A"`
	EptA                 string `json:"EPT_A,omitempty" db:"EPT_A" csv:"EPT_A"`
	TcA                  string `json:"TC_A,omitempty" db:"TC_A" csv:"TC_A"`
	TcEA                 string `json:"TC_E_A,omitempty" db:"TC_E_A" csv:"TC_E_A"`
	TCFeeA               string `json:"TC_fee_A,omitempty" db:"TC_fee_A" csv:"TC_fee_A"`
	TcPkgA               string `json:"TC_PKG_A,omitempty" db:"TC_PKG_A" csv:"TC_PKG_A"`
	PaxA                 string `json:"PAX_A,omitempty" db:"PAX_A" csv:"PAX_A"`
	EtCO2A               string `json:"EtCO2_A,omitempty" db:"EtCO2_A" csv:"EtCO2_A"`
	EyCO2A               string `json:"EyCO2_A,omitempty" db:"EyCO2_A" csv:"EyCO2_A"`
	ECyrA                string `json:"ECyr_A,omitempty" db:"ECyr_A" csv:"ECyr_A"`
	IvttB                string `json:"IVTT_B,omitempty" db:"IVTT_B" csv:"IVTT_B"`
	IvttAB               string `json:"IVTT_A_B,omitempty" db:"IVTT_A_B" csv:"IVTT_A_B"`
	IVTTPkupB            string `json:"IVTT_pkup_B,omitempty" db:"IVTT_pkup_B" csv:"IVTT_pkup_B"`
	IvttPrkgB            string `json:"IVTT_PRKG_B,omitempty" db:"IVTT_PRKG_B" csv:"IVTT_PRKG_B"`
	OvttB                string `json:"OVTT_B,omitempty" db:"OVTT_B" csv:"OVTT_B"`
	OVTTwalkB            string `json:"OVTTwalk_B,omitempty" db:"OVTTwalk_B" csv:"OVTTwalk_B"`
	OVTTbikeB            string `json:"OVTTbike_B,omitempty" db:"OVTTbike_B" csv:"OVTTbike_B"`
	OVTTwaitB            string `json:"OVTTwait_B,omitempty" db:"OVTTwait_B" csv:"OVTTwait_B"`
	TttB                 string `json:"TTT_B,omitempty" db:"TTT_B" csv:"TTT_B"`
	EptB                 string `json:"EPT_B,omitempty" db:"EPT_B" csv:"EPT_B"`
	TcB                  string `json:"TC_B,omitempty" db:"TC_B" csv:"TC_B"`
	TcEB                 string `json:"TC_E_B,omitempty" db:"TC_E_B" csv:"TC_E_B"`
	TCFeeB               string `json:"TC_fee_B,omitempty" db:"TC_fee_B" csv:"TC_fee_B"`
	TcPkgB               string `json:"TC_PKG_B,omitempty" db:"TC_PKG_B" csv:"TC_PKG_B"`
	PaxB                 string `json:"PAX_B,omitempty" db:"PAX_B" csv:"PAX_B"`
	ECtripB              string `json:"ECtrip_B,omitempty" db:"ECtrip_B" csv:"ECtrip_B"`
	ECyrB                string `json:"ECyr_B,omitempty" db:"ECyr_B" csv:"ECyr_B"`
	IvttC                string `json:"IVTT_C,omitempty" db:"IVTT_C" csv:"IVTT_C"`
	IvttAC               string `json:"IVTT_A_C,omitempty" db:"IVTT_A_C" csv:"IVTT_A_C"`
	IVTTPkupC            string `json:"IVTT_pkup_C,omitempty" db:"IVTT_pkup_C" csv:"IVTT_pkup_C"`
	IvttPrkgC            string `json:"IVTT_PRKG_C,omitempty" db:"IVTT_PRKG_C" csv:"IVTT_PRKG_C"`
	OvttC                string `json:"OVTT_C,omitempty" db:"OVTT_C" csv:"OVTT_C"`
	OVTTwalkC            string `json:"OVTTwalk_C,omitempty" db:"OVTTwalk_C" csv:"OVTTwalk_C"`
	OVTTbikeC            string `json:"OVTTbike_C,omitempty" db:"OVTTbike_C" csv:"OVTTbike_C"`
	OVTTwaitC            string `json:"OVTTwait_C,omitempty" db:"OVTTwait_C" csv:"OVTTwait_C"`
	TttC                 string `json:"TTT_C,omitempty" db:"TTT_C" csv:"TTT_C"`
	EptC                 string `json:"EPT_C,omitempty" db:"EPT_C" csv:"EPT_C"`
	TcC                  string `json:"TC_C,omitempty" db:"TC_C" csv:"TC_C"`
	TcEC                 string `json:"TC_E_C,omitempty" db:"TC_E_C" csv:"TC_E_C"`
	TCFeeC               string `json:"TC_fee_C,omitempty" db:"TC_fee_C" csv:"TC_fee_C"`
	TcPkgC               string `json:"TC_PKG_C,omitempty" db:"TC_PKG_C" csv:"TC_PKG_C"`
	PaxC                 string `json:"PAX_C,omitempty" db:"PAX_C" csv:"PAX_C"`
	ECtripC              string `json:"ECtrip_C,omitempty" db:"ECtrip_C" csv:"ECtrip_C"`
	ECyrC                string `json:"ECyr_C,omitempty" db:"ECyr_C" csv:"ECyr_C"`
	IvttD                string `json:"IVTT_D,omitempty" db:"IVTT_D" csv:"IVTT_D"`
	IvttAD               string `json:"IVTT_A_D,omitempty" db:"IVTT_A_D" csv:"IVTT_A_D"`
	IVTTPkupD            string `json:"IVTT_pkup_D,omitempty" db:"IVTT_pkup_D" csv:"IVTT_pkup_D"`
	IvttPrkgD            string `json:"IVTT_PRKG_D,omitempty" db:"IVTT_PRKG_D" csv:"IVTT_PRKG_D"`
	OvttD                string `json:"OVTT_D,omitempty" db:"OVTT_D" csv:"OVTT_D"`
	OVTTwalkD            string `json:"OVTTwalk_D,omitempty" db:"OVTTwalk_D" csv:"OVTTwalk_D"`
	OVTTbikeD            string `json:"OVTTbike_D,omitempty" db:"OVTTbike_D" csv:"OVTTbike_D"`
	OVTTwaitD            string `json:"OVTTwait_D,omitempty" db:"OVTTwait_D" csv:"OVTTwait_D"`
	TttD                 string `json:"TTT_D,omitempty" db:"TTT_D" csv:"TTT_D"`
	EptD                 string `json:"EPT_D,omitempty" db:"EPT_D" csv:"EPT_D"`
	TcED                 string `json:"TC_E_D,omitempty" db:"TC_E_D" csv:"TC_E_D"`
	TCFeeD               string `json:"TC_fee_D,omitempty" db:"TC_fee_D" csv:"TC_fee_D"`
	TcPkgD               string `json:"TC_PKG_D,omitempty" db:"TC_PKG_D" csv:"TC_PKG_D"`
	PaxD                 string `json:"PAX_D,omitempty" db:"PAX_D" csv:"PAX_D"`
	EtCO2D               string `json:"EtCO2_D,omitempty" db:"EtCO2_D" csv:"EtCO2_D"`
	EyCO2D               string `json:"EyCO2_D,omitempty" db:"EyCO2_D" csv:"EyCO2_D"`
	ECtripD              string `json:"ECtrip_D,omitempty" db:"ECtrip_D" csv:"ECtrip_D"`
	ECyrD                string `json:"ECyr_D,omitempty" db:"ECyr_D" csv:"ECyr_D"`
	IvttE                string `json:"IVTT_E,omitempty" db:"IVTT_E" csv:"IVTT_E"`
	IvttAE               string `json:"IVTT_A_E,omitempty" db:"IVTT_A_E" csv:"IVTT_A_E"`
	IVTTPkupE            string `json:"IVTT_pkup_E,omitempty" db:"IVTT_pkup_E" csv:"IVTT_pkup_E"`
	IvttPrkgE            string `json:"IVTT_PRKG_E,omitempty" db:"IVTT_PRKG_E" csv:"IVTT_PRKG_E"`
	OvttE                string `json:"OVTT_E,omitempty" db:"OVTT_E" csv:"OVTT_E"`
	OVTTwalkE            string `json:"OVTTwalk_E,omitempty" db:"OVTTwalk_E" csv:"OVTTwalk_E"`
	OVTTbikeE            string `json:"OVTTbike_E,omitempty" db:"OVTTbike_E" csv:"OVTTwait_E"`
	OVTTwaitE            string `json:"OVTTwait_E,omitempty" db:"OVTTwait_E" csv:"OVTTwait_E"`
	TttE                 string `json:"TTT_E,omitempty" db:"TTT_E" csv:"TTT_E"`
	EptE                 string `json:"EPT_E,omitempty" db:"EPT_E" csv:"EPT_E"`
	TcE                  string `json:"TC_E,omitempty" db:"TC_E" csv:"TC_E"`
	TcEE                 string `json:"TC_E_E,omitempty" db:"TC_E_E" csv:"TC_E_E"`
	TCFeeE               string `json:"TC_fee_E,omitempty" db:"TC_fee_E" csv:"TC_fee_E"`
	TcPkgE               string `json:"TC_PKG_E,omitempty" db:"TC_PKG_E" csv:"TC_PKG_E"`
	PaxE                 string `json:"PAX_E,omitempty" db:"PAX_E" csv:"PAX_E"`
	EtCO2E               string `json:"EtCO2_E,omitempty" db:"EtCO2_E" csv:"EtCO2_E"`
	EyCO2E               string `json:"EyCO2_E,omitempty" db:"EyCO2_E" csv:"EyCO2_E"`
	ECtripE              string `json:"ECtrip_E,omitempty" db:"ECtrip_E" csv:"ECtrip_E"`
	ECyrE                string `json:"ECyr_E,omitempty" db:"ECyr_E" csv:"ECyr_E"`
	IvttF                string `json:"IVTT_F,omitempty" db:"IVTT_F" csv:"IVTT_F"`
	IvttAF               string `json:"IVTT_A_F,omitempty" db:"IVTT_A_F" csv:"IVTT_A_F"`
	IVTTPkupF            string `json:"IVTT_pkup_F,omitempty" db:"IVTT_pkup_F" csv:"IVTT_pkup_F"`
	IvttPrkgF            string `json:"IVTT_PRKG_F,omitempty" db:"IVTT_PRKG_F" csv:"IVTT_PRKG_F"`
	OvttF                string `json:"OVTT_F,omitempty" db:"OVTT_F" csv:"OVTT_F"`
	OVTTwalkF            string `json:"OVTTwalk_F,omitempty" db:"OVTTwalk_F" csv:"OVTTwalk_F"`
	OVTTbikeF            string `json:"OVTTbike_F,omitempty" db:"OVTTbike_F" csv:"OVTTbike_F"`
	OVTTwaitF            string `json:"OVTTwait_F,omitempty" db:"OVTTwait_F" csv:"OVTTwait_F"`
	TttF                 string `json:"TTT_F,omitempty" db:"TTT_F" csv:"TTT_F"`
	EptF                 string `json:"EPT_F,omitempty" db:"EPT_F" csv:"EPT_F"`
	TcF                  string `json:"TC_F,omitempty" db:"TC_F" csv:"TC_F"`
	TcEF                 string `json:"TC_E_F,omitempty" db:"TC_E_F" csv:"TC_E_F"`
	TCFeeF               string `json:"TC_fee_F,omitempty" db:"TC_fee_F" csv:"TC_fee_F"`
	TcPkgF               string `json:"TC_PKG_F,omitempty" db:"TC_PKG_F" csv:"TC_PKG_F"`
	PaxF                 string `json:"PAX_F,omitempty" db:"PAX_F" csv:"PAX_F"`
	EtCO2F               string `json:"EtCO2_F,omitempty" db:"EtCO2_F" csv:"EtCO2_F"`
	EyCO2F               string `json:"EyCO2_F,omitempty" db:"EyCO2_F" csv:"EyCO2_F"`
	ECtripF              string `json:"ECtrip_F,omitempty" db:"ECtrip_F" csv:"ECtrip_F"`
	ECyrF                string `json:"ECyr_F,omitempty" db:"ECyr_F" csv:"ECyr_F"`
	IvttG                string `json:"IVTT_G,omitempty" db:"IVTT_G" csv:"IVTT_G"`
	IvttAG               string `json:"IVTT_A_G,omitempty"  db:"IVTT_A_G" csv:"IVTT_A_G"`
	IVTTPkupG            string `json:"IVTT_pkup_G,omitempty"  db:"IVTT_pkup_G" csv:"IVTT_pkup_G"`
	IvttPrkgG            string `json:"IVTT_PRKG_G,omitempty"  db:"IVTT_PRKG_G" csv:"IVTT_PRKG_G"`
	OvttG                string `json:"OVTT_G,omitempty"  db:"OVTT_G" csv:"OVTT_G"`
	OVTTwalkG            string `json:"OVTTwalk_G,omitempty"  db:"OVTTwalk_G" csv:"OVTTwalk_G"`
	OVTTbikeG            string `json:"OVTTbike_G,omitempty"  db:"OVTTbike_G" csv:"OVTTbike_G"`
	OVTTwaitG            string `json:"OVTTwait_G,omitempty"  db:"OVTTwait_G" csv:"OVTTwait_G"`
	TttG                 string `json:"TTT_G,omitempty"  db:"TTT_G" csv:"TTT_G"`
	EptG                 string `json:"EPT_G,omitempty"  db:"EPT_G" csv:"EPT_G"`
	TcG                  string `json:"TC_G,omitempty"  db:"TC_G" csv:"TC_G"`
	TcEG                 string `json:"TC_E_G,omitempty"  db:"TC_E_G" csv:"TC_E_G"`
	TCFeeG               string `json:"TC_fee_G,omitempty"  db:"TC_fee_G" csv:"TC_fee_G"`
	TcPkgG               string `json:"TC_PKG_G,omitempty"  db:"TC_PKG_G" csv:"TC_PKG_G"`
	PaxG                 string `json:"PAX_G,omitempty"  db:"PAX_G" csv:"PAX_G"`
	EtCO2G               string `json:"EtCO2_G,omitempty"  db:"EtCO2_G" csv:"EtCO2_G"`
	EyCO2G               string `json:"EyCO2_G,omitempty"  db:"EyCO2_G" csv:"EyCO2_G"`
	ECtripG              string `json:"ECtrip_G,omitempty"  db:"ECtrip_G" csv:"ECtrip_G"`
	ECyrG                string `json:"ECyr_G,omitempty"  db:"ECyr_G" csv:"ECyr_G"`
	IvttH                string `json:"IVTT_H,omitempty"  db:"IVTT_H" csv:"IVTT_H"`
	IvttAH               string `json:"IVTT_A_H,omitempty"  db:"IVTT_A_H" csv:"IVTT_A_H"`
	IVTTPkupH            string `json:"IVTT_pkup_H,omitempty"  db:"IVTT_pkup_H" csv:"IVTT_pkup_H"`
	IvttPrkgH            string `json:"IVTT_PRKG_H,omitempty"  db:"IVTT_PRKG_H" csv:"IVTT_PRKG_H"`
	OvttH                string `json:"OVTT_H,omitempty"  db:"OVTT_H" csv:"OVTT_H"`
	OVTTwalkH            string `json:"OVTTwalk_H,omitempty"  db:"OVTTwalk_H" csv:"OVTTwalk_H"`
	OVTTbikeH            string `json:"OVTTbike_H,omitempty"  db:"OVTTbike_H" csv:"OVTTbike_H"`
	OVTTwaitH            string `json:"OVTTwait_H,omitempty"  db:"OVTTwait_H" csv:"OVTTwait_H"`
	TttH                 string `json:"TTT_H,omitempty"  db:"TTT_H" csv:"TTT_H"`
	EptH                 string `json:"EPT_H,omitempty"  db:"EPT_H" csv:"EPT_H"`
	TcH                  string `json:"TC_H,omitempty"  db:"TC_H" csv:"TC_H"`
	TcEH                 string `json:"TC_E_H,omitempty"  db:"TC_E_H" csv:"TC_E_H"`
	TCFeeH               string `json:"TC_fee_H,omitempty"  db:"TC_fee_H" csv:"TC_fee_H"`
	TcPkgH               string `json:"TC_PKG_H,omitempty"  db:"TC_PKG_H" csv:"TC_PKG_H"`
	PaxH                 string `json:"PAX_H,omitempty"  db:"PAX_H" csv:"PAX_H"`
	EtCO2H               string `json:"EtCO2_H,omitempty"  db:"EtCO2_H" csv:"EtCO2_H"`
	EyCO2H               string `json:"EyCO2_H,omitempty"  db:"EyCO2_H" csv:"EyCO2_H"`
	ECtripH              string `json:"ECtrip_H,omitempty"  db:"ECtrip_H" csv:"ECtrip_H"`
	ECyrH                string `json:"ECyr_H,omitempty"  db:"ECyr_H" csv:"ECyr_H"`
	WTAModeBnpChoice     string `json:"WTA_Mode_Bnp_choice,omitempty"  db:"WTA_Mode_Bnp_choice" csv:"WTA_Mode_Bnp_choice"`
	WTAModeBnp           string `json:"WTA_Mode_Bnp,omitempty"  db:"WTA_Mode_Bnp" csv:"WTA_Mode_Bnp"`
	WTAModeCnpChoice     string `json:"WTA_Mode_Cnp_choice,omitempty"  db:"WTA_Mode_Cnp_choice" csv:"WTA_Mode_Cnp_choice"`
	WTAModeCnp           string `json:"WTA_Mode_Cnp,omitempty"  db:"WTA_Mode_Cnp" csv:"WTA_Mode_Cnp"`
	WTAModeDnpChoice     string `json:"WTA_Mode_Dnp_choice,omitempty"  db:"WTA_Mode_Dnp_choice" csv:"WTA_Mode_Dnp_choice"`
	WTAModeDnp           string `json:"WTA_Mode_Dnp,omitempty"  db:"WTA_Mode_Dnp" csv:"WTA_Mode_Dnp"`
	WTAModeEnpChoice     string `json:"WTA_Mode_Enp_choice,omitempty"  db:"WTA_Mode_Enp_choice" csv:"WTA_Mode_Enp_choice"`
	WTAModeEnp           string `json:"WTA_Mode_Enp,omitempty"  db:"WTA_Mode_Enp" csv:"WTA_Mode_Enp"`
	WTAModeFnpChoice     string `json:"WTA_Mode_Fnp_choice,omitempty"  db:"WTA_Mode_Fnp_choice" csv:"WTA_Mode_Fnp_choice"`
	WTAModeFnp           string `json:"WTA_Mode_Fnp,omitempty"  db:"WTA_Mode_Fnp" csv:"WTA_Mode_Fnp"`
	WTAModeGnpChoice     string `json:"WTA_Mode_Gnp_choice,omitempty"  db:"WTA_Mode_Gnp_choice" csv:"WTA_Mode_Gnp_choice"`
	WTAModeGnp           string `json:"WTA_Mode_Gnp,omitempty"  db:"WTA_Mode_Gnp" csv:"WTA_Mode_Gnp"`
	WTAModeHnpChoice     string `json:"WTA_Mode_Hnp_choice,omitempty"  db:"WTA_Mode_Hnp_choice" csv:"WTA_Mode_Hnp_choice"`
	WTAModeHnp           string `json:"WTA_Mode_Hnp,omitempty"  db:"WTA_Mode_Hnp" csv:"WTA_Mode_Hnp"`
	WTAModeZero          string `json:"WTA_Mode_zero,omitempty"  db:"WTA_Mode_zero" csv:"WTA_Mode_zero"`
	WTAModeZerorsn       string `json:"WTA_Mode_zerorsn,omitempty"  db:"WTA_Mode_zerorsn" csv:"WTA_Mode_zerorsn"`
	WTAModeRMN           string `json:"WTA_Mode_RMN,omitempty"  db:"WTA_Mode_RMN" csv:"WTA_Mode_RMN"`
	WTAModeRMNRsn        string `json:"WTA_Mode_RMN_rsn,omitempty"  db:"WTA_Mode_RMN_rsn" csv:"WTA_Mode_RMN_rsn"`
	COVIDModeChoice      string `json:"COVID_Mode_Choice,omitempty"  db:"COVID_Mode_Choice" csv:"COVID_Mode_Choice"`
	WTAAttract           string `json:"WTA_Attract,omitempty"  db:"WTA_Attract" csv:"WTA_Attract"`
	WTAAttractRsn        string `json:"WTA_Attract_rsn,omitempty"  db:"WTA_Attract_rsn" csv:"WTA_Attract_rsn"`
	Gender               string `json:"gender,omitempty"  db:"gender" csv:"gender"`
	Age                  string `json:"age,omitempty"  db:"age" csv:"age"`
	Education            string `json:"education,omitempty"  db:"education" csv:"education"`
	IncomePers           string `json:"income_pers,omitempty"  db:"income_pers" csv:"income_pers"`
	WorkPeriod           string `json:"work_period,omitempty"  db:"work_period" csv:"work_period"`
	WorkPeriod2          string `json:"work_period_2,omitempty"  db:"work_period_2" csv:"work_period_2"`
	ResidentStatus       string `json:"resident_status,omitempty"  db:"resident_status" csv:"resident_status"`
	ResidentHH           string `json:"resident_HH,omitempty"  db:"resident_HH" csv:"resident_HH"`
	IncomePersValue      string `json:"income_pers_value,omitempty"  db:"income_pers_value" csv:"income_pers_value"`
	IncomeHH             string `json:"income_HH,omitempty"  db:"income_HH" csv:"income_HH"`
	ResidentKid          string `json:"resident_kid,omitempty"  db:"resident_kid" csv:"resident_kid"`
	IncomeHHValue        string `json:"income_HH_value,omitempty"  db:"income_HH_value" csv:"income_HH_value"`
	DrivingEXPC          string `json:"driving_EXPC,omitempty"  db:"driving_EXPC" csv:"driving_EXPC"`
	ScooterHH            string `json:"Scooter_HH,omitempty"  db:"Scooter_HH" csv:"Scooter_HH"`
	ScooterHHUse         string `json:"Scooter_HH_use,omitempty"  db:"Scooter_HH_use" csv:"Scooter_HH_use"`
	CarHH                string `json:"Car_HH,omitempty"  db:"Car_HH" csv:"Car_HH"`
	CarHHUse             string `json:"Car_HH_use,omitempty"  db:"Car_HH_use" csv:"Car_HH_use"`
	VehicleINCR          string `json:"Vehicle_INCR,omitempty"  db:"Vehicle_INCR" csv:"Vehicle_INCR"`
	VehicleDECR          string `json:"Vehicle_DECR,omitempty"  db:"Vehicle_DECR" csv:"Vehicle_DECR"`
	VehicleNnChange      string `json:"Vehicle_nn_change,omitempty"  db:"Vehicle_nn_change" csv:"Vehicle_nn_change"`
	WFHday               string `json:"WFHday,omitempty"  db:"WFHday" csv:"WFHday"`
	Sr04                 string `json:"SR04,omitempty" db:"SR04" csv:"SR04"`
	Sp05                 string `json:"SP05,omitempty" db:"SP05" csv:"SP05"`
	Sp09                 string `json:"SP09,omitempty" db:"SP09" csv:"SP09"`
	Sr10                 string `json:"SR10,omitempty" db:"SR10" csv:"SR10"`
	Sr12                 string `json:"SR12,omitempty" db:"SR12" csv:"SR12"`
	Sr14                 string `json:"SR14,omitempty" db:"SR14" csv:"SR14"`
	Sp15                 string `json:"SP15,omitempty" db:"SP15" csv:"SP15"`
	Sr18                 string `json:"SR18,omitempty" db:"SR18" csv:"SR18"`
	Sp19                 string `json:"SP19,omitempty" db:"SP19" csv:"SP19"`
	Sp21                 string `json:"SP21,omitempty" db:"SP21" csv:"SP21"`
	Sr24                 string `json:"SR24,omitempty" db:"SR24" csv:"SR24"`
	Sp31                 string `json:"SP31,omitempty" db:"SP31" csv:"SP31"`
	Sp35                 string `json:"SP35,omitempty" db:"SP35" csv:"SP35"`
	Sp37                 string `json:"SP37,omitempty" db:"SP37" csv:"SP37"`
	Sr38                 string `json:"SR38,omitempty" db:"SR38" csv:"SR38"`
	Sp39                 string `json:"SP39,omitempty" db:"SP39" csv:"SP39"`
	Sr44                 string `json:"SR44,omitempty" db:"SR44" csv:"SR44"`
	Sr46                 string `json:"SR46,omitempty" db:"SR46" csv:"SR46"`
	Sp47                 string `json:"SP47,omitempty" db:"SP47" csv:"SP47"`
	Sp48                 string `json:"SP48,omitempty" db:"SP48" csv:"SP48"`
	FilterCity           string `json:"filter_city,omitempty" db:"FilterCity" csv:"FilterCity"`
	FilterLicience       string `json:"filter_licence,omitempty" db:"FilterLicience" csv:"FilterLicience"`
	FilterMode           string `json:"filter_mode,omitempty" db:"FilterMode" csv:"FilterMode"`
}
