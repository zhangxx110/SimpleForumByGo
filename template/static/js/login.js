window.onload = main;
var socket;
function main() {
    var oBtn = document.getElementById('btn1');
    oBtn.onclick = OnButton1;
    //socket
    initsocket();
    var osocketBtn = document.getElementById('socket_btn');
    osocketBtn.onclick = sendMsg;
}

function OnButton1() {
	var oUserNameTxt = document.getElementById('usernametxt');
	var oPasswTxt = document.getElementById('passwtxt');
	var oTokenTxt = document.getElementById('tokenId');
	alert("oUserNameTxt:"+oUserNameTxt.value+" "+oPasswTxt.value+" "+oTokenTxt.value)
    var xhr = new XMLHttpRequest();
    xhr.open('post', '/ajax', true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) { 
            if (xhr.status == 200) {
                var oTxt = document.getElementById('txt1');
                oTxt.value = xhr.responseText;
				jumpPage('/upload')
            }
        }
    }
    //post请求要自己设置请求头  
    xhr.setRequestHeader("Content-Type","application/x-www-form-urlencoded");  
     //发送数据，开始与服务器进行交互  
     //要用send发送消息，必须使用post方式，并且设置  RequestHeader
    xhr.send("username="+oUserNameTxt.value+"&password="+oPasswTxt.value+"&token="+oTokenTxt.value);
    
}
function jumpPage(page){       
    self.location=page;
}
function sendMsg(){
	var oUserNameTxt = document.getElementById('socket_txt');
	socket.send(oUserNameTxt.value)
	alert("send message:"+oUserNameTxt.value+" success")
}
function initsocket(){	
//创建一个Socket实例
socket = new WebSocket('ws://localhost:8888'); 

// 打开Socket 
socket.onopen = function(event) { 

  // 发送一个初始化消息
  socket.send('I am the client and I\'m listening!'); 
};
  // 监听消息
  socket.onmessage = function(event) { 
    console.log('Client received a message',event); 
  }; 

  // 监听Socket的关闭
  socket.onclose = function(event) { 
    console.log('Client notified socket has closed',event); 
  }; 

  // 关闭Socket.... 
  //socket.close() 

}