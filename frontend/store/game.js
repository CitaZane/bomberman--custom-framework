export default {
  state: {
    players: [],
    map: [],
    powerUps: [],
  },
  mutations: {
    updatePlayers(state, players) {
      state.players = players;
    },
    updateMap(state, map) {
      state.map = map;
    },

    updatePowerUps(state, powerUps) {
      state.powerUps = powerUps;
    },
  },
  actions: {
    registerPlayer({ state, commit }, player) {
      let players = state.players;
      players.push(player);
      commit("updatePlayers", players);
    },
  },
};
