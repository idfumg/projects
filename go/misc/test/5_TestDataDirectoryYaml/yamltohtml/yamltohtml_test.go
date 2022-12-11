package yamltohtml

import "testing"

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestYamlToHtml(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "Test case 1",
			path:     "testdata/test_01.yaml",
			expected: "<html><head><title>My Awesome Page</title></head><body>This is my awesome content</body></html>",
		},
		{
			desc:     "Test case 2",
			path:     "testdata/test_02.yaml",
			expected: "<html><head><title>My Awesome Page 2</title></head><body>This is my awesome content 2</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := 	YamlToHtml(test.path)
			if err != nil {
				t.Fail()
			}

			t.Log(result)

			if result != test.expected {
				t.Fail()
			}
		})
	}
}
