package fs_test

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/walmartdigital/commit-message-formatter/fs"
	"github.com/walmartdigital/commit-message-formatter/mocks"
)

var ctrl *gomock.Controller

func TestAll(t *testing.T) {
	ctrl = gomock.NewController(t)
	defer ctrl.Finish()

	RegisterFailHandler(Fail)
	RunSpecs(t, "fs tests")
}

var _ = Describe("fs package", func() {
	var (
		fakeVFS *mocks.MockVFS
	)

	BeforeEach(func() {
		fakeVFS = mocks.NewMockVFS(ctrl)
	})

	It("should return a file from VirtualFS", func() {
		fakeVFS.EXPECT().ReadFile(gomock.Any().String()).Return([]byte{}, nil).Times(1)

		vfs := fs.NewFs(fakeVFS)
		file, err := vfs.GetFileFromVirtualFS(gomock.Any().String())

		Expect(vfs).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(file).To(Equal(""))
	})

	It("should not return a file from VirtualFS when readFile fails", func() {
		fakeVFS.EXPECT().ReadFile(gomock.Any().String()).Return([]byte(""), errors.New(gomock.Any().String())).Times(1)

		vfs := fs.NewFs(fakeVFS)
		file, err := vfs.GetFileFromVirtualFS(gomock.Any().String())

		Expect(vfs).ToNot(BeNil())
		Expect(file).To(Equal(""))
		Expect(err).To(Equal(errors.New(fs.GetFileFromVirtualFSError)))
	})

	It("should return a file from user FS", func() {

		vfs := fs.NewFs(fakeVFS)
		folder, _ := os.Getwd()
		filePath := folder + "/fs.go"
		file, err := vfs.GetFileFromFS(filePath)

		Expect(vfs).ToNot(BeNil())
		Expect(err).To(BeNil())
		Expect(file).ToNot(Equal(""))
	})
})
