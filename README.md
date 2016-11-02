# terraform-provider-destroy

A simple Terraform provider to add the ability to insert custom behaviour "on destroy" like clean-up and tear-down tasks.

Failed hooks, by default, will **not** fail the destroy. In the majority of use-cases, this is not the desired behaviour. If you wish to fail the destroy on an error in your hooks then you may set `fail_on_error = true`. All errors will be logged to `destroy_provider_errors.log` in your working directory.

Here is an example making use of all implemented functionality:
```
provider "destroy" {}

resource "destroy_hook" "script" {
  name = "script"
  commands = [
    "${file("script.sh")}",
    "echo \"you've failed me for the last time, commander.\" > darth_vader_says.txt"
  ]
  fail_on_error = true
  retries = 4
  retry_period = 10
}
```
Further examples, including the use of dependency to force ordering, can be found in the `examples/` directory.

## Installation

```
make install

# Do this if you don't currently have a .terraformrc. If you do, add the provider entry manually.
cp terraformrc ~/.terraformrc
```

Credit to [ContainerSolutions](https://github.com/ContainerSolutions) whose [template](https://github.com/ContainerSolutions/terraform-provider-template) I based this provider on.
