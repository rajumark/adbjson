# shell acpi

Show ACPI power and thermal information.

## Command
```bash
adbjson shell acpi [flags]
```

## Description
Executes `adb shell acpi` and outputs the result as structured JSON. Shows power sources and thermal device information.

## Examples

### Show cooling device state
```bash
./adbjson shell acpi --cooling
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Cooling Devices",
      "content": "Cooling 0: mdmss-0 no state information\nCooling 0: pause-cpu3 0 of 1\nCooling 0: battery no state information\nCooling 0: sdr0_lte_dsc_kr 0 of 255\nCooling 0: aoss-1 no state information\nCooling 0: cpuss-1 no state information\nCooling 0: socd no state information\nCooling 0: cpu-cluster1 0 of 11\nCooling 0: mmw-ific0 no state information\nCooling 0: pm7250b-ibat-lvl0 no state information\nCooling 0: cpu-hotplug1 0 of 1\nCooling 0: xo-therm no state information\nCooling 0: thermal-pause-40 0 of 1\nCooling 0: sub1_scg_fr1_cc no state information\nCooling 0: video no state information\nCooling 0: pause-cpu2 0 of 1\nCooling 0: usb no state information\nCooling 0: zeroc-0 no state information\nCooling 0: aoss-0 no state information\nCooling 0: 5G-pa-therm no state information\nCooling 0: mmw2 no state information\nCooling 0: pm7250b-vbat-lvl2 no state information\nCooling 0: pm6150l_tz no state information\nCooling 0: sub1_mcg_fr1_cc no state information\nCooling 0: pa_nr_sdr1_dsc_kr 0 of 255\nCooling 0: nspss-1 no state information\nCooling 0: pm6150l-bcl-lvl2 no state information\nCooling 0: conn-therm no state information\nCooling 0: mmodem_lte_dsc_kr 0 of 255\nCooling 0: camera-0 no state information\nCooling 0: thermal-pause-1 0 of 1\nCooling 0: quiet-therm no state information\nCooling 0: cpufreq-cpu4 0 of 11\nCooling 0: mmw1 no state information\nCooling 0: cpu-1-5 no state information\nCooling 0: wlan 0 of 4\nCooling 0: zeroc-1 no state information\nCooling 0: pause-cpu6 0 of 1\nCooling 0: sub1-lte-cc no state information\nCooling 0: pa_lte_sdr1_dsc_kr 0 of 255\nCooling 0: cpu-0-3 no state information\nCooling 0: cpu-1-2 no state information\nCooling 0: cam-flash-therm no state information\nCooling 0: secondary_charge 0 of 0\nCooling 0: gpuss-1 no state information\nCooling 0: cpu-hotplug7 0 of 1\nCooling 0: fuse_temp no state information\nCooling 0: display-fps 0 of 3\nCooling 0: sdr1 no state information\nCooling 0: cpu-1-4 no state information\nCooling 0: mmw3_dsc_kr 0 of 255\nCooling 0: mdmss-3 no state information\nCooling 0: thermal-pause-E0 0 of 1\nCooling 0: bcl-warn no state information\nCooling 0: sdr1_nr_dsc_kr 0 of 255\nCooling 0: cpu-0-2 no state information\nCooling 0: cpu-1-1 no state information\nCooling 0: rec-therm no state information\nCooling 0: battery 7 of 11\nCooling 0: gpuss-0 no state information\nCooling 0: cpu-hotplug6 0 of 1\nCooling 0: back_tmo no state information\nCooling 0: ufs 0 of 2\nCooling 0: sdr1-pa no state information\nCooling 0: mmw1_dsc_kr 0 of 255\nCooling 0: mdmss-1 no state information\nCooling 0: thermal-pause-8 0 of 1\nCooling 0: pa no state information\nCooling 0: sdr1_lte_dsc_kr 0 of 255\nCooling 0: cpu-0-0 no state information\nCooling 0: cpu-1-0 no state information\nCooling 0: pa-therm2 no state information\nCooling 0: cpu-cluster0 0 of 7\nCooling 0: cpu-1-7 no state information\nCooling 0: cpu-hotplug2 0 of 1\nCooling 0: front_temp no state information\nCooling 0: thermal-pause-80 0 of 1\nCooling 0: sub1_scg_fr2_cc no state information\nCooling 0: pa_nr_sdr1_scg_kr 0 of 255\nCooling 0: ddr no state information\nCooling 0: thermal-pause-4 0 of 1\nCooling 0: wireless no state information\nCooling 0: modem_nr_scg_dsc_kr 0 of 255\nCooling 0: pm7250b-bcl-lvl2 no state information\nCooling 0: cpuss-0 no state information\nCooling 0: cam-therm no state information\nCooling 0: panel0-backlight 0 of 255\nCooling 0: mmw3 no state information\nCooling 0: cpu-1-6 no state information\nCooling 0: cpu-hotplug0 0 of 1\nCooling 0: pm7250b_tz no state information\nCooling 0: thermal-pause-C0 0 of 1\nCooling 0: sub1_mcg_fr2_cc no state information\nCooling 0: pa_nr_sdr0_scg_kr 0 of 255\nCooling 0: nspss-2 no state information\nCooling 0: thermal-pause-2 0 of 1\nCooling 0: chg-therm no state information\nCooling 0: modem_nr_dsc_kr 0 of 255\nCooling 0: pm7250b-bcl-lvl1 no state information\nCooling 0: pause-cpu1 0 of 1\nCooling 0: mmi_battery no state information\nCooling 0: kgsl 0 of 6\nCooling 0: pm7250b-vbat-lvl1 no state information\nCooling 0: pm6450_tz no state information\nCooling 0: pause-cpu7 0 of 1\nCooling 0: pa_nr_sdr0_dsc_kr 0 of 255\nCooling 0: nspss-0 no state information\nCooling 0: cpu-1-3 no state information\nCooling 0: modem_vdd 0 of 1\nCooling 0: pm7250b-bcl-lvl0 no state information\nCooling 0: pause-cpu0 0 of 1\nCooling 0: cpufreq-cpu0 0 of 7\nCooling 0: mmw0 no state information\nCooling 0: pm7250b-vbat-lvl0 no state information\nCooling 0: mmw_ific_dsc_kr 0 of 255\nCooling 0: thermal-pause-20 0 of 1\nCooling 0: sub1-modem-cfg no state information\nCooling 0: pa_lte_sdr0_dsc_kr 0 of 255\nCooling 0: pm6150l-bcl-lvl1 no state information\nCooling 0: msm-therm no state information\nCooling 0: primary_charge 0 of 0\nCooling 0: cpu-hotplug3 0 of 1\nCooling 0: back_tmo1 no state information\nCooling 0: ddr-cdev 0 of 1\nCooling 0: sdr0 no state information\nCooling 0: mmw2_dsc_kr 0 of 255\nCooling 0: mdmss-2 no state information\nCooling 0: pause-cpu5 0 of 1\nCooling 0: pa1 no state information\nCooling 0: sdr0_nr_dsc_kr 0 of 255\nCooling 0: cpu-0-1 no state information\nCooling 0: pm6150l-bcl-lvl0 no state information\nCooling 0: pa-therm1 no state information\nCooling 0: gpu 0 of 6\nCooling 0: pm7250b-ibat-lvl1 no state information\nCooling 0: cpu-hotplug5 0 of 1\nCooling 0: back_temp no state information\nCooling 0: cx_cdev 0 of 1\nCooling 0: sdr0-pa no state information\nCooling 0: mmw0_dsc_kr 0 of 255"
    }
  ],
  "count": 1
}
```

