package books_test

//import (
//	. "github.com/YunCaiCaoYuan/test_golang/library/books"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//)
//
//var _ = Describe("Book", func() {
//	var (
//		longBook  Book
//		shortBook Book
//		smallBook Book
//	)
//
//	//BeforeEach是每个测试例执行前执行该段代码
//	//比如创建数据库连接就可以使用BeforeEach ，每个BeforeEach只在当前域内起作用。
//	//执行顺序是同一层级的顺序执行，不同层级的从外层到里层以此执行。类似与 全局变量和局部变量的区别
//	BeforeEach(func() {
//		longBook = Book{
//			Title:  "Les Miserables",
//			Author: "Victor Hugo",
//			Pages:  1488,
//		}
//
//		shortBook = Book{
//			Title:  "Fox In Socks",
//			Author: "Dr. Seuss",
//			Pages:  240,
//		}
//		smallBook = Book{
//			Title:  "go program",
//			Author: "caochunhui",
//			Pages:  20,
//		}
//	})
//
//
//	Describe("Categorizing book length", func() {
//		Context("With more than 300 pages", func() {
//			//It是测试例的基本单位，即It包含的代码就算一个测试用例
//			//It可以理解为测试代码的最小单元
//			It("should be a novel", func() {
//				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
//			})
//		})
//
//		Context("With fewer than 300 pages", func() {
//			It("should be a short story", func() {
//				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
//			})
//		})
//		Context("With fewer than 100 pages", func() {
//			It("should be a small story", func() {
//				Expect(smallBook.CategoryByLength()).To(Equal("SMALL STORY"))
//			})
//		})
//	})
//})
