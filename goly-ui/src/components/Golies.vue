<script setup>
    import { onMounted, watch } from "vue-demi"
    import { useGoliesStore } from "../stores/golies"
    const goliesStore = useGoliesStore()

    console.log({...goliesStore.golies});

    watch(goliesStore,(state) => { 
        console.log({...state.golies[0]});
    })

    onMounted(() => {
        goliesStore.getAll()
    })
</script>

<template>
    <div class="flex flex-col gap-2" >
        <div class="p-4 text-xl border-2 border-red-300 rounded-md" v-for="(item, index) in goliesStore.golies" :key="index" >
            <a :href="item.redirect" class="block " >redirect: {{item.redirect}}</a>
            <a :href="item.goly"  class="block " >goly: {{item.goly}}</a>
            <span class="" >used count: {{item.clicked}}</span>
        </div>
    </div>
</template>
