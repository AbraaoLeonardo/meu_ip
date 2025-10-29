<script setup></script>

<template>
  <div id="app">
    <h1>{{ message }}</h1>
    <button @click="fetchMessage">Obter Mensagem</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
  data() {
    return {
      message: ''
    };
  },
  methods: {
    async fetchMessage() {
      try {
        const response = await axios.get('http://localhost:8080/');
        this.message = response.data.content;
      } catch (error) {
        console.error('Erro ao buscar a mensagem:', error);
      }
    }
  }
}
</script>


<style scoped></style>
