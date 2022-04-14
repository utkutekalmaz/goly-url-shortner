import { defineStore } from 'pinia'
import { createGolyService, getAllGolies } from '../api/golies'

// export const useCounterStore = defineStore({
//   id: 'counter',
//   state: () => ({
//     counter: 0
//   }),
//   getters: {
//     doubleCount: (state) => state.counter * 2
//   },
//   actions: {
//     increment() {
//       this.counter++
//     }
//   }
// })
export const useGoliesStore = defineStore({
  id: 'golies',
  state: () => ({
    golies: []
  }),
  getters: {},
  actions: {
    getAll() {
      console.log("GET GOLY CALLED");
      getAllGolies().then( g => {
        console.log("GET GOLIES RESP");
        this.golies = g
      } )
    },
    createGoly(url){
      createGolyService(url).then(
        (r) => { 
          this.getAll()
         }
      )
    }
  }
})
