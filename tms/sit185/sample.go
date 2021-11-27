package sit185

const sitChile = `
/00005 00000/7250/18 346 1201
1.  DISTRESS COSPAS-SARSAT POSITION CONFIRMED ALERT
2.  MSG NO 00005  CHMCC REF NO 01924
3.  DETECTED AT 12 DEC 18 1201 UTC BY MEOSAR
4.  DETECTION FREQUENCY  406.0278 MHZ
5.  COUNTRY OF BEACON REGISTRATION  SPAIN
6.  USER CLASS  DISTRESS
7.  EMERGENCY CODE  NIL
8.  POSITIONS
        CONFIRMED - 20 31.8S  070 14.0W
        DOPPLER A - NIL
        DOPPLER B - NIL
        DOA       - 20 31.8S  070 14.0W
        EXPECTED ACCURACY UNKNOWN
        ALTITUDE 00617 METRES
        ENCODED   - 20 32.00S  070 12.00W
        TIME OF UPDATE UNKNOWN
9.  ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS TIMES
        CONFIRMED - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES
        DOPPLER A - UNKNOWN
        DOPPLER B - UNKNOWN
        DOA       - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES
        ENCODED   - UNKNOWN
11. HEX ID  DAA6492495805C1    HOMING SIGNAL  121.5
12. ACTIVATION TYPE  UNKNOWN
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NO  000
14. OTHER ENCODED INFORMATION  
    A.  ENCODED POSITION UNCERTAINTY PLUS-MINUS 2 MINUTES OF
        LATITUDE AND LONGITUDE
15. OPERATIONAL INFORMATION  
    A.  MEOSAR ALERT LAST DETECTED AT 12 DEC 18 1201 UTC
16. REMARKS  NIL
END OF MESSAGE
`

// DopConfAlt SAMPLE 406 MHz DOPPLER CONFIRMED ALERT
const DopConfAlt = `
1. DISTRESS COSPAS-SARSAT POSITION CONFIRMED ALERT
2. MSG NO 00932 AUMCC REF 9D064BED62EAFE1
3. DETECTED AT 10 MAY 07 0654 UTC BY SARSAT S11
4. DETECTION FREQUENCY 406.0246 MHz
5. COUNTRY OF BEACON REGISTRATION 232/ G. BRITAIN
6. USER CLASS ELT USER
AIRCRAFT REGISTRATION VP-CGK
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - 25 13.1 N 055 22.2 E
DOPPLER A - 25 17.1 N 055 23.2 E 
DOPPLER B – NIL 
DOA/ALTITUDE - NIL 
ENCODED - NIL
9. ENCODED POSITION PROVIDED BY NIL
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL DOPPLER B – NIL DOA - NIL ENCODED - NIL
11. HEX ID 9D064BED62EAFE1 HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE MANUAL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL 0
14. OTHER ENCODED INFORMATION NIL
15. OPERATIONAL INFORMATION NIL
16. REMARKS NIL
END OF MESSAGE
`

// DopInitAlt SAMPLE 406 MHz DOPPLER INITIAL ALERT
const DopInitAlt = `
1. DISTRESS COSPAS-SARSAT INITIAL ALERT
2. MSG NO 01087 AUMCC REF ADCE402FA80028D
3. DETECTED AT 20 MAY 07 1613 UTC BY SARSAT S08
4. DETECTION FREQUENCY 406.0266 MHz
5. COUNTRY OF BEACON REGISTRATION 366/ USA
6. USER CLASS SERIAL USER – EPIRB (NON-FLOAT FREE) SERIAL NO 0003050
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - NIL
DOPPLER A - 36 38.1 S 168 58.2 E PROB 50 PERCENT 
DOPPLER B - 36 39.1 S 169 01.2 E PROB 50 PERCENT 
DOA/ALTITUDE - NIL
ENCODED - NIL
9. ENCODED POSITION PROVIDED BY NIL
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - 21 MAY 07 0812 UTC DOPPLER B - 21 MAY 07 0812 UTC DOA - NIL
ENCODED - NIL
11. HEX ID ADCE402FA80028D HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE MANUAL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION CSTA CERTIFICATE NO 0163
BEACON MODEL - MCMURDO LTD G5 OR E5 SMARTFIND
15. OPERATIONAL INFORMATION
RELIABILITY OF DOPPLER POSITION DATA - SUSPECT
LUT ID AULUTW ALBANY, AUSTRALIA
16. REMARKS NIL
END OF MESSAGE
`

