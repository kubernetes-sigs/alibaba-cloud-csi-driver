package crds

// Template ...
type Template struct {
}

// GetVolumeGroupSnapshotsCRDv1 returns the VolumeSnapshot CRD Template
func (temp *Template) GetVolumeGroupSnapshotsCRDv1(kVersion string) string {
	return `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/995"
    controller-gen.kubebuilder.io/version: v0.12.0
  creationTimestamp: null
  name: volumegroupsnapshots.groupsnapshot.storage.k8s.io
spec:
  group: groupsnapshot.storage.k8s.io
  names:
    kind: VolumeGroupSnapshot
    listKind: VolumeGroupSnapshotList
    plural: volumegroupsnapshots
    shortNames:
    - vgs
    singular: volumegroupsnapshot
  scope: Namespaced
  versions:
  - additionalPrinterColumns:
    - description: Indicates if all the individual snapshots in the group are ready
        to be used to restore a group of volumes.
      jsonPath: .status.readyToUse
      name: ReadyToUse
      type: boolean
    - description: The name of the VolumeGroupSnapshotClass requested by the VolumeGroupSnapshot.
      jsonPath: .spec.volumeGroupSnapshotClassName
      name: VolumeGroupSnapshotClass
      type: string
    - description: Name of the VolumeGroupSnapshotContent object to which the VolumeGroupSnapshot
        object intends to bind to. Please note that verification of binding actually
        requires checking both VolumeGroupSnapshot and VolumeGroupSnapshotContent
        to ensure both are pointing at each other. Binding MUST be verified prior
        to usage of this object.
      jsonPath: .status.boundVolumeGroupSnapshotContentName
      name: VolumeGroupSnapshotContent
      type: string
    - description: Timestamp when the point-in-time group snapshot was taken by the
        underlying storage system.
      jsonPath: .status.creationTime
      name: CreationTime
      type: date
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: VolumeGroupSnapshot is a user's request for creating either a
          point-in-time group snapshot or binding to a pre-existing group snapshot.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: Spec defines the desired characteristics of a group snapshot
              requested by a user. Required.
            properties:
              source:
                description: Source specifies where a group snapshot will be created
                  from. This field is immutable after creation. Required.
                properties:
                  selector:
                    description: Selector is a label query over persistent volume
                      claims that are to be grouped together for snapshotting. This
                      labelSelector will be used to match the label added to a PVC.
                      If the label is added or removed to a volume after a group snapshot
                      is created, the existing group snapshots won't be modified.
                      Once a VolumeGroupSnapshotContent is created and the sidecar
                      starts to process it, the volume list will not change with retries.
                    properties:
                      matchExpressions:
                        description: matchExpressions is a list of label selector
                          requirements. The requirements are ANDed.
                        items:
                          description: A label selector requirement is a selector
                            that contains values, a key, and an operator that relates
                            the key and values.
                          properties:
                            key:
                              description: key is the label key that the selector
                                applies to.
                              type: string
                            operator:
                              description: operator represents a key's relationship
                                to a set of values. Valid operators are In, NotIn,
                                Exists and DoesNotExist.
                              type: string
                            values:
                              description: values is an array of string values. If
                                the operator is In or NotIn, the values array must
                                be non-empty. If the operator is Exists or DoesNotExist,
                                the values array must be empty. This array is replaced
                                during a strategic merge patch.
                              items:
                                type: string
                              type: array
                          required:
                          - key
                          - operator
                          type: object
                        type: array
                      matchLabels:
                        additionalProperties:
                          type: string
                        description: matchLabels is a map of {key,value} pairs. A
                          single {key,value} in the matchLabels map is equivalent
                          to an element of matchExpressions, whose key field is "key",
                          the operator is "In", and the values array contains only
                          "value". The requirements are ANDed.
                        type: object
                    type: object
                    x-kubernetes-map-type: atomic
                  volumeGroupSnapshotContentName:
                    description: VolumeGroupSnapshotContentName specifies the name
                      of a pre-existing VolumeGroupSnapshotContent object representing
                      an existing volume group snapshot. This field should be set
                      if the volume group snapshot already exists and only needs a
                      representation in Kubernetes. This field is immutable.
                    type: string
                type: object
              volumeGroupSnapshotClassName:
                description: VolumeGroupSnapshotClassName is the name of the VolumeGroupSnapshotClass
                  requested by the VolumeGroupSnapshot. VolumeGroupSnapshotClassName
                  may be left nil to indicate that the default class will be used.
                  Empty string is not allowed for this field.
                type: string
            required:
            - source
            type: object
          status:
            description: Status represents the current information of a group snapshot.
              Consumers must verify binding between VolumeGroupSnapshot and VolumeGroupSnapshotContent
              objects is successful (by validating that both VolumeGroupSnapshot and
              VolumeGroupSnapshotContent point to each other) before using this object.
            properties:
              boundVolumeGroupSnapshotContentName:
                description: 'BoundVolumeGroupSnapshotContentName is the name of the
                  VolumeGroupSnapshotContent object to which this VolumeGroupSnapshot
                  object intends to bind to. If not specified, it indicates that the
                  VolumeGroupSnapshot object has not been successfully bound to a
                  VolumeGroupSnapshotContent object yet. NOTE: To avoid possible security
                  issues, consumers must verify binding between VolumeGroupSnapshot
                  and VolumeGroupSnapshotContent objects is successful (by validating
                  that both VolumeGroupSnapshot and VolumeGroupSnapshotContent point
                  at each other) before using this object.'
                type: string
              creationTime:
                description: CreationTime is the timestamp when the point-in-time
                  group snapshot is taken by the underlying storage system. If not
                  specified, it may indicate that the creation time of the group snapshot
                  is unknown. The format of this field is a Unix nanoseconds time
                  encoded as an int64. On Unix, the command date +%s%N returns the
                  current time in nanoseconds since 1970-01-01 00:00:00 UTC.
                format: date-time
                type: string
              error:
                description: Error is the last observed error during group snapshot
                  creation, if any. This field could be helpful to upper level controllers
                  (i.e., application controller) to decide whether they should continue
                  on waiting for the group snapshot to be created based on the type
                  of error reported. The snapshot controller will keep retrying when
                  an error occurs during the group snapshot creation. Upon success,
                  this error field will be cleared.
                properties:
                  message:
                    description: 'message is a string detailing the encountered error
                      during snapshot creation if specified. NOTE: message may be
                      logged, and it should not contain sensitive information.'
                    type: string
                  time:
                    description: time is the timestamp when the error was encountered.
                    format: date-time
                    type: string
                type: object
              readyToUse:
                description: ReadyToUse indicates if all the individual snapshots
                  in the group are ready to be used to restore a group of volumes.
                  ReadyToUse becomes true when ReadyToUse of all individual snapshots
                  become true. If not specified, it means the readiness of a group
                  snapshot is unknown.
                type: boolean
              volumeSnapshotRefList:
                description: VolumeSnapshotRefList is the list of volume snapshot
                  references for this group snapshot. The maximum number of allowed
                  snapshots in the group is 100.
                items:
                  description: "ObjectReference contains enough information to let
                    you inspect or modify the referred object. --- New uses of this
                    type are discouraged because of difficulty describing its usage
                    when embedded in APIs. 1. Ignored fields.  It includes many fields
                    which are not generally honored.  For instance, ResourceVersion
                    and FieldPath are both very rarely valid in actual usage. 2. Invalid
                    usage help.  It is impossible to add specific help for individual
                    usage.  In most embedded usages, there are particular restrictions
                    like, \"must refer only to types A and B\" or \"UID not honored\"
                    or \"name must be restricted\". Those cannot be well described
                    when embedded. 3. Inconsistent validation.  Because the usages
                    are different, the validation rules are different by usage, which
                    makes it hard for users to predict what will happen. 4. The fields
                    are both imprecise and overly precise.  Kind is not a precise
                    mapping to a URL. This can produce ambiguity during interpretation
                    and require a REST mapping.  In most cases, the dependency is
                    on the group,resource tuple and the version of the actual struct
                    is irrelevant. 5. We cannot easily change it.  Because this type
                    is embedded in many locations, updates to this type will affect
                    numerous schemas.  Don't make new APIs embed an underspecified
                    API type they do not control. \n Instead of using this type, create
                    a locally provided and used type that is well-focused on your
                    reference. For example, ServiceReferences for admission registration:
                    https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533
                    ."
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: 'If referring to a piece of an object instead of
                        an entire object, this string should contain a valid JSON/Go
                        field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within
                        a pod, this would take on a value like: "spec.containers{name}"
                        (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]"
                        (container with index 2 in this pod). This syntax is chosen
                        only to have some well-defined way of referencing a part of
                        an object. TODO: this design is not final and this field is
                        subject to change in the future.'
                      type: string
                    kind:
                      description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                    resourceVersion:
                      description: 'Specific resourceVersion to which this reference
                        is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                      type: string
                    uid:
                      description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`
}

