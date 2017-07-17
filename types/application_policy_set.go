//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/aniketgawade/contrail-go-api"
)

const (
	application_policy_set_is_global uint64 = 1 << iota
	application_policy_set_id_perms
	application_policy_set_perms2
	application_policy_set_annotations
	application_policy_set_display_name
	application_policy_set_firewall_policy_refs
	application_policy_set_global_vrouter_config_refs
	application_policy_set_tag_refs
)

type ApplicationPolicySet struct {
        contrail.ObjectBase
	is_global bool
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	firewall_policy_refs contrail.ReferenceList
	global_vrouter_config_refs contrail.ReferenceList
	tag_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ApplicationPolicySet) GetType() string {
        return "application-policy-set"
}

func (obj *ApplicationPolicySet) GetDefaultParent() []string {
        name := []string{"default-policy-management"}
        return name
}

func (obj *ApplicationPolicySet) GetDefaultParentType() string {
        return "policy-management"
}

func (obj *ApplicationPolicySet) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ApplicationPolicySet) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ApplicationPolicySet) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ApplicationPolicySet) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ApplicationPolicySet) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ApplicationPolicySet) GetIsGlobal() bool {
        return obj.is_global
}

func (obj *ApplicationPolicySet) SetIsGlobal(value bool) {
        obj.is_global = value
        obj.modified |= application_policy_set_is_global
}

func (obj *ApplicationPolicySet) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ApplicationPolicySet) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= application_policy_set_id_perms
}

func (obj *ApplicationPolicySet) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *ApplicationPolicySet) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= application_policy_set_perms2
}

func (obj *ApplicationPolicySet) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *ApplicationPolicySet) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified |= application_policy_set_annotations
}

func (obj *ApplicationPolicySet) GetDisplayName() string {
        return obj.display_name
}

func (obj *ApplicationPolicySet) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= application_policy_set_display_name
}

