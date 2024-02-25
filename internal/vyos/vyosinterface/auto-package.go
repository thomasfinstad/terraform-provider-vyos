// Package vyosinterface generated by Makefile on 2024-02-25T08:44:41+00:00. DO NOT EDIT.
package vyosinterface

import "github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"

// GetInterfaces returns all autogenerated interface definitions
func GetInterfaces() []interfacedefinition.InterfaceDefinition {
	return []interfacedefinition.InterfaceDefinition{
		bcastrelay(),
		container(),
		cron(),
		dhcprelay(),
		dhcpserver(),
		dnsdomainname(),
		dnsdynamic(),
		dnsforwarding(),
		firewall(),
		flowaccountingconf(),
		highavailability(),
		https(),
		igmpproxy(),
		interfacesbonding(),
		interfacesbridge(),
		interfacesdummy(),
		interfacesethernet(),
		interfacesgeneve(),
		interfacesinput(),
		interfacesloopback(),
		interfacesmacsec(),
		interfacesopenvpn(),
		interfacespppoe(),
		interfacespseudoethernet(),
		interfacessstpc(),
		interfacestunnel(),
		interfacesvirtualethernet(),
		interfacesvti(),
		interfacesvxlan(),
		interfaceswireguard(),
		interfaceswireless(),
		interfaceswwan(),
		lldp(),
		loadbalancingwan(),
		nat(),
		netns(),
		ntp(),
		pki(),
		policy(),
		policylocalroute(),
		policyroute(),
		protocolsbabel(),
		protocolsbfd(),
		protocolsbgp(),
		protocolseigrp(),
		protocolsfailover(),
		protocolsigmp(),
		protocolsisis(),
		protocolsmpls(),
		protocolsmulticast(),
		protocolsnhrp(),
		protocolsospf(),
		protocolspim(),
		protocolsrip(),
		protocolsripng(),
		protocolsrpki(),
		protocolsstatic(),
		protocolsstaticarp(),
		qos(),
		saltminion(),
		serviceconntracksync(),
		serviceconsoleserver(),
		serviceeventhandler(),
		serviceidsddosprotection(),
		serviceipoeserver(),
		servicemdnsrepeater(),
		servicemonitoringtelegraf(),
		servicepppoeserver(),
		servicerouteradvert(),
		servicesla(),
		serviceupnp(),
		servicewebproxy(),
		snmp(),
		ssh(),
		systemaccelerationqat(),
		systemconfigmgmt(),
		systemconntrack(),
		systemconsole(),
		systemfrr(),
		systemip(),
		systemlcd(),
		systemlogin(),
		systemloginbanner(),
		systemlogs(),
		systemoption(),
		systemproxy(),
		systemsflow(),
		systemsysctl(),
		systemsyslog(),
		systemtimezone(),
		systemupdatecheck(),
		tftpserver(),
		vpnipsec(),
		vpnopenconnect(),
		vpnpptp(),
		vpnsstp(),
		vrf(),
	}
}