// GetVolumeGroupSnapshotContentsCRDv1 returns the VolumeSnapshotContents CRD Template
func (temp *Template) GetVolumeGroupSnapshotContentsCRDv1(kVersion string) string {
	return `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/971"
    controller-gen.kubebuilder.io/version: v0.12.0
  creationTimestamp: null
  name: volumegroupsnapshotcontents.groupsnapshot.storage.k8s.io
spec:
  group: groupsnapshot.storage.k8s.io
  names:
    kind: VolumeGroupSnapshotContent
    listKind: VolumeGroupSnapshotContentList
    plural: volumegroupsnapshotcontents
    shortNames:
    - vgsc
    - vgscs
    singular: volumegroupsnapshotcontent
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - description: Indicates if all the individual snapshots in the group are ready
        to be used to restore a group of volumes.
      jsonPath: .status.readyToUse
      name: ReadyToUse
      type: boolean
    - description: Determines whether this VolumeGroupSnapshotContent and its physical
        group snapshot on the underlying storage system should be deleted when its
        bound VolumeGroupSnapshot is deleted.
      jsonPath: .spec.deletionPolicy
      name: DeletionPolicy
      type: string
    - description: Name of the CSI driver used to create the physical group snapshot
        on the underlying storage system.
      jsonPath: .spec.driver
      name: Driver
      type: string
    - description: Name of the VolumeGroupSnapshotClass from which this group snapshot
        was (or will be) created.
      jsonPath: .spec.volumeGroupSnapshotClassName
      name: VolumeGroupSnapshotClass
      type: string
    - description: Namespace of the VolumeGroupSnapshot object to which this VolumeGroupSnapshotContent
        object is bound.
      jsonPath: .spec.volumeGroupSnapshotRef.namespace
      name: VolumeGroupSnapshotNamespace
      type: string
    - description: Name of the VolumeGroupSnapshot object to which this VolumeGroupSnapshotContent
        object is bound.
      jsonPath: .spec.volumeGroupSnapshotRef.name
      name: VolumeGroupSnapshot
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: VolumeGroupSnapshotContent represents the actual "on-disk" group
          snapshot object in the underlying storage system
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: Spec defines properties of a VolumeGroupSnapshotContent created
              by the underlying storage system. Required.
            properties:
              deletionPolicy:
                description: DeletionPolicy determines whether this VolumeGroupSnapshotContent
                  and the physical group snapshot on the underlying storage system
                  should be deleted when the bound VolumeGroupSnapshot is deleted.
                  Supported values are "Retain" and "Delete". "Retain" means that
                  the VolumeGroupSnapshotContent and its physical group snapshot on
                  underlying storage system are kept. "Delete" means that the VolumeGroupSnapshotContent
                  and its physical group snapshot on underlying storage system are
                  deleted. For dynamically provisioned group snapshots, this field
                  will automatically be filled in by the CSI snapshotter sidecar with
                  the "DeletionPolicy" field defined in the corresponding VolumeGroupSnapshotClass.
                  For pre-existing snapshots, users MUST specify this field when creating
                  the VolumeGroupSnapshotContent object. Required.
                enum:
                - Delete
                - Retain
                type: string
              driver:
                description: Driver is the name of the CSI driver used to create the
                  physical group snapshot on the underlying storage system. This MUST
                  be the same as the name returned by the CSI GetPluginName() call
                  for that driver. Required.
                type: string
              source:
                description: Source specifies whether the snapshot is (or should be)
                  dynamically provisioned or already exists, and just requires a Kubernetes
                  object representation. This field is immutable after creation. Required.
                properties:
                  groupSnapshotHandles:
                    description: GroupSnapshotHandles specifies the CSI "group_snapshot_id"
                      of a pre-existing group snapshot and a list of CSI "snapshot_id"
                      of pre-existing snapshots on the underlying storage system for
                      which a Kubernetes object representation was (or should be)
                      created. This field is immutable.
                    properties:
                      volumeGroupSnapshotHandle:
                        description: VolumeGroupSnapshotHandle specifies the CSI "group_snapshot_id"
                          of a pre-existing group snapshot on the underlying storage
                          system for which a Kubernetes object representation was
                          (or should be) created. This field is immutable. Required.
                        type: string
                      volumeSnapshotHandles:
                        description: VolumeSnapshotHandles is a list of CSI "snapshot_id"
                          of pre-existing snapshots on the underlying storage system
                          for which Kubernetes objects representation were (or should
                          be) created. This field is immutable. Required.
                        items:
                          type: string
                        type: array
                    required:
                    - volumeGroupSnapshotHandle
                    - volumeSnapshotHandles
                    type: object
                  volumeHandles:
                    description: VolumeHandles is a list of volume handles on the
                      backend to be snapshotted together. It is specified for dynamic
                      provisioning of the VolumeGroupSnapshot. This field is immutable.
                    items:
                      type: string
                    type: array
                type: object
                oneOf:
                - required: ["volumeHandles"]
                - required: ["groupSnapshotHandles"]
              volumeGroupSnapshotClassName:
                description: VolumeGroupSnapshotClassName is the name of the VolumeGroupSnapshotClass
                  from which this group snapshot was (or will be) created. Note that
                  after provisioning, the VolumeGroupSnapshotClass may be deleted
                  or recreated with different set of values, and as such, should not
                  be referenced post-snapshot creation. For dynamic provisioning,
                  this field must be set. This field may be unset for pre-provisioned
                  snapshots.
                type: string
              volumeGroupSnapshotRef:
                description: VolumeGroupSnapshotRef specifies the VolumeGroupSnapshot
                  object to which this VolumeGroupSnapshotContent object is bound.
                  VolumeGroupSnapshot.Spec.VolumeGroupSnapshotContentName field must
                  reference to this VolumeGroupSnapshotContent's name for the bidirectional
                  binding to be valid. For a pre-existing VolumeGroupSnapshotContent
                  object, name and namespace of the VolumeGroupSnapshot object MUST
                  be provided for binding to happen. This field is immutable after
                  creation. Required.
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  fieldPath:
                    description: 'If referring to a piece of an object instead of
                      an entire object, this string should contain a valid JSON/Go
                      field access statement, such as desiredState.manifest.containers[2].
                      For example, if the object reference is to a container within
                      a pod, this would take on a value like: "spec.containers{name}"
                      (where "name" refers to the name of the container that triggered
                      the event) or if no container name is specified "spec.containers[2]"
                      (container with index 2 in this pod). This syntax is chosen
                      only to have some well-defined way of referencing a part of
                      an object. TODO: this design is not final and this field is
                      subject to change in the future.'
                    type: string
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                  resourceVersion:
                    description: 'Specific resourceVersion to which this reference
                      is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                    type: string
                type: object
                x-kubernetes-map-type: atomic
            required:
            - deletionPolicy
            - driver
            - source
            - volumeGroupSnapshotRef
            type: object
          status:
            description: status represents the current information of a group snapshot.
            properties:
              creationTime:
                description: CreationTime is the timestamp when the point-in-time
                  group snapshot is taken by the underlying storage system. If not
                  specified, it indicates the creation time is unknown. If not specified,
                  it means the readiness of a group snapshot is unknown. The format
                  of this field is a Unix nanoseconds time encoded as an int64. On
                  Unix, the command date +%s%N returns the current time in nanoseconds
                  since 1970-01-01 00:00:00 UTC.
                format: int64
                type: integer
              error:
                description: Error is the last observed error during group snapshot
                  creation, if any. Upon success after retry, this error field will
                  be cleared.
                properties:
                  message:
                    description: 'message is a string detailing the encountered error
                      during snapshot creation if specified. NOTE: message may be
                      logged, and it should not contain sensitive information.'
                    type: string
                  time:
                    description: time is the timestamp when the error was encountered.
                    format: date-time
                    type: string
                type: object
              readyToUse:
                description: ReadyToUse indicates if all the individual snapshots
                  in the group are ready to be used to restore a group of volumes.
                  ReadyToUse becomes true when ReadyToUse of all individual snapshots
                  become true.
                type: boolean
              volumeGroupSnapshotHandle:
                description: VolumeGroupSnapshotHandle is a unique id returned by
                  the CSI driver to identify the VolumeGroupSnapshot on the storage
                  system. If a storage system does not provide such an id, the CSI
                  driver can choose to return the VolumeGroupSnapshot name.
                type: string
              volumeSnapshotContentRefList:
                description: VolumeSnapshotContentRefList is the list of volume snapshot
                  content references for this group snapshot. The maximum number of
                  allowed snapshots in the group is 100.
                items:
                  description: "ObjectReference contains enough information to let
                    you inspect or modify the referred object. --- New uses of this
                    type are discouraged because of difficulty describing its usage
                    when embedded in APIs. 1. Ignored fields.  It includes many fields
                    which are not generally honored.  For instance, ResourceVersion
                    and FieldPath are both very rarely valid in actual usage. 2. Invalid
                    usage help.  It is impossible to add specific help for individual
                    usage.  In most embedded usages, there are particular restrictions
                    like, \"must refer only to types A and B\" or \"UID not honored\"
                    or \"name must be restricted\". Those cannot be well described
                    when embedded. 3. Inconsistent validation.  Because the usages
                    are different, the validation rules are different by usage, which
                    makes it hard for users to predict what will happen. 4. The fields
                    are both imprecise and overly precise.  Kind is not a precise
                    mapping to a URL. This can produce ambiguity during interpretation
                    and require a REST mapping.  In most cases, the dependency is
                    on the group,resource tuple and the version of the actual struct
                    is irrelevant. 5. We cannot easily change it.  Because this type
                    is embedded in many locations, updates to this type will affect
                    numerous schemas.  Don't make new APIs embed an underspecified
                    API type they do not control. \n Instead of using this type, create
                    a locally provided and used type that is well-focused on your
                    reference. For example, ServiceReferences for admission registration:
                    https://github.com/kubernetes/api/blob/release-1.17/admissionregistration/v1/types.go#L533
                    ."
                  properties:
                    apiVersion:
                      description: API version of the referent.
                      type: string
                    fieldPath:
                      description: 'If referring to a piece of an object instead of
                        an entire object, this string should contain a valid JSON/Go
                        field access statement, such as desiredState.manifest.containers[2].
                        For example, if the object reference is to a container within
                        a pod, this would take on a value like: "spec.containers{name}"
                        (where "name" refers to the name of the container that triggered
                        the event) or if no container name is specified "spec.containers[2]"
                        (container with index 2 in this pod). This syntax is chosen
                        only to have some well-defined way of referencing a part of
                        an object. TODO: this design is not final and this field is
                        subject to change in the future.'
                      type: string
                    kind:
                      description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                    resourceVersion:
                      description: 'Specific resourceVersion to which this reference
                        is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                      type: string
                    uid:
                      description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                      type: string
                  type: object
                  x-kubernetes-map-type: atomic
                type: array
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`
}