### Show temperatures
```bash
./adbjson shell acpi --temperatures
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Thermal Information",
      "content": "Thermal 0: 38.5 degrees C\nThermal 0: 35.0 degrees C\nThermal 0: 38.9 degrees C\nThermal 0: 38.9 degrees C\nThermal 0: 1.2 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: -46.-8 degrees C\nThermal 0: 3835.0 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.2 degrees C\nThermal 0: -40.0 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 37.8 degrees C\nThermal 0: 3857.0 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 433.4 degrees C\nThermal 0: 4086.3 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: -40.0 degrees C\nThermal 0: 37.8 degrees C\nThermal 0: 3839.3 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 39.5 degrees C\nThermal 0: 37.8 degrees C\nThermal 0: -40.0 degrees C\nThermal 0: 37.5 degrees C\nThermal 0: 3692.5 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.2 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.9 degrees C\nThermal 0: 37.8 degrees C\nThermal 0: 3703.1 degrees C\nThermal 0: 37.5 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 3690.2 degrees C\nThermal 0: 38.9 degrees C\nThermal 0: 36.8 degrees C\nThermal 0: 3703.1 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 36.1 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.2 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 38.2 degrees C\nThermal 0: 3823.9 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 37.8 degrees C\nThermal 0: 3862.5 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 3802.0 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 35.0 degrees C\nThermal 0: 433.4 degrees C\nThermal 0: 37.0 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 38.2 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 433.4 degrees C\nThermal 0: -273.0 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 3802.0 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 37.0 degrees C\nThermal 0: 38.5 degrees C\nThermal 0: 3709.6 degrees C\nThermal 0: 38.9 degrees C\nThermal 0: 0.0 degrees C\nThermal 0: 3682.2 degrees C\nThermal 0: -46.-8 degrees C\nThermal 0: 3863.1 degrees C\nThermal 0: -273.0 degrees C"
    },
    {
      "name": "Permission Errors",
      "content": "acpi: /sys/class/android_usb: Permission denied\nacpi: /sys/class/scsi_generic/sg6: Permission denied\nacpi: /sys/class/scsi_generic/sg4: Permission denied\nacpi: /sys/class/scsi_generic/sg0: Permission denied\nacpi: /sys/class/scsi_generic/sg7: Permission denied\nacpi: /sys/class/scsi_generic/sg5: Permission denied\nacpi: /sys/class/scsi_generic/sg3: Permission denied\nacpi: /sys/class/scsi_generic/sg8: Permission denied\nacpi: /sys/class/sensors: Permission denied\nacpi: /sys/class/wakeup: Permission denied\nacpi: /sys/class/devcoredump/disabled: Permission denied\nacpi: /sys/class/backlight/panel0-backlight: Permission denied\nacpi: /sys/class/uio: Permission denied\nacpi: /sys/class/leds: Permission denied\nacpi: /sys/class/extcon: Permission denied\nacpi: /sys/class/devfreq-event: Permission denied\nacpi: /sys/class/kgsl: Permission denied\nacpi: /sys/class/typec: Permission denied\nacpi: /sys/class/capsense: Permission denied\nacpi: /sys/class/typec_mux: Permission denied\nacpi: /sys/class/devfreq: Permission denied\nacpi: /sys/class/firmware/timeout: Permission denied\nacpi: /sys/class/fingerprint: Permission denied\nacpi: /sys/class/frsm_amp/calre: Permission denied\nacpi: /sys/class/frsm_amp/monitor: Permission denied\nacpi: /sys/class/frsm_amp/livedata: Permission denied\nacpi: /sys/class/frsm_amp/scene: Permission denied\nacpi: /sys/class/frsm_amp/ndev: Permission denied\nacpi: /sys/class/frsm_amp/batt: Permission denied\nacpi: /sys/class/frsm_amp/func: Permission denied\nacpi: /sys/class/frsm_amp/fsalgo: Permission denied\nacpi: /sys/class/frsm_amp/state: Permission denied\nacpi: /sys/class/frsm_amp/spkon: Permission denied\nacpi: /sys/class/frsm_amp/tunings: Permission denied\nacpi: /sys/class/frsm_amp/init: Permission denied\nacpi: /sys/class/frsm_amp/regs: Permission denied\nacpi: /sys/class/qcom-battery/restrict_cur: Permission denied\nacpi: /sys/class/qcom-battery/fake_soc: Permission denied\nacpi: /sys/class/qcom-battery/battery_parallel_cell_count: Permission denied\nacpi: /sys/class/qcom-battery/wireless_fw_update: Permission denied\nacpi: /sys/class/qcom-battery/ship_mode_en: Permission denied\nacpi: /sys/class/qcom-battery/charge_control_en: Permission denied\nacpi: /sys/class/qcom-battery/flash_active: Permission denied\nacpi: /sys/class/qcom-battery/moisture_detection_status: Permission denied\nacpi: /sys/class/qcom-battery/wireless_boost_en: Permission denied\nacpi: /sys/class/qcom-battery/restrict_chg: Permission denied\nacpi: /sys/class/qcom-battery/wireless_fw_crc: Permission denied\nacpi: /sys/class/qcom-battery/wireless_fw_version: Permission denied\nacpi: /sys/class/qcom-battery/wireless_fw_force_update: Permission denied\nacpi: /sys/class/qcom-battery/usb_typec_compliant: Permission denied\nacpi: /sys/class/qcom-battery/wireless_type: Permission denied\nacpi: /sys/class/qcom-battery/moisture_detection_en: Permission denied\nacpi: /sys/class/qcom-battery/wireless_fw_update_time_ms: Permission denied\nacpi: /sys/class/qcom-battery/usb_real_type: Permission denied\nacpi: /sys/class/input/input9: Permission denied\nacpi: /sys/class/input/event9: Permission denied\nacpi: /sys/class/input/input7: Permission denied\nacpi: /sys/class/input/event7: Permission denied\nacpi: /sys/class/input/input5: Permission denied\nacpi: /sys/class/input/event5: Permission denied\nacpi: /sys/class/input/input3: Permission denied\nacpi: /sys/class/input/event3: Permission denied\nacpi: /sys/class/input/input6: Permission denied\nacpi: /sys/class/input/event6: Permission denied\nacpi: /sys/class/input/input4: Permission denied\nacpi: /sys/class/input/event4: Permission denied\nacpi: /sys/class/zram-control/hot_remove: Permission denied\nacpi: /sys/class/zram-control/hot_add: Permission denied\nacpi: /sys/class/rtc: Permission denied\nacpi: /sys/class/drm/version: Permission denied"
    }
  ],
  "count": 2
}
```

