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
// gen_go_data -package beta -var YAML_reservation blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/reservation.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/beta/reservation.yaml
var YAML_reservation = []byte("info:\n  title: Compute/Reservation\n  description: DCL Specification for the Compute Reservation resource\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Reservation\n    parameters:\n    - name: Reservation\n      required: true\n      description: A full instance of a Reservation\n  apply:\n    description: The function used to apply information about a Reservation\n    parameters:\n    - name: Reservation\n      required: true\n      description: A full instance of a Reservation\n  delete:\n    description: The function used to delete a Reservation\n    parameters:\n    - name: Reservation\n      required: true\n      description: A full instance of a Reservation\n  deleteAll:\n    description: The function used to delete all Reservation\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: zone\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Reservation\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: zone\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Reservation:\n      title: Reservation\n      x-dcl-id: projects/{{project}}/zones/{{zone}}/reservations/{{name}}\n      x-dcl-locations:\n      - zone\n      x-dcl-parent-container: project\n      type: object\n      required:\n      - zone\n      - name\n      - project\n      properties:\n        commitment:\n          type: string\n          x-dcl-go-name: Commitment\n          description: Full or partial URL to a parent commitment. This field displays\n            for reservations that are tied to a commitment.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n          x-kubernetes-immutable: true\n        id:\n          type: integer\n          format: int64\n          x-dcl-go-name: Id\n          description: The unique identifier for the resource. This identifier is\n            defined by the server.\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the resource, provided by the client when initially\n            creating the resource. The resource name must be 1-63 characters long,\n            and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically,\n            the name must be 1-63 characters long and match the regular expression\n            `)?` which means the first character must be a lowercase letter, and all\n            following characters must be a dash, lowercase letter, or digit, except\n            the last character, which cannot be a dash.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Server-defined fully-qualified URL for this resource.\n          x-kubernetes-immutable: true\n        specificReservation:\n          type: object\n          x-dcl-go-name: SpecificReservation\n          x-dcl-go-type: ReservationSpecificReservation\n          x-kubernetes-immutable: true\n          properties:\n            count:\n              type: integer\n              format: int64\n              x-dcl-go-name: Count\n              description: Specifies the number of resources that are allocated.\n              x-kubernetes-immutable: true\n            inUseCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: InUseCount\n              description: Indicates how many instances are in use.\n              x-kubernetes-immutable: true\n            instanceProperties:\n              type: object\n              x-dcl-go-name: InstanceProperties\n              x-dcl-go-type: ReservationSpecificReservationInstanceProperties\n              description: The instance properties for the reservation.\n              x-kubernetes-immutable: true\n              properties:\n                guestAccelerators:\n                  type: array\n                  x-dcl-go-name: GuestAccelerators\n                  x-kubernetes-immutable: true\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: ReservationSpecificReservationInstancePropertiesGuestAccelerators\n                    properties:\n                      acceleratorCount:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: AcceleratorCount\n                        description: The number of the guest accelerator cards exposed\n                          to this instance.\n                        x-kubernetes-immutable: true\n                      acceleratorType:\n                        type: string\n                        x-dcl-go-name: AcceleratorType\n                        description: 'Full or partial URL of the accelerator type\n                          resource to attach to this instance. For example: `projects/my-project/zones/us-central1-c/acceleratorTypes/nvidia-tesla-p100`\n                          If you are creating an instance template, specify only the\n                          accelerator name. See [GPUs on Compute Engine](/compute/docs/gpus/#introduction)\n                          for a full list of accelerator types.'\n                        x-kubernetes-immutable: true\n                localSsds:\n                  type: array\n                  x-dcl-go-name: LocalSsds\n                  x-kubernetes-immutable: true\n                  x-dcl-send-empty: true\n                  x-dcl-list-type: list\n                  items:\n                    type: object\n                    x-dcl-go-type: ReservationSpecificReservationInstancePropertiesLocalSsds\n                    properties:\n                      diskSizeGb:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: DiskSizeGb\n                        description: Specifies the size of the disk in base-2 GB.\n                        x-kubernetes-immutable: true\n                      interface:\n                        type: string\n                        x-dcl-go-name: Interface\n                        x-dcl-go-type: ReservationSpecificReservationInstancePropertiesLocalSsdsInterfaceEnum\n                        description: Specifies the disk interface to use for attaching\n                          this disk, which is either `SCSI` or `NVME`. The default\n                          is `SCSI`. For performance characteristics of SCSI over\n                          NVMe, see [Local SSD performance](/compute/docs/disks#localssds).\n                        x-kubernetes-immutable: true\n                        enum:\n                        - SCSI\n                        - NVME\n                        - NVDIMM\n                machineType:\n                  type: string\n                  x-dcl-go-name: MachineType\n                  description: Specifies type of machine (name only) which has fixed\n                    number of vCPUs and fixed amount of memory. This also includes\n                    specifying custom machine type following custom-NUMBER_OF_CPUS-AMOUNT_OF_MEMORY\n                    pattern.\n                  x-kubernetes-immutable: true\n                minCpuPlatform:\n                  type: string\n                  x-dcl-go-name: MinCpuPlatform\n                  description: Minimum cpu platform the reservation.\n                  x-kubernetes-immutable: true\n        specificReservationRequired:\n          type: boolean\n          x-dcl-go-name: SpecificReservationRequired\n          x-kubernetes-immutable: true\n        status:\n          type: string\n          x-dcl-go-name: Status\n          x-dcl-go-type: ReservationStatusEnum\n          readOnly: true\n          description: 'The status of the reservation. Possible values: PENDING, RUNNING,\n            DONE'\n          x-kubernetes-immutable: true\n          enum:\n          - PENDING\n          - RUNNING\n          - DONE\n        zone:\n          type: string\n          x-dcl-go-name: Zone\n          description: Zone in which the reservation resides. A zone must be provided\n            if the reservation is created within a commitment.\n          x-kubernetes-immutable: true\n")

// 8824 bytes
// MD5: a2a0dc4370e6f05e99e6c29b9b3c1c8f
