<template>
  <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" :for="data">
        {{data}}
      </label>
      <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" :id="data" type="text" :placeholder="data">
    </div>
</template>

<script>
export default {
    props: ['data'],
}
</script>

<style>

</style>