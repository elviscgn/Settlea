<template>
  <div class="container">
    <div class="baseGame prevent-select">
      <ChatBox
        class="chat-box"
        :messages="messages"
        :sys-messages="sysMessages"
        :send-message="sendMessage"
      />
      <StatusBox
        class="status-box"
        :connection-status="connectionStatus"
        :ping="ping"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { Application } from "pixi.js";
import { init } from "../scripts/game";

import StatusBox from "./StatusBox.vue";
import ChatBox from "./ChatBox.vue";

import config from "@/config";
import { nanoid } from "nanoid";
import { MessageResponse } from "@/library/types";
import { ApiClient } from "@/library/api";

export default {
  name: "BaseGame",
  data() {
    return {
      app: new Application(),
      api: new ApiClient(config.apiUrl),
      connectionStatus: "Disconnected",
      ping: 0,
      username: "",
      lastPingTime: 0,
      wsConn: null as WebSocket | null,
      messages: [] as {
        author: string;
        message: string;
      }[],
      sysMessages: [] as {
        message: string;
      }[],
    };
  },
  components: {
    ChatBox,
    StatusBox,
  },
  async mounted() {
    await init(this);
    // await this.createRoom();
    this.connectWebsocket();

    window.addEventListener("beforeunload", () => {
      if (this.wsConn) {
        this.wsConn.close();
      }
    });
  },
  methods: {
    async createRoom() {
      const res = await this.api.post("ws/createRoom", {
        id: "testID",
        name: "testRoom",
      });
      console.log(res);
    },
    connectWebsocket() {
      const roomID = "testID";

      const storedUsername = localStorage.getItem("username");

      this.username = storedUsername || nanoid(15);
      if (!storedUsername) {
        localStorage.setItem("username", this.username);
      }

      const ws = new WebSocket(
        config.wsUrl +
          `ws/joinRoom/${roomID}?userId=${this.username}&username=${this.username}`
      );

      this.wsConn = ws;

      ws.onopen = () => {
        this.connectionStatus = "Connected";

        setInterval(() => {
          this.lastPingTime = Date.now();
          const pingMessage = {
            action: "ping",
            content: "",
            username: this.username,
            roomID: "testID",
          };
          ws.send(JSON.stringify(pingMessage));
        }, 5000);
      };

      ws.onmessage = (event) => {
        this.handleMessage(event.data);
      };

      // ws.onclose = () => {
      //   this.connectionStatus = "Disconnected";
      //   setTimeout(() => {
      //     this.connectWebsocket(); // Reconnect after 3 seconds
      //   }, 3000);
      // };

      ws.onerror = (error) => {
        console.error("WebSocket error:", error);
        this.connectionStatus = "Error";
      };
    },

    sendMessage(message: string) {
      if (this.wsConn && this.wsConn.readyState === WebSocket.OPEN) {
        const chatMessage = {
          action: "send_message",
          content: message,
          username: this.username,
          roomID: "testID",
        };

        this.wsConn.send(JSON.stringify(chatMessage));
      } else {
        console.warn("WebSocket is not connected.");
      }
    },

    handleMessage(data: string) {
      try {
        const message: MessageResponse = JSON.parse(data);

        switch (message.action) {
          case "pong":
            this.ping = Date.now() - this.lastPingTime;

            break;

          case "send_message":
            this.messages.push({
              author: message.username,
              message: message.content,
            });
            break;

          case "join_room":
            this.sysMessages.push({
              message: message.content,
            });
            break;

          case "leave_room":
            this.sysMessages.push({
              message: message.content,
            });
            break;

          default:
            console.log("Unimplemented action", message);
            break;
        }
      } catch (error) {
        console.error("Error parsing message", error);
      }
    },
  },
};
</script>

<style>
html,
body {
  margin: 0;
  padding: 0;

  overflow: hidden; /* Prevent scrolling */
}

.container {
  width: 100%;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  overflow: hidden; /* Lock content within the viewport */
}

.baseGame {
  background-color: #1199bb;
  width: 100%;
  height: 100%;
  position: relative;
}

.chat-box {
  height: 40%;
  position: absolute;
  bottom: 0;
  right: 0;
  z-index: 10;
}

.status-box {
  position: absolute;
  top: 0;
  left: 0;
  z-index: 10;
}

.prevent-select {
  -webkit-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
