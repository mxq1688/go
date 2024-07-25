package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt" // 导入内置 fmt 包

	"github.com/pion/webrtc/v3"
)

var g peerConnection

func main() { // main函数，是程序执行的入口
	fmt.Println("Hello 这是你第一个GO程序，开始进入GO世界吧") // 在终端打印
}

// 创建PeerConnection：使用Pion提供的API创建一个新的PeerConnection对象，用于处理与对等方的连接。
func connection() {
	config := webrtc.Configuration{}
	peerConnection, err := webrtc.NewPeerConnection(config)
}

// 获取本地媒体流：使用Pion的media API获取本地的音视频流，例如摄像头和麦克风。
func getUserMedia() {
	mediaDevices := media.GetUserMedia(constraints)
	localStream, err := mediaDevices.GetVideoTrack()
}

// 添加本地媒体流到PeerConnection：通过AddTrack方法将本地媒体流添加到PeerConnection对象中。
func addTrack() {
	peerConnection.AddTrack(localStream)
}

// 连接到对等方：使用Pion提供的信令方式与对等方建立连接。
// func createOffer() {
// 	offer, err := peerConnection.CreateOffer(nil)
// 	// 发送offer给对方...
// 	// 接收对方的answer...
// 	peerConnection.SetRemoteDescription(answer)

// }
