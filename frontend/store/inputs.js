export default {
  state: {
    inputs: {},
    movement: { move: "", stop: "", bombDropped: false },
  },
  mutations: {
    updateInputs(state, inputs) {
      state.inputs = inputs;
    },
    updateMovement(state, movement) {
      state.movement = movement;
    },
  },
  actions: {
    registerKeyUp({ state, commit }, key) {
      let inputs = state.inputs;
      inputs[key] = false;

      let movement = state.movement;
      if (key == "ArrowLeft") {
        if (movement.move == "LEFT") {
          movement.move = "";
          movement.stop = "LEFT-STOP";
        }
      } else if (key == "ArrowDown") {
        if (movement.move == "DOWN") {
          movement.move = "";
          movement.stop = "DOWN-STOP";
        }
      } else if (key == "ArrowRight") {
        if (movement.move == "RIGHT") {
          movement.move = "";
          movement.stop = "RIGHT-STOP";
        }
      } else if (key == "ArrowUp") {
        if (movement.move == "UP") {
          movement.move = "";
          movement.stop = "UP-STOP";
        }
      }

      commit("updateInputs", inputs);
      commit("updateMovement", movement);
    },
    registerKeyDown({ state, commit }, key) {
      let inputs = state.inputs;
      inputs[key] = true;

      let movement = state.movement;
      if (key == "ArrowLeft") {
        movement.stop = "";
        movement.move = "LEFT";
      } else if (key == "ArrowDown") {
        movement.stop = "";
        movement.move = "DOWN";
      } else if (key == "ArrowRight") {
        movement.stop = "";
        movement.move = "RIGHT";
      } else if (key == "ArrowUp") {
        movement.stop = "";
        movement.move = "UP";
      }
      commit("updateInputs", inputs);
      commit("updateMovement", movement);
    },
    clearStopMovement({ state, commit }) {
      let movement = state.movement;
      movement.stop = "";
      commit("updateMovement", movement);
    },
  },
};
/* ------------------------- Input plan for movement ------------------------ */
// Input is registered inn {"Key":true}
// set true when key down
// set false when key up

// If key down == one of the movements
// let variable movement.move = movement name

// on key up if on of the movements
// check if last registered movement is the one that we want to stop
// if yes
//  let variable movement = ""
// and set variable movement.stop = stop movement
// why? if movement we want to stop != movement last registered,
// the user pressed next movement before letting go of previos,
// so stopping shouldn't be registerd in this case

// On animation loop we check if movement has value
// if has, send it
// if not check if we have stop movement
// if has send and clear the stop movement, so it gets sent only once
