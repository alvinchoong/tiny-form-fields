package tinyformfields

import (
	"errors"
	"net/url"
	"testing"
)

func TestValidFormValues(t *testing.T) {
	formFieldsJSON := `[
          {
            "label": "Question 1",
            "name": "question_1",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "Dropdown",
              "choices": [
                "Red",
                "Orange",
                "Yellow",
                "Green",
                "Blue",
                "Indigo",
                "Violet"
              ]
            }
          },
          {
            "label": "Question 2",
            "name": "question_2",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ChooseOne",
              "choices": [
                "Yes",
                "No"
              ]
            }
          },
          {
            "label": "Question 3",
            "name": "question_3",
            "presence": "Optional",
            "description": "",
            "type": {
              "type": "ChooseMultiple",
              "choices": [
                "Apple",
                "Banana",
                "Cantaloupe",
                "Durian"
              ]
            }
          },
          {
            "label": "Question 4",
            "name": "question_4",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "LongText",
              "maxLength": 160
            }
          },
          {
            "label": "Question 5",
            "name": "question_5",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Single-line free text",
              "attributes": {
                "type": "text"
              }
            }
          },
          {
            "label": "Question 6",
            "name": "question_6",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Email",
              "attributes": {
                "type": "email"
              }
            }
          },
          {
            "label": "Question 7",
            "name": "question_7",
            "presence": "Optional",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Emails",
              "attributes": {
                "multiple": "true",
                "type": "email"
              }
            }
          },
          {
            "label": "Question 8",
            "name": "question_8",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "NRIC",
              "attributes": {
                "maxlength": "9",
                "minlength": "9",
                "pattern": "^[STGM][0-9]{7}[ABCDEFGHIZJ]$",
                "type": "text"
              }
            }
          },
          {
            "label": "Question 9",
            "name": "question_9",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Telephone",
              "attributes": {
                "type": "tel"
              }
            }
          },
          {
            "label": "Question 10",
            "name": "question_10",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "URL",
              "attributes": {
                "type": "url"
              }
            }
          },
          {
            "label": "Question 11",
            "name": "question_11",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Color",
              "attributes": {
                "type": "color"
              }
            }
          },
          {
            "label": "Question 12",
            "name": "question_12",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Date",
              "attributes": {
                "type": "date"
              }
            }
          },
          {
            "label": "Question 13",
            "name": "question_13",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Time",
              "attributes": {
                "type": "time"
              }
            }
          },
          {
            "label": "Question 14",
            "name": "question_14",
            "presence": "Required",
            "description": "",
            "type": {
              "type": "ShortText",
              "inputType": "Date & Time",
              "attributes": {
                "type": "datetime-local"
              }
            }
          }
        ]`

	formValues := url.Values{
		"question_1":  {"Red"},
		"question_2":  {"No"},
		"question_3":  {"Apple", "Banana", "Cantaloupe", "Durian"},
		"question_4":  {"multiple lines\r\nare accepted\r\nhere"},
		"question_5":  {"single line only"},
		"question_6":  {"alice@example.com"},
		"question_7":  {"alice@example.com,bob@example.com"},
		"question_8":  {"S1234567A"},
		"question_9":  {"123"},
		"question_10": {"ftp://example.com"},
		"question_11": {"#000000"},
		"question_12": {"2024-09-19"},
		"question_13": {"19:18"},
		"question_14": {"2024-09-19T21:01"},
	}

	err := ValidFormValues([]byte(formFieldsJSON), formValues)
	if err != nil {
		t.Errorf("Validation Error: %v", err)
	}
}

func TestChoiceParsingAndValidation(t *testing.T) {
	formFieldsJSON := `[
		{
			"label": "Question 1",
			"name": "question_1",
			"presence": "Required",
			"type": {
				"type": "Dropdown",
				"choices": [
					"Yes",
					"Maybe | I might want to go!",
					"No"
				]
			}
		},
		{
			"label": "Question 2",
			"name": "question_2",
			"presence": "Required",
			"type": {
				"type": "ChooseMultiple",
				"choices": [
					"Option1 | First Option",
					"Option2 | Second Option",
					"Option3"
				]
			}
		}
	]`

	formValues := url.Values{
		"question_1": {"Maybe"},
		"question_2": {"Option1", "Option3"},
	}

	err := ValidFormValues([]byte(formFieldsJSON), formValues)
	if err != nil {
		t.Errorf("Validation Error: %v", err)
	}

	// Test invalid value
	formValuesInvalid := url.Values{
		"question_1": {"I might want to go!"},
	}

	err = ValidFormValues([]byte(formFieldsJSON), formValuesInvalid)
	expectedError := "invalid choice: question_1 has invalid value 'I might want to go!'. Valid choices are: [Yes Maybe No]"
	if err == nil {
		t.Errorf("Expected error but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error %q but got %q", expectedError, err.Error())
	}
}

