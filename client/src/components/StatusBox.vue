<template>
  <div class="status-box">
    <p>
      Ping: <span :class="ping > 100 ? 'red' : 'green'">{{ ping }} ms</span>
    </p>
    <p>
      Status:
      <span :class="connectionStatus === 'Connected' ? 'green' : 'red'">{{
        connectionStatus
      }}</span>
    </p>
  </div>
</template>

<script lang="ts">
export default {
  name: "StatusBox",
  props: {
    ping: {
      type: Number,
      required: true,
    },
    connectionStatus: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
.status-box {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  padding: 0;
  color: #177e89;
  width: 20%;
  height: 10%;
  background-color: #f8f2dc;
  border: 4px solid #183a37;
  border-left: 0px;
  font-family: "Londrina Solid", sans-serif;
  font-size: 24px;
  border-top-right-radius: 10px;
  border-bottom-right-radius: 10px;
  text-align: left;
}

.status-box p {
  margin: 0;
  padding-left: 10px;
}

.red {
  color: #ff0000;
}

.green {
  color: #27ba1d;
}
</style>
