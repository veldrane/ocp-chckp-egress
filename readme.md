### chckpapi service

#### Description

Simple api provides checkpoint json rule format based on the annotation
in Openshift 4.x (3.x will be probably work as well). Kubernetes are not 
supported because 

THe whole solution and description of the checkpoint side is here: 
https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk167210.



#### Instalation

```make install```

then make your container image as you want (Dockerfile is not included yet).

Service can be started outside of the kubernetes (k8s cfg must be present) or inside of the cluster. The security role
ensures access to the namespace and netnamespaces object must be assigned to the external or service account.


#### Configuration

For right publishing the rules two annotation have to be added to the namespace object:

checkpoint.com/egress-rules
    - this annotation publishes rules for egress ip 

checkpoint.com/ingress-rules
    - this annotation publishrs rules for egress ip

```
example:
apiVersion: v1
kind: Namespace
metadata:
  annotations:
    checkpoint.com/egress-rules: '["default-https-443", "extended-https-8443", "pods-to-redis"]'
    checkpoint.com/ingress-rules: '[{"ng-ingress-http": ["default-https-443", "extended-https-8443"]},
      {"default": ["pods-to-redis", "extended-https-8443"]}]'
    openshift.io/description: some project
    openshift.io/display-name: some project
  creationTimestamp: null
  labels:
    kubernetes.io/metadata.name: some-ns
  name: some-ns
spec:
  finalizers:
  - kubernetes
```

Rulesets must be present on the checkpoint side.

##### Egress

Ip adress is get from netnamespace object, other solution is not supported

##### Ingress

Ip address is get from external ip of the svc, usualy binded with keepalived ip adresses. Other solution is not supported.

#### How it works

After start chckpapi ask kubernetes api for all namespace object with annotation details and netnamespace objects,
join all information together and present that in the json format supported by checkpoint part. This part is done by background process
and it keeps repeating each 5 mins. Json is published from internal structers so its not possible to DoS k8s api by huge calling of 
this svc. 


#### Is the chckpapi production ready ?

Ehm probably no :). This tool was part of PoC project of integration between ocp and checkpoint infrastructure, but lots of things is
missing. Code is not clean, helm chart and Dockerfile is missing as well (it's somewhere, but i am leazy to look for it ). Also definitiot
of the svc account and better role for running this service is missing (some basic role object - probably sufficient is part of role directory).


#### Dev notes

for api definition has been used goa fromework and design file can be found in design subdir. Directory local/checkpoint/
contains main code for interacting with k8s api. Swagger client and definition is part of this project as well.

https://github.com/goadesign/goa
