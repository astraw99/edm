package main

/*
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	batchLength = 20
	wg          sync.WaitGroup
	finish      = make(chan bool)
)

func main() {
	startTime := time.Now().UnixNano()

	for i := 1; i <= 3; i++ {
		filename := "./task/edm" + strconv.Itoa(i) + ".txt"
		start := 60

		go RunTask(filename, start, batchLength)
	}

	// main 阻塞等待goroutine执行完成
	fmt.Println(<-finish)

	fmt.Println("finished all tasks.")

	endTime := time.Now().UnixNano()
	fmt.Println("Total cost(ms):", (endTime-startTime)/1e6)
}

// 单任务
func RunTask(filename string, start, length int) (retErr error) {
	for {
		readLine, err := ReadLines(filename, start, length)
		if err == io.EOF {
			fmt.Println("Read EOF:", filename)
			retErr = err
			break
		}
		if err != nil {
			fmt.Println(err)
			retErr = err
			break
		}

		fmt.Println("current line:", readLine)

		start += length

		// 等待一批完成才进入下一批
		wg.Wait()
	}

	//wg.Wait()
	finish <- true

	return retErr
}

// 读取指定行数据
func ReadLines(filename string, start, length int) (line int, retErr error) {
	fmt.Println("current file:", filename)

	fileObj, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()

	// 跳过开始行之前的行-scanner方式
	//startLine := 1
	//scanner := bufio.NewScanner(fileObj)
	//for scanner.Scan() {
	//	if startLine > start && startLine <= start + length {
	//		line := scanner.Text()
	//		fmt.Println(line)
	//	}
	//	startLine++
	//}

	// 跳过开始行之前的行-ReadString方式
	startLine := 1
	endLine := start + length
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err == io.EOF {
			fmt.Println("Read EOF:", filename)
			retErr = err
			break
		}
		if err != nil {
			log.Fatal(err)
			retErr = err
			break
		}

		if startLine > start && startLine <= endLine {
			wg.Add(1)
			// go并发执行
			go SendEmail(line)
			if startLine == endLine {
				break
			}
		}

		startLine++
	}

	return startLine, retErr
}

// 模拟邮件发送
func SendEmail(email string) error {
	defer wg.Done()

	time.Sleep(time.Second * 1)
	fmt.Println(email)

	return nil
}
*/