### Show everything (default)
```bash
./adbjson shell acpi --all
```

**Output:**
```json
{
  "sections": [
    {
      "name": "Cooling Devices",
      "content": "Cooling 0: mdmss-0 no state information\nCooling 0: pause-cpu3 0 of 1\nCooling 0: battery no state information\n..."
    },
    {
      "name": "Thermal Information", 
      "content": "Thermal 0: 39.5 degrees C\nThermal 0: 35.0 degrees C\nThermal 0: 39.9 degrees C\n..."
    },
    {
      "name": "Permission Errors",
      "content": "acpi: /sys/class/android_usb: Permission denied\nacpi: /sys/class/scsi_generic/sg6: Permission denied\n..."
    }
  ],
  "count": 3
}
```

## Response Fields

- **sections** (array): List of ACPI information sections
- **count** (number): Total number of sections

### Section Fields

- **name** (string): Section name (e.g., "Cooling Devices", "Thermal Information", "Permission Errors")
- **content** (string): Raw content of the section

## Flags

- `-b, --batteries`: Show batteries
- `-a, --adapters`: Show power adapters  
- `-c, --cooling`: Show cooling device state
- `-t, --temperatures`: Show temperatures
- `-V, --all`: Show everything (default if no flags specified)

## Common Sections

### Cooling Devices
- CPU cooling states and throttling information
- GPU cooling device status
- Battery cooling management
- Thermal pause mechanisms
- Display and camera cooling
- Wireless and modem thermal management

### Thermal Information
- Temperature sensor readings from various components
- CPU, GPU, and modem temperatures
- Battery and charging thermal data
- Ambient and case temperature measurements
- Invalid sensor readings (-273°C indicates sensor issues)

### Permission Errors
- System files that require root access to read
- Normal for non-root access on Android devices
- Does not affect the validity of thermal and cooling data

## Notes

- Requires Android device with ACPI support
- Permission errors are normal for non-root access
- Temperature readings may include invalid sensors (-273°C)
- Cooling device states show current thermal management activity
- Useful for monitoring device thermal performance and cooling behavior
- Some devices may not support all ACPI options
- Battery and adapter information may not be available on all devices

## Related Commands

- `adbjson shell dumpsys battery` - Battery information
- `adbjson shell dumpsys thermal` - Thermal service information  
- `adbjson shell dumpsys power` - Power management information
