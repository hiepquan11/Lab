package config

import (
	"document-management/models"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"log"
)

var Enforcer *casbin.Enforcer

func CasbinConfig(db *gorm.DB) {
	// initialize casbin with gorm adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Error creating casbin gorm adapter, %s", err)
	}

	// load casbin model config
	enforcer, err := casbin.NewEnforcer("share/casbin/model.conf", adapter)
	if err != nil {
		log.Fatalf("Error initialize casbin enforcer, %s", err)
	}
	// load policy from db
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Error loading casbin policy, %s", err)
	}
	Enforcer = enforcer
}

// AddPolicy adds multiple policies (read, edit, delete) into Casbin
func AddPolicy(policy *models.CasbinPolicy) error {
	if Enforcer == nil {
		return fmt.Errorf("casbin enforcer is not initialized")
	}

	// permissions
	var perms = []models.Permission{
		models.Read,
		models.Edit,
		models.Del,
	}

	for _, perm := range perms {
		success, err := Enforcer.AddPolicy(policy.V0, policy.V1, string(perm))
		if err != nil {
			log.Printf("Failed to add policy: %s", err)
			return err
		}
		if !success {
			log.Printf("Policy already exists: %s, %s, %s", policy.V0, policy.V1, perm)
		}
	}

	return nil
}

// AddGroupPolicy create group policy
func AddGroupPolicy(sub string, act string, domain string) error {
	if Enforcer == nil {
		return fmt.Errorf("casbin enforcer is not initialized")
	}

	success, err := Enforcer.AddGroupingPolicy(sub, act, domain)
	if err != nil {
		log.Printf("Failed to add group policy: %s", err)
		return err
	}
	if !success {
		log.Printf("Group policy already exists: %s, %s, %s", sub, act, domain)
	}
	return nil
}
