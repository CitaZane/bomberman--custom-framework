import gameState from "./game"

export default {
  state: {
    userQueueCount: 0,
    messages: [],
    currentPlayerName: "",
    ...gameState.state
  },

  mutations: {
    updateUserQueueCount(state, count) {
      state.userQueueCount = count;
    },
    updateMessages(state, messages) {
      state.messages = messages;
    },
    changePlayerName(state, name) {
      state.currentPlayerName = name;
    },
    ...gameState.mutations
  },

  actions: {
    addNewMessage({ state, commit }, message) {
      const messages = state.messages;
      messages.push(message);
      commit("updateMessages", messages);
    },
    savePlayerName({ state, commit }, newName) {
      let name = state.currentPlayerName;
      name = newName
      commit("changePlayerName", name);
    },

    ...gameState.actions
  },
};
