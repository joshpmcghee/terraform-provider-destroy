# Copyright 2015 Container Solutions
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
