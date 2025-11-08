<template>
  <div id="chat-container">
    <nav>
      <button
        class="nav-button"
        @click="activeTab = 'chat'"
        :class="{ active: activeTab === 'chat' }"
      >
        Chat
      </button>
      <button
        class="nav-button"
        @click="activeTab = 'log'"
        :class="{ active: activeTab === 'log' }"
      >
        Logs (10)
      </button>
    </nav>

    <div class="chat-wrapper">
      <div v-if="activeTab === 'chat'" class="tab-content">
        <TabChat :messages="allMessages" />
      </div>
      <div v-else-if="activeTab === 'log'" class="tab-content">
        <TabLog />
      </div>

      <div class="chat-input" v-if="activeTab === 'chat'">
        <input
          v-model="message"
          @keydown.enter="sendMessageToServer"
          type="text"
          placeholder="Type a message..."
          class="input-text"
        />
        <button @click="sendMessageToServer" class="send-button">
          <i class="fa-solid fa-paper-plane"></i>
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, PropType, computed } from "vue";
import TabChat from "./TabChat.vue";
import TabLog from "./TabLog.vue";

export default {
  name: "ChatBox",
  components: {
    TabChat,
    TabLog,
  },
  props: {
    sendMessage: {
      type: Function as PropType<(message: string) => void>,
      required: true,
    },
    messages: {
      type: Array as PropType<{ author: string; message: string }[]>,
      required: true,
    },
    sysMessages: {
      type: Array as PropType<{ message: string }[]>,
      required: true,
    },
  },
  setup(props) {
    const activeTab = ref("chat");
    const message = ref("");

    const sendMessageToServer = () => {
      if (message.value.trim()) {
        props.sendMessage(message.value);
        message.value = ""; // Clear input after sending
      }
    };

    const allMessages = computed(() => {
      const sysMessages = props.sysMessages.map((msg) => ({
        ...msg,
        type: "system",
      }));
      const userMessages = props.messages.map((msg) => ({
        ...msg,
        type: "user",
      }));

      const combinedMessages = [];
      const maxLength = Math.max(userMessages.length, sysMessages.length);

      for (let i = 0; i < maxLength; i++) {
        if (i < sysMessages.length) {
          combinedMessages.push(sysMessages[i]);
        }
        if (i < userMessages.length) {
          combinedMessages.push(userMessages[i]);
        }
      }

      return combinedMessages;
    });

    return {
      activeTab,
      message,
      sendMessageToServer,
      allMessages,
    };
  },
};
</script>

<style scoped>
#chat-container {
  margin-top: 20px;
  border-left: 4px solid #183a37;
  border-top: 4px solid #183a37;
  border-right: 4px solid #183a37;
  border-radius: 10px 0px 0px 0px;
  background-color: #f8f2dc;
  width: 20%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  height: 50vh;
}

nav {
  display: flex;
}

.nav-button {
  width: 50%;
  padding: 10px;
  font-size: 20px;
  cursor: pointer;
  border: 0px;
  text-align: left;
  color: #183a37;
  transition: background-color 0.2s ease;
  font-family: "Londrina Solid", sans-serif;
  background-color: #e9d690;

  border-bottom: 4px solid #183a37;
}

.nav-button:hover {
  background-color: #f5eccc;
}

.nav-button.active {
  background-color: #f8f2dc;
  border-bottom: 0px;
}

.nav-button.active:first-child {
  border-top-left-radius: 10px;
  border-right: 4px solid #183a37;
  /* border-bottom: 0px; */
}

.nav-button.active:last-child {
  border-top-left-radius: 0%;
  border-left: 4px solid #183a37;
  /* border-bottom: 0px; */
}

.chat-wrapper {
  display: flex;
  flex-direction: column;
  flex: 1; /* Allow it to grow and fill the available space */
  overflow: hidden; /* Prevent overflow */
}

.tab-content {
  padding-left: 10px;
  padding-top: 5px;
  padding-right: 5px;
  background-color: #f8f2dc;
  font-family: "Outfit", sans-serif;
  height: 100%; /* Take full height */
  overflow-y: auto; /* Allow scrolling */
}

.chat-input {
  display: flex;
  padding: 1px;
  background-color: #f5eccc;
}

.input-text {
  font-family: "Outfit", sans-serif;
  width: 85%;
  margin-left: 5px;
  padding-left: 10px;
  border: none;
  font-size: 16px;
  outline: none;
  border-radius: 10px;
  border: 2px solid #183a37;
  color: #177e89;
}

.input-text::placeholder {
  color: #177e89;
}

.send-button {
  width: 15%;
  padding: 10px;
  margin-left: 5px;
  margin-right: 5px;
  background-color: #183a37;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.send-button:hover {
  background-color: #145a50;
}
</style>