// DopPosConfAlt SAMPLE 406 MHz DOPPLER POSITION CONFLICT ALERT
const DopPosConfAlt = `
1. DISTRESS COSPAS-SARSAT POSITION CONFLICT ALERT
2. MSG NO 02698 AUMCC REF C1ADE28809C0185
3. DETECTED AT 06 APR 07 1440 UTC BY SARSAT S11
4. DETECTION FREQUENCY 406.0246 MHz
5. COUNTRY OF BEACON REGISTRATION 525/ INDONESIA
6. USER CLASS SERIAL USER-LOCATION - ELT
AIRCRAFT 24-BIT ADDRESS 8A2027
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - NIL
DOPPLER A - 07 00.1 S 098 42.2 E PROB 50 PERCENT 
DOPPLER B - 05 42.1 S 107 20.2 E PROB 50 PERCENT 
DOA/ALTITUDE - NIL
ENCODED - NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY INTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - 06 APR 07 1805 UTC AULUTW ALBANY LUT AUSTRALIA DOPPLER B - 06 APR 07 1956 UTC AULUTW ALBANY LUT AUSTRALIA DOA - NIL
ENCODED - NIL
11. HEX ID C1ADE28809C0185 HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL 00
14. OTHER ENCODED INFORMATION CSTA CERTIFICATE NO 0097
BEACON MODEL - TECHTEST, UK 503-1
AIRCRAFT 24-BIT ADDRESS ASSIGNED TO INDONESIA
15. OPERATIONAL INFORMATION
RELIABILITY OF DOPPLER POSITION DATA - SUSPECT
LUT ID INLUT1 BANGALORE, INDIA
16. REMARKS
THIS POSITION 51 KILOMETRES FROM PREVIOUS ALERT
END OF MESSAGE
`

// PosAlt SAMPLE 406 MHz POSITION ALERT
const PosAlt = `
1. DISTRESS TRACKING COSPAS-SARSAT ALERT
2. MSG NO 00192 AUMCC REF 3266E2019CFFBFF
3. DETECTED AT 03 MAY 19 085310 UTC BY MEOSAR
4. DETECTION FREQUENCY 406.0276 MHz
5. COUNTRY OF BEACON REGISTRATION 403 / SAUDI
6. USER CLASS SGB –ELT(DT)
AIRCRAFT 24 BIT ADDRESS 7100CE
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - 02 15.1 N 046 00.2 E 
DOPPLER A - NIL 
DOPPLER B – NIL
DOA - 02 25.1 N 046 06.2 E
ENCODED - 01 54.40 N - 045 37.53 E
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES
CONFIRMED - NIL 
DOPPLER A - NIL 
DOPPLER B – NIL 
DOA - NIL 
ENCODED - NIL
11. HEX ID 3266E2019CFFBFF
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL 00
14. OTHER ENCODED INFORMATION NIL
TIME SINCE ENCODED LOCATION 4 SECONDS
ALTITUDE OF ENCODED LOCATION 125 METRES
GNSS RECEIVER STATUS 3D LOCATION
AIRCRAFT 24-BIT ADDRESS ASSIGNED TO SAUDI ARABIA
15. OPERATIONAL INFORMATION
MEOSAR ALERT LAST DETECTED AT 03 MAY 19 085310 UTC
ELAPSED TIME SINCE ACTIVATION 13 MINUTES
REMAINING BATTERY CAPACITY BETWEEN 75 AND 100 PERCENT%
BEACON CHARACTERISTICS PER TAC DATABASE – MANUFACTURER APPLIED TECHNOLOGY CORP. MODEL BEACON MODEL XXXYYY-01234
BEACON SUBTYPE FLOAT-FREE
TEMPERATURE RANGE-40C +55C
HOMING 406=25 MW, AIS=20 MW
STROBE BRIGHTNESS=0.75 CANDELA, DUTY-CYCLE=15 FLASH/MINUTE
16. REMARKS
THIS DISTRESS TRACKING MESSAGE IS BEING SENT TO APPROPRIATE SAR AUTHORITIES.
PROCESS THIS ALERT ACCORDING TO RELEVANT REQUIREMENTS
END OF MESSAGE
`

