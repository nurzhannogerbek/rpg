<template>
  <div class="order-form">
    <label for="order">Enter the number of items:</label>
    <input v-model="order" type="number" id="order" />
    <label for="packSizes">Enter available pack sizes (comma-separated):</label>
    <input v-model="packSizes" type="text" id="packSizes" />
    <button @click="calculate">Calculate</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      order: 0,
      packSizes: '',
    };
  },
  methods: {
    calculate() {
      // Convert the comma-separated pack sizes to an array of integers.
      const packSizesArray = this.packSizes.split(',').map(size => parseInt(size.trim(), 10));

      // Make the API request to calculate packs.
      this.$emit("calculate", {order: this.order, packSizes: packSizesArray});
    },
  },
};
</script>

<style scoped>
.order-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
}

label {
  margin-bottom: 5px;
}

input {
  margin-bottom: 10px;
}

button {
  background-color: #4caf50;
  color: white;
  padding: 10px;
  border: none;
  cursor: pointer;
  border-radius: 5px;
}
</style>
