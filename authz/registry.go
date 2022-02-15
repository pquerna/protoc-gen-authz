package authz

import (
	"errors"
	"fmt"
	"sort"
)

// PermissionId is an in-process identifier for a Permission
type PermissionId uint32

type Permission struct {
	ID   PermissionId
	Name string
}

type MessageRequiresPermission interface {
	AuthzRequiresPermission() PermissionId
}

var (
	roleRegistry       map[string][]PermissionId
	permissionMap      map[string]PermissionId
	permissionRegistry []Permission
)

func init() {
	roleRegistry = make(map[string][]PermissionId)
	permissionMap = make(map[string]PermissionId)
	permissionRegistry = make([]Permission, 0, 64)
	permissionRegistry = append(permissionRegistry, Permission{
		ID:   0,
		Name: "invalid",
	})
}

// RolePermissions returns a list of permissions from a service Role.
//
// The returned PermissionId is only valid for within this process. Use
// PermissionLookup to find a portable string representation of an Permission
func RolePermissions(roleName string) []PermissionId {
	return roleRegistry[roleName]
}

// PermissionFromID converts a in-memory PermissionId to a Permission
func PermissionFromID(id PermissionId) Permission {
	return permissionRegistry[id]
}

var ErrInvalidPermissionName = errors.New("authz: invalid permission name")

// PermissionFromName finds a Permission based on its string identifier.
func PermissionFromName(name string) (Permission, error) {
	pid, ok := permissionMap[name]
	if !ok {
		return Permission{}, fmt.Errorf("authz: invalid permission: '%s': %w",
			name, ErrInvalidPermissionName)
	}
	return permissionRegistry[pid], nil
}

// RegisterRole stores a role to permission mapping.
//
// NOTE: RegisterRole is not goroutine safe, and is intended to be called
// from init() functions in generated code.
func RegisterRole(roleName string, permissions []PermissionId) {
	roleRegistry[roleName] = append(roleRegistry[roleName], permissions...)
}

// RegisterPermission returns a in-memory integer identifier for a permission.
//
// NOTE: RegisterPermission is not goroutine safe, and is intended to be called
// from init() functions in generated code.
func RegisterPermission(permissionName string) Permission {
	id := PermissionId(len(permissionRegistry))
	permissionRegistry = append(permissionRegistry, Permission{
		ID:   id,
		Name: permissionName,
	})
	permissionMap[permissionName] = id
	return permissionRegistry[id]
}

// AllPermissions provides a copy of all registered Permissions.
func AllPermissions() []Permission {
	rv := make([]Permission, 0, len(permissionRegistry))
	rv = append(rv, permissionRegistry...)
	return rv
}

// AllRoles provides a copy of all registered Role strings.
func AllRoles() []string {
	rv := make([]string, 0, len(roleRegistry))
	for k, _ := range roleRegistry {
		rv = append(rv, k)
	}
	sort.Strings(rv)
	return rv
}