// ConfUpdPosAlt SAMPLE 406 MHz CONFIRMED UPDATE POSITION ALERT
const ConfUpdPosAlt = `
1. SHIP SECURITY COSPAS-SARSAT POSITION CONFIRMED UPDATE ALERT
2. MSG NO 00192 AUMCC REF 2AB82AF800FFBFF
3. DETECTED AT 03 MAY 07 0853 UTC BY SARSAT S09
4. DETECTION FREQUENCY 406.0276 MHz
5. COUNTRY OF BEACON REGISTRATION 341/ ST KITTS
6. USER CLASS STANDARD LOCATION – SHIP SECURITY MMSI LAST 6 DIGITS 088000
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - 02 15.1 N 046 00.2 E
DOPPLER A - 02 25.1 N 046 06.2 E
DOPPLER B – NIL
DOA/ALTITUDE - NIL
ENCODED - 01 54.40 N - 045 37.53 E
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL DOPPLER B – NIL DOA - NIL ENCODED - NIL
11. HEX ID 2AB82AF800FFBFF
HOMING SIGNAL OTHER (NOT 121.5 MHZ) OR NIL
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL 00
14. OTHER ENCODED INFORMATION NIL
15. OPERATIONAL INFORMATION
LUT ID NZLUT WELLINGTON, NEW ZEALAND
16. REMARKS
THIS IS A SHIP SECURITY ALERT.
PROCESS THIS ALERT ACCORDING TO RELEVANT SECURITY REQUIREMENTS
END OF MESSAGE
`

// AltUnrBeaMssg SAMPLE 406 MHz ALERT WITH UNRELIABLE BEACON MESSAGE
const AltUnrBeaMssg = `
1. DISTRESS COSPAS-SARSAT INITIAL ALERT
2. MSG NO 00506 AUMCC REF 12345
3. DETECTED AT 01 APR 07 0610 UTC BY SARSAT S08
4. DETECTION FREQUENCY 406.0315 MHz
5. COUNTRY OF BEACON REGISTRATION NIL
6. USER CLASS NIL
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - NIL
DOPPLER A - 07 23.1 S 136 46.2 E PROB 92 PERCENT 
DOPPLER B - 03 00.1 S 155 08.2 E PROB 08 PERCENT 
DOA/ALTITUDE - NIL
ENCODED - NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY NIL
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL DOPPLER B – NIL DOA - NIL ENCODED - NIL
11. HEX ID 4C4B4E007688888
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION
15. OPERATIONAL INFORMATION
DATA DECODED FROM THE BEACON MESSAGE IS NOT RELIABLE
16. REMARKS NIL
END OF MESSAGE
`

// InitAlt SAMPLE 406 MHz INITIAL ALERT
const InitAlt = `
1 DISTRESS COSPAS-SARSAT POSITION CONFLICT ALERT
2. MSG NO 00308 USMCC REF 12345
3. DETECTED AT 18 DEC 10 1630 UTC BY SARSAT S09
4. DETECTION FREQUENCY 406.0370 MHz
5. COUNTRY OF BEACON REGISTRATION 227/ FRANCE
6. USER CLASS PLB (RETURN LINK) SERIAL NO 00029
7. EMERGENCY CODE NIL
8. POSITIONS CONFIRMED - NIL
DOPPLER A - NIL
DOPPLER B - NIL
DOA/ALTITUDE - NIL
ENCODED - 17 44.1 N 087 26.3 E
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL
DOPPLER B - NIL
DOA - NIL
ENCODED - 18 DEC 10 1655 UTC
11. HEX ID 1C7B000EBF81FE0 HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION NIL
15. OPERATIONAL INFORMATION
BEACON REGISTRATION AT FMCC
16. REMARKS
THIS BEACON HAS GALILEO RETURN LINK CAPABILITY TYPE 1 CAPABILITY (AUTOMATIC ACKNOWLEDGEMENT)
END OF MESSAGE`

