export default {
  state: {
    players: [],
    setup: {
      stagger: 5,
      frameCount: 5,
    },
    monsterStates: {
      ArrowLeft: 0,
      ArrowDown: 1,
      ArrowRight: 2,
      ArrowUp: 3,
    },
  },
  mutations: {
    updatePlayers(state, players) {
      state.players = players;
    },
    updateCurrentPlayerIndex(state, index) {
      state.currentPlayerIndex = index;
    },
  },
  actions: {
    registerPlayer({ state, commit }, player) {
      let players = state.players;
      players.push(player);
      commit("updatePlayers", players);
    },

    movePlayerDown({ state, commit }, { index, gameFrame }) {
      let players = state.players;
      players[index].y = players[index].y + players[index].speed;
      players[index].state = 1;

      if (gameFrame % state.setup.stagger == 0) {
        players[index].frame =
          players[index].frame >= players[index].frameCount
            ? 0
            : (players[index].frame += 1);
      }
      commit("updatePlayers", players);
    },
  },
};

// save player objects in case

// {
//         id: 0,
//         name: "",
//         state: 0,
//         frame: 0,
//         x: 0,
//         y: 0,
//         speed: 1,
//       },
//       {
//         id: 1,
//         name: "",
//         state: 0,
//         frame: 0,
//         x: 768,
//         y: 0,
//         speed: 1,
//       },
//       {
//         id: 2,
//         name: "",
//         state: 0,
//         frame: 0,
//         x: 0,
//         y: 768,
//         speed: 1,
//       },
//       {
//         id: 3,
//         name: "",
//         state: 0,
//         frame: 0,
//         x: 768,
//         y: 768,
//         speed: 1,
//       },

// registerCurrentPlayer({ state, commit }, name) {
//   let currentPlayerIndex = state.players.findIndex(
//     (player) => player.name === name
//   );

//   commit("updateCurrentPlayerIndex", currentPlayerIndex);
// },
// movePlayerLeft({ state, commit }, { index, gameFrame }) {
//   let players = state.players;
//   players[index].x = players[index].x - players[index].speed;
//   players[index].state = 0;

//   if (gameFrame % state.setup.stagger == 0) {
//     players[index].frame =
//       players[index].frame >= players[index].frameCount
//         ? 0
//         : (players[index].frame += 1);
//   }
//   commit("updatePlayers", players);
// },
// movePlayerRight({ state, commit }, data) {
//   let players = state.players;
//   let index = Number(data.creator);
//   players[index].x = data.players[index].x;
//   players[index].state = 2;

//   if (gameFrame % state.setup.stagger == 0) {
//     players[index].frame =
//       players[index].frame >= players[index].frameCount
//         ? 0
//         : (players[index].frame += 1);
//   }
//   commit("updatePlayers", players);
// },
// movePlayerUp({ state, commit }, { index, gameFrame }) {
//   let players = state.players;
//   players[index].y = players[index].y - players[index].speed;
//   players[index].state = 3;

//   if (gameFrame % state.setup.stagger == 0) {
//     players[index].frame =
//       players[index].frame >= players[index].frameCount
//         ? 0
//         : (players[index].frame += 1);
//   }
//   commit("updatePlayers", players);
// },
