export default {
    state: {
        inputs: {}
    },
    mutations: {
        updateInputs(state, inputs) {
            state.inputs = inputs
        }
    },
    actions: {
        registerKeyUp({ state, commit }, key) {
            let inputs = state.inputs;
            inputs[key] = false
            commit('updateInputs', inputs)
        },
        registerKeyDown({ state, commit }, key) {
            let inputs = state.inputs;
            inputs[key] = true
            commit('updateInputs', inputs)
        }
    }
}
