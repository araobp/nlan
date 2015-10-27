# Dockerfile to create NLAN agent image
#
# Original image "router" have already openvswitch
# preinstalled.

FROM router
ENV ETCD_ADDRESS "http://172.17.42.1:2379"
ENV NO_PROXY 172.17.42.1,localhost
ADD agent /root/bin/agent
CMD service ssh start
CMD service openvswitch-switch start

