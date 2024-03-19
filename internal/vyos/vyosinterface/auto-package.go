// Package vyosinterface generated by Makefile. DO NOT EDIT.
package vyosinterface

import "github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"

// GetInterfaces returns all autogenerated interface definitions
func GetInterfaces() []interfacedefinition.InterfaceDefinition {
	return []interfacedefinition.InterfaceDefinition{
		container(),
		firewall(),
		highavailability(),
		nat(),
		netns(),
		pki(),
		policy(),
		qos(),
		vrf(),
	}
}
