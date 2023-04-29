<template>
  <div
    class="max-w-sm rounded-lg overflow-hidden  bg-white border-gray-200 shadow-md hover:shadow-lg px-6 py-4 flex flex-col justify-center items-center cursor-pointer">

    <div class="font-bold text-xl mb-2 text-center  flex-auto">{{header}}</div>
    <p class="text-gray-700 text-base text-center bg-gray-200 rounded-full px-3 py-1 font-semibold w-32   flex-auto">
     {{description}}
    </p>


  </div>
</template>

<script>
export default {
  props: ['header','description'],
}
</script>

<style></style>