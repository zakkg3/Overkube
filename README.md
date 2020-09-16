# Overkube
Automated k8s clusters on demand


# Workflow.

1- Apply overkube CRD

```
Kind: OverKube
apiVersion: ovkb.io/v1
metadata:
  name: overcluster
  namespace: overkube
cluster-namespace: tenant1
workers:
  count: 3
  memory: 4Gb
  cores: 4
master-replicas: 2
apiServer:
  certSANs:
    - "129.132.98.170"
controlPlaneEndpoint: "129.132.98.170:6443"
networking:
  dnsDomain: cluster.local
  podSubnet: 10.206.176.0/20
  serviceSubnet: 10.206.146.0/24
```

2- The operator will start the new k8s overcluster control plane on the namespace "cluster-namespace". (https://github.com/kvaps/kubernetes-in-kubernetes)

At the same time will spawn N (workers.count) Kubevirt vm's using a jinja template wich reads from a configfile (so we can change workers size and way to deploy them)

3- Once the control-plane is in place, and the vms are up and running the operator will run a job (kopf.adopt) this job will use @asteven scripts to bootstrap the cluster using the kubeadm clusterconfiguration provided in the CRD

4- Everithing is up so the operator updates the crd, or create a secret with this info:


```
status:
 certificate-authority-data: 
 client-certificate-data:
 client-key-data:
```

Now, we are ready configure kubectl and start using the Overkube.



