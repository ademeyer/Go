package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Request is a bank transactiona
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Timi   []int   `json:"-"`
}

type DeviceInfo_Gen2 struct {
	// Data
	Name        string `json:"name"`
	ID          string `json:"id"`
	Mac         string `json:"mac"`
	Slot        int    `json:"slot"`
	Model       string `json:"model"`
	API_Version int    `json:"gen"`
	FW_ID       string `json:"fw_id"`
	FW_Version  string `json:"ver"`
	App         string `json:"app"`
	Auth_EN     bool   `json:"auth_en"`
	Auth_Domain string `json:"auth_domain"`
	Profile     string `json:"profile"`
}

// var data = `{
// 	"user": "Scrooge McDuck",
// 	"type": "deposit",
// 	"amount": 1234.4
// }
// `

// var data = "{"name":null,"id":"shellypro3em-e05a1b3331fc","mac":"E05A1B3331FC","slot":0,"model":"SPEM-003CEBEU","gen":2,"fw_id":"20240822-121054/1.4.3-g143ff62","ver":"1.4.3","app":"Pro3EM","auth_en":false,"auth_domain":null,"profile":"triphase"}"
var data = "{{\"name\":null,\"id\":\"shellypro3em-e05a1b3331fc\",\"mac\":\"E05A1B3331FC\",\"slot\":0,\"model\":\"SPEM-003CEBEU\",\"gen\":2,\"fw_id\":\"20240822-121054/1.4.3-g143ff62\",\"ver\":\"1.4.3\",\"app\":\"Pro3EM\",\"auth_en\":false,\"auth_domain\":null,\"profile\":\"triphase\"}}"

type capacityStatus struct {
	SerialNumbers []struct {
		Device_ID []struct {
			Capacity uint32 `json:"encharge_capacity"`
		} `json:"492416000033"`
	} `json:"serial_nums"`
}