// InitDoaPosAlt SAMPLE 406 MHz INITIAL DOA POSITION ALERT
const InitDoaPosAlt = `
1 DISTRESS COSPAS-SARSAT INITIAL ALERT
2. MSG NO 00306 BRMCC REF 12345
3. DETECTED AT 17 DEC 10 1627 UTC BY MEOSAR
4. DETECTION FREQUENCY 406.0371 MHz
5. COUNTRY OF BEACON REGISTRATION 316/ CANADA
6. USER CLASS STANDARD LOCATION - EPIRB SERIAL NO 05918
7. EMERGENCY CODE NIL
8. POSITIONS CONFIRMED - NIL
DOPPLER A - NIL
DOPPLER B - NIL
DOA - 05 10.1 S 178 01.4 E EXPECTED ACCURACY 15 NMS ALTITUDE 45 METRES
ENCODED - NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL DOPPLER B - NIL DOA - NIL ENCODED - NIL
11. HEX ID 278C362E3CFFBFF HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION CSTA CERTIFICATE NO 0108
BEACON MODEL - ACR, RLB-33
15. OPERATIONAL INFORMATION
BEACON REGISTRATION AT CMCC
MEOSAR ALERT LAST DETECTED AT 17 DEC 10 1627 UTC
16. REMARKS NIL
END OF MESSAGE`

// InitDopPosAlt SAMPLE 406 MHz INITIAL DOPPLER POSITION ALERT
const InitDopPosAlt = `
1. DISTRESS COSPAS-SARSAT INITIAL ALERT
2. MSG NO 00741 AUMCC REF 3266E2019CFFBFF
3. DETECTED AT 22 APR 07 0912 UTC BY SARSAT S10
4. DETECTION FREQUENCY 406.0247 MHz
5. COUNTRY OF BEACON REGISTRATION 403/ SAUDI
6. USER CLASS STANDARD LOCATION - ELT AIRCRAFT 24 BIT ADDRESS 7100CE
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - NIL
DOPPLER A - 32 49.1 N 081 54.2 E PROB 69 PERCENT  
DOPPLER B - 24 18.1 N 041 18.2 E PROB 31 PERCENT 
DOA/ALTITUDE - NIL
ENCODED - NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL DOPPLER B – NIL DOA - NIL ENCODED - NIL
11. HEX ID 3266E2019CFFBFF HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION
AIRCRAFT 24-BIT ADDRESS ASSIGNED TO SAUDI ARABIA
15. OPERATIONAL INFORMATION
LUT ID INLUT1 BANGALORE, INDIA
16. REMARKS NIL
END OF MESSAGE
`

// NocrEncPosAlt SAMPLE 406 MHz NOCR ENCODED POSITION ALERT
const NocrEncPosAlt = `
1. DISTRESS COSPAS-SARSAT NOTIFICATION OF COUNTRY OF BEACON REGISTRATION ALERT
2. MSG NO 01737 AUMCC REF 3EF6C34FBF81FE0
3. DETECTED AT 20 MAR 07 0504 UTC BY SARSAT S08
4. DETECTION FREQUENCY 406.0216 MHz
5. COUNTRY OF BEACON REGISTRATION 503/ AUSTRALIA
6. USER CLASS NATIONAL LOCATION - PLB SERIAL NO 099999
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - NIL
DOPPLER A - NIL
DOPPLER B – NIL
DOA/ALTITUDE - NIL
ENCODED - 28 06.00 S 153 40.00 E
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED - NIL
DOPPLER A - NIL
DOPPLER B – NIL
DOA - NIL
ENCODED - 20 MAR 07 1417 UTC NZLUT WELLINGTON LUT NEW ZEALAND
11. HEX ID 3EF6C34FBF81FE0
HOMING SIGNAL OTHER (NOT 121.5 MHZ) OR NIL
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION
ENCODED POSITION UNCERTAINTY
PLUS-MINUS 4 SECONDS IN LATITUDE AND LONGITUDE
15. OPERATIONAL INFORMATION
LUT ID ASLUT CAPE TOWN, SOUTH AFRICA
16. REMARKS NIL
END OF MESSAGE
`

