// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Group_To_user_Group,
		Convert_user_Group_To_v1_Group,
		Convert_v1_GroupList_To_user_GroupList,
		Convert_user_GroupList_To_v1_GroupList,
		Convert_v1_Identity_To_user_Identity,
		Convert_user_Identity_To_v1_Identity,
		Convert_v1_IdentityList_To_user_IdentityList,
		Convert_user_IdentityList_To_v1_IdentityList,
		Convert_v1_User_To_user_User,
		Convert_user_User_To_v1_User,
		Convert_v1_UserIdentityMapping_To_user_UserIdentityMapping,
		Convert_user_UserIdentityMapping_To_v1_UserIdentityMapping,
		Convert_v1_UserList_To_user_UserList,
		Convert_user_UserList_To_v1_UserList,
	)
}

func autoConvert_v1_Group_To_user_Group(in *Group, out *user.Group, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	return nil
}

func Convert_v1_Group_To_user_Group(in *Group, out *user.Group, s conversion.Scope) error {
	return autoConvert_v1_Group_To_user_Group(in, out, s)
}

func autoConvert_user_Group_To_v1_Group(in *user.Group, out *Group, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Users == nil {
		out.Users = make(OptionalNames, 0)
	} else {
		out.Users = *(*OptionalNames)(unsafe.Pointer(&in.Users))
	}
	return nil
}

func Convert_user_Group_To_v1_Group(in *user.Group, out *Group, s conversion.Scope) error {
	return autoConvert_user_Group_To_v1_Group(in, out, s)
}

func autoConvert_v1_GroupList_To_user_GroupList(in *GroupList, out *user.GroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]user.Group)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_GroupList_To_user_GroupList(in *GroupList, out *user.GroupList, s conversion.Scope) error {
	return autoConvert_v1_GroupList_To_user_GroupList(in, out, s)
}

func autoConvert_user_GroupList_To_v1_GroupList(in *user.GroupList, out *GroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Group, 0)
	} else {
		out.Items = *(*[]Group)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_user_GroupList_To_v1_GroupList(in *user.GroupList, out *GroupList, s conversion.Scope) error {
	return autoConvert_user_GroupList_To_v1_GroupList(in, out, s)
}

func autoConvert_v1_Identity_To_user_Identity(in *Identity, out *user.Identity, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Extra = *(*map[string]string)(unsafe.Pointer(&in.Extra))
	return nil
}

func Convert_v1_Identity_To_user_Identity(in *Identity, out *user.Identity, s conversion.Scope) error {
	return autoConvert_v1_Identity_To_user_Identity(in, out, s)
}

func autoConvert_user_Identity_To_v1_Identity(in *user.Identity, out *Identity, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	out.Extra = *(*map[string]string)(unsafe.Pointer(&in.Extra))
	return nil
}

func Convert_user_Identity_To_v1_Identity(in *user.Identity, out *Identity, s conversion.Scope) error {
	return autoConvert_user_Identity_To_v1_Identity(in, out, s)
}

func autoConvert_v1_IdentityList_To_user_IdentityList(in *IdentityList, out *user.IdentityList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]user.Identity, len(*in))
		for i := range *in {
			if err := Convert_v1_Identity_To_user_Identity(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_IdentityList_To_user_IdentityList(in *IdentityList, out *user.IdentityList, s conversion.Scope) error {
	return autoConvert_v1_IdentityList_To_user_IdentityList(in, out, s)
}

func autoConvert_user_IdentityList_To_v1_IdentityList(in *user.IdentityList, out *IdentityList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Identity, len(*in))
		for i := range *in {
			if err := Convert_user_Identity_To_v1_Identity(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = make([]Identity, 0)
	}
	return nil
}

func Convert_user_IdentityList_To_v1_IdentityList(in *user.IdentityList, out *IdentityList, s conversion.Scope) error {
	return autoConvert_user_IdentityList_To_v1_IdentityList(in, out, s)
}

func autoConvert_v1_User_To_user_User(in *User, out *user.User, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.FullName = in.FullName
	out.Identities = *(*[]string)(unsafe.Pointer(&in.Identities))
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	return nil
}

func Convert_v1_User_To_user_User(in *User, out *user.User, s conversion.Scope) error {
	return autoConvert_v1_User_To_user_User(in, out, s)
}

func autoConvert_user_User_To_v1_User(in *user.User, out *User, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.FullName = in.FullName
	if in.Identities == nil {
		out.Identities = make([]string, 0)
	} else {
		out.Identities = *(*[]string)(unsafe.Pointer(&in.Identities))
	}
	if in.Groups == nil {
		out.Groups = make([]string, 0)
	} else {
		out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	}
	return nil
}

func Convert_user_User_To_v1_User(in *user.User, out *User, s conversion.Scope) error {
	return autoConvert_user_User_To_v1_User(in, out, s)
}

func autoConvert_v1_UserIdentityMapping_To_user_UserIdentityMapping(in *UserIdentityMapping, out *user.UserIdentityMapping, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.Identity, &out.Identity, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ObjectReference_To_api_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_UserIdentityMapping_To_user_UserIdentityMapping(in *UserIdentityMapping, out *user.UserIdentityMapping, s conversion.Scope) error {
	return autoConvert_v1_UserIdentityMapping_To_user_UserIdentityMapping(in, out, s)
}

func autoConvert_user_UserIdentityMapping_To_v1_UserIdentityMapping(in *user.UserIdentityMapping, out *UserIdentityMapping, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.Identity, &out.Identity, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ObjectReference_To_v1_ObjectReference(&in.User, &out.User, s); err != nil {
		return err
	}
	return nil
}

func Convert_user_UserIdentityMapping_To_v1_UserIdentityMapping(in *user.UserIdentityMapping, out *UserIdentityMapping, s conversion.Scope) error {
	return autoConvert_user_UserIdentityMapping_To_v1_UserIdentityMapping(in, out, s)
}

func autoConvert_v1_UserList_To_user_UserList(in *UserList, out *user.UserList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]user.User)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1_UserList_To_user_UserList(in *UserList, out *user.UserList, s conversion.Scope) error {
	return autoConvert_v1_UserList_To_user_UserList(in, out, s)
}

func autoConvert_user_UserList_To_v1_UserList(in *user.UserList, out *UserList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]User, 0)
	} else {
		out.Items = *(*[]User)(unsafe.Pointer(&in.Items))
	}
	return nil
}

func Convert_user_UserList_To_v1_UserList(in *user.UserList, out *UserList, s conversion.Scope) error {
	return autoConvert_user_UserList_To_v1_UserList(in, out, s)
}
