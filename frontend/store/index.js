const storeObj = {
  state: {
    monster:{
        type:0,
    },
    inputs:{}
  },
  mutations:{
    updateMonster(state, monster){
        state.type = monster
    },
    updateInputs(state, inputs){
        state.inputs = inputs
    }
  },
  actions:{
    updateMonsterType({ state, commit }, type){
        let monster = state.monster;
        monster.type = type;
        commit("updateMonster", monster)
    },
    registerKeyUp({ state, commit }, key){
      let inputs = state.inputs;
      inputs[key] = false
      commit('updateInputs', inputs)
    },
    registerKeyDown({ state, commit }, key){
      let inputs = state.inputs;
      inputs[key] = true
      commit('updateInputs', inputs)
    }
  }
}

export default storeObj;