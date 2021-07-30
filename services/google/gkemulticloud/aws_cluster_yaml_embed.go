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
// gen_go_data -package gkemulticloud -var YAML_aws_cluster blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkemulticloud/aws_cluster.yaml

package gkemulticloud

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/gkemulticloud/aws_cluster.yaml
var YAML_aws_cluster = []byte("info:\n  title: Gkemulticloud/AwsCluster\n  description: DCL Specification for the Gkemulticloud AwsCluster resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a AwsCluster\n    parameters:\n    - name: AwsCluster\n      required: true\n      description: A full instance of a AwsCluster\n  apply:\n    description: The function used to apply information about a AwsCluster\n    parameters:\n    - name: AwsCluster\n      required: true\n      description: A full instance of a AwsCluster\n  delete:\n    description: The function used to delete a AwsCluster\n    parameters:\n    - name: AwsCluster\n      required: true\n      description: A full instance of a AwsCluster\n  deleteAll:\n    description: The function used to delete all AwsCluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many AwsCluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    AwsCluster:\n      title: AwsCluster\n      x-dcl-id: projects/{{project}}/locations/{{location}}/awsClusters/{{name}}\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - name\n      - networking\n      - awsRegion\n      - controlPlane\n      - authorization\n      - project\n      - location\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: 'Optional. Annotations on the cluster. This field has the same\n            restrictions as Kubernetes annotations. The total size of all keys and\n            values combined is limited to 256k. Key can have 2 segments: prefix (optional)\n            and name (required), separated by a slash (/). Prefix must be a DNS subdomain.\n            Name must be 63 characters or less, begin and end with alphanumerics,\n            with dashes (-), underscores (_), dots (.), and alphanumerics between.'\n          x-kubernetes-immutable: true\n        authorization:\n          type: object\n          x-dcl-go-name: Authorization\n          x-dcl-go-type: AwsClusterAuthorization\n          description: Required. Configuration related to the cluster RBAC settings.\n          x-kubernetes-immutable: true\n          required:\n          - adminUsers\n          properties:\n            adminUsers:\n              type: array\n              x-dcl-go-name: AdminUsers\n              description: Required. Users to perform operations as a cluster admin.\n                A managed ClusterRoleBinding will be created to grant the `cluster-admin`\n                ClusterRole to the users. At most one user can be specified. For more\n                info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: AwsClusterAuthorizationAdminUsers\n                required:\n                - username\n                properties:\n                  username:\n                    type: string\n                    x-dcl-go-name: Username\n                    description: Required. The name of the user, e.g. `my-gcp-id@gmail.com`.\n                    x-kubernetes-immutable: true\n        awsRegion:\n          type: string\n          x-dcl-go-name: AwsRegion\n          description: Required. The AWS region where the cluster runs. Each Google\n            Cloud region supports a subset of nearby AWS regions. You can call to\n            list all supported AWS regions within a given Google Cloud region.\n          x-kubernetes-immutable: true\n        controlPlane:\n          type: object\n          x-dcl-go-name: ControlPlane\n          x-dcl-go-type: AwsClusterControlPlane\n          description: Required. Configuration related to the cluster control plane.\n          x-kubernetes-immutable: true\n          required:\n          - version\n          - subnetIds\n          - iamInstanceProfile\n          - databaseEncryption\n          - awsServicesAuthentication\n          properties:\n            awsServicesAuthentication:\n              type: object\n              x-dcl-go-name: AwsServicesAuthentication\n              x-dcl-go-type: AwsClusterControlPlaneAwsServicesAuthentication\n              description: Required. Authentication configuration for management of\n                AWS resources.\n              x-kubernetes-immutable: true\n              required:\n              - roleArn\n              properties:\n                roleArn:\n                  type: string\n                  x-dcl-go-name: RoleArn\n                  description: Required. The Amazon Resource Name (ARN) of the role\n                    that the Anthos Multi-Cloud API will assume when managing AWS\n                    resources on your account.\n                  x-kubernetes-immutable: true\n                roleSessionName:\n                  type: string\n                  x-dcl-go-name: RoleSessionName\n                  description: Optional. An identifier for the assumed role session.\n                    When unspecified, it defaults to `multicloud-service-agent`.\n                  x-kubernetes-immutable: true\n            databaseEncryption:\n              type: object\n              x-dcl-go-name: DatabaseEncryption\n              x-dcl-go-type: AwsClusterControlPlaneDatabaseEncryption\n              description: Required. The ARN of the AWS KMS key used to encrypt cluster\n                secrets.\n              x-kubernetes-immutable: true\n              required:\n              - kmsKeyArn\n              properties:\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: Required. The ARN of the AWS KMS key used to encrypt\n                    cluster secrets.\n                  x-kubernetes-immutable: true\n            iamInstanceProfile:\n              type: string\n              x-dcl-go-name: IamInstanceProfile\n              description: Required. The name of the AWS IAM instance pofile to assign\n                to each control plane replica.\n              x-kubernetes-immutable: true\n            instanceType:\n              type: string\n              x-dcl-go-name: InstanceType\n              description: Optional. The AWS instance type. When unspecified, it defaults\n                to `t3.medium`.\n              x-kubernetes-immutable: true\n            mainVolume:\n              type: object\n              x-dcl-go-name: MainVolume\n              x-dcl-go-type: AwsClusterControlPlaneMainVolume\n              description: Optional. Configuration related to the main volume provisioned\n                for each control plane replica. The main volume is in charge of storing\n                all of the cluster's etcd state. Volumes will be provisioned in the\n                availability zone associated with the corresponding subnet. When unspecified,\n                it defaults to 8 GiB with the GP2 volume type.\n              x-kubernetes-immutable: true\n              properties:\n                iops:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Iops\n                  description: Optional. The number of I/O operations per second (IOPS)\n                    to provision for GP3 volume.\n                  x-kubernetes-immutable: true\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: Optional. The Amazon Resource Name (ARN) of the Customer\n                    Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified,\n                    the default Amazon managed key associated to the AWS region where\n                    this cluster runs will be used.\n                  x-kubernetes-immutable: true\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the volume, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-kubernetes-immutable: true\n                volumeType:\n                  type: string\n                  x-dcl-go-name: VolumeType\n                  x-dcl-go-type: AwsClusterControlPlaneMainVolumeVolumeTypeEnum\n                  description: 'Optional. Type of the EBS volume. When unspecified,\n                    it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED,\n                    GP2, GP3'\n                  x-kubernetes-immutable: true\n                  enum:\n                  - VOLUME_TYPE_UNSPECIFIED\n                  - GP2\n                  - GP3\n            rootVolume:\n              type: object\n              x-dcl-go-name: RootVolume\n              x-dcl-go-type: AwsClusterControlPlaneRootVolume\n              description: Optional. Configuration related to the root volume provisioned\n                for each control plane replica. Volumes will be provisioned in the\n                availability zone associated with the corresponding subnet. When unspecified,\n                it defaults to 32 GiB with the GP2 volume type.\n              x-kubernetes-immutable: true\n              properties:\n                iops:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Iops\n                  description: Optional. The number of I/O operations per second (IOPS)\n                    to provision for GP3 volume.\n                  x-kubernetes-immutable: true\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: Optional. The Amazon Resource Name (ARN) of the Customer\n                    Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified,\n                    the default Amazon managed key associated to the AWS region where\n                    this cluster runs will be used.\n                  x-kubernetes-immutable: true\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the volume, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-kubernetes-immutable: true\n                volumeType:\n                  type: string\n                  x-dcl-go-name: VolumeType\n                  x-dcl-go-type: AwsClusterControlPlaneRootVolumeVolumeTypeEnum\n                  description: 'Optional. Type of the EBS volume. When unspecified,\n                    it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED,\n                    GP2, GP3'\n                  x-kubernetes-immutable: true\n                  enum:\n                  - VOLUME_TYPE_UNSPECIFIED\n                  - GP2\n                  - GP3\n            securityGroupIds:\n              type: array\n              x-dcl-go-name: SecurityGroupIds\n              description: Optional. The IDs of additional security groups to add\n                to control plane replicas. The Anthos Multi-Cloud API will automatically\n                create and manage security groups with the minimum rules needed for\n                a functioning cluster.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            sshConfig:\n              type: object\n              x-dcl-go-name: SshConfig\n              x-dcl-go-type: AwsClusterControlPlaneSshConfig\n              description: Optional. SSH configuration for how to access the underlying\n                control plane machines.\n              x-kubernetes-immutable: true\n              required:\n              - ec2KeyPair\n              properties:\n                ec2KeyPair:\n                  type: string\n                  x-dcl-go-name: Ec2KeyPair\n                  description: Required. The name of the EC2 key pair used to login\n                    into cluster machines.\n                  x-kubernetes-immutable: true\n            subnetIds:\n              type: array\n              x-dcl-go-name: SubnetIds\n              description: Required. The list of subnets where control plane replicas\n                will run. A replica will be provisioned on each subnet and up to three\n                values can be provided. Each subnet must be in a different AWS Availability\n                Zone (AZ).\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            tags:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Tags\n              description: Optional. A set of AWS resource tags to propagate to all\n                underlying managed AWS resources. Specify at most 50 pairs containing\n                alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127\n                Unicode characters. Values can be up to 255 Unicode characters.\n              x-kubernetes-immutable: true\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Required. The Kubernetes version to run on control plane\n                replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions\n                on a given Google Cloud region by calling .\n              x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this cluster was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A human readable description of this cluster. Cannot\n            be longer than 255 UTF-8 encoded bytes.\n          x-kubernetes-immutable: true\n        endpoint:\n          type: string\n          x-dcl-go-name: Endpoint\n          readOnly: true\n          description: Output only. The endpoint of the cluster's API server.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Allows clients to perform consistent read-modify-writes through\n            optimistic concurrency control. May be sent on update and delete requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of this resource. Cluster names are formatted as `projects/<project-number>/locations/<region>/awsClusters/<cluster-id>`.\n            See (https://cloud.google.com/apis/design/resource_names) for more details\n            on GCP resource names.\n          x-kubernetes-immutable: true\n        networking:\n          type: object\n          x-dcl-go-name: Networking\n          x-dcl-go-type: AwsClusterNetworking\n          description: Required. Cluster-wide networking configuration.\n          x-kubernetes-immutable: true\n          required:\n          - vpcId\n          - podAddressCidrBlocks\n          - serviceAddressCidrBlocks\n          - serviceLoadBalancerSubnetIds\n          properties:\n            podAddressCidrBlocks:\n              type: array\n              x-dcl-go-name: PodAddressCidrBlocks\n              description: Required. All pods in the cluster are assigned an RFC1918\n                IPv4 address from these ranges. Only a single range is supported.\n                This field cannot be changed after creation.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            serviceAddressCidrBlocks:\n              type: array\n              x-dcl-go-name: ServiceAddressCidrBlocks\n              description: Required. All services in the cluster are assigned an RFC1918\n                IPv4 address from these ranges. Only a single range is supported.\n                This field cannot be changed after creation.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            serviceLoadBalancerSubnetIds:\n              type: array\n              x-dcl-go-name: ServiceLoadBalancerSubnetIds\n              description: Required. When creating Kubernetes services of type 'Load\n                Balancer', the load balancer will be created in these subnets.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            vpcId:\n              type: string\n              x-dcl-go-name: VPCId\n              description: Required. The VPC associated with the cluster. All component\n                clusters (i.e. control plane and node pools) run on a single VPC.\n                This field cannot be changed after creation.\n              x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        reconciling:\n          type: boolean\n          x-dcl-go-name: Reconciling\n          readOnly: true\n          description: Output only. If set, there are currently changes in flight\n            to the cluster.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: AwsClusterStateEnum\n          readOnly: true\n          description: 'Output only. The current state of the cluster. Possible values:\n            STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,\n            DEGRADED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PROVISIONING\n          - RUNNING\n          - RECONCILING\n          - STOPPING\n          - ERROR\n          - DEGRADED\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. A globally unique identifier for the cluster.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time at which this cluster was last updated.\n          x-kubernetes-immutable: true\n        workloadIdentityConfig:\n          type: object\n          x-dcl-go-name: WorkloadIdentityConfig\n          x-dcl-go-type: AwsClusterWorkloadIdentityConfig\n          readOnly: true\n          description: Output only. Workload Identity settings.\n          x-kubernetes-immutable: true\n          properties:\n            identityProvider:\n              type: string\n              x-dcl-go-name: IdentityProvider\n              description: The ID of the OIDC Identity Provider (IdP) associated to\n                the Workload Identity Pool.\n              x-kubernetes-immutable: true\n            issuerUri:\n              type: string\n              x-dcl-go-name: IssuerUri\n              description: The OIDC issuer URL for this cluster.\n              x-kubernetes-immutable: true\n            workloadPool:\n              type: string\n              x-dcl-go-name: WorkloadPool\n              description: The Workload Identity Pool associated to the cluster.\n              x-kubernetes-immutable: true\n")

// 20666 bytes
// MD5: 0f0a6e7febeab5b9adae414a6a8f3718
