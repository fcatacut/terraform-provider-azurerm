package guestconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionAfterReboot enumerates the values for action after reboot.
type ActionAfterReboot string

const (
	// ContinueConfiguration ...
	ContinueConfiguration ActionAfterReboot = "ContinueConfiguration"
	// StopConfiguration ...
	StopConfiguration ActionAfterReboot = "StopConfiguration"
)

// PossibleActionAfterRebootValues returns an array of possible values for the ActionAfterReboot const type.
func PossibleActionAfterRebootValues() []ActionAfterReboot {
	return []ActionAfterReboot{ContinueConfiguration, StopConfiguration}
}

// ComplianceStatus enumerates the values for compliance status.
type ComplianceStatus string

const (
	// Compliant ...
	Compliant ComplianceStatus = "Compliant"
	// NonCompliant ...
	NonCompliant ComplianceStatus = "NonCompliant"
	// Pending ...
	Pending ComplianceStatus = "Pending"
)

// PossibleComplianceStatusValues returns an array of possible values for the ComplianceStatus const type.
func PossibleComplianceStatusValues() []ComplianceStatus {
	return []ComplianceStatus{Compliant, NonCompliant, Pending}
}

// ConfigurationMode enumerates the values for configuration mode.
type ConfigurationMode string

const (
	// ApplyAndAutoCorrect ...
	ApplyAndAutoCorrect ConfigurationMode = "ApplyAndAutoCorrect"
	// ApplyAndMonitor ...
	ApplyAndMonitor ConfigurationMode = "ApplyAndMonitor"
	// ApplyOnly ...
	ApplyOnly ConfigurationMode = "ApplyOnly"
)

// PossibleConfigurationModeValues returns an array of possible values for the ConfigurationMode const type.
func PossibleConfigurationModeValues() []ConfigurationMode {
	return []ConfigurationMode{ApplyAndAutoCorrect, ApplyAndMonitor, ApplyOnly}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// DSC ...
	DSC Kind = "DSC"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{DSC}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Created ...
	Created ProvisioningState = "Created"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Created, Failed, Succeeded}
}

// Type enumerates the values for type.
type Type string

const (
	// Consistency ...
	Consistency Type = "Consistency"
	// Initial ...
	Initial Type = "Initial"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{Consistency, Initial}
}