var enphase = `
{
    "inventory": {
        "serial_nums": {
            "202419011131": {
                "device_type": 22,
                "com_interface_str": "USB",
                "device_id": "202419011131",
                "admin_state": 43,
                "admin_state_str": "ENS_DEVICE_STATE_READY",
                "msg_retry_count": 0,
                "part_number": "834-01927-r03",
                "assembly_number": "861-00802-r08",
                "app_fw_version": "3.0.369_rel/31.33",
                "ibl_fw_version": "",
                "swift_asic_fw_version": "",
                "bmu_fw_version": "",
                "submodule_count": 1,
                "submodules": {
                    "202419011131": {
                        "device_type": 24,
                        "admin_state": 1,
                        "part_number": "800-01821-r05",
                        "assembly_number": "880-01821-r05",
                        "dmir": {
                            "part_number": "546-00003-01",
                            "assembly_number": "01"
                        },
                        "procload": {
                            "part_number": "522-00003-01",
                            "assembly_number": "3.0.369_rel/31.33"
                        }
                    }
                }
            },
            "492416000033": {
                "device_type": 13,
                "com_interface_str": "CAN",
                "device_id": "492416000033",
                "admin_state": 6,
                "admin_state_str": "ENCHG_STATE_READY",
                "reported_grid_mode": "grid-tied",
                "phase": "ph-b",
                "der_index": 2,
                "encharge_revision": 3,
                "encharge_capacity": 5000,
                "encharge_rated_power": 3840,
                "reported_enc_grid_state": "grid-tied",
                "msg_retry_count": 2,
                "part_number": "836-01890-r19",
                "assembly_number": "892-00040-r19",
                "app_fw_version": "3.0.6582_rel/31.11",
                "ibl_fw_version": "3.1.813-8c003b",
                "swift_asic_fw_version": "001.002.1.7.2",
                "bmu_fw_version": "3.8.63",
                "submodule_count": 7,
                "submodules": {
                    "122410011120": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.31.1-D31"
                        },
                        "procload": {
                            "part_number": "521-00008-r01",
                            "assembly_number": "4.31.1-D31"
                        }
                    },
                    "122410011224": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.31.1-D31"
                        },
                        "procload": {
                            "part_number": "521-00008-r01",
                            "assembly_number": "4.31.1-D31"
                        }
                    },
                    "122410011261": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.31.1-D31"
                        },
                        "procload": {
                            "part_number": "521-00008-r01",
                            "assembly_number": "4.31.1-D31"
                        }
                    },
                    "122410012048": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.31.1-D31"
                        },
                        "procload": {
                            "part_number": "521-00008-r01",
                            "assembly_number": "4.31.1-D31"
                        }
                    },
                    "122410017270": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.31.1-D31"
                        },
                        "procload": {
                            "part_number": "521-00008-r01",
                            "assembly_number": "4.31.1-D31"
                        }
                    },
                    "122410017532": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.31.1-D31"
                        },
                        "procload": {
                            "part_number": "521-00008-r01",
                            "assembly_number": "4.31.1-D31"
                        }
                    },
                    "492415004103": {
                        "device_type": 15,
                        "admin_state": 1,
                        "part_number": "800-00331-r01",
                        "assembly_number": "880-00331-r01",
                        "dmir": {
                            "part_number": "546-00002-01",
                            "assembly_number": "01"
                        },
                        "procload": {
                            "part_number": "522-00002-01",
                            "assembly_number": "3.0.6582_rel/31.11"
                        }
                    }
                }
            },
            "492416000038": {
                "device_type": 13,
                "com_interface_str": "CAN",
                "device_id": "492416000038",
                "admin_state": 6,
                "admin_state_str": "ENCHG_STATE_READY",
                "reported_grid_mode": "grid-tied",
                "phase": "ph-a",
                "der_index": 1,
                "encharge_revision": 3,
                "encharge_capacity": 5000,
                "encharge_rated_power": 3840,
                "reported_enc_grid_state": "grid-tied",
                "msg_retry_count": 0,
                "part_number": "836-01890-r19",
                "assembly_number": "892-00040-r19",
                "app_fw_version": "3.0.7790_rel/31.33",
                "ibl_fw_version": "3.1.813-8c003b",
                "swift_asic_fw_version": "001.002.1.7.2",
                "bmu_fw_version": "3.8.69",
                "submodule_count": 7,
                "submodules": {
                    "122410018234": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.60.1-D60"
                        },
                        "procload": {
                            "part_number": "52100008R000",
                            "assembly_number": "4.60.1-D60"
                        }
                    },
                    "122410018383": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.60.1-D60"
                        },
                        "procload": {
                            "part_number": "52100008R000",
                            "assembly_number": "4.60.1-D60"
                        }
                    },
                    "122410018412": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.60.1-D60"
                        },
                        "procload": {
                            "part_number": "52100008R000",
                            "assembly_number": "4.60.1-D60"
                        }
                    },
                    "122411005846": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.60.1-D60"
                        },
                        "procload": {
                            "part_number": "52100008R000",
                            "assembly_number": "4.60.1-D60"
                        }
                    },
                    "122411005904": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.60.1-D60"
                        },
                        "procload": {
                            "part_number": "52100008R000",
                            "assembly_number": "4.60.1-D60"
                        }
                    },
                    "122411005909": {
                        "device_type": 14,
                        "admin_state": 1,
                        "part_number": "800-01729-r02",
                        "assembly_number": "880-01691-r42",
                        "dmir": {
                            "part_number": "549-00057-r00",
                            "assembly_number": "4.60.1-D60"
                        },
                        "procload": {
                            "part_number": "52100008R000",
                            "assembly_number": "4.60.1-D60"
                        }
                    },
                    "492415004075": {
                        "device_type": 15,
                        "admin_state": 1,
                        "part_number": "800-00331-r01",
                        "assembly_number": "880-00331-r01",
                        "dmir": {
                            "part_number": "546-00002-01",
                            "assembly_number": "01"
                        },
                        "procload": {
                            "part_number": "522-00002-01",
                            "assembly_number": "3.0.7790_rel/31.33"
                        }
                    }
                }
            }
        }
    },
    "secctrl": {
        "shutdown": false,
        "freq_bias_hz": -0.062062397599220279,
        "voltage_bias_v": -0.743945837020874,
        "freq_bias_hz_q8": -99,
        "voltage_bias_v_q5": -23,
        "freq_bias_hz_phaseb": -0.062062397599220279,
        "voltage_bias_v_phaseb": -0.743945837020874,
        "freq_bias_hz_q8_phaseb": -99,
        "voltage_bias_v_q5_phaseb": -23,
        "freq_bias_hz_phasec": 0.0,
        "voltage_bias_v_phasec": 0.0,
        "freq_bias_hz_q8_phasec": 0,
        "voltage_bias_v_q5_phasec": 0,
        "configured_backup_soc": 0,
        "adjusted_backup_soc": 0,
        "agg_soc": 54,
        "Max_energy": 10000,
        "ENC_agg_soc": 54,
        "ENC_agg_soh": 100,
        "ENC_agg_backup_energy": 0,
        "ENC_agg_avail_energy": 5400,
        "Enc_commissioned_capacity": 10000,
        "Enc_max_available_capacity": 10000,
        "ACB_agg_soc": 0,
        "ACB_agg_energy": 0,
        "VLS_Limit": 5,
        "soc_rec_enabled": true,
        "soc_recovery_entry": 0,
        "soc_recovery_exit": 10,
        "Commission_in_progress": false,
        "ESS_in_progress": false
    },
    "relay": {
        "mains_admin_state": "closed",
        "mains_oper_state": "closed",
        "srt12_state": "open",
        "srt13_state": "open",
        "der1_state": 0,
        "der2_state": 0,
        "der3_state": 0,
        "Enchg_grid_mode": "grid-tied",
        "Solar_grid_mode": "unknown"
    }
}
`

func main() {

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(enphase), &data); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	serialNums := data["inventory"].(map[string]interface{})["serial_nums"].(map[string]interface{})
	device := serialNums["492416000033"].(map[string]interface{})
	enchargeCapacity := device["encharge_capacity"].(float64)

	fmt.Printf("Encharge Capacity for device 492416000033: %d\n", uint64(enchargeCapacity))

	// rdr := strings.NewReader(enphase) // Simulate a file/socket reader

	// // Decode request
	// dec := json.NewDecoder(rdr)

	// var req capacityStatus
	// if err := dec.Decode(&req); err != nil {
	// 	log.Fatalf("error: can't decode - %s", err)
	// }

	// print the parsed struct
	//fmt.Printf("%s\n%s\n%s\n%d\n%s\n%d\n%s\n%s\n%s\n", req.Name, req.ID, req.Mac, req.Slot, req.Model, req.API_Version, req.FW_ID, req.FW_Version, req.App)
	//fmt.Printf("got: %+v\n", req)

	// // Create response
	// prevBalance := 1_000_000.0 // Loaded from database
	// resp := map[string]interface{}{
	// 	"ok": true,
	// 	"balance": prevBalance + req.Amount,
	// }

	// // Encode response
	// enc := json.NewDecoder(os.Stdout)
	// if err := enc.Encode(resp); err != nil {
	// 	log.Fatalf("errpr: can't encode - %s", err)
	// }

	// fmt.Printf("got: %+v\n", resp)

}
