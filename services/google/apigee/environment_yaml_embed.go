// Copyright 2021 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package apigee -var YAML_environment blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/environment.yaml

package apigee

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/apigee/environment.yaml
var YAML_environment = []byte("info:\n  title: Apigee/Environment\n  description: DCL Specification for the Apigee Environment resource\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Environment\n    parameters:\n    - name: Environment\n      required: true\n      description: A full instance of a Environment\n  apply:\n    description: The function used to apply information about a Environment\n    parameters:\n    - name: Environment\n      required: true\n      description: A full instance of a Environment\n  delete:\n    description: The function used to delete a Environment\n    parameters:\n    - name: Environment\n      required: true\n      description: A full instance of a Environment\n  deleteAll:\n    description: The function used to delete all Environment\n    parameters:\n    - name: organization\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Environment\n    parameters:\n    - name: organization\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Environment:\n      title: Environment\n      x-dcl-id: organizations/{{organization}}/environments/{{name}}\n      x-dcl-parent-container: organization\n      type: object\n      required:\n      - name\n      - organization\n      properties:\n        createdAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: CreatedAt\n          readOnly: true\n          description: Output only. Creation time of this environment as milliseconds\n            since epoch.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. Description of the environment.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Optional. Display name for this environment.\n          x-kubernetes-immutable: true\n        lastModifiedAt:\n          type: integer\n          format: int64\n          x-dcl-go-name: LastModifiedAt\n          readOnly: true\n          description: Output only. Last modification time of this environment as\n            milliseconds since epoch.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the environment. Values must match the regular\n            expression ^[.\\p{Alnum}-_]{1,255}$\n          x-kubernetes-immutable: true\n        organization:\n          type: string\n          x-dcl-go-name: Organization\n          description: The organization for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Organization\n            field: name\n            parent: true\n        properties:\n          type: object\n          x-dcl-go-name: Properties\n          x-dcl-go-type: EnvironmentProperties\n          description: Optional. Key-value pairs that may be used for customizing\n            the environment.\n          x-kubernetes-immutable: true\n          properties:\n            property:\n              type: array\n              x-dcl-go-name: Property\n              description: List of all properties in the object\n              x-dcl-send-empty: true\n              x-dcl-list-type: set\n              items:\n                type: object\n                x-dcl-go-type: EnvironmentPropertiesProperty\n                properties:\n                  name:\n                    type: string\n                    x-dcl-go-name: Name\n                    description: The property key\n                    x-kubernetes-immutable: true\n                  value:\n                    type: string\n                    x-dcl-go-name: Value\n                    description: The property value\n                    x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: EnvironmentStateEnum\n          readOnly: true\n          description: 'Output only. State of the environment. Values other than ACTIVE\n            means the resource is not ready to use. Possible values: STATE_UNSPECIFIED,\n            CREATING, ACTIVE, DELETING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - ACTIVE\n          - DELETING\n")

// 4380 bytes
// MD5: 094c1aa879c942917bbc2936cd85b570
