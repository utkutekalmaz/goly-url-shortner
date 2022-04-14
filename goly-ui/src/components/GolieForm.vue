<script setup>
import { useGoliesStore } from "../stores/golies"
import { ref } from 'vue'

const golyStore = useGoliesStore()
const longURL = ref('')


const submit = () => { 
    console.log({asd:longURL.value});
    if (longURL.value.length) {
        golyStore.createGoly(longURL.value)
    }
 }
    
</script>

<template>
    <div class="flex gap-2" >
        <input v-model="longURL" class="w-auto px-2 py-4 mx-auto border-2 border-red-300 rounded-md" type="text" name="longUrl" placeholder="put your long url here" >
        <button @click="submit" class="w-auto px-2 py-4 mx-auto border-2 border-red-300 rounded-md cursor-pointer">shorten</button>
        {{longURL}}
    </div>
</template>
