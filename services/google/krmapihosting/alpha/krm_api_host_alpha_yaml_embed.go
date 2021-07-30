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
// gen_go_data -package alpha -var YAML_krm_api_host blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/krmapihosting/alpha/krm_api_host.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/krmapihosting/alpha/krm_api_host.yaml
var YAML_krm_api_host = []byte("info:\n  title: Krmapihosting/KrmApiHost\n  description: DCL Specification for the Krmapihosting KrmApiHost resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a KrmApiHost\n    parameters:\n    - name: KrmApiHost\n      required: true\n      description: A full instance of a KrmApiHost\n  apply:\n    description: The function used to apply information about a KrmApiHost\n    parameters:\n    - name: KrmApiHost\n      required: true\n      description: A full instance of a KrmApiHost\n  delete:\n    description: The function used to delete a KrmApiHost\n    parameters:\n    - name: KrmApiHost\n      required: true\n      description: A full instance of a KrmApiHost\n  deleteAll:\n    description: The function used to delete all KrmApiHost\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many KrmApiHost\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    KrmApiHost:\n      title: KrmApiHost\n      x-dcl-id: projects/{{project}}/locations/{{location}}/krmApiHosts/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      - location\n      properties:\n        bundlesConfig:\n          type: object\n          x-dcl-go-name: BundlesConfig\n          x-dcl-go-type: KrmApiHostBundlesConfig\n          description: Configuration for the bundles that are enabled on the KrmApiHost.\n          properties:\n            configControllerConfig:\n              type: object\n              x-dcl-go-name: ConfigControllerConfig\n              x-dcl-go-type: KrmApiHostBundlesConfigConfigControllerConfig\n              description: Configuration for the Config Controller bundle.\n              properties:\n                enabled:\n                  type: boolean\n                  x-dcl-go-name: Enabled\n                  description: Whether the Config Controller bundle is enabled on\n                    the KrmApiHost.\n        gkeResourceLink:\n          type: string\n          x-dcl-go-name: GkeResourceLink\n          readOnly: true\n          description: Output only. KrmApiHost GCP self link used for identifying\n            the underlying endpoint (GKE cluster currently).\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels are used for additional information for a KrmApiHost.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        managementConfig:\n          type: object\n          x-dcl-go-name: ManagementConfig\n          x-dcl-go-type: KrmApiHostManagementConfig\n          description: Configuration of the cluster management\n          properties:\n            standardManagementConfig:\n              type: object\n              x-dcl-go-name: StandardManagementConfig\n              x-dcl-go-type: KrmApiHostManagementConfigStandardManagementConfig\n              description: Configuration of the standard (GKE) cluster management\n              properties:\n                clusterCidrBlock:\n                  type: string\n                  x-dcl-go-name: ClusterCidrBlock\n                  description: The IP address range for the cluster pod IPs. Set to\n                    blank to have a range chosen with the default size. Set to /netmask\n                    (e.g. /14) to have a range chosen with a specific netmask. Set\n                    to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private\n                    networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick\n                    a specific range to use.\n                clusterNamedRange:\n                  type: string\n                  x-dcl-go-name: ClusterNamedRange\n                  description: The name of the existing secondary range in the cluster's\n                    subnetwork to use for pod IP addresses. Alternatively, cluster_cidr_block\n                    can be used to automatically create a GKE-managed one.\n                manBlock:\n                  type: string\n                  x-dcl-go-name: ManBlock\n                  description: Master Authorized Network. Allows access to the k8s\n                    master from this block.\n                masterIPv4CidrBlock:\n                  type: string\n                  x-dcl-go-name: MasterIPv4CidrBlock\n                  description: The /28 network that the masters will use.\n                network:\n                  type: string\n                  x-dcl-go-name: Network\n                  description: Existing VPC Network to put the GKE cluster and nodes\n                    in.\n                servicesCidrBlock:\n                  type: string\n                  x-dcl-go-name: ServicesCidrBlock\n                  description: The IP address range for the cluster service IPs. Set\n                    to blank to have a range chosen with the default size. Set to\n                    /netmask (e.g. /14) to have a range chosen with a specific netmask.\n                    Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private\n                    networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick\n                    a specific range to use.\n                servicesNamedRange:\n                  type: string\n                  x-dcl-go-name: ServicesNamedRange\n                  description: The name of the existing secondary range in the cluster's\n                    subnetwork to use for service ClusterIPs. Alternatively, services_cidr_block\n                    can be used to automatically create a GKE-managed one.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. The name of this KrmApiHost resource in the format:\n            ''projects/{project_id}/locations/{location}/krmApiHosts/{krm_api_host_id}''.'\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n          x-dcl-references:\n          - resource: Krmapihosting/KrmApiHost\n            field: selfLink\n            parent: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: KrmApiHostStateEnum\n          readOnly: true\n          description: 'Output only. The current state of the internal state machine\n            for the KrmApiHost. Possible values: STATE_UNSPECIFIED, CREATING, RUNNING,\n            DELETING, SUSPENDED, READ_ONLY'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - CREATING\n          - RUNNING\n          - DELETING\n          - SUSPENDED\n          - READ_ONLY\n        usePrivateEndpoint:\n          type: boolean\n          x-dcl-go-name: UsePrivateEndpoint\n          description: Only allow access to the master's private endpoint IP.\n")

// 7484 bytes
// MD5: 4a2e1c02952edcacfede0513e4c73c21
