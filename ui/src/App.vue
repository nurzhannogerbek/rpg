<template>
  <div id="app">
    <div class="form-container">
      <h1>Pack Calculator</h1>
      <form @submit.prevent="calculatePacks">
        <div class="input-container">
          <label for="order">Items Ordered:</label>
          <input
              v-model="order"
              type="text"
              id="order"
              @input="clearResult; validateOrder"
          />
          <div v-if="orderError" class="error-message">{{ orderError }}</div>
        </div>
        <div class="input-container">
          <label for="packSizes">Pack Sizes:</label>
          <input
              v-model="packSizes"
              type="text"
              id="packSizes"
              placeholder="e.g., 250, 500, 1000"
              @input="clearResult; validatePackSizes"
          />
          <div v-if="packSizesError" class="error-message">{{ packSizesError }}</div>
        </div>
        <button type="submit" :disabled="!isValidInput">Calculate Packs</button>
      </form>
      <div v-if="result" class="result-container">
        <h2>Result:</h2>
        <div class="pack-item" v-for="pack in result.packs" :key="pack.pack_size">
          <p>{{ pack.quantity }} x {{ pack.pack_size }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      order: "",
      packSizes: "",
      result: null,
      orderError: null,
      packSizesError: null,
    };
  },
  computed: {
    isValidInput() {
      return !this.orderError && !this.packSizesError && this.order && this.packSizes;
    },
  },
  watch: {
    order: 'clearResult',
    packSizes: 'clearResult',
  },
  methods: {
    async calculatePacks() {
      try {
        const packSizesArray = this.packSizes.split(',').map(size => parseInt(size.trim()));
        const formattedPackSizes = packSizesArray.filter(size => size !== "");
        const formattedPayload = {
          order: parseInt(this.order),
          pack_sizes: formattedPackSizes,
        };
        const apiUrl = process.env.RPG_BACKEND_API_URL || "http://localhost:8080";
        const response = await axios.post(`${apiUrl}/calculate`, formattedPayload);
        this.result = response.data;
      } catch (error) {
        if (error.response) {
          console.error("Server responded with an error:", error.response.data);
        } else if (error.request) {
          console.error("No response received from the server.");
        } else {
          console.error("Error setting up the request:", error.message);
        }
      }
    },
    validateOrder() {
      const pattern = /^[1-9]\d*$/;
      this.orderError = pattern.test(this.order) ? null : "Please enter a valid positive integer.";
    },
    validatePackSizes() {
      const pattern = /^[0-9\s]*(?:,\s?[0-9\s]*)*$/;
      this.packSizesError = pattern.test(this.packSizes)
          ? null
          : "Please enter valid numbers separated by commas.";
    },
    clearResult() {
      this.result = null;
    },
  },
};
</script>

<style>
body {
  margin: 0;
  font-family: 'Arial', sans-serif;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

#app {
  text-align: center;
  color: #2c3e50;
}

.form-container {
  width: 400px;
  padding: 20px;
  border: 1px solid #bdc3c7;
  border-radius: 8px;
}

.input-container {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
  font-size: 14px;
  font-weight: bold;
  text-align: left;
}

input {
  width: calc(100% - 16px);
  padding: 8px;
  font-size: 14px;
  border: 1px solid #bdc3c7;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  background-color: #2ecc71;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:disabled {
  background-color: #bdc3c7;
  cursor: not-allowed;
}

.error-message {
  color: #e74c3c;
  font-size: 12px;
  margin-top: 5px;
}
</style>