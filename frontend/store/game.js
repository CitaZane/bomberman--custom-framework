export default {
  state: {
    players: [],
    map: [],
    powerUps: [],
    explosionTime: {},
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
    updateExplosionTime(state, explosionTime) {
      state.explosionTime = explosionTime;
    },
  },
  actions: {
    addStartTime({ state, commit }, { time, explosionId }) {
      let explosionTime = state.explosionTime;
      explosionTime[explosionId] = time;
      commit("updateExplosionTime", explosionTime);
    },
    removeStartTime({ state, commit }, explosionId) {
      let explosionTime = state.explosionTime;
      delete explosionTime[explosionId];
      commit("updateExplosionTime", explosionTime);
    },
  },
};
