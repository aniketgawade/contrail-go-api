//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/aniketgawade/contrail-go-api"
)

const (
	firewall_policy_id_perms uint64 = 1 << iota
	firewall_policy_perms2
	firewall_policy_annotations
	firewall_policy_display_name
	firewall_policy_firewall_rule_refs
	firewall_policy_tag_refs
	firewall_policy_application_policy_set_back_refs
)

type FirewallPolicy struct {
        contrail.ObjectBase
	id_perms IdPermsType
	perms2 PermType2
	annotations KeyValuePairs
	display_name string
	firewall_rule_refs contrail.ReferenceList
	tag_refs contrail.ReferenceList
	application_policy_set_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *FirewallPolicy) GetType() string {
        return "firewall-policy"
}

func (obj *FirewallPolicy) GetDefaultParent() []string {
        name := []string{"default-policy-management"}
        return name
}

func (obj *FirewallPolicy) GetDefaultParentType() string {
        return "policy-management"
}

func (obj *FirewallPolicy) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *FirewallPolicy) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *FirewallPolicy) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *FirewallPolicy) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *FirewallPolicy) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *FirewallPolicy) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *FirewallPolicy) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= firewall_policy_id_perms
}

func (obj *FirewallPolicy) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *FirewallPolicy) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= firewall_policy_perms2
}

func (obj *FirewallPolicy) GetAnnotations() KeyValuePairs {
        return obj.annotations
}

func (obj *FirewallPolicy) SetAnnotations(value *KeyValuePairs) {
        obj.annotations = *value
        obj.modified |= firewall_policy_annotations
}

func (obj *FirewallPolicy) GetDisplayName() string {
        return obj.display_name
}

func (obj *FirewallPolicy) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= firewall_policy_display_name
}

