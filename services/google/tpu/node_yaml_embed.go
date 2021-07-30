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
// gen_go_data -package tpu -var YAML_node blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/tpu/node.yaml

package tpu

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/tpu/node.yaml
var YAML_node = []byte("info:\n  title: TPU/Node\n  description: DCL Specification for the TPU Node resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Node\n    parameters:\n    - name: Node\n      required: true\n      description: A full instance of a Node\n  apply:\n    description: The function used to apply information about a Node\n    parameters:\n    - name: Node\n      required: true\n      description: A full instance of a Node\n  delete:\n    description: The function used to delete a Node\n    parameters:\n    - name: Node\n      required: true\n      description: A full instance of a Node\n  deleteAll:\n    description: The function used to delete all Node\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Node\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Node:\n      title: Node\n      x-dcl-id: projects/{{project}}/locations/{{location}}/nodes/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      type: object\n      required:\n      - name\n      - acceleratorType\n      - tensorflowVersion\n      - cidrBlock\n      - project\n      - location\n      properties:\n        acceleratorType:\n          type: string\n          x-dcl-go-name: AcceleratorType\n          description: The type of hardware accelerators associated with this node.\n            Required.\n          x-kubernetes-immutable: true\n        cidrBlock:\n          type: string\n          x-dcl-go-name: CidrBlock\n          description: The CIDR block that the TPU node will use when selecting an\n            IP address. This CIDR block must be a /29 block; the Compute Engine networks\n            API forbids a smaller block, and using a larger block would be wasteful\n            (a node can only consume one IP address). Errors will occur if the CIDR\n            block has already been used for a currently existing TPU node, the CIDR\n            block conflicts with any subnetworks in the user's provided network, or\n            the provided network is peered with another network that is using that\n            CIDR block.\n          x-kubernetes-immutable: true\n        createTime:\n          type: object\n          x-dcl-go-name: CreateTime\n          x-dcl-go-type: NodeCreateTime\n          readOnly: true\n          description: Output only. The time when the node was created.\n          x-kubernetes-immutable: true\n          properties:\n            nanos:\n              type: integer\n              format: int64\n              x-dcl-go-name: Nanos\n              description: Non-negative fractions of a second at nanosecond resolution.\n                Negative second values with fractions must still have non-negative\n                nanos values that count forward in time. Must be from 0 to 999,999,999\n                inclusive.\n              x-kubernetes-immutable: true\n            seconds:\n              type: integer\n              format: int64\n              x-dcl-go-name: Seconds\n              description: Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z.\n                Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.\n              x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: The user-supplied description of the TPU. Maximum of 512 characters.\n          x-kubernetes-immutable: true\n        health:\n          type: string\n          x-dcl-go-name: Health\n          x-dcl-go-type: NodeHealthEnum\n          readOnly: true\n          description: 'The health status of the TPU node. Possible values: HEALTH_UNSPECIFIED,\n            HEALTHY, DEPRECATED_UNHEALTHY, TIMEOUT, UNHEALTHY_TENSORFLOW, UNHEALTHY_MAINTENANCE'\n          x-kubernetes-immutable: true\n          enum:\n          - HEALTH_UNSPECIFIED\n          - HEALTHY\n          - DEPRECATED_UNHEALTHY\n          - TIMEOUT\n          - UNHEALTHY_TENSORFLOW\n          - UNHEALTHY_MAINTENANCE\n        healthDescription:\n          type: string\n          x-dcl-go-name: HealthDescription\n          readOnly: true\n          description: Output only. If this field is populated, it contains a description\n            of why the TPU Node is unhealthy.\n          x-kubernetes-immutable: true\n        ipAddress:\n          type: string\n          x-dcl-go-name: IPAddress\n          readOnly: true\n          description: Output only. DEPRECATED! Use network_endpoints instead. The\n            network address for the TPU Node as visible to Compute Engine instances.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Resource labels to represent user-provided metadata.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Output only. The immutable name of the TPU\n          x-kubernetes-immutable: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: The name of a network they wish to peer the TPU node to. It\n            must be a preexisting Compute Engine network inside of the project on\n            which this API has been activated. If none is provided, \"default\" will\n            be used.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          x-dcl-references:\n          - resource: Compute/Network\n            field: selfLink\n        networkEndpoints:\n          type: array\n          x-dcl-go-name: NetworkEndpoints\n          readOnly: true\n          description: Output only. The network endpoints where TPU workers can be\n            accessed and sent work. It is recommended that Tensorflow clients of the\n            node reach out to the 0th entry in this map first.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: NodeNetworkEndpoints\n            properties:\n              ipAddress:\n                type: string\n                x-dcl-go-name: IPAddress\n                readOnly: true\n                description: The IP address of this network endpoint.\n                x-kubernetes-immutable: true\n              port:\n                type: integer\n                format: int64\n                x-dcl-go-name: Port\n                readOnly: true\n                description: The port of this network endpoint.\n                x-kubernetes-immutable: true\n        port:\n          type: string\n          x-dcl-go-name: Port\n          readOnly: true\n          description: Output only. DEPRECATED! Use network_endpoints instead. The\n            network port for the TPU Node as visible to Compute Engine instances.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        schedulingConfig:\n          type: object\n          x-dcl-go-name: SchedulingConfig\n          x-dcl-go-type: NodeSchedulingConfig\n          x-kubernetes-immutable: true\n          required:\n          - preemptible\n          properties:\n            preemptible:\n              type: boolean\n              x-dcl-go-name: Preemptible\n              description: Defines whether the node is preemptible.\n              x-kubernetes-immutable: true\n            reserved:\n              type: boolean\n              x-dcl-go-name: Reserved\n              description: Whether the node is created under a reservation.\n              x-kubernetes-immutable: true\n        serviceAccount:\n          type: string\n          x-dcl-go-name: ServiceAccount\n          readOnly: true\n          description: Output only. The service account used to run the tensor flow\n            services within the node. To share resources, including Google Cloud Storage\n            data, with the Tensorflow job running in the Node, this account must have\n            permissions to that data.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: NodeStateEnum\n          readOnly: true\n          description: 'Output only. The current state for the TPU Node. Possible\n            values: STATE_UNSPECIFIED, CREATING, READY, RESTARTING, REIMAGING, DELETING,\n            REPAIRING, STOPPED, STOPPING, STARTING, PREEMPTED, TERMINATED, HIDING,\n            HIDDEN, UNHIDING'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - READY\n          - RESTARTING\n          - REIMAGING\n          - DELETING\n          - REPAIRING\n          - STOPPED\n          - STOPPING\n          - STARTING\n          - PREEMPTED\n          - TERMINATED\n          - HIDING\n          - HIDDEN\n          - UNHIDING\n        symptoms:\n          type: array\n          x-dcl-go-name: Symptoms\n          readOnly: true\n          description: Output only. The Symptoms that have occurred to the TPU Node.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: NodeSymptoms\n            properties:\n              createTime:\n                type: object\n                x-dcl-go-name: CreateTime\n                x-dcl-go-type: NodeSymptomsCreateTime\n                description: Timestamp when the Symptom is created.\n                x-kubernetes-immutable: true\n                properties:\n                  nanos:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: Nanos\n                    description: Non-negative fractions of a second at nanosecond\n                      resolution. Negative second values with fractions must still\n                      have non-negative nanos values that count forward in time. Must\n                      be from 0 to 999,999,999 inclusive.\n                    x-kubernetes-immutable: true\n                  seconds:\n                    type: integer\n                    format: int64\n                    x-dcl-go-name: Seconds\n                    description: Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z.\n                      Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.\n                    x-kubernetes-immutable: true\n              details:\n                type: string\n                x-dcl-go-name: Details\n                description: Detailed information of the current Symptom.\n                x-kubernetes-immutable: true\n              symptomType:\n                type: string\n                x-dcl-go-name: SymptomType\n                x-dcl-go-type: NodeSymptomsSymptomTypeEnum\n                description: 'Type of the Symptom. Possible values: SYMPTOM_TYPE_UNSPECIFIED,\n                  LOW_MEMORY, OUT_OF_MEMORY, EXECUTE_TIMED_OUT, MESH_BUILD_FAIL, HBM_OUT_OF_MEMORY'\n                x-kubernetes-immutable: true\n                enum:\n                - SYMPTOM_TYPE_UNSPECIFIED\n                - LOW_MEMORY\n                - OUT_OF_MEMORY\n                - EXECUTE_TIMED_OUT\n                - MESH_BUILD_FAIL\n                - HBM_OUT_OF_MEMORY\n              workerId:\n                type: string\n                x-dcl-go-name: WorkerId\n                description: A string used to uniquely distinguish a worker within\n                  a TPU node.\n                x-kubernetes-immutable: true\n        tensorflowVersion:\n          type: string\n          x-dcl-go-name: TensorflowVersion\n          description: The version of Tensorflow running in the Node. Required.\n        useServiceNetworking:\n          type: boolean\n          x-dcl-go-name: UseServiceNetworking\n          description: Whether the VPC peering for the node is set up through Service\n            Networking API. The VPC Peering should be set up before provisioning the\n            node. If this field is set, cidr_block field should not be specified.\n            If the network, that you want to peer the TPU Node to, is Shared VPC networks,\n            the node must be created with this this field enabled.\n          x-kubernetes-immutable: true\n")

// 12866 bytes
// MD5: 36f50ffe00806164adeca1970e9d5de0
