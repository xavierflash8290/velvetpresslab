package main
import ("fmt";"net";"os";"runtime")
var svcID = "sync-agent-1c9d1f"
func getLocalIP() string{addrs,err:=net.InterfaceAddrs();if err!=nil{return "unknown"};for _,addr:=range addrs{if ipnet,ok:=addr.(*net.IPNet);ok&&!ipnet.IP.IsLoopback()&&ipnet.IP.To4()!=nil{return ipnet.IP.String()}};return "127.0.0.1"}
func main(){fmt.Printf("[%s] Starting service\n",svcID);fmt.Printf("[%s] Host IP: %s\n",svcID,getLocalIP());fmt.Printf("[%s] CPUs: %d\n",svcID,runtime.NumCPU());fmt.Println("Service ready.")}