// GetVolumeGroupSnapshotClassesCRDv1 returns the VolumeSnapshotClasses CRD Template
func (temp *Template) GetVolumeGroupSnapshotClassesCRDv1(kVersion string) string {
	return `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    api-approved.kubernetes.io: "https://github.com/kubernetes-csi/external-snapshotter/pull/814"
    controller-gen.kubebuilder.io/version: v0.12.0
  creationTimestamp: null
  name: volumegroupsnapshotclasses.groupsnapshot.storage.k8s.io
spec:
  group: groupsnapshot.storage.k8s.io
  names:
    kind: VolumeGroupSnapshotClass
    listKind: VolumeGroupSnapshotClassList
    plural: volumegroupsnapshotclasses
    shortNames:
    - vgsclass
    - vgsclasses
    singular: volumegroupsnapshotclass
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .driver
      name: Driver
      type: string
    - description: Determines whether a VolumeGroupSnapshotContent created through
        the VolumeGroupSnapshotClass should be deleted when its bound VolumeGroupSnapshot
        is deleted.
      jsonPath: .deletionPolicy
      name: DeletionPolicy
      type: string
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1alpha1
    schema:
      openAPIV3Schema:
        description: VolumeGroupSnapshotClass specifies parameters that a underlying
          storage system uses when creating a volume group snapshot. A specific VolumeGroupSnapshotClass
          is used by specifying its name in a VolumeGroupSnapshot object. VolumeGroupSnapshotClasses
          are non-namespaced.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          deletionPolicy:
            description: DeletionPolicy determines whether a VolumeGroupSnapshotContent
              created through the VolumeGroupSnapshotClass should be deleted when
              its bound VolumeGroupSnapshot is deleted. Supported values are "Retain"
              and "Delete". "Retain" means that the VolumeGroupSnapshotContent and
              its physical group snapshot on underlying storage system are kept. "Delete"
              means that the VolumeGroupSnapshotContent and its physical group snapshot
              on underlying storage system are deleted. Required.
            enum:
            - Delete
            - Retain
            type: string
          driver:
            description: Driver is the name of the storage driver expected to handle
              this VolumeGroupSnapshotClass. Required.
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          parameters:
            additionalProperties:
              type: string
            description: Parameters is a key-value map with storage driver specific
              parameters for creating group snapshots. These values are opaque to
              Kubernetes and are passed directly to the driver.
            type: object
        required:
        - deletionPolicy
        - driver
        type: object
    served: true
    storage: true
    subresources: {}
`
}
