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

resource "destroy_hook" "broken_script" {
  name = "failure"
  commands = [
    "${file("broken_script.sh")}"
  ]
  fail_on_error = true
  retries = 4
  retry_period = 10
}