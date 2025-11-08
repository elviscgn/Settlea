<template>
  <div class="messages-container">
    <div v-for="(msg, index) in messages" :key="index" class="message-item">
      <div v-if="msg.type === 'user'" class="user-message">
        <img
          src="/assets/icons/profile_icon.svg"
          alt="Profile Icon"
          class="icon"
        />
        <span class="text-content">
          <b>{{ msg.author }}</b
          >: {{ msg.message }}
        </span>
      </div>
      <div v-else-if="msg.type === 'system'" class="sys-message">
        {{ msg.message }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { PropType } from "vue";

export default {
  name: "TabChat",
  props: {
    messages: {
      type: Array as PropType<
        { author?: string; message: string; type: string }[]
      >,
      required: true,
    },
  },
};
</script>

<style>
.user-message {
  display: flex;
  align-items: flex-start;
}

.sys-message {
  display: flex;
  align-items: flex-start;
  color: #555;
}

.icon {
  width: 24px;
  height: 24px;
  margin-right: 5px;
}

.text-content {
  padding-top: 2px;
  flex: 1;
  word-wrap: break-word;
  white-space: normal;
}

.messages-container {
  padding: 5px;
}

/* .user-message {
  margin-bottom: 8px;
  padding: 5px;
  background-color: #f9f9f9;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
} */

/* .sys-message {
  margin-bottom: 8px;
  padding: 5px;
  background-color: #e0e0e0;
  color: #555;
  font-style: italic;
  border-radius: 4px;
} */
</style>
