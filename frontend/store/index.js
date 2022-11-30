import inputs from "./inputs";
import gameState from "./game";

export default {
  state: {
    userQueueCount: 0,
    messages: [],
    lobbyPlayersNames: [],
    timer: 0,
    gameTimerActive: false,
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
    updateTimer(state, timer) {
      state.timer = timer;
    },
    updateGameTimer(state, timer) {
      state.gameTimerActive = timer;
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
    initializeGameTimer({commit}){
      commit("updateGameTimer", true);
      commit("updateTimer", 10);
    },
    ...inputs.actions,
    ...gameState.actions,
  },
};
