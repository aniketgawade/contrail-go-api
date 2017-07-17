
package types

import (
        "reflect"

        "github.com/aniketgawade/contrail-go-api"
)

var (
        TypeMap = map[string]reflect.Type {
		"domain": reflect.TypeOf(Domain{}),
		"global-vrouter-config": reflect.TypeOf(GlobalVrouterConfig{}),
		"address-group": reflect.TypeOf(AddressGroup{}),
		"instance-ip": reflect.TypeOf(InstanceIp{}),
		"floating-ip-pool": reflect.TypeOf(FloatingIpPool{}),
		"loadbalancer-pool": reflect.TypeOf(LoadbalancerPool{}),
		"virtual-DNS-record": reflect.TypeOf(VirtualDnsRecord{}),
		"route-target": reflect.TypeOf(RouteTarget{}),
		"bridge-domain": reflect.TypeOf(BridgeDomain{}),
		"discovery-service-assignment": reflect.TypeOf(DiscoveryServiceAssignment{}),
		"floating-ip": reflect.TypeOf(FloatingIp{}),
		"alias-ip": reflect.TypeOf(AliasIp{}),
		"network-policy": reflect.TypeOf(NetworkPolicy{}),
		"physical-router": reflect.TypeOf(PhysicalRouter{}),
		"bgp-router": reflect.TypeOf(BgpRouter{}),
		"api-access-list": reflect.TypeOf(ApiAccessList{}),
		"virtual-router": reflect.TypeOf(VirtualRouter{}),
		"config-root": reflect.TypeOf(ConfigRoot{}),
		"subnet": reflect.TypeOf(Subnet{}),
		"alarm": reflect.TypeOf(Alarm{}),
		"policy-management": reflect.TypeOf(PolicyManagement{}),
		"global-system-config": reflect.TypeOf(GlobalSystemConfig{}),
		"service-appliance": reflect.TypeOf(ServiceAppliance{}),
		"routing-policy": reflect.TypeOf(RoutingPolicy{}),
		"namespace": reflect.TypeOf(Namespace{}),
		"forwarding-class": reflect.TypeOf(ForwardingClass{}),
		"service-instance": reflect.TypeOf(ServiceInstance{}),
		"service-group": reflect.TypeOf(ServiceGroup{}),
		"route-table": reflect.TypeOf(RouteTable{}),
		"application-policy-set": reflect.TypeOf(ApplicationPolicySet{}),
		"physical-interface": reflect.TypeOf(PhysicalInterface{}),
		"access-control-list": reflect.TypeOf(AccessControlList{}),
		"bgp-as-a-service": reflect.TypeOf(BgpAsAService{}),
		"port-tuple": reflect.TypeOf(PortTuple{}),
		"security-logging-object": reflect.TypeOf(SecurityLoggingObject{}),
		"analytics-node": reflect.TypeOf(AnalyticsNode{}),
		"virtual-DNS": reflect.TypeOf(VirtualDns{}),
		"customer-attachment": reflect.TypeOf(CustomerAttachment{}),
		"service-appliance-set": reflect.TypeOf(ServiceApplianceSet{}),
		"config-node": reflect.TypeOf(ConfigNode{}),
		"qos-queue": reflect.TypeOf(QosQueue{}),
		"virtual-machine": reflect.TypeOf(VirtualMachine{}),
		"interface-route-table": reflect.TypeOf(InterfaceRouteTable{}),
		"service-template": reflect.TypeOf(ServiceTemplate{}),
		"dsa-rule": reflect.TypeOf(DsaRule{}),
		"firewall-policy": reflect.TypeOf(FirewallPolicy{}),
		"firewall-rule": reflect.TypeOf(FirewallRule{}),
		"bgpvpn": reflect.TypeOf(Bgpvpn{}),
		"global-qos-config": reflect.TypeOf(GlobalQosConfig{}),
		"virtual-ip": reflect.TypeOf(VirtualIp{}),
		"loadbalancer-member": reflect.TypeOf(LoadbalancerMember{}),
		"security-group": reflect.TypeOf(SecurityGroup{}),
		"service-health-check": reflect.TypeOf(ServiceHealthCheck{}),
		"qos-config": reflect.TypeOf(QosConfig{}),
		"provider-attachment": reflect.TypeOf(ProviderAttachment{}),
		"virtual-machine-interface": reflect.TypeOf(VirtualMachineInterface{}),
		"loadbalancer-healthmonitor": reflect.TypeOf(LoadbalancerHealthmonitor{}),
		"loadbalancer-listener": reflect.TypeOf(LoadbalancerListener{}),
		"virtual-network": reflect.TypeOf(VirtualNetwork{}),
		"project": reflect.TypeOf(Project{}),
		"logical-interface": reflect.TypeOf(LogicalInterface{}),
		"tag": reflect.TypeOf(Tag{}),
		"loadbalancer": reflect.TypeOf(Loadbalancer{}),
		"database-node": reflect.TypeOf(DatabaseNode{}),
		"routing-instance": reflect.TypeOf(RoutingInstance{}),
		"alias-ip-pool": reflect.TypeOf(AliasIpPool{}),
		"network-ipam": reflect.TypeOf(NetworkIpam{}),
		"route-aggregate": reflect.TypeOf(RouteAggregate{}),
		"logical-router": reflect.TypeOf(LogicalRouter{}),

        }
)

func init() {
        contrail.RegisterTypeMap(TypeMap)
}
