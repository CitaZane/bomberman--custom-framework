export default {
  state: {
    userQueueCount: 0,
    messages: [],
  },
  mutations: {
    updateUserQueueCount(state, count) {
      state.userQueueCount = count;
    },
    updateMessages(state, messages) {
      state.messages = messages;
    },
  },
  actions: {
    addNewMessage({ state, commit }, message) {
      const messages = state.messages;
      messages.push(message);
      commit("updateMessages", messages);
    },
  },
};