func (obj *FirewallPolicy) readFirewallRuleRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & firewall_policy_firewall_rule_refs == 0) {
                err := obj.GetField(obj, "firewall_rule_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FirewallPolicy) GetFirewallRuleRefs() (
        contrail.ReferenceList, error) {
        err := obj.readFirewallRuleRefs()
        if err != nil {
                return nil, err
        }
        return obj.firewall_rule_refs, nil
}

func (obj *FirewallPolicy) AddFirewallRule(
        rhs *FirewallRule, data FirewallSequence) error {
        err := obj.readFirewallRuleRefs()
        if err != nil {
                return err
        }

        if obj.modified & firewall_policy_firewall_rule_refs == 0 {
                obj.storeReferenceBase("firewall-rule", obj.firewall_rule_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
        obj.firewall_rule_refs = append(obj.firewall_rule_refs, ref)
        obj.modified |= firewall_policy_firewall_rule_refs
        return nil
}

func (obj *FirewallPolicy) DeleteFirewallRule(uuid string) error {
        err := obj.readFirewallRuleRefs()
        if err != nil {
                return err
        }

        if obj.modified & firewall_policy_firewall_rule_refs == 0 {
                obj.storeReferenceBase("firewall-rule", obj.firewall_rule_refs)
        }

        for i, ref := range obj.firewall_rule_refs {
                if ref.Uuid == uuid {
                        obj.firewall_rule_refs = append(
                                obj.firewall_rule_refs[:i],
                                obj.firewall_rule_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= firewall_policy_firewall_rule_refs
        return nil
}

func (obj *FirewallPolicy) ClearFirewallRule() {
        if (obj.valid & firewall_policy_firewall_rule_refs != 0) &&
           (obj.modified & firewall_policy_firewall_rule_refs == 0) {
                obj.storeReferenceBase("firewall-rule", obj.firewall_rule_refs)
        }
        obj.firewall_rule_refs = make([]contrail.Reference, 0)
        obj.valid |= firewall_policy_firewall_rule_refs
        obj.modified |= firewall_policy_firewall_rule_refs
}

func (obj *FirewallPolicy) SetFirewallRuleList(
        refList []contrail.ReferencePair) {
        obj.ClearFirewallRule()
        obj.firewall_rule_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.firewall_rule_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *FirewallPolicy) readTagRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & firewall_policy_tag_refs == 0) {
                err := obj.GetField(obj, "tag_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FirewallPolicy) GetTagRefs() (
        contrail.ReferenceList, error) {
        err := obj.readTagRefs()
        if err != nil {
                return nil, err
        }
        return obj.tag_refs, nil
}

func (obj *FirewallPolicy) AddTag(
        rhs *Tag) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if obj.modified & firewall_policy_tag_refs == 0 {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.tag_refs = append(obj.tag_refs, ref)
        obj.modified |= firewall_policy_tag_refs
        return nil
}

func (obj *FirewallPolicy) DeleteTag(uuid string) error {
        err := obj.readTagRefs()
        if err != nil {
                return err
        }

        if obj.modified & firewall_policy_tag_refs == 0 {
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
        obj.modified |= firewall_policy_tag_refs
        return nil
}

func (obj *FirewallPolicy) ClearTag() {
        if (obj.valid & firewall_policy_tag_refs != 0) &&
           (obj.modified & firewall_policy_tag_refs == 0) {
                obj.storeReferenceBase("tag", obj.tag_refs)
        }
        obj.tag_refs = make([]contrail.Reference, 0)
        obj.valid |= firewall_policy_tag_refs
        obj.modified |= firewall_policy_tag_refs
}

func (obj *FirewallPolicy) SetTagList(
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


func (obj *FirewallPolicy) readApplicationPolicySetBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & firewall_policy_application_policy_set_back_refs == 0) {
                err := obj.GetField(obj, "application_policy_set_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FirewallPolicy) GetApplicationPolicySetBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readApplicationPolicySetBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.application_policy_set_back_refs, nil
}

func (obj *FirewallPolicy) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & firewall_policy_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & firewall_policy_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & firewall_policy_annotations != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified & firewall_policy_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.firewall_rule_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.firewall_rule_refs)
                if err != nil {
                        return nil, err
                }
                msg["firewall_rule_refs"] = &value
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

func (obj *FirewallPolicy) UnmarshalJSON(body []byte) error {
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
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= firewall_policy_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= firewall_policy_perms2
                        }
                        break
                case "annotations":
                        err = json.Unmarshal(value, &obj.annotations)
                        if err == nil {
                                obj.valid |= firewall_policy_annotations
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= firewall_policy_display_name
                        }
                        break
                case "tag_refs":
                        err = json.Unmarshal(value, &obj.tag_refs)
                        if err == nil {
                                obj.valid |= firewall_policy_tag_refs
                        }
                        break
                case "firewall_rule_refs": {
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
                        obj.valid |= firewall_policy_firewall_rule_refs
                        obj.firewall_rule_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.firewall_rule_refs = append(obj.firewall_rule_refs, ref)
                        }
                        break
                }
                case "application_policy_set_back_refs": {
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
                        obj.valid |= firewall_policy_application_policy_set_back_refs
                        obj.application_policy_set_back_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.application_policy_set_back_refs = append(obj.application_policy_set_back_refs, ref)
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

func (obj *FirewallPolicy) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & firewall_policy_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & firewall_policy_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & firewall_policy_annotations != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.annotations)
                if err != nil {
                        return nil, err
                }
                msg["annotations"] = &value
        }

        if obj.modified & firewall_policy_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & firewall_policy_firewall_rule_refs != 0 {
                if len(obj.firewall_rule_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["firewall_rule_refs"] = &value
                } else if !obj.hasReferenceBase("firewall-rule") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.firewall_rule_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["firewall_rule_refs"] = &value
                }
        }


        if obj.modified & firewall_policy_tag_refs != 0 {
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

func (obj *FirewallPolicy) UpdateReferences() error {

        if (obj.modified & firewall_policy_firewall_rule_refs != 0) &&
           len(obj.firewall_rule_refs) > 0 &&
           obj.hasReferenceBase("firewall-rule") {
                err := obj.UpdateReference(
                        obj, "firewall-rule",
                        obj.firewall_rule_refs,
                        obj.baseMap["firewall-rule"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & firewall_policy_tag_refs != 0) &&
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

func FirewallPolicyByName(c contrail.ApiClient, fqn string) (*FirewallPolicy, error) {
    obj, err := c.FindByName("firewall-policy", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*FirewallPolicy), nil
}

func FirewallPolicyByUuid(c contrail.ApiClient, uuid string) (*FirewallPolicy, error) {
    obj, err := c.FindByUuid("firewall-policy", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*FirewallPolicy), nil
}
