package websocket



//func Run() {
//	go src.ReceivingResults(ResponseRsCh) //统计处理
//	WgTask.Add(1)
//	LaunchWebsocket()
//}

//func LaunchWebsocket() {
//
//	for i := 1; i <= userNum; i++ {
//		WgUser.Add(1)
//		go func(i int) {
//
//			defer func() {
//				WgUser.Done()
//			}()
//
//			conn, err := StartWsConn(wsUrl)
//
//			if err != nil {
//				fmt.Println(fmt.Sprintf("websocket err --------->%d", i))
//			} else {
//				defer func() {
//					err := conn.Close()
//					if err != nil {
//						fmt.Println(fmt.Sprintf("关闭长连接时出现错误：%s", err.Error()))
//					}
//				}()
//
//				timer := time.NewTimer(1 * time.Second) //一秒后激活时间
//				n := 0
//
//				for {
//					select {
//					case <-timer.C:
//						timer.Reset(1 * time.Second) //重置倒计时
//						n++
//
//						WebSocketRequest(conn, src.WsRequestData(&src.WsRequest{
//							Id:   strconv.Itoa(n),
//							Path: "/",
//							Data: map[string]string{"queue": strconv.Itoa(n), "user": strconv.Itoa(i)},
//						}), ResponseRsCh)
//
//						if n >= userRunNum {
//							timer.Stop()                //到达指定次数结束时间
//							time.Sleep(2 * time.Second) //让信息处理缓一会儿
//							return
//						}
//					}
//				}
//			}
//		}(i)
//	}
//
//	WgUser.Wait()
//	close(ResponseRsCh)
//	WgTask.Wait()
//	fmt.Println("-------success-------")
//}
