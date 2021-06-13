package templaterunner_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/walmartdigital/commit-message-formatter/mocks"
	"github.com/walmartdigital/commit-message-formatter/templaterunner"
)

var ctrl *gomock.Controller

type promptCallback func([]rune) error

func TestAll(t *testing.T) {
	ctrl = gomock.NewController(t)
	defer ctrl.Finish()

	RegisterFailHandler(Fail)
	RunSpecs(t, "templaterunner tests")
}

var _ = Describe("templaterunner package", func() {
	var (
		fakePromptManager *mocks.MockPromptManager
	)
	BeforeEach(func() {
		fakePromptManager = mocks.NewMockPromptManager(ctrl)
	})

	It("should return a template instance", func() {
		template := templaterunner.NewTemplateRunner(fakePromptManager)

		Expect(template).ToNot(BeNil())
	})

	It("should return a message from a 1 input template", func() {
		label := "label test:"
		template := templaterunner.NewTemplateRunner(fakePromptManager)
		yaml := `PROMPT:
- KEY: "TEST"
  LABEL: "label test"
TEMPLATE: "{{TEST}}"`

		fakePromptManager.EXPECT().ReadValue(label, "empty value", "").Return("ok").Times(1)
		messageFromTemplate, err := template.Run(yaml, nil)

		Expect(template).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(messageFromTemplate).To(Equal("ok"))
	})

	It("should return a message from a 1 input template with all options available filled", func() {
		label := "label test (test value):"
		defaultValue := "test value"
		template := templaterunner.NewTemplateRunner(fakePromptManager)
		yaml := `PROMPT:
- KEY: "TEST"
  LABEL: "label test"
  ERROR_LABEL: "error value"
  DEFAULT_VALUE: "test value"
TEMPLATE: "{{TEST}}"`

		fakePromptManager.EXPECT().ReadValue(label, "error value", defaultValue).Return("test value").Times(1)
		messageFromTemplate, err := template.Run(yaml, nil)

		Expect(template).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(messageFromTemplate).To(Equal(defaultValue))
	})

	It("should return a message from a 1 multi selection item template", func() {
		label := "label test"
		template := templaterunner.NewTemplateRunner(fakePromptManager)
		yaml := `PROMPT:
- KEY: "TEST"
  LABEL: "label test"
  OPTIONS:
  - VALUE: "opt1"
    DESC: "Test option1"
  - VALUE: "opt2"
    DESC: "Test option2"
TEMPLATE: "{{TEST}}"`

		options := []templaterunner.Options{
			{Value: "opt1", Description: "Test option1"},
			{Value: "opt2", Description: "Test option2"},
		}

		fakePromptManager.EXPECT().ReadValueFromList(label, options).Return("Test option1").Times(1)
		messageFromTemplate, err := template.Run(yaml, nil)

		Expect(template).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(messageFromTemplate).To(Equal("Test option1"))
	})

	It("should return a message from a 1 input template with extra variables", func() {
		label := "label test:"
		template := templaterunner.NewTemplateRunner(fakePromptManager)
		yaml := `PROMPT:
- KEY: "TEST"
  LABEL: "label test"
TEMPLATE: "{{TEST}}: {{SAMPLE}}"`
		injectedVariables := map[string]string{
			"SAMPLE": "sample value",
		}

		fakePromptManager.EXPECT().ReadValue(label, "empty value", "").Return("ok").Times(1)
		messageFromTemplate, err := template.Run(yaml, injectedVariables)

		Expect(template).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(messageFromTemplate).To(Equal("ok: sample value"))
	})

	It("should fail if parseYaml fail", func() {
		template := templaterunner.NewTemplateRunner(fakePromptManager)
		// this yaml fails because it is bad format
		yaml := `PROMPT:
	- KEY: "TEST"
  	  LABEL: "label test"`

		_, err := template.Run(yaml, nil)

		Expect(template).ToNot(BeNil())
		Expect(err).To(Not(BeNil()))
	})
})
