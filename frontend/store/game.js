export default {
  state: {
    players: [],
    map:[],
    explosions:[],
  },
  mutations: {
    updatePlayers(state, players) {
      state.players = players;
    },
    updateMap(state, map) {
      state.map = map;
    },
    updateExplosions(state, explosions){
      state.explosions = explosions
    }
  },
  actions: {
    registerPlayer({ state, commit }, player) {
      let players = state.players;
      players.push(player);
      commit("updatePlayers", players);
    },
    addExplosion({state, commit}, explosion){
      let explosions = state.explosions
      explosions.push(explosion)
      commit('updateExplosions', explosions)
    }
  },
};
