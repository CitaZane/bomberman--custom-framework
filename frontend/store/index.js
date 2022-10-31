const storeObj = {
  state: {
    monster:{
        type:0
    }
  },
  mutations:{
    updateMonster(state, monster){
        state.type = monster
    }
  },
  actions:{
    updateMonsterType({ state, commit }, type){
        let monster = state.monster;
        monster.type = type;
        commit("updateMonster", monster)
    }
  }
}

export default storeObj;