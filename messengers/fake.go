/* walter: a deployment pipeline template
 * Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
 * (see CONTRIBUTORS.md)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package messengers defines clients for the messaging services such as Slack.
// The clients reports the results of the pipeline stages and whole results.
package messengers

// FakeMessenger is a Messenger type used only for testing.
type FakeMessenger struct {
	BaseMessenger
}

//Post just returns true indicating the messages was posted (faked)
func (fakeMsg *FakeMessenger) Post(messege string) bool {
	return true
}