// DoaPosConfAlt SAMPLE 406 MHz DOA POSITION CONFIRMATION ALERT
const DoaPosConfAlt = `
1. DISTRESS COSPAS-SARSAT POSITION CONFIRMED ALERT
2. MSG NO 00306 BRMCC REF 12345
3. DETECTED AT 17 DEC 10 163040 UTC BY MEOSAR
4. DETECTION FREQUENCY 406.0371 MHz
5. COUNTRY OF BEACON REGISTRATION 316/ CANADA
6. USER CLASS STANDARD LOCATION - EPIRB SERIAL NO 05918
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - 05 10.1 S 178 01.3 E
DOPPLER A - NIL
DOPPLER B - NIL
DOA - 05 10.2 S 178 01.2 E EXPECTED ACCURACY 03 NMS 
ALTITUDE 45 METRES
ENCODED - NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES
CONFIRMED – MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES DOPPLER A – NIL
DOPPLER B – NIL
DOA - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES ENCODED - NIL
11. HEX ID 278C362E3CFFBFF HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION CSTA CERTIFICATE NO 0108
BEACON MODEL - ACR, RLB-33
15. OPERATIONAL INFORMATION
BEACON REGISTRATION AT CMCC
MEOSAR ALERT LAST DETECTED AT 17 DEC 10 164610 UTC
16. REMARKS NIL
END OF MESSAGE
`

// PosConfAltPLB SAMPLE 406 MHz POSITION CONFIRMATION ALERT (SGB, LOCATION - PLB)
const PosConfAltPLB = `

1. DISTRESS COSPAS-SARSAT POSITION CONFIRMED ALERT
2. MSG NO 00812 AUMCC REF 2DD747073F81FE0
3. DETECTED AT 28 APR 19 092045 UTC BY MEOSAR
4. DETECTION FREQUENCY 406.0278 MHz
5. COUNTRY OF BEACON REGISTRATION 366/ USA
6. USER CLASS SGB LOCATION – PLB SERIAL NO 167438
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - 33 27.1 N 038 56.2 E
DOPPLER A – NIL
DOPPLER B – NIL
DOA - 33 27.1 N 038 56.2 E EXPECTED ACCURACY 3 NMS ALTITUDE 140 METRES
ENCODED - 33 26.93 N 038 55.67 E
9. ENCODED POSITION PROVIDED BY INTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES
CONFIRMED - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES DOPPLER A – NIL
DOPPLER B – NIL
DOA - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES ENCODED - NIL
11. HEX ID 2DD747073F81FE0 HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE MANUAL ACTIVATION BY USER
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION
TIME SINCE ENCODED LOCATION 23 MINUTES ALTITUDE OF ENCODED LOCATION 125 METRES
15. OPERATIONAL INFORMATION
MEOSAR ALERT LAST DETECTED AT 28 APR 19 092405 UTC
ELAPSED TIME SINCE ACTIVATION 35 MINUTES
REMAINING BATTERY CAPACITY BETWEEN 75AND 100 PERCENT BEACON CHARACTERISTICS PER TAC DATABASE –
MANUFACTURER APPLIED TECHNOLOGY CORP.
MODEL BEACON MODEL XXXYYY-01234
BEACON SUBTYPE FLOAT-FREE
TEMPERATURE RANGE-40C +55C
HOMING 406=25 MW, AIS=20 MW
NAV DEVICE GALILEO, GLONASS
STROBE BRIGHTNESS=0.75 CANDELA, DUTY-CYCLE=15 FLASH/MINUTE
16. REMARKS NIL
END OF MESSAGE
`

// PosConfAlt SAMPLE 406 MHz POSITION CONFIRMATION ALERT
const PosConfAlt = `
1. DISTRESS COSPAS-SARSAT POSITION CONFIRMED ALERT
2. MSG NO 00812 AUMCC REF 2DD747073F81FE0
3. DETECTED AT 28 APR 07 0920 UTC BY SARSAT S11
4. DETECTION FREQUENCY 406.0278 MHz
5. COUNTRY OF BEACON REGISTRATION 366/ USA
6. USER CLASS NATIONAL LOCATION – PLB SERIAL NO 167438
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - 33 27.1 N 038 56.2 E
DOPPLER A - 33 27.1 N 038 56.2 E
DOPPLER B – NIL
DOA/ALTITUDE – NIL
ENCODED - 33 25.93 N 038 55.67 E UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY INTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED – NIL
DOPPLER A – NIL DOPPLER B – NIL DOA – NIL ENCODED - NIL
11. HEX ID 2DD747073F81FE0 HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION NIL
15. OPERATIONAL INFORMATION
LUT ID FRLUT2 TOULOUSE, FRANCE
16. REMARKS NIL
END OF MESSAGE
`

