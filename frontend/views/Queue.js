/* @jsx jsx */
import jsx from "../framework/vDom/jsx";
import { store } from "../app";
import { ChatRoom } from "../components/ChatRoom";
import { ws } from "../websocket";
function startGame() {
  ws.send(
    JSON.stringify({
      type: "START_GAME",
    })
  );
}

export function QueueView() {
  let userQueueCount = store.state.userQueueCount;

  return {
    template: (
      <div id="lobby-layout">
        <h3>Users joined: {String(userQueueCount)}</h3>
        <button onClick={startGame}>Start Game</button>
        <h2>
          {/* Game starts in<span id="timer">10</span> */}
          Need at least 2 players to start timer
        </h2>
        <div id="lobby-players">
          <h3>Waiting players to join...</h3>
          <ul id="lobby-players_list">
            <li>
              <div class="player-avatar" id="monster-1"></div>
              <p>Player 1</p>
            </li>

            <li>
              <div class="player-avatar" id="monster-1"></div>
              <p>Player 2</p>
            </li>

            <li>
              <div class="player-avatar" id="monster-1"></div>
              <p>Player 3</p>
            </li>
          </ul>

          <button class="btn">Leave Lobby</button>
        </div>
        <ChatRoom />
      </div>
    ),
  };
}
