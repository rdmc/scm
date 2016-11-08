package main


// ios Version 12.2(33)SCG7
// assume NO dual ip (ip v4 only), No cable modem remote query
// show cable modem xxxx.xxxx.xxxx

// fields
"MAC Addres"
"IP Address"
"I/F    - interface"
"MAC State"
"Prim Sid"
"RxPwr (dBmv)"
"Timing Oddset"
"Num CPE"
"DIP"   - "Dual IP"


// MAC States:
// simple version
"offline", "Cable modem considered offline"

"init (r1)", "Cable modem sent initial ranging"
"init (r2)", "Cable modem is ranging"
"init (rc)", "Cable modem ranging complete"

"init (d)", "DHCP request received"
"init (i)", "DHCP reply received; IP address assigned"

"init (t)", "TOD exchange started"
"init (o)", "Option file transfer started"

"online", "Cable modem registered, enabled for data"
"online(d)", "Cable modem registered, but network access for the cable modem is disabled"
"online(pk)", "Cable modem registered, BPI enabled and KEK assigned"
"online(pt)", "Cable modem registered, BPI enabled and TEK assigned"

"reject (pk)", "KEK modem key assignment rejected"
"reject (pt)" , "TEK modem key assignment rejected"
"reject (m)", "Cable modem did attempt to register; registration was refused due to bad MIC (Message Integrity Check)"
"reject (c)", "Cable modem did attempt to register; registration was refused due to bad COS (Class of Service)"
"reject (r)",  "CM did attempt to register, registration was refused due to unavailable resource."

// MAC States:

//Registration and Provisioning Status Conditions

"init(r1)" "The CM sent initial ranging."
"init(r2)" "The CM is ranging."
"init(rc)" "Ranging has completed."

"init(d)" "The DHCP request was received, as DHCPDISCOVER."
"init(dr)" "The DHCP request has been sent to the cable modem."
"init(i)" "The cable modem has received the DHCPOFFER reply (DHCPACK)."
"init(io)" "The Cisco CMTS has seen the DHCP offer as sent to the cable modem."
"init(o)" "The CM has begun to download the option file (DOCSIS configuration file)"
"init(t) " "Time-of-day (TOD) exchange has started."
"resetting"  "The CM is being reset and will shortly restart the registration process."

//Non-Error Status Conditions
"cc(r1)" "CM received a Downstream Channel Change (DCC) or Upstream Channel Change (UCC) request message from the CMTS."
"cc(r2)" "This state should normally follow cc(r1) and indicates that the CM has finished its initial ranging on the new channel."

"offline" "The CM is considered offline (disconnected or powered down)."

"online" "The CM has registered and is enabled to pass data on the network."

"online(d)" "The CM registered, but network access for CPE devices using this CM has been disabled through the DOCSIS configuration file." 
"online(pkd)" "This state is equivalent to the online(d) and onlike(pk) states"
"online(ptd)" "This state is equivalent to the online(d) and onlike(pt) states"
"online(pk)" "The CM registered, BPI is enabled and KEK is assigned."
"online(pt)" "The CM registered, BPI is enabled and TEK is assigned. BPI encryption is now being performed."

Note If an exclamation point (!) appears in front of one of the online states, it indicates that the cable
dynamic-secret command has been used with either the mark or reject option, and that the cable modem
has failed the dynamic secret authentication check.

"expire(pk)" "The CM registered, BPI is enabled, KEK was assigned, but the current KEK expired before the CM could successfully renew a new KEK value."
"expire(pkd)" "This state is equivalente to online(d) and expire(pk) states."
"expire(pt)" "The CM registered, BPI is enabled, TEK was assigned, but the current TEK expired before the CM could successfully renew a new KEK value."
"expire(ptd)" "This state is equivalente to online(d) and expire(pt) states."

//Error Status Conditions
"reject(m)" "The CM attempted to register but registration was refused due to a bad Message Integrity Check (MIC) value."
"reject(c)" "The CM attempted to register, but registration was refused due to a a number of possible errors:...."

"reject(pk)" "KEK key assignment is rejected, BPI encryption has not been established."
"reject(pkd)" "This state is equivalent to the online(d) and reject(pk) states."
"reject(pt)" "TEK key assignment is rejected, BPI encryption has not been established."
"reject(ptd)" "This state is equivalent to the online(d) and reject(pt) states."

"reject(ts)" "The CM attempted to register, but registration failed because the TFTP server timestamp in the CM registration request did not match the timestamp maintained by the CMTS."
"reject(ip)" "The CM attempted to register, but registration failed because the IP address in the CM request did not match the IP address that the TFTP server recorded when it sent the DOCSIS configuration file to the CM."
"reject(na)" "The CM attempted to register, but registration failed because the CM did not send a Registration-Acknowledgement (REG-ACK) message in reply to the Registration-Response (REG-RSP) message sent by the CMTS. A RegistrationNonAcknowledgement (REG-NACK) is assumed."

// semo for wideband:
"w-online" "The WCM has registered and is enabled to pass data on the network."
"w-online(d)" "The WCM registered, but network access for CPE devices using this WCM has been disabled through the DOCSIS configuration file."
"w-online(pkd)" "This state is equivalent to the w-online(d) and w-online(pk) states."
"w-online(pt)" "The WCM registered, BPI is enabled and TEK is assigned. BPI encryption is now being performed."
"w-online(ptd)" "This state is equivalent to the w-online(d) and w-online(pt) states."
"w-online(pk)" "The WCM registered, BPI is enabled and KEK is assigned."

"w-expire(pk)" "The WCM registered, BPI is enabled, KEK was assigned, but the current KEK expired before the WCM could successfully renew a new KEK value."
"w-expire(pkd)" "This state is equivalent to the w-online(d) and w-expire(pk) states."
"w-expire(pt)" "The WCM registered, BPI is enabled, TEK was assigned, but the current TEK expired before the WCM could successfully renew a new KEK value."
"w-expire(ptd)" "This state is equivalent to the w-online(d) and w-expire(pt) states." 

"w-reject(pk)" "KEK key assignment is rejected, BPI encryption has not been established."
"w-reject(pkd)" "This state is equivalent to the w-online(d) and w-reject(pk) states."
"w-reject(pt)" "TEK key assignment is rejected, BPI encryption has not been established."
"w-reject(ptd)" "This state is equivalent to the w-online(d) and w-reject(pt) states." 

