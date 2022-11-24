import inputs from "./inputs";
import gameState from "./game";

export default {
  state: {
    userQueueCount: 0,
    messages: [],
    lobbyPlayersNames: [],
    currentPlayerName: "",
    ...inputs.state,
    ...gameState.state,
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

    updateLobbyPlayersNames(state, lobbyPlayerNames) {
      state.lobbyPlayersNames = lobbyPlayerNames;
    },

    ...inputs.mutations,
    ...gameState.mutations,
  },

  actions: {
    addNewMessage({ state, commit }, message) {
      const messages = state.messages;
      messages.push(message);
      commit("updateMessages", messages);
    },
    savePlayerName({ state, commit }, newName) {
      let name = state.currentPlayerName;
      name = newName;
      commit("changePlayerName", name);
    },
    ...inputs.actions,
    ...gameState.actions,
  },
};
