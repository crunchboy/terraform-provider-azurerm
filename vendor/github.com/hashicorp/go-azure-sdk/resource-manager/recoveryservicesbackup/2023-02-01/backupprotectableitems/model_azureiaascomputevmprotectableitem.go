package backupprotectableitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkloadProtectableItem = AzureIaaSComputeVMProtectableItem{}

type AzureIaaSComputeVMProtectableItem struct {
	ResourceGroup         *string `json:"resourceGroup,omitempty"`
	VirtualMachineId      *string `json:"virtualMachineId,omitempty"`
	VirtualMachineVersion *string `json:"virtualMachineVersion,omitempty"`

	// Fields inherited from WorkloadProtectableItem

	BackupManagementType *string           `json:"backupManagementType,omitempty"`
	FriendlyName         *string           `json:"friendlyName,omitempty"`
	ProtectableItemType  string            `json:"protectableItemType"`
	ProtectionState      *ProtectionStatus `json:"protectionState,omitempty"`
	WorkloadType         *string           `json:"workloadType,omitempty"`
}

func (s AzureIaaSComputeVMProtectableItem) WorkloadProtectableItem() BaseWorkloadProtectableItemImpl {
	return BaseWorkloadProtectableItemImpl{
		BackupManagementType: s.BackupManagementType,
		FriendlyName:         s.FriendlyName,
		ProtectableItemType:  s.ProtectableItemType,
		ProtectionState:      s.ProtectionState,
		WorkloadType:         s.WorkloadType,
	}
}

var _ json.Marshaler = AzureIaaSComputeVMProtectableItem{}

func (s AzureIaaSComputeVMProtectableItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureIaaSComputeVMProtectableItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureIaaSComputeVMProtectableItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureIaaSComputeVMProtectableItem: %+v", err)
	}

	decoded["protectableItemType"] = "Microsoft.Compute/virtualMachines"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureIaaSComputeVMProtectableItem: %+v", err)
	}

	return encoded, nil
}
