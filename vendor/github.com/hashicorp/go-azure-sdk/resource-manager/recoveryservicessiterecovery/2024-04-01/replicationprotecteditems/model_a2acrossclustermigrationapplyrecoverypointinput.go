package replicationprotecteditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplyRecoveryPointProviderSpecificInput = A2ACrossClusterMigrationApplyRecoveryPointInput{}

type A2ACrossClusterMigrationApplyRecoveryPointInput struct {

	// Fields inherited from ApplyRecoveryPointProviderSpecificInput

	InstanceType string `json:"instanceType"`
}

func (s A2ACrossClusterMigrationApplyRecoveryPointInput) ApplyRecoveryPointProviderSpecificInput() BaseApplyRecoveryPointProviderSpecificInputImpl {
	return BaseApplyRecoveryPointProviderSpecificInputImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = A2ACrossClusterMigrationApplyRecoveryPointInput{}

func (s A2ACrossClusterMigrationApplyRecoveryPointInput) MarshalJSON() ([]byte, error) {
	type wrapper A2ACrossClusterMigrationApplyRecoveryPointInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2ACrossClusterMigrationApplyRecoveryPointInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2ACrossClusterMigrationApplyRecoveryPointInput: %+v", err)
	}

	decoded["instanceType"] = "A2ACrossClusterMigration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2ACrossClusterMigrationApplyRecoveryPointInput: %+v", err)
	}

	return encoded, nil
}
