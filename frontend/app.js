import createRouter from "./framework/router";
import createStore from "./framework/store";
import routes from "./router";
import storeObj from "./store/index";
import { SendWsMessage } from "./websocket";
import { reactive } from "./framework/reactive";
const store = createStore(storeObj);
const router = createRouter(routes);
let gameStarted = false;

// add keyup and KeyDown listeners
export function setupGame() {
  let throttleKeyDown = throttle((e) => {
    store.dispatch("registerKeyDown", e.code);
  });
  document.addEventListener("keyup", (e) => {
    store.dispatch("registerKeyUp", e.code);
  });
  document.addEventListener("keydown", (e) => {
    if (e.code == "Space") {
      throttleKeyDown(e);
    } else {
      store.dispatch("registerKeyDown", e.code);
    }
  });
  if (!gameStarted) {
    gameStarted = true;
    animate();
  }
}

// main game loop
let lastUpdate;

function animate(timestamp) {
  let delta = lastUpdate ? (timestamp - lastUpdate) / 10 : 0; //approx 1.6 from 16 ms

  let movement = store.state.movement;
  let currentPlayerName = store.state.currentPlayerName;

  // sends all 4 movements
  if (movement.move) {
    SendWsMessage("PLAYER_MOVE", currentPlayerName, movement.move, delta);
  } else if (movement.stop) {
    SendWsMessage("PLAYER_MOVE", currentPlayerName, movement.stop);
    //send movement stop only once, so clear the variable after sending
    store.dispatch("clearStopMovement");
  }
  if (movement.bomb) {
    SendWsMessage("PLAYER_DROPPED_BOMB", currentPlayerName);
    store.dispatch("clearBombDrop");
  }

  lastUpdate = timestamp;
  requestAnimationFrame(animate);
}

export { store, router };

function throttle(cb, delay = 250) {
  let shouldWait = false;

  return (...args) => {
    if (shouldWait) return;

    cb(...args);
    shouldWait = true;
    setTimeout(() => {
      shouldWait = false;
    }, delay);
  };
}
