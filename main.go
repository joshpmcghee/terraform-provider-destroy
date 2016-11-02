package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

type Destroy struct{}

type Hook struct {
	Name     string
	Commands []string
}

func (m *Hook) Id() string {
	return m.Name
}

func (c *Destroy) CreateHook(m *Hook) error {
	return nil
}

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"destroy_hook": &schema.Resource{
			SchemaVersion: 1,
			Create:        createFunc,
			Read:          readFunc,
			Update:        updateFunc,
			Delete:        deleteFunc,
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"fail_on_error": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"retries": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  3,
				},
				"retry_period": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  60,
				},
				"commands": &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Required: true,
				},
			},
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	destroy := Destroy{}

	return &destroy, nil
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("name").(string))
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	var cmd *exec.Cmd
	var err error
	commands := d.Get("commands")
	name := d.Get("name")
	shouldFail := d.Get("fail_on_error")
	retries := d.Get("retries")
	period := d.Get("retry_period")

	for _, command := range commands.([]interface{}) {
		attempt := 1

		for attempt <= retries.(int) {
			cmd = exec.Command("sh", "-c", command.(string))
			err = cmd.Run()

			if err != nil {
				out, _ := cmd.Output()
				writeErr(fmt.Sprintf("Timestamp: %v\nAttempt: %v\nHook: %v\nCommand:\n%v\nError: %v\nOutput: %v\n", time.Now(), attempt, name.(string), command.(string), err, out))
			}
			attempt += 1
			time.Sleep(time.Duration(period.(int)) * time.Second)
		}

		// Deliberately will NOT fail unless specified.
		if err != nil && shouldFail.(bool) {
			return err
		}
	}

	return nil
}

// Errors are ignored because we don't really care and definitely do not want it to fail.
func writeErr(log string) {
	errLog, _ := os.OpenFile("destroy_hook_error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	defer errLog.Close()

	_, _ = errLog.WriteString(log)
}