// InitAltNoLoc SAMPLE 406 MHz INITIAL ALERT WITH NO LOCATION
const InitAltNoLoc = `
1. DISTRESS COSPAS-SARSAT INITIAL ALERT
2. MSG NO 00141 SPMCC REF 12345
3. DETECTED AT 21 FEB 07 0646 UTC BY MSG-2
4. DETECTION FREQUENCY 406.0249 MHz
5. COUNTRY OF BEACON REGISTRATION 408/ BAHRAIN
6. USER CLASS NATIONAL LOCATION – ELT SERIAL NO 000006
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED – NIL
DOPPLER A – NIL
DOPPLER B – NIL
DOA/ALTITUDE – NIL
ENCODED - NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED – NIL
DOPPLER A – NIL DOPPLER B – NIL DOA – NIL ENCODED - NIL
11. HEX ID 331000033F81FE0 HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION NIL
15. OPERATIONAL INFORMATION
BEACON REGISTRATION AT WWW.406REGISTRATION.COM
16. REMARKS NIL
END OF MESSAGE
`

// InitEncPosAlert SAMPLE 406 MHz INITIAL ENCODED POSITION ALERT
const InitEncPosAlert = `
1. DISTRESS COSPAS-SARSAT INITIAL ALERT
2. MSG NO 00306 AUMCC REF 12345
3. DETECTED AT 17 APR 07 1627 UTC BY GOES 11
4. DETECTION FREQUENCY 406.0250 MHz
5. COUNTRY OF BEACON REGISTRATION 316/ CANADA
6. USER CLASS STANDARD LOCATION – EPIRB SERIAL NO 05918
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED - NIL
DOPPLER A – NIL
DOPPLER B – NIL
DOA/ALTITUDE – NIL
ENCODED - 05 00.00 S 178 00.00 E
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED – NIL
DOPPLER A – NIL DOPPLER B – NIL DOA – NIL ENCODED - NIL
11. HEX ID 278C362E3CFFBFF HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION CSTA CERTIFICATE NO 0108
BEACON MODEL - ACR, RLB-33
ENCODED POSITION UNCERTAINTY PLUS-MINUS 30 MINUTES OF LATITUDE AND LONGITUDE
15. OPERATIONAL INFORMATION
LUT ID NZGEO1 WELLINGTON GEOLUT, NEW ZEALAND (GOES 11) BEACON REGISTRATION AT [CMCC]
16. REMARKS NIL
END OF MESSAGE
`

// UnresolvedDoppPos SAMPLE 406 MHz UNRESOLVED DOPPLER POSITION MATCH
const UnresolvedDoppPos = `
1. DISTRESS COSPAS-SARSAT UNRESOLVED DOPPLER POSITION MATCH
2. MSG NO 00741 AUMCC REF 1C04273BC0FFBFF
3. DETECTED AT 19 MAR 09 0514 UTC BY SARSAT S08
4. DETECTION FREQUENCY 406.0250 MHz
5. COUNTRY OF BEACON REGISTRATION 224/ SPAIN
6. USER CLASS STANDARD LOCATION – EPIRB MMSI LAST 6 DIGIT 080350
7. EMERGENCY CODE NIL
8. POSITIONS
CONFIRMED – NIL
DOPPLER A - 41 07.1 N 001 12.7 E PROB 69 PERCENT DOPPLER B - 36 48.4 N 022 20.2 E PROB 31 PERCENT DOA – NIL
ENCODED – NIL
UPDATE TIME WITHIN 4 HOURS OF DETECTION TIME
9. ENCODED POSITION PROVIDED BY EXTERNAL DEVICE
10. NEXT PASS/EXPECTED DATA TIMES CONFIRMED – NIL
DOPPLER A – NIL DOPPLER B – NIL DOA – NIL ENCODED - NIL
11. HEX ID1C04273BC0FFBFF HOMING SIGNAL 121.5 MHZ
12. ACTIVATION TYPE NIL
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NIL
14. OTHER ENCODED INFORMATION NIL
15. OPERATIONAL INFORMATION
WARNING AMBIGUITY IS NOT RESOLVED
16. REMARKS NIL
END OF MESSAGE`

