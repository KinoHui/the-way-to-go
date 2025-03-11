package stringio

import (
	"fmt"
	"io"
)

type StringReader struct {
	data string
	pos  int
}

// 实现 io.Reader 接口
func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.pos >= len(sr.data) {
		return 0, io.EOF // 达到输入末尾
	}

	n = copy(p, sr.data[sr.pos:]) // 从 data 中复制到 p
	sr.pos += n                   // 更新当前位置
	return n, nil                 // 返回读取的字节数和 nil 错误
}

func TestReader() {
	sr := &StringReader{data: "Hello, World!"}
	buf := make([]byte, 8) // 创建一个字节切片作为缓冲区

	for {
		n, err := sr.Read(buf) // 从 StringReader 读取数据
		if err == io.EOF {
			break // 到达输入末尾，退出循环
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Print(string(buf[:n])) // 打印读取的内容
	}
}

type StringWriter struct {
	data string
}

// 实现 io.Writer 接口
func (sw *StringWriter) Write(p []byte) (n int, err error) {
	sw.data += string(p) // 将字节切片转换为字符串并追加到 data
	return len(p), nil   // 返回实际写入的字节数和 nil 错误
}

func testWriter() {
	sw := &StringWriter{}
	msg := []byte("Hello, Go!")

	n, err := sw.Write(msg) // 向 StringWriter 写入数据
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Wrote %d bytes: %s\n", n, sw.data) // 输出写入的字节数和内容
}