func (obj *ApplicationPolicySet) readFirewallPolicyRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & application_policy_set_firewall_policy_refs == 0) {
                err := obj.GetField(obj, "firewall_policy_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ApplicationPolicySet) GetFirewallPolicyRefs() (
        contrail.ReferenceList, error) {
        err := obj.readFirewallPolicyRefs()
        if err != nil {
                return nil, err
        }
        return obj.firewall_policy_refs, nil
}

func (obj *ApplicationPolicySet) AddFirewallPolicy(
        rhs *FirewallPolicy, data FirewallSequence) error {
        err := obj.readFirewallPolicyRefs()
        if err != nil {
                return err
        }

        if obj.modified & application_policy_set_firewall_policy_refs == 0 {
                obj.storeReferenceBase("firewall-policy", obj.firewall_policy_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
        obj.firewall_policy_refs = append(obj.firewall_policy_refs, ref)
        obj.modified |= application_policy_set_firewall_policy_refs
        return nil
}

func (obj *ApplicationPolicySet) DeleteFirewallPolicy(uuid string) error {
        err := obj.readFirewallPolicyRefs()
        if err != nil {
                return err
        }

        if obj.modified & application_policy_set_firewall_policy_refs == 0 {
                obj.storeReferenceBase("firewall-policy", obj.firewall_policy_refs)
        }

        for i, ref := range obj.firewall_policy_refs {
                if ref.Uuid == uuid {
                        obj.firewall_policy_refs = append(
                                obj.firewall_policy_refs[:i],
                                obj.firewall_policy_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= application_policy_set_firewall_policy_refs
        return nil
}

func (obj *ApplicationPolicySet) ClearFirewallPolicy() {
        if (obj.valid & application_policy_set_firewall_policy_refs != 0) &&
           (obj.modified & application_policy_set_firewall_policy_refs == 0) {
                obj.storeReferenceBase("firewall-policy", obj.firewall_policy_refs)
        }
        obj.firewall_policy_refs = make([]contrail.Reference, 0)
        obj.valid |= application_policy_set_firewall_policy_refs
        obj.modified |= application_policy_set_firewall_policy_refs
}

func (obj *ApplicationPolicySet) SetFirewallPolicyList(
        refList []contrail.ReferencePair) {
        obj.ClearFirewallPolicy()
        obj.firewall_policy_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.firewall_policy_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *ApplicationPolicySet) readGlobalVrouterConfigRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & application_policy_set_global_vrouter_config_refs == 0) {
                err := obj.GetField(obj, "global_vrouter_config_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ApplicationPolicySet) GetGlobalVrouterConfigRefs() (
        contrail.ReferenceList, error) {
        err := obj.readGlobalVrouterConfigRefs()
        if err != nil {
                return nil, err
        }
        return obj.global_vrouter_config_refs, nil
}

func (obj *ApplicationPolicySet) AddGlobalVrouterConfig(
        rhs *GlobalVrouterConfig) error {
        err := obj.readGlobalVrouterConfigRefs()
        if err != nil {
                return err
        }

        if obj.modified & application_policy_set_global_vrouter_config_refs == 0 {
                obj.storeReferenceBase("global-vrouter-config", obj.global_vrouter_config_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.global_vrouter_config_refs = append(obj.global_vrouter_config_refs, ref)
        obj.modified |= application_policy_set_global_vrouter_config_refs
        return nil
}

func (obj *ApplicationPolicySet) DeleteGlobalVrouterConfig(uuid string) error {
        err := obj.readGlobalVrouterConfigRefs()
        if err != nil {
                return err
        }

        if obj.modified & application_policy_set_global_vrouter_config_refs == 0 {
                obj.storeReferenceBase("global-vrouter-config", obj.global_vrouter_config_refs)
        }

        for i, ref := range obj.global_vrouter_config_refs {
                if ref.Uuid == uuid {
                        obj.global_vrouter_config_refs = append(
                                obj.global_vrouter_config_refs[:i],
                                obj.global_vrouter_config_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= application_policy_set_global_vrouter_config_refs
        return nil
}

func (obj *ApplicationPolicySet) ClearGlobalVrouterConfig() {
        if (obj.valid & application_policy_set_global_vrouter_config_refs != 0) &&
           (obj.modified & application_policy_set_global_vrouter_config_refs == 0) {
                obj.storeReferenceBase("global-vrouter-config", obj.global_vrouter_config_refs)
        }
        obj.global_vrouter_config_refs = make([]contrail.Reference, 0)
        obj.valid |= application_policy_set_global_vrouter_config_refs
        obj.modified |= application_policy_set_global_vrouter_config_refs
}

func (obj *ApplicationPolicySet) SetGlobalVrouterConfigList(
        refList []contrail.ReferencePair) {
        obj.ClearGlobalVrouterConfig()
        obj.global_vrouter_config_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.global_vrouter_config_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *ApplicationPolicySet) readTagRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & application_policy_set_tag_refs == 0) {
                err := obj.GetField(obj, "tag_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ApplicationPolicySet) GetTagRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTagRefs()
        if err != nil {
                return nil, err
        }
        return obj.tag_refs, nil
}

func (obj *ApplicationPolicySet) AddTag(
        rhs *Tag) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if obj.modified & application_policy_set_tag_refs == 0 {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.tag_refs = append(obj.tag_refs, ref)
        obj.modified |= application_policy_set_tag_refs
        return nil
}

func (obj *ApplicationPolicySet) DeleteTag(uuid string) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if obj.modified & application_policy_set_tag_refs == 0 {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        for i, ref := range obj.tag_refs {
                if ref.Uuid == uuid {
                        obj.tag_refs = append(
                                obj.tag_refs[:i],
                                obj.tag_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= application_policy_set_tag_refs
        return nil
}

func (obj *ApplicationPolicySet) ClearTag() {
        if (obj.valid & application_policy_set_tag_refs != 0) &&
           (obj.modified & application_policy_set_tag_refs == 0) {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }
        obj.tag_refs = make([]contrail.Reference, 0)
        obj.valid |= application_policy_set_tag_refs
        obj.modified |= application_policy_set_tag_refs
}

func (obj *ApplicationPolicySet) SetTagList(
        refList []contrail.ReferencePair) {
        obj.ClearTag()
        obj.tag_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.tag_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *ApplicationPolicySet) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & application_policy_set_is_global != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.is_global)
                if err != nil {
                        return nil, err
                }
                msg["is_global"] = &value
        }

        if obj.modified & application_policy_set_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & application_policy_set_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & application_policy_set_annotations != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified & application_policy_set_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.firewall_policy_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.firewall_policy_refs)
                if err != nil {
                        return nil, err
                }
                msg["firewall_policy_refs"] = &value
        }

        if len(obj.global_vrouter_config_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.global_vrouter_config_refs)
                if err != nil {
                        return nil, err
                }
                msg["global_vrouter_config_refs"] = &value
        }

        if len(obj.tag_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.tag_refs)
                if err != nil {
                        return nil, err
                }
                msg["tag_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ApplicationPolicySet) UnmarshalJSON(body []byte) error {
        var m map[string]json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
                return err
        }
        err = obj.UnmarshalCommon(m)
        if err != nil {
                return err
        }
        for key, value := range m {
                switch key {
                case "is_global":
                        err = json.Unmarshal(value, &obj.is_global)
                        if err == nil {
                                obj.valid |= application_policy_set_is_global
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= application_policy_set_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= application_policy_set_perms2
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid |= application_policy_set_annotations
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= application_policy_set_display_name
                        }
                        break
                case "global_vrouter_config_refs":
                        err = json.Unmarshal(value, &obj.global_vrouter_config_refs)
                        if err == nil {
                                obj.valid |= application_policy_set_global_vrouter_config_refs
                        }
                        break
                case "tag_refs":
                        err = json.Unmarshal(value, &obj.tag_refs)
                        if err == nil {
                                obj.valid |= application_policy_set_tag_refs
                        }
                        break
                case "firewall_policy_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr FirewallSequence
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= application_policy_set_firewall_policy_refs
                        obj.firewall_policy_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.firewall_policy_refs = append(obj.firewall_policy_refs, ref)
                        }
                        break
                }
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ApplicationPolicySet) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & application_policy_set_is_global != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.is_global)
                if err != nil {
                        return nil, err
                }
                msg["is_global"] = &value
        }

        if obj.modified & application_policy_set_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & application_policy_set_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & application_policy_set_annotations != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified & application_policy_set_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & application_policy_set_firewall_policy_refs != 0 {
                if len(obj.firewall_policy_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["firewall_policy_refs"] = &value
                } else if !obj.hasReferenceBase("firewall-policy") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.firewall_policy_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["firewall_policy_refs"] = &value
                }
        }


        if obj.modified & application_policy_set_global_vrouter_config_refs != 0 {
                if len(obj.global_vrouter_config_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["global_vrouter_config_refs"] = &value
                } else if !obj.hasReferenceBase("global-vrouter-config") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.global_vrouter_config_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["global_vrouter_config_refs"] = &value
                }
        }


        if obj.modified & application_policy_set_tag_refs != 0 {
                if len(obj.tag_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["tag_refs"] = &value
                } else if !obj.hasReferenceBase("tag") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.tag_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["tag_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *ApplicationPolicySet) UpdateReferences() error {

        if (obj.modified & application_policy_set_firewall_policy_refs != 0) &&
           len(obj.firewall_policy_refs) > 0 &&
           obj.hasReferenceBase("firewall-policy") {
                err := obj.UpdateReference(
                        obj, "firewall-policy",
                        obj.firewall_policy_refs,
                        obj.baseMap["firewall-policy"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & application_policy_set_global_vrouter_config_refs != 0) &&
           len(obj.global_vrouter_config_refs) > 0 &&
           obj.hasReferenceBase("global-vrouter-config") {
                err := obj.UpdateReference(
                        obj, "global-vrouter-config",
                        obj.global_vrouter_config_refs,
                        obj.baseMap["global-vrouter-config"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & application_policy_set_tag_refs != 0) &&
           len(obj.tag_refs) > 0 &&
           obj.hasReferenceBase("tag") {
                err := obj.UpdateReference(
                        obj, "tag",
                        obj.tag_refs,
                        obj.baseMap["tag"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func ApplicationPolicySetByName(c contrail.ApiClient, fqn string) (*ApplicationPolicySet, error) {
    obj, err := c.FindByName("application-policy-set", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ApplicationPolicySet), nil
}

func ApplicationPolicySetByUuid(c contrail.ApiClient, uuid string) (*ApplicationPolicySet, error) {
    obj, err := c.FindByUuid("application-policy-set", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ApplicationPolicySet), nil
}
