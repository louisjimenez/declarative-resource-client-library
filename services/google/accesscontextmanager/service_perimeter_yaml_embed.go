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
// gen_go_data -package accesscontextmanager -var YAML_service_perimeter blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/accesscontextmanager/service_perimeter.yaml

package accesscontextmanager

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/accesscontextmanager/service_perimeter.yaml
var YAML_service_perimeter = []byte("info:\n  title: AccessContextManager/ServicePerimeter\n  description: DCL Specification for the AccessContextManager ServicePerimeter resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a ServicePerimeter\n    parameters:\n    - name: ServicePerimeter\n      required: true\n      description: A full instance of a ServicePerimeter\n  apply:\n    description: The function used to apply information about a ServicePerimeter\n    parameters:\n    - name: ServicePerimeter\n      required: true\n      description: A full instance of a ServicePerimeter\n  delete:\n    description: The function used to delete a ServicePerimeter\n    parameters:\n    - name: ServicePerimeter\n      required: true\n      description: A full instance of a ServicePerimeter\n  deleteAll:\n    description: The function used to delete all ServicePerimeter\n    parameters:\n    - name: policy\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ServicePerimeter\n    parameters:\n    - name: policy\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ServicePerimeter:\n      title: ServicePerimeter\n      x-dcl-id: accessPolicies/{{policy}}/servicePerimeters/{{name}}\n      type: object\n      required:\n      - title\n      - policy\n      - name\n      properties:\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: 'Time the AccessPolicy was created in UTC. '\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: 'Description of the ServicePerimeter and its use. Does not\n            affect behavior. '\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Resource name for the ServicePerimeter. The short_name component\n            must begin with a letter and only include alphanumeric and ''_''. Format:\n            accessPolicies/{policy_id}/servicePerimeters/{short_name} '\n          x-kubernetes-immutable: true\n        perimeterType:\n          type: string\n          x-dcl-go-name: PerimeterType\n          x-dcl-go-type: ServicePerimeterPerimeterTypeEnum\n          description: 'Specifies the type of the Perimeter. There are two types:\n            regular and bridge. Regular Service Perimeter contains resources, access\n            levels, and restricted services. Every resource can be in at most ONE\n            regular Service Perimeter.  In addition to being in a regular service\n            perimeter, a resource can also be in zero or more perimeter bridges. A\n            perimeter bridge only contains resources. Cross project operations are\n            permitted if all effected resources share some perimeter (whether bridge\n            or regular). Perimeter Bridge does not contain access levels or services:\n            those are governed entirely by the regular perimeter that resource is\n            in.  Perimeter Bridges are typically useful when building more complex\n            topologies with many independent perimeters that need to share some data\n            with a common perimeter, but should not be able to share data among themselves.  Possible\n            values: PERIMETER_TYPE_REGULAR, PERIMETER_TYPE_BRIDGE'\n          x-kubernetes-immutable: true\n          default: PERIMETER_TYPE_REGULAR\n          enum:\n          - PERIMETER_TYPE_REGULAR\n          - PERIMETER_TYPE_BRIDGE\n        policy:\n          type: string\n          x-dcl-go-name: Policy\n          description: 'The AccessPolicy this ServicePerimeter lives in. Format: accessPolicies/{policy_id} '\n          x-kubernetes-immutable: true\n        spec:\n          type: object\n          x-dcl-go-name: Spec\n          x-dcl-go-type: ServicePerimeterSpec\n          description: 'Proposed (or dry run) ServicePerimeter configuration. This\n            configuration allows to specify and test ServicePerimeter configuration\n            without enforcing actual access restrictions. Only allowed to be set when\n            the \"useExplicitDryRunSpec\" flag is set. '\n          properties:\n            accessLevels:\n              type: array\n              x-dcl-go-name: AccessLevels\n              description: 'A list of AccessLevel resource names that allow resources\n                within the ServicePerimeter to be accessed from the internet. AccessLevels\n                listed must be in the same policy as this ServicePerimeter. Referencing\n                a nonexistent AccessLevel is a syntax error. If no AccessLevel names\n                are listed, resources within the perimeter can only be accessed via\n                GCP calls with request origins within the perimeter. For Service Perimeter\n                Bridge, must be empty.  Format: accessPolicies/{policy_id}/accessLevels/{access_level_name} '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            resources:\n              type: array\n              x-dcl-go-name: Resources\n              description: 'A list of GCP resources that are inside of the service\n                perimeter. Currently only projects are allowed. Format: projects/{project_number} '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            restrictedServices:\n              type: array\n              x-dcl-go-name: RestrictedServices\n              description: 'GCP services that are subject to the Service Perimeter\n                restrictions. Must contain a list of services. For example, if `storage.googleapis.com`\n                is specified, access to the storage buckets inside the perimeter must\n                meet the perimeter''s access restrictions. '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            vpcAccessibleServices:\n              type: object\n              x-dcl-go-name: VPCAccessibleServices\n              x-dcl-go-type: ServicePerimeterSpecVPCAccessibleServices\n              description: Configuration for APIs allowed within Perimeter.\n              properties:\n                allowedServices:\n                  type: array\n                  x-dcl-go-name: AllowedServices\n                  description: The list of APIs usable within the Service Perimeter.\n                    Must be empty unless 'enableRestriction' is True.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                enableRestriction:\n                  type: boolean\n                  x-dcl-go-name: EnableRestriction\n                  description: Whether to restrict API calls within the Service Perimeter\n                    to the list of APIs specified in 'allowedServices'.\n        status:\n          type: object\n          x-dcl-go-name: Status\n          x-dcl-go-type: ServicePerimeterStatus\n          description: 'ServicePerimeter configuration. Specifies sets of resources,\n            restricted services and access levels that determine perimeter content\n            and boundaries. '\n          properties:\n            accessLevels:\n              type: array\n              x-dcl-go-name: AccessLevels\n              description: 'A list of AccessLevel resource names that allow resources\n                within the ServicePerimeter to be accessed from the internet. AccessLevels\n                listed must be in the same policy as this ServicePerimeter. Referencing\n                a nonexistent AccessLevel is a syntax error. If no AccessLevel names\n                are listed, resources within the perimeter can only be accessed via\n                GCP calls with request origins within the perimeter. For Service Perimeter\n                Bridge, must be empty.  Format: accessPolicies/{policy_id}/accessLevels/{access_level_name} '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            resources:\n              type: array\n              x-dcl-go-name: Resources\n              description: 'A list of GCP resources that are inside of the service\n                perimeter. Currently only projects are allowed. Format: projects/{project_number} '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            restrictedServices:\n              type: array\n              x-dcl-go-name: RestrictedServices\n              description: 'GCP services that are subject to the Service Perimeter\n                restrictions. Must contain a list of services. For example, if `storage.googleapis.com`\n                is specified, access to the storage buckets inside the perimeter must\n                meet the perimeter''s access restrictions. '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            vpcAccessibleServices:\n              type: object\n              x-dcl-go-name: VPCAccessibleServices\n              x-dcl-go-type: ServicePerimeterStatusVPCAccessibleServices\n              description: Configuration for APIs allowed within Perimeter.\n              properties:\n                allowedServices:\n                  type: array\n                  x-dcl-go-name: AllowedServices\n                  description: The list of APIs usable within the Service Perimeter.\n                    Must be empty unless 'enableRestriction' is True.\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: string\n                    x-dcl-go-type: string\n                enableRestriction:\n                  type: boolean\n                  x-dcl-go-name: EnableRestriction\n                  description: Whether to restrict API calls within the Service Perimeter\n                    to the list of APIs specified in 'allowedServices'.\n        title:\n          type: string\n          x-dcl-go-name: Title\n          description: 'Human readable title. Must be unique within the Policy. '\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: 'Time the AccessPolicy was updated in UTC. '\n          x-kubernetes-immutable: true\n        useExplicitDryRunSpec:\n          type: boolean\n          x-dcl-go-name: UseExplicitDryRunSpec\n          description: Use explicit dry run spec flag. Ordinarily, a dry-run spec\n            implicitly exists for all Service Perimeters, and that spec is identical\n            to the status for those Service Perimeters. When this flag is set, it\n            inhibits the generation of the implicit spec, thereby allowing the user\n            to explicitly provide a configuration (\"spec\") to use in a dry-run version\n            of the Service Perimeter. This allows the user to test changes to the\n            enforced config (\"status\") without actually enforcing them. This testing\n            is done through analyzing the differences between currently enforced and\n            suggested restrictions. useExplicitDryRunSpec must be set to True if any\n            of the fields in the spec are set to non-default values.\n          x-kubernetes-immutable: true\n")

// 11799 bytes
// MD5: 709a657f7c2bbaadbe526138478be8d3
