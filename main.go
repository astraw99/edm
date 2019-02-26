package main

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
	batchLength = 10
	outerWg     sync.WaitGroup
)

func main() {
	startTime := time.Now().UnixNano()

	for i := 1; i <= 5; i++ {
		filename := "./task/edm" + strconv.Itoa(i) + ".txt"
		start := 60

		outerWg.Add(1)
		go func() {
			err := RunTask(filename, start, batchLength)
			if err != nil && err != io.EOF {
				log.Fatalln(err)
			}
		}()
	}

	// main 阻塞等待goroutine执行完成
	outerWg.Wait()

	fmt.Println("finished all tasks.")

	endTime := time.Now().UnixNano()
	fmt.Println("Total cost(ms):", (endTime-startTime)/1e6)
}

// RunTask 单任务
func RunTask(filename string, start, length int) (retErr error) {
	for {
		isFinish := make(chan bool, 1)
		readLine, err := ReadLines(filename, start, length, isFinish)
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

		// 等待一批完成才进入下一批
		fmt.Println("current line:", readLine)
		start += length
		<-isFinish

		// 关闭channel，释放资源
		close(isFinish)
	}

	outerWg.Done()

	return retErr
}

// ReadLines 读取指定行数据
func ReadLines(filename string, start, length int, isFinish chan bool) (line int, retErr error) {
	fmt.Println("current file:", filename)

	// 控制每一批发完再下一批
	var wg sync.WaitGroup

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
			go func() {
				err := SendEmail(line, &wg)
				if err != nil {
					log.Fatalln(err)
				}
			}()
			if startLine == endLine {
				isFinish <- true
				break
			}
		}

		startLine++
	}

	wg.Wait()

	return startLine, retErr
}

// SendEmail 模拟邮件发送
func SendEmail(email string, wg *sync.WaitGroup) error {
	defer wg.Done()

	time.Sleep(time.Second * 1)
	fmt.Println(email)

	return nil
}
