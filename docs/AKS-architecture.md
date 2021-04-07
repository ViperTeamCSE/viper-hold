# Introduction

The API component(s) will need to be run in the cloud.  This can be done in multiple ways.  Two options are Azure Functions or as a container in AKS.  The prelimary decision for this engagement is to use AKS.

The AKS design must take into account the following:
* Secure by default
* Massive global scale in future
* Appropriate update / upgrade strategy

There are [best practice guidelines and recommendations](https://docs.microsoft.com/en-us/azure/cloud-adoption-framework/innovate/kubernetes/) for cluster design and security that should be reviewed and followed.


## Security

AKS security involves many areas which includes (but not limited to):
* Control access to clusters using group membership. Configure Kubernetes role-based access control (Kubernetes RBAC) to limit access to cluster resources based on user identity or group membership.
* Create a secrets management policy. Securely deploy and manage sensitive information, such as passwords and certificates, using secrets management in Kubernetes.
* Secure intra-pod network traffic with network policies. Apply the principle of least privilege to control network traffic flow between pods in the cluster.
* Restrict access to the API server using authorized IPs. Improve cluster security and minimize attack surface by limiting access to the API server to a limited set of IP address ranges.
> A private AKS cluster can also be deployed to further secure access to the API server to private (internal) IP addresses.
* Restrict cluster egress traffic. Learn what ports and addresses to allow if you restrict egress traffic for the cluster. You can use Azure Firewall or a third-party firewall appliance to secure your egress traffic and define these required ports and addresses.
* Secure traffic with Web Application Firewall (WAF). Use Azure Application Gateway as an ingress controller for Kubernetes clusters.
* Apply security and kernel updates to worker nodes. Understand the AKS node update experience. To protect your clusters, security updates are automatically applied to Linux nodes in AKS. These updates include OS security fixes or kernel updates. Some of these updates require a node reboot to complete the process.
* Configure a container and cluster scanning solution. Scan containers pushed into Azure Container Registry and gain deeper visibility to your cluster nodes, cloud traffic, and security controls.
* Limit access to cluster configuration file (kubeconfig)
* Implement cluster certificate rotation policy.

## Scale

To acheive scale and availability, some key items need to be considered:

* Automatically scale a cluster to meet application demands. To keep up with application demands, you may need to adjust the number of nodes that run your workloads automatically using the cluster autoscaler.
* Plan for business continuity and disaster recovery. Plan for multiregion deployment, create a storage migration plan, and enable geo-replication for container images.
* Configure monitoring and troubleshooting at scale. Set up alerting and monitoring for applications in Kubernetes

## Update / Upgrade

Kuberenetes is updated on a regular basis.  The AKS cluster needs to be kept up to date to stay secure and within support guidelines.  The worker nodes are set to auto update by default but any required reboots will have to be handled separately.  Updates also need to be done in a manner that does not affect the availability of the cluster as a whole or the applications running within the cluster.


# Deployment

The deployment of the AKS cluster will be done using Terraform templates / scripts.  A CI / CD pipeline will be used to automatically deploy the AKS cluster and supporting components.


# Proposed Architecture

The recommendation for AKS configuration:

* AAD Integration (AKS Managed)
* System assigned managed identity
* RBAC enabled
* Encryption at-rest with platform managed key
* Azure policies
* Advanced Networking
* VMSS
* Availability Zones enabled
* Private cluster
* Azure monitor integration enabled
* AAD pod identity
* Key Vault integration via CSI driver
* 3 nodes (to start)
* Cluster autoscaling

The following items will need to be determined prior to the start of the engagement:

* Node VM size
* OS disk size
* Kubernetes version

# **APPENDIX**

The containers that will be run on the devices in the car will need to be tested as well as part of end to end testing.  These containers will need to be run in the cloud.  Two viable options are AKS or VMs (or VMSS).  Using AKS would not be the same as running the containers on the device if the architecture for the device is that one of the containers will be the management component for managing the other containers on the device.  For this reason, the recommendation is to use VMs or VMSS to host the testing of containers to simulate the device environment.
