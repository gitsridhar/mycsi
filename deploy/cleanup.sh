oc delete daemonset my-csi-plugin
oc delete statefulset my-csi-attacher-plugin
oc delete statefulset my-csi-provisioner-plugin
oc delete statefulset my-csi-resizer-plugin

oc delete service my-csi-attacher-plugin
oc delete service my-csi-provisioner-plugin

oc delete clusterrolebinding my-csi-node-role
oc delete clusterrolebinding my-csi-provisioner-role
oc delete clusterrolebinding my-csi-attacher-role
oc delete clusterrolebinding my-csi-resizer-role

oc delete clusterrole my-csi-node
oc delete clusterrole my-csi-provisioner
oc delete clusterrole my-csi-attacher
oc delete clusterrole my-csi-resizer

oc delete serviceaccount my-csi-node
oc delete serviceaccount my-csi-provisioner
oc delete serviceaccount my-csi-attacher
oc delete serviceaccount my-csi-resizer

oc delete storageclass my-csi-volume-default
oc delete configmap my-config
oc delete csidriver my-csi

oc delete securitycontextconstraints mycsiaccess

oc delete csinode okd4-master1
oc delete csinode okd4-master2
oc delete csinode okd4-master3
oc delete csinode okd4-worker1
oc delete csinode okd4-worker2
oc delete csinode okd4-worker3
