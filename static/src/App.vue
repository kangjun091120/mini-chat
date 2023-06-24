<template>
  <div class="sidebar">
    <li v-for="site in messages">
      {{ site.text }}
    </li>
  </div>
  <input id="messageInput" type="text" placeholder="输入你的消息" />
  <input id="idSelect" type="text" placeholder="输入ID" />
  <button @click="sendMessage">发送消息</button>
</template>

<script>
export default {

  data() {
    return {
      messages: [
        { text: 'Google' },
        { text: 'Runoob' },
        { text: 'Taobao' }
      ]
    }
  },

  methods: {
    sendMessage() {
      let messageInput = document.getElementById("messageInput");
      let message = messageInput.value;
      var url = "http://localhost:81/send";
      var data = {
        id: '11',
        message: message
      };

      $.ajax({
        url: url,
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
        success: function (response) {
          console.log(response);
        },
      });
      messageInput.value = '';
    }
  },

  mounted() {
    let socket = new WebSocket("ws://localhost:81");
    let id = '11'
    socket.onopen = function (e) {
      console.log("[开启连接] 连接到服务器");
      if (typeof id === 'string' && id !== '') {
        socket.send(JSON.stringify({ id: id }));
      } else {
        console.log("无效的 ID");
      }
    };
    socket.onmessage = function (event) {
      console.log(`[服务器消息] 数据: ${event.data}`);
    };

    socket.onclose = function (event) {
      if (event.wasClean) {
        console.log(`[关闭连接] 连接已经关闭，代码=${event.code} 原因=${event.reason}`);
      } else {
        console.log('[关闭连接] 连接意外关闭');
      }
    };

    socket.onerror = function (error) {
      console.log(`[错误] ${error.message}`);
    };
  }
}
</script>


<style>
.sidebar {
  height: 100%;
  width: 10%;
  background-color: rgba(0, 0, 0, 0.8);
}
</style>