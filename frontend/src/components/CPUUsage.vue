<template>
  <div class="container">
    <h1>{{message}}</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: " "
    };
  },
  mounted: () => {
    wails.events.on('cpu_usage', cpu_usage => {
      this.message = cpu_usage.avg
    })
  },/*
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.Stats.GetCPUUsage().then(cpu_usage => {
        self.message = cpu_usage.avg;
      });
    }
  }*/
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
