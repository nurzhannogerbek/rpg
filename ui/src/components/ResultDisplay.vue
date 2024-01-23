<template>
  <div class="result-display">
    <h2>Result:</h2>
    <p v-if="result">Number of Packs: {{ result.packs }}</p>
    <p v-if="result">Packs: {{ formattedPacks(result.packSizes) }}</p>
  </div>
</template>

<script>
export default {
  props: {
    result: Object,
  },
  methods: {
    formattedPacks(packSizes) {
      return packSizes.map(pack => `${pack.packSize} x ${pack.quantity}`).join(", ");
    },
  },
};
</script>

<style scoped>
.result-display {
  display: flex;
  flex-direction: column;
  align-items: center;
}

h2 {
  margin-bottom: 10px;
}
</style>
