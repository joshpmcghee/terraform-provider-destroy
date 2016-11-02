provider "destroy" {}

resource "destroy_hook" "broken_script" {
  name = "failure"
  commands = [
    "${file("broken_script.sh")}"
  ]
  fail_on_error = true
  retries = 4
  retry_period = 10
}
