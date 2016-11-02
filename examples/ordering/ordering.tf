provider "destroy" {}

resource "destroy_hook" "a" {
  name = "test_b"
  commands = [
    "sleep 20"
  ]

  depends_on = ["null_resource.a"]
}

resource null_resource "a" {

  depends_on = ["destroy_hook.b"]
}

resource "destroy_hook" "b" {
  name = "test_b"
  commands = [
    "sleep 20"
  ]

  depends_on = ["null_resource.b"]
}

resource null_resource "b" {

}
