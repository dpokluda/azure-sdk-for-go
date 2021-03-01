package quantum

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ProvisioningStatus enumerates the values for provisioning status.
type ProvisioningStatus string

const (
	// Failed ...
	Failed ProvisioningStatus = "Failed"
	// ProviderDeleting ...
	ProviderDeleting ProvisioningStatus = "ProviderDeleting"
	// ProviderLaunching ...
	ProviderLaunching ProvisioningStatus = "ProviderLaunching"
	// ProviderProvisioning ...
	ProviderProvisioning ProvisioningStatus = "ProviderProvisioning"
	// ProviderUpdating ...
	ProviderUpdating ProvisioningStatus = "ProviderUpdating"
	// Succeeded ...
	Succeeded ProvisioningStatus = "Succeeded"
)

// PossibleProvisioningStatusValues returns an array of possible values for the ProvisioningStatus const type.
func PossibleProvisioningStatusValues() []ProvisioningStatus {
	return []ProvisioningStatus{Failed, ProviderDeleting, ProviderLaunching, ProviderProvisioning, ProviderUpdating, Succeeded}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusDeleted ...
	StatusDeleted Status = "Deleted"
	// StatusDeleting ...
	StatusDeleting Status = "Deleting"
	// StatusFailed ...
	StatusFailed Status = "Failed"
	// StatusLaunching ...
	StatusLaunching Status = "Launching"
	// StatusSucceeded ...
	StatusSucceeded Status = "Succeeded"
	// StatusUpdating ...
	StatusUpdating Status = "Updating"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusDeleted, StatusDeleting, StatusFailed, StatusLaunching, StatusSucceeded, StatusUpdating}
}

// UsableStatus enumerates the values for usable status.
type UsableStatus string

const (
	// No ...
	No UsableStatus = "No"
	// Partial ...
	Partial UsableStatus = "Partial"
	// Yes ...
	Yes UsableStatus = "Yes"
)

// PossibleUsableStatusValues returns an array of possible values for the UsableStatus const type.
func PossibleUsableStatusValues() []UsableStatus {
	return []UsableStatus{No, Partial, Yes}
}