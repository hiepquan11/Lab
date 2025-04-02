package utils

import (
	"document-management/core/config"
	"document-management/models"
	"fmt"
	"log"
)

func AssignWorkspacePermission(userId string, workspaceId string) {

	var rule = models.CasbinPolicy{
		V0: fmt.Sprintf("user"),
		V1: fmt.Sprintf("%s", workspaceId),
		//V2: "read",
	}

	err := config.AddPolicy(&rule)
	if err != nil {
		log.Fatal("add policy err: ", err)
	}

	err = config.AddGroupPolicy(userId, "user", workspaceId)
}
