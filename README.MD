kubectl create -f my-first-pod.yaml
kubectl describe pod my-pod 
kubectl port-forward my-pod 6379:6379
kubectl expose pod my-pod --type=NodePort --name=my-pod-service
minikube service my-pod-service --url

---

kubectl apply -f ingress.deployment.yaml 
kubectl get pods
kubectl delete deployment --all
kubectl exec gin-pod-77f6586747-bmxh9 env

kubectl get replicasets
kubectl describe replicasets

kubectl apply -f ingress.service.yaml
kubectl get service
kubectl describe svc

kubectl apply -f ingress.yaml
kubectl get ingress

kubectl run -i --tty alpine --image=alpine --restart=Never -- sh
apk add --no-cache curl
curl http://10.97.6.77:8080
nslookup redis-service.default.svc.cluster.local
cat /etc/resolv.conf

kubectl exec gin-pod-77f6586747-5bmfc env 
kubectl exec --stdin --tty gin-pod-77f6586747-c6ngp -- /bin/bash
kubectl exec --stdin --tty redis-pod-856b7f4f4f-l8ztb -- /bin/bash
apt install curl