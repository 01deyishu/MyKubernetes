#!/bin/bash
#get somp snmp info
snmpwalk -v 2c -c 7Niuread 183.136.239.33 1.3.6.1.2.1.4.34.1.3.1.4 | cut -d "." -f 13-16 | cut -d " " -f 1,4
