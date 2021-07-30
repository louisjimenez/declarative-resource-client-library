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
// gen_go_data -package beta -var YAML_ssl_cert blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/sql/beta/ssl_cert.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/sql/beta/ssl_cert.yaml
var YAML_ssl_cert = []byte("info:\n  title: Sql/SslCert\n  description: DCL Specification for the Sql SslCert resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a SslCert\n    parameters:\n    - name: SslCert\n      required: true\n      description: A full instance of a SslCert\n  apply:\n    description: The function used to apply information about a SslCert\n    parameters:\n    - name: SslCert\n      required: true\n      description: A full instance of a SslCert\n  delete:\n    description: The function used to delete a SslCert\n    parameters:\n    - name: SslCert\n      required: true\n      description: A full instance of a SslCert\n  deleteAll:\n    description: The function used to delete all SslCert\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: instance\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many SslCert\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: instance\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    SslCert:\n      title: SslCert\n      x-dcl-id: projects/{{project}}/instances/{{instance}}/sslCerts/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - commonName\n      - instance\n      properties:\n        cert:\n          type: string\n          x-dcl-go-name: Cert\n          readOnly: true\n          description: PEM representation.\n          x-kubernetes-immutable: true\n        certSerialNumber:\n          type: string\n          x-dcl-go-name: CertSerialNumber\n          readOnly: true\n          description: Serial number, as extracted from the certificate.\n          x-kubernetes-immutable: true\n        commonName:\n          type: string\n          x-dcl-go-name: CommonName\n          description: User supplied name.  Constrained to [a-zA-Z.-_ ]+.\n          x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339)\n            format, for example **2012-11-15T16:19:00.094Z**\n          x-kubernetes-immutable: true\n        expirationTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: ExpirationTime\n          readOnly: true\n          description: The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339)\n            format, for example **2012-11-15T16:19:00.094Z**.\n          x-kubernetes-immutable: true\n        instance:\n          type: string\n          x-dcl-go-name: Instance\n          description: Name of the database instance.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Sha1 Fingerprint.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n")

// 3397 bytes
// MD5: 1d35efd461c04bd441e6a0ac34b28062
