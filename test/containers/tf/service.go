package tf

import (
	"log"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/testing"
)

const (
	scriptDir = "../terraform"
)

type (
	TearDownFunc func()
)

func Deploy(t testing.TestingT, cfg Config) TearDownFunc {
	tfOptions := &terraform.Options{
		TerraformDir: scriptDir,
		Vars:         cfg.Vars(),
		NoColor:      true,
	}

	out := terraform.InitAndApply(t, tfOptions)
	log.Printf("Terraform deploy: %s", out)

	return func() {
		terraform.Destroy(t, tfOptions)
	}
}
