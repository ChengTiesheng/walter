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

package stages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatorWithSimpleCommand(t *testing.T) {
	validator := NewResourceValidator()
	validator.AddCommandName("ls")
	result := validator.Validate()
	assert.Equal(t, true, result)
}

func TestValidatorWithNoExistCommand(t *testing.T) {
	validator := NewResourceValidator()
	validator.AddCommandName("xxxxx")
	result := validator.Validate()
	assert.Equal(t, false, result)
}

func TestValidatorWithFile(t *testing.T) {
	validator := NewResourceValidator()
	validator.AddFile("./test_sample.sh")
	result := validator.Validate()
	assert.Equal(t, true, result)
}

func TestValidatorWithNoExistFile(t *testing.T) {
	validator := NewResourceValidator()
	validator.AddFile("xxxxx")
	result := validator.Validate()
	assert.Equal(t, false, result)
}

func TestValidatorWithFileAndNoExistCommand(t *testing.T) {
	validator := NewResourceValidator()
	validator.AddFile("./test_sample.sh")
	validator.AddCommandName("xxxxx")
	result := validator.Validate()
	assert.Equal(t, false, result)
}

func TestValidatorWithNoExistFileAndExistCommand(t *testing.T) {
	validator := NewResourceValidator()
	validator.AddFile("xxxxx")
	validator.AddCommandName("ls")
	result := validator.Validate()
	assert.Equal(t, false, result)
}
