<template>
  <div class="message-item">
    <img src="/assets/icons/profile_icon.svg" alt="Profile Icon" class="icon" />
    <span class="text-content"> {{ author }}: {{ message }} </span>
  </div>
</template>

<script lang="ts">
export default {
  name: "MessageItem",
  props: {
    author: {
      type: String,
      default: "Author",
    },
    message: {
      type: String,
      default: "Message",
    },
  },
};
</script>

<style scoped>
.message-item {
  display: flex;
  align-items: flex-start;
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
</style>