const sit185DCSA = `
/00042 00000/5030/17 304 1918
1.  DISTRESS COSPAS-SARSAT ALERT
2.  MSG NO 00042  AUMCC REF NO 00890
3.  DETECTED AT 28 JUL 2016 2338 UTC BY GOES 15
4.  DETECTION FREQUENCY  406.0399 MHZ
5.  COUNTRY OF BEACON REGISTRATION  512/NEWZEALAND
6.  USER CLASS  STANDARD LOCATION - PLB
    SERIAL NO: 4350
    IDENTIFICATION  239/4350
7.  EMERGENCY CODE  NIL
8.  POSITIONS
        CONFIRMED - 36 36.20S  174 43.00E
        DOPPLER A - NIL
        DOPPLER B - NIL
        DOA       - 05 10.2 S 178 01.2 E EXPECTED ACCURACY 03 NMS
                    ALTITUDE 45 METRES
        ENCODED   - NIL
        TIME OF UPDATE UNKNOWN
9.  ENCODED POSITION PROVIDED BY INTERNAL DEVICE
10. NEXT PASS TIMES
        CONFIRMED - UNKNOWN
        DOPPLER A - UNKNOWN
        DOPPLER B - UNKNOWN
        DOA       - UNKNOWN
        ENCODED   - UNKNOWN
11. HEX ID  400E77A1FCFFBFF    HOMING SIGNAL  121.5
12. ACTIVATION TYPE  UNKNOWN
13. BEACON NUMBER ON AIRCRAFT OR VESSEL NO  NIL
14. OTHER ENCODED INFORMATION  
    A.  ENCODED POSITION UNCERTAINTY PLUS-MINUS 2 SECONDS OF
        LATITUDE AND LONGITUDE
15. OPERATIONAL INFORMATION  NIL
16. REMARKS  NIL
END OF MESSAGE
`

const xmlinput = `/00001 00000/5030/17 299 1532
/122/503A/012/01
<?xml version="1.0" ?>
<topMessage>
    <header dest="503A" orig="5030" number="00001" date="2017-10-26T15:32:45Z" />
    <message>
        <unlocatedAlertMessage>
            <header>
                <siteId>20</siteId>
                <beacon>3EF43F8ABF81FE0</beacon>
            </header>
            <tca>2016-07-24T06:44:24.640Z</tca>
            <satellite>12</satellite>
            <orbitNumber>0</orbitNumber>
        </unlocatedAlertMessage>
      </message>
</topMessage>
/LASSIT
/ENDMSG`

const MalaysiaSit185CONFIRMED = `
/00015 00000 /5330 /19 227 2105
1.  DISTRESS COSPAS-SARSAT POSITION CONFIRMED ALERT

2.  MSG NO 00015  MLYMCC REF NO 00302

3.  DETECTED AT 15 AUG 19 2105 UTC BY MEOSAR

4.  DETECTION FREQUENCY  406.0280 MHZ

5.  COUNTRY OF BEACON REGISTRATION  533/MALAYSIA

6.  USER CLASS  USER - EPIRB USER
    MMSI - LAST 6 DIGITS: 130898
    IDENTIFICATION  130898/0

7.  EMERGENCY CODE  NIL

8.  POSITIONS
        CONFIRMED - 01 10.5N  104 08.8E
        DOPPLER A - NIL
        DOPPLER B - NIL
        DOA       - 01 10.5N  104 08.8E
        EXPECTED ACCURACY UNKNOWN
        ALTITUDE 00005 METRES
        ENCODED   - NIL
        TIME OF UPDATE UNKNOWN

9.  ENCODED POSITION PROVIDED BY NIL

10. NEXT PASS TIMES
        CONFIRMED - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES
        DOPPLER A - UNKNOWN
        DOPPLER B - UNKNOWN
        DOA       - MEOSAR DATA USUALLY SENT WITHIN 15 MINUTES
        ENCODED   - UNKNOWN

11. HEX ID  C2A9D40D30330D1    HOMING SIGNAL  121.5

12. ACTIVATION TYPE  AUTOMATIC OR MANUAL

13. BEACON NUMBER ON AIRCRAFT OR VESSEL NO  000

14. OTHER ENCODED INFORMATION  NIL

15. OPERATIONAL INFORMATION
    A.  MEOSAR ALERT LAST DETECTED AT 15 AUG 19 2105 UTC

16. REMARKS  NIL`
