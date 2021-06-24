package linkedlist_test

import (
	"testing"

	linkedlist "github.com/golanshabi/LinkedList-golang"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ListLen = 5

func TestAdder(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adder Suite")
}

var _ = Describe("Linked List", func() {
	IndexList := [...]int{ListLen - 1, 2, 0}
	Describe("TestPushBack", func() {
		It("test the pushBack function directly", func() {
			list := linkedlist.NewLinkedList()
			pushBackToList(&list)
			cur := list.Head()
			for i := 0; cur != nil; cur, i = cur.GetNext(), i+1 {
				Expect(cur.GetVal()).To(Equal(i))
			}
		})
	})

	Describe("TestPushFront", func() {
		It("test the pushFront function directly", func() {
			list := linkedlist.NewLinkedList()
			pushFrontToList(&list, ListLen)
			cur := list.Head()
			for i := ListLen - 1; cur != nil; cur, i = cur.GetNext(), i-1 {
				Expect(cur.GetVal()).To(Equal(i))
			}
		})
	})

	Describe("TestPopBack", func() {
		It("test the PopBack function", func() {
			list := linkedlist.NewLinkedList()
			pushBackToList(&list)
			for i := ListLen - 1; i <= 0; i-- {
				Expect(list.PopBack()).To(Equal(i))
				Expect(list.Len()).To(Equal(i + 1))
			}
		})
	})

	Describe("TestPopFront", func() {
		It("test the PopFront function", func() {
			list := linkedlist.NewLinkedList()
			pushBackToList(&list)
			for i := 0; i < ListLen; i++ {
				Expect(list.PopFront()).To(Equal(i))
				Expect(list.Len()).To(Equal(ListLen - i - 1))
			}
		})
	})

	Describe("TestPeek", func() {
		It("test the Peek function", func() {
			list := linkedlist.NewLinkedList()
			pushBackToList(&list)
			for i := 0; i < ListLen; i++ {
				Expect(list.Peek(i)).To(Equal(i))
			}
		})
	})

	Describe("TestEraseList", func() {
		It("test the eraseList function", func() {
			list := linkedlist.NewLinkedList()
			pushBackToList(&list)
			err := list.EraseIndexList(IndexList[0:3])
			Expect(err).To(BeNil())
			Expect(list.Len()).To(Equal(2))
			cur := list.Head()
			for i := 0; i < ListLen; i++ {
				if !contains(IndexList[0:3], i) {
					Expect(cur.GetVal()).To(Equal(i))
					cur = cur.GetNext()
				}
			}
		})
	})

	Describe("TestError", func() {
		It("asserts that the linked list returns error for illegal input", func() {
			list := linkedlist.NewLinkedList()
			_, err := list.PopBack()
			Expect(err).NotTo(BeNil())
			_, err = list.PopFront()
			Expect(err).NotTo(BeNil())
			_, err = list.Peek(-1)
			Expect(err).NotTo(BeNil())
			pushBackToList(&list)
			_, err = list.Peek(ListLen)
			Expect(err).NotTo(BeNil())
			err = list.EraseIndexList([]int{-1})
			Expect(err).NotTo(BeNil())
			err = list.EraseIndexList([]int{0, 0})
			Expect(err).NotTo(BeNil())
			err = list.EraseIndexList([]int{ListLen})
			Expect(err).NotTo(BeNil())
		})
	})
})

func pushBackToList(list *linkedlist.LinkedList) {
	for i := 0; i < ListLen; i++ {
		list.PushBack(i)
		Expect(list.Len()).To(Equal(i + 1))
		Expect(list.Tail().GetVal()).To(Equal(i))
	}
}

func pushFrontToList(list *linkedlist.LinkedList, amountToPush int) {
	for i := 0; i < amountToPush; i++ {
		list.PushFront(i)
		Expect(list.Len()).To(Equal(i + 1))
		Expect(list.Head().GetVal()).To(Equal(i))
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
