# Section 2: Core Concepts

## 19. Kube Proxy

- Pod Networking
- process that run each node in the k8s cluster
- using iptables ...

## 20. Recap - Pods

- k8s! deploy our application in the form of containers ...
- container are encapsulated into a k8s object known as pods
- smallest object in k8s
- single container in pods / multi containers in pods

## 21. PODs with YAML

- develop pods using YAML

```yaml
# pod-definition.yml

apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
      app: myapp
      type: front-end
spec:
  containers: <- List/Array
    - name: nginx-container
      image: nginx
    - name: database-container
      image: mariadb

```

```bash
$ kubectl create -f pod-definition.yml

$ kubectl get pods
```
