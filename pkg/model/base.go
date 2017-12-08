// Copyright 2017 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements the common data structure.

*/

package model

import (
	"encoding/json"
)

type Modeler interface {
	GetId() string

	GetCreatedTime() string

	SetId(uuid string)

	SetCreatedTime(createdAt string)

	SetUpdatedTime(updatedAt string)
}

type BaseModel struct {
	// +readOnly:true
	Id string `json:"id"`
	// +readOnly:true
	CreatedAt string `json:"createdAt"`
	// +readOnly:true
	UpdatedAt string `json:"updatedAt"`
}

func (b *BaseModel) GetId() string {
	return b.Id
}

func (b *BaseModel) GetCreatedTime() string {
	return b.CreatedAt
}

func (b *BaseModel) SetId(uuid string) {
	b.Id = uuid
}

func (b *BaseModel) SetCreatedTime(createdAt string) {
	b.CreatedAt = createdAt
}

func (b *BaseModel) SetUpdatedTime(updatedAt string) {
	b.UpdatedAt = updatedAt
}

// ExtraSpec is a dictionary object that contains unique keys and json
// objects.
type ExtraSpec map[string]interface{}

func (ext ExtraSpec) Encode() []byte {
	parmBody, _ := json.Marshal(&ext)
	return parmBody
}