func TestInvalidFormValues(t *testing.T) {
	scenarios := []struct {
		name        string
		formFields  string
		formValues  url.Values
		expectError error
	}{
		{
			name: "Invalid Choice with Label",
			formFields: `[
			  {
				"label": "Question 1",
				"name": "question_1",
				"presence": "Required",
				"type": {
				  "type": "Dropdown",
				  "choices": [
					  "Yes",
					  "Maybe | I might want to go!",
					  "No"
				  ]
				}
			  }
			]`,
			formValues:  url.Values{"question_1": {"I might want to go!"}},
			expectError: ErrInvalidChoice,
		},
		{
			name: "Valid Choice with Value",
			formFields: `[
			  {
				"label": "Question 1",
				"name": "question_1",
				"presence": "Required",
				"type": {
				  "type": "Dropdown",
				  "choices": [
					  "Yes",
					  "Maybe | I might want to go!",
					  "No"
				  ]
				}
			  }
			]`,
			formValues:  url.Values{"question_1": {"Maybe"}},
			expectError: nil,
		},
		{
			name: "Invalid Choice in Dropdown",
			formFields: `[
			  {
				"label": "Question 1",
				"name": "question_1",
				"presence": "Required",
				"type": {
				  "type": "Dropdown",
				  "choices": ["Red", "Green", "Blue"]
				}
			  }
			]`,
			formValues:  url.Values{"question_1": {"Purple"}},
			expectError: ErrInvalidChoice,
		},
		{
			name: "Missing Required Field",
			formFields: `[
			  {
				"label": "Question 1",
				"name": "question_1",
				"presence": "Required",
				"type": {
				  "type": "ShortText",
				  "attributes": {"type": "text"}
				}
			  }
			]`,
			formValues:  url.Values{},
			expectError: ErrRequiredFieldMissing,
		},
		{
			name: "Valid Optional Field Missing",
			formFields: `[
			  {
				"label": "Question 1",
				"name": "question_1",
				"presence": "Optional",
				"type": {
				  "type": "ShortText",
				  "attributes": {"type": "text"}
				}
			  }
			]`,
			formValues:  url.Values{},
			expectError: nil,
		},
		{
			name: "Invalid Email",
			formFields: `[
			  {
				"label": "Email",
				"name": "email",
				"presence": "Required",
				"type": {
				  "type": "ShortText",
				  "attributes": {"type": "email"}
				}
			  }
			]`,
			formValues:  url.Values{"email": {"invalid-email"}},
			expectError: ErrInvalidEmail,
		},
		{
			name: "Invalid Pattern",
			formFields: `[
			  {
				"label": "Pattern Field",
				"name": "pattern_field",
				"presence": "Required",
				"type": {
				  "type": "ShortText",
				  "attributes": {
					"type": "text",
					"pattern": "^[0-9]{3}$"
				  }
				}
			  }
			]`,
			formValues:  url.Values{"pattern_field": {"12a"}},
			expectError: ErrInvalidPattern,
		},
		{
			name: "Line Break Not Allowed",
			formFields: `[
			  {
				"label": "Short Text Field",
				"name": "short_text",
				"presence": "Required",
				"type": {
				  "type": "ShortText",
				  "attributes": {
					"type": "text"
				  }
				}
			  }
			]`,
			formValues:  url.Values{"short_text": {"Line1\nLine2"}},
			expectError: ErrLineBreakNotAllowed,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			err := ValidFormValues([]byte(scenario.formFields), scenario.formValues)
			if scenario.expectError == nil && err != nil {
				t.Errorf("Unexpected error: %v", err)
			} else if scenario.expectError != nil {
				if err == nil {
					t.Errorf("Expected error %v but got nil", scenario.expectError)
				} else if !errors.Is(err, scenario.expectError) {
					t.Errorf("Expected error %v but got %v", scenario.expectError, err)
				}
			}
		})
	}
}