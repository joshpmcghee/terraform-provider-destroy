provider "destroy" {}

resource "destroy_hook" "test_script" {
  name = "test_file"
  commands = [
    "${file("test_script.sh")}"
  ]

}
