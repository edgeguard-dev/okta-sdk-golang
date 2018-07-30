/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
)

type SecurityQuestionFactorProfile struct {
	Answer string `json:"answer,omitempty"`
	Question string `json:"question,omitempty"`
	QuestionText string `json:"questionText,omitempty"`
}

func (m *SecurityQuestionFactorProfile) WithAnswer(v string) *SecurityQuestionFactorProfile {
	m.Answer = v
	return m
}

func (m *SecurityQuestionFactorProfile) WithQuestion(v string) *SecurityQuestionFactorProfile {
	m.Question = v
	return m
}

func (m *SecurityQuestionFactorProfile) WithQuestionText(v string) *SecurityQuestionFactorProfile {
	m.QuestionText = v
	return m
}

