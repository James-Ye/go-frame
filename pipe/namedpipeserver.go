package pipe

// var ServerSocket net.Conn
// var lockSend sync.Mutex
// var pipeFile = "\\\\.\\pipe\\wtreelee_pipe"

// //注册函数，一个函数对应一个口令
// func register() map[string]returnFunc {
// 	return map[string]returnFunc{
// 		"entry": pipeEntry, // 获取入口数据
// 		"main":  startMain, // 进入
// 		"load":  load,      // 加载
// 	}
// }

// /**
// 开启管道通信监听
// 通过一个命名管道文件 wtreelee_pipe 进行管道通信进行数据交互
// 通过封装数据交互：方法&&&参数
// */
// func OpenPipeServer() {
// 	log.Println("开启管道通信")
// 	l, err := gw.ListenPipe(pipeFile, &config)
// 	if err != nil {
// 		log.Println("管道通信，开启管道错误")
// 		log.Println(err)
// 		return
// 	}
// 	for {
// 		conn, err := l.Accept()
// 		if err != nil {
// 			log.Println("管道通信，开启监听错误")
// 			log.Println(err)
// 			continue
// 		}
// 		go doServer(conn)
// 	}
// }

// func doServer(conn net.Conn) {
// 	ServerSocket = conn
// 	var cmdChan = make(chan string)  //接收请求数据
// 	var stopChan = make(chan string) //发送、接收一起停止
// 	var methods = register()
// 	go func() {
// 		for {
// 			cmd := <-cmdChan
// 			var message []byte
// 			if len(cmd) > 0 {
// 				log.Println("管道通信，接收到数据")
// 				// 口令格式和数据格式: cmd&&&json
// 				data := map[string]interface{}{}
// 				if strings.Contains(cmd, "&&&") {
// 					params := strings.Split(cmd, "&&&")
// 					cmd = params[0]
// 					log.Println("read ", cmd+" ", params[1])
// 					err := json.Unmarshal([]byte(params[1]), &data)
// 					if err != nil {
// 						log.Println("read2:", err)
// 						stopChan <- err.Error()
// 						break
// 					}
// 				}
// 				cmdFunc, ok := methods[cmd]
// 				if !ok {
// 					log.Println("socket", "没有这个函数", cmd)
// 					stopChan <- fmt.Errorf("没有这个函数" + cmd).Error()
// 					break
// 				}
// 				message, _ = json.Marshal(cmdFunc(data))
// 				//log.Println(string(message))
// 				err := sendMessage(fmt.Sprintf("%s&&&%s", cmd, string(message)))
// 				if err != nil {
// 					log.Println("write2:", err)
// 					stopChan <- err.Error()
// 					break
// 				}
// 				stopChan <- ""
// 			}
// 		}
// 	}()

// 	for {
// 		r := bufio.NewReader(conn)
// 		msg, err := r.ReadString('\n')
// 		if err != nil {
// 			log.Println("readErr", err)
// 			return
// 		}
// 		log.Println("msg:", msg)
// 		cmdChan <- msg
// 	}
// }

// /**
// 管道通信:发送数据
// */
// func sendMessage(message string) error {
// 	lockSend.Lock()
// 	defer lockSend.Unlock()
// 	if ServerSocket != nil {
// 		_, err := fmt.Fprintln(ServerSocket, message)
// 		if err != nil {
// 			log.Println("sendMessage:", err)
// 			return err
// 		}
// 	}
// 	return nil
// }
