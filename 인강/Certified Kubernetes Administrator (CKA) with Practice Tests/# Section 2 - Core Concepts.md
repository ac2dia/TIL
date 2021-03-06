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

## 26. Practice Test - Pods

- create pods
- get pods / describe pods
- kubectl apply -f nginx-pod.yaml
- delete pods

## 28. Recap - ReplicaSets

- replica controller
- pods crush!! ㅠㅠ => so more than one pods at the same time.
- Load Balancing & Scaling

```yaml
# replicaset-definition.yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: myapp-replicaset
  labels:
    app: myapp
    type: front-end
spec:
  template:
    metadata:
      name: myapp-pod
      labels:
        app: myapp
        type: front-end
    spec:
      containers:
        - name: nginx-container
          image: nginx
  replicas: 3
  selector:
    matchLabels:
      tpye: front-end
```

```shell
$ kubectl apply -f replicaset-definition.yaml
$ kubectl get replicaset (rs)
$ kubectl get pods
```

- Labels and Selectors (라벨링 작업)
- scale (replica 갯수 조절 가능))

## 29. Practice Test - ReplicaSets

- create replicaset
- get replicaset / describe replicaset
- kubectl apply -f new-replica-set.yaml
- delete replicaset
- scale --replicas=n rs new-replica-set (scale-up / scale-down)

## 31. Deployments

-

## 33. Practice Test - Deployments

## 35. Services

- Services Types
  - NodePort, ClusterIP, Loadbalancer
  - NodePort = Access to Node IP and Port by Direct!

## 36. Services - Cluster IP

## 37. Services - Loadbalancer

## 38. Practice Test - Services

- service, deployment ...

## 40. Namespaces

- Isolate!
- kube-system / default / kube-public
- DNS
- Resource Quota

## 41. Practice Test - Namespaces

## 43. Imperative vs Declarative

- Infrastructure as Code

## 44. Certification Tips - Imperative Commands with Kubectl

- --dry-run: By default as soon as the command is run, the resource will be created. If you simply want to test your command , use the
- --dry-run=client option. This will not create the resource, instead, tell you whether the resource can be created and if your command is right.
- -o yaml: This will output the resource definition in YAML format on screen.

## 45. Practice Test - Imperative Commands

## 47. Kubectl Apply Command

- Create Objects
  - kubectl apply -f nginx.yaml
- Update Objects
  - kubectl apply -f nginx.yaml

```yaml
# nginx.yaml

apiVersion: v1
kind: Pod

metadata:
  name: myapp-pod
  labels:
    app: myapp
    type: front-end-service
sepc:
  containers:
    - name: nginx-container
      image: nginx:1.18
```

$ kubectl apply -f nginx.yaml
