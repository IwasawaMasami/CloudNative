

**外部客户端访问==> iptable如何处理:**

```
#kubectl get svc -A -o wide | grep nginx-basic

default      nginx-basic                ClusterIP   10.97.35.60   <none>        80/TCP                80m   app=nginx
```



```
10.97.35.60:80 ->
192.168.166.157:80
192.168.166.166:80
192.168.166.162:80
```

```
#  kubectl get pod -A -o wide | grep nginx-deployment
default   nginx-deployment-58c445c754-5qdqm    1/1  Running  0     91m   192.168.166.157   hecs-364953   <none>    <none>
default   nginx-deployment-58c445c754-6lpqd     1/1  Running  0     91m   192.168.166.166  hecs-364953   <none>    <none>
default   nginx-deployment-58c445c754-xqrtb     1/1  Running  0     91m   192.168.166.162   hecs-364953  <none>    <none>

```

所有进来的包都要去kubeservice这条链去处理

```
-A PREROUTING -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
```

离开的包都要去kubeservice这条链处理

```
-A OUTPUT -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
```

访问目的地址为node port ip，协议为TCP，目标端口为30207就交给KUBE-SVC-WWRFY3PZ7W3FGMQW这条链去处理数据包

```
-A KUBE-NODEPORTS -p tcp -m comment --comment "default/nginx-basic:http" -m tcp --dport 30207 -j KUBE-SVC-WWRFY3PZ7W3FGMQW
```

访问目的地址为10.97.35.60/32，协议为TCP，目标端口为80的数据包就交给KUBE-SVC-WWRFY3PZ7W3FGMQW这条链去处理

```
-A KUBE-SERVICES -d 10.97.35.60/32 -p tcp -m comment --comment "default/nginx-basic:http cluster IP" -m tcp --dport 80 -j KUBE-SVC-WWRFY3PZ7W3FGMQW
```

**交给KUBE-SVC-WWRFY3PZ7W3FGMQW有33%概率去处理数据包，并交给KUBE-SEP-2422Y4UT42X3LJXA转到最终POD的地址**

```
-A KUBE-SVC-WWRFY3PZ7W3FGMQW -m comment --comment "default/nginx-basic:http" -m statistic --mode random --probability 0.33333333349 -j KUBE-SEP-2422Y4UT42X3LJXA
-A KUBE-SEP-2422Y4UT42X3LJXA -p tcp -m comment --comment "default/nginx-basic:http" -m tcp -j DNAT --to-destination 192.168.166.157:80
```

交给KUBE-SVC-WWRFY3PZ7W3FGMQW有50%概率去处理数据包，并交给KUBE-SEP-STI77FUROETI63JE转到最终POD的地址

```
-A KUBE-SVC-WWRFY3PZ7W3FGMQW -m comment --comment "default/nginx-basic:http" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-STI77FUROETI63JE
-A KUBE-SEP-STI77FUROETI63JE -p tcp -m comment --comment "default/nginx-basic:http" -m tcp -j DNAT --to-destination 192.168.166.162:80
```

**交给KUBE-SVC-WWRFY3PZ7W3FGMQW有100%概率去处理数据包并交给KUBE-SEP-FGYAQDFXQTLSKPFT转到最终POD的地址**

```
-A KUBE-SVC-WWRFY3PZ7W3FGMQW -m comment --comment "default/nginx-basic:http" -j KUBE-SEP-FGYAQDFXQTLSKPFT
-A KUBE-SEP-FGYAQDFXQTLSKPFT -p tcp -m comment --comment "default/nginx-basic:http" -m tcp -j DNAT --to-destination 192.168.166.166:80
```

