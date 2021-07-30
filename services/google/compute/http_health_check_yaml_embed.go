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
// gen_go_data -package compute -var YAML_http_health_check blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/http_health_check.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/http_health_check.yaml
var YAML_http_health_check = []byte("info:\n  title: Compute/HttpHealthCheck\n  description: DCL Specification for the Compute HttpHealthCheck resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a HttpHealthCheck\n    parameters:\n    - name: HttpHealthCheck\n      required: true\n      description: A full instance of a HttpHealthCheck\n  apply:\n    description: The function used to apply information about a HttpHealthCheck\n    parameters:\n    - name: HttpHealthCheck\n      required: true\n      description: A full instance of a HttpHealthCheck\n  delete:\n    description: The function used to delete a HttpHealthCheck\n    parameters:\n    - name: HttpHealthCheck\n      required: true\n      description: A full instance of a HttpHealthCheck\n  deleteAll:\n    description: The function used to delete all HttpHealthCheck\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many HttpHealthCheck\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    HttpHealthCheck:\n      title: HttpHealthCheck\n      x-dcl-id: projects/{{project}}/global/httpHealthChecks/{{name}}\n      x-dcl-locations:\n      - global\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        checkIntervalSec:\n          type: integer\n          format: int64\n          x-dcl-go-name: CheckIntervalSec\n          description: How often (in seconds) to send a health check. The default\n            value is `5` seconds.\n          default: 5\n        creationTimestamp:\n          type: string\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: Creation timestamp in RFC3339 text format.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n        healthyThreshold:\n          type: integer\n          format: int64\n          x-dcl-go-name: HealthyThreshold\n          description: A so-far unhealthy instance will be marked healthy after this\n            many consecutive successes. The default value is `2`.\n          default: 2\n        host:\n          type: string\n          x-dcl-go-name: Host\n          description: The value of the host header in the HTTP health check request.\n            If left empty (default value), the public IP on behalf of which this health\n            check is performed will be used.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource. Provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash.\n          x-kubernetes-immutable: true\n        port:\n          type: integer\n          format: int64\n          x-dcl-go-name: Port\n          description: The TCP port number for the HTTP health check request. The\n            default value is `80`.\n          default: 80\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        requestPath:\n          type: string\n          x-dcl-go-name: RequestPath\n          description: The request path of the HTTP health check request. The default\n            value is `/`. This field does not support query parameters.\n          default: /\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: This health check's self-reference.\n          x-kubernetes-immutable: true\n        timeoutSec:\n          type: integer\n          format: int64\n          x-dcl-go-name: TimeoutSec\n          description: How long (in seconds) to wait before claiming failure. The\n            default value is `5` seconds. It is invalid for `timeoutSec` to have greater\n            value than `checkIntervalSec`.\n          default: 5\n        unhealthyThreshold:\n          type: integer\n          format: int64\n          x-dcl-go-name: UnhealthyThreshold\n          description: A so-far healthy instance will be marked unhealthy after this\n            many consecutive failures. The default value is `2`.\n          default: 2\n")

// 4958 bytes
// MD5: 21398a37d39ed658c3830e5dfc7d72bc
