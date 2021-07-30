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
// gen_go_data -package beta -var YAML_target_pool blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/target_pool.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/target_pool.yaml
var YAML_target_pool = []byte("info:\n  title: Compute/TargetPool\n  description: DCL Specification for the Compute TargetPool resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a TargetPool\n    parameters:\n    - name: TargetPool\n      required: true\n      description: A full instance of a TargetPool\n  apply:\n    description: The function used to apply information about a TargetPool\n    parameters:\n    - name: TargetPool\n      required: true\n      description: A full instance of a TargetPool\n  delete:\n    description: The function used to delete a TargetPool\n    parameters:\n    - name: TargetPool\n      required: true\n      description: A full instance of a TargetPool\n  deleteAll:\n    description: The function used to delete all TargetPool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many TargetPool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: region\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    TargetPool:\n      title: TargetPool\n      x-dcl-id: projects/{{project}}/regions/{{region}}/targetPools/{{name}}\n      x-dcl-locations:\n      - region\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - project\n      properties:\n        backupPool:\n          type: string\n          x-dcl-go-name: BackupPool\n          description: 'This field is applicable only when the containing target pool\n            is serving a forwarding rule as the primary pool, and its `failoverRatio`\n            field is properly set to a value between [0, 1]. `backupPool` and `failoverRatio`\n            together define the fallback behavior of the primary target pool: if the\n            ratio of the healthy instances in the primary pool is at or below `failoverRatio`,\n            traffic arriving at the load-balanced IP will be directed to the backup\n            pool. In case where `failoverRatio` and `backupPool` are not set, or all\n            the instances in the backup pool are unhealthy, the traffic will be directed\n            back to the primary pool in the \"force\" mode, where traffic will be spread\n            to the healthy instances with the best effort, or to all instances when\n            no instance is healthy.'\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n          x-kubernetes-immutable: true\n        failoverRatio:\n          type: number\n          format: double\n          x-dcl-go-name: FailoverRatio\n          description: 'This field is applicable only when the containing target pool\n            is serving a forwarding rule as the primary pool (i.e., not as a backup\n            pool to some other target pool). The value of the field must be in [0,\n            1]. If set, `backupPool` must also be set. They together define the fallback\n            behavior of the primary target pool: if the ratio of the healthy instances\n            in the primary pool is at or below this number, traffic arriving at the\n            load-balanced IP will be directed to the backup pool. In case where `failoverRatio`\n            is not set or all the instances in the backup pool are unhealthy, the\n            traffic will be directed back to the primary pool in the \"force\" mode,\n            where traffic will be spread to the healthy instances with the best effort,\n            or to all instances when no instance is healthy.'\n          x-kubernetes-immutable: true\n        healthChecks:\n          type: array\n          x-dcl-go-name: HealthChecks\n          description: The URL of the HttpHealthCheck resource. A member instance\n            in this pool is considered healthy if and only if the health checks pass.\n            An empty list means all member instances will be considered healthy at\n            all times. Only HttpHealthChecks are supported. Only one health check\n            may be specified.\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        instances:\n          type: array\n          x-dcl-go-name: Instances\n          description: A list of resource URLs to the virtual machine instances serving\n            this pool. They must live in zones contained in the same region as this\n            pool.\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for this target pool.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          description: '[Output Only] URL of the region where the target pool resides.'\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          description: '[Output Only] Server-defined URL for the resource.'\n          x-kubernetes-immutable: true\n        sessionAffinity:\n          type: string\n          x-dcl-go-name: SessionAffinity\n          x-dcl-go-type: TargetPoolSessionAffinityEnum\n          description: \"Session affinity option, must be one of the following values:\n            \\ \\n`NONE`: Connections from the same client IP may go to any instance\n            in the pool.  \\n`CLIENT_IP`: Connections from the same client IP will\n            go to the same instance in the pool while that instance remains healthy.\n            \\ \\n`CLIENT_IP_PROTO`: Connections from the same client IP with the same\n            IP protocol will go to the same instance in the pool while that instance\n            remains healthy.\"\n          x-kubernetes-immutable: true\n          enum:\n          - NONE\n          - CLIENT_IP\n          - CLIENT_IP_PROTO\n          - GENERATED_COOKIE\n          - CLIENT_IP_PORT_PROTO\n          - HTTP_COOKIE\n          - HEADER_FIELD\n")

// 7055 bytes
// MD5: b94bb783a6ab175be60f3daa51de778f
