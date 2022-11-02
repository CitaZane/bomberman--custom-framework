export default {
    state: {
        monster: {
            type: 0,
            x: 0,
            y: 0,
        },
        bomb: {
            drop: false
        },
        inputs: {}
    },
    mutations: {
        updateMonster(state, monster) {
            state.type = monster
        },
        updateInputs(state, inputs) {
            state.inputs = inputs
        },
        updateBomb(state, bomb) {
            state.bomb = bomb
        }
    },
    actions: {
        updateMonsterType({ state, commit }, type) {
            let monster = state.monster;
            monster.type = type;
            commit("updateMonster", monster)
        },
        updateMonsterX({ state, commit }, x) {
            let monster = state.monster;
            monster.x = x;
            commit("updateMonster", monster)
        },
        updateMonsterY({ state, commit }, y) {
            let monster = state.monster;
            monster.y = y;
            commit("updateMonster", monster)
        },
        registerKeyUp({ state, commit }, key) {
            let inputs = state.inputs;
            inputs[key] = false
            commit('updateInputs', inputs)
        },
        registerKeyDown({ state, commit }, key) {
            let inputs = state.inputs;
            inputs[key] = true
            commit('updateInputs', inputs)
        },
        updateBombDrop({ state, commit }, drop) {
            let bomb = state.bomb;
            bomb.drop = drop;
            commit("updateBomb", bomb)
        },
    }
}
