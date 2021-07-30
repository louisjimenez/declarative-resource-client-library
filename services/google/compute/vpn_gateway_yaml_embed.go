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
// gen_go_data -package compute -var YAML_vpn_gateway blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/vpn_gateway.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/vpn_gateway.yaml
var YAML_vpn_gateway = []byte("info:\n  title: Compute/VpnGateway\n  description: DCL Specification for the Compute VpnGateway resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a VpnGateway\n    parameters:\n    - name: VpnGateway\n      required: true\n      description: A full instance of a VpnGateway\n  apply:\n    description: The function used to apply information about a VpnGateway\n    parameters:\n    - name: VpnGateway\n      required: true\n      description: A full instance of a VpnGateway\n  delete:\n    description: The function used to delete a VpnGateway\n    parameters:\n    - name: VpnGateway\n      required: true\n      description: A full instance of a VpnGateway\n  deleteAll:\n    description: The function used to delete all VpnGateway\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many VpnGateway\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    VpnGateway:\n      title: VpnGateway\n      x-dcl-id: projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}\n      x-dcl-locations:\n      - region\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      type: object\n      required:\n      - name\n      - network\n      - project\n      properties:\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: |-\n            An optional description of this resource. Provide this property when you\n            create the resource.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          readOnly: true\n          description: |-\n            [Output Only] The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels for this resource. These can only be added or modified\n            by the `setLabels` method. Each label key/value pair must comply with\n            [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: |-\n            Name of the resource. Provided by the client when the resource is created.\n            The name must be 1-63 characters long, and comply with\n            <a href=\"https://www.ietf.org/rfc/rfc1035.txt\" target=\"_blank\">RFC1035</a>.\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first\n            character must be a lowercase letter, and all following characters must be\n            a dash, lowercase letter, or digit, except the last character, which cannot\n            be a dash.\n          x-kubernetes-immutable: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: |-\n            URL of the network to which this VPN gateway is attached. Provided by the\n            client when the VPN gateway is created.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Compute/Network\n            field: selfLink\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          description: |-\n            URL of the region where the target VPN gateway resides.\n            You must specify this field as part of the HTTP request URL. It is\n            not settable as a field in the request body.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: '[Output Only] Server-defined URL for the resource.'\n          x-kubernetes-immutable: true\n        vpnInterface:\n          type: array\n          x-dcl-go-name: VpnInterface\n          readOnly: true\n          description: A list of interfaces on this VPN gateway.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: VpnGatewayVpnInterface\n            properties:\n              id:\n                type: integer\n                format: int64\n                x-dcl-go-name: Id\n                readOnly: true\n                description: The numeric ID of this VPN gateway interface.\n                x-kubernetes-immutable: true\n              ipAddress:\n                type: string\n                x-dcl-go-name: IPAddress\n                readOnly: true\n                description: '[Output Only] The external IP address for this VPN gateway\n                  interface.'\n                x-kubernetes-immutable: true\n")

// 5409 bytes
// MD5: 8068b56ef12fcbf399c1e537f592b1c1
