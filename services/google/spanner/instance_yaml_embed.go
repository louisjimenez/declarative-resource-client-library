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
// gen_go_data -package spanner -var YAML_instance blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/spanner/instance.yaml

package spanner

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/spanner/instance.yaml
var YAML_instance = []byte("info:\n  title: Spanner/Instance\n  description: DCL Specification for the Spanner Instance resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  apply:\n    description: The function used to apply information about a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  delete:\n    description: The function used to delete a Instance\n    parameters:\n    - name: Instance\n      required: true\n      description: A full instance of a Instance\n  deleteAll:\n    description: The function used to delete all Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Instance\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Instance:\n      title: Instance\n      x-dcl-id: projects/{{project}}/instances/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      - config\n      - displayName\n      properties:\n        config:\n          type: string\n          x-dcl-go-name: Config\n          description: Required. The name of the instance's configuration. Values\n            are of the form `projects/<project>/instanceConfigs/<configuration>`.\n          x-dcl-references:\n          - resource: Spanner/InstanceConfig\n            field: name\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Required. The descriptive name for this instance as it appears\n            in UIs. Must be unique per project and between 4 and 30 characters in\n            length.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: 'Cloud Labels are a flexible and lightweight mechanism for\n            organizing cloud resources into groups that reflect a customer''s organizational\n            needs and deployment strategies. Cloud Labels can be used to filter collections\n            of resources. They can be used to control how resource metrics are aggregated.\n            And they can be used as arguments to policy management rules (e.g. route,\n            firewall, load balancing, etc.).   * Label keys must be between 1 and\n            63 characters long and must conform to    the following regular expression:\n            `[a-z]([-a-z0-9]*[a-z0-9])?`.  * Label values must be between 0 and 63\n            characters long and must conform    to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.  *\n            No more than 64 labels can be associated with a given resource.  See https://goo.gl/xmQnxf\n            for more information on and examples of labels.  If you plan to use labels\n            in your own code, please note that additional characters may be allowed\n            in the future. And so you are advised to use an internal label representation,\n            such as JSON, which doesn''t rely upon specific characters being disallowed.  For\n            example, representing labels as the string:  name + \"_\" + value  would\n            prove problematic if we were to allow \"_\" in a future release.'\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: A unique identifier for the instance, which cannot be changed\n            after the instance is created. Values are of the form `[a-z][-a-z0-9]*[a-z0-9]`.\n          x-kubernetes-immutable: true\n        nodeCount:\n          type: integer\n          format: int64\n          x-dcl-go-name: NodeCount\n          description: The number of nodes allocated to this instance. See [the documentation](https://cloud.google.com/spanner/docs/instances#node_count)\n            for more information about nodes.\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: InstanceStateEnum\n          readOnly: true\n          description: 'The current instance state. Possible values: CREATING, READY'\n          x-kubernetes-immutable: true\n          enum:\n          - CREATING\n          - READY\n")

// 4634 bytes
// MD5: d9d38aa9883bd16e3be907e3bf1828c8
