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
        {/* <h1>Queue</h1> */}
        {/* <h3>Users joined: {String(userQueueCount)}</h3> */}
        {/* <button onClick={startGame}>Start Game</button> */}
        {/* <h2>
          Game starts in<span id="timer">10</span>
        </h2>
        <div id="lobby-players">
          <h3>Waiting players to join...</h3>
          <div id="lobby-players_list">
            <div>
              <div class="player-avatar"></div>
              <p>Player 1</p>
            </div>
            <div>
              <div class="player-avatar"></div>
              <p>Player 2</p>
            </div>
            <div>
              <div class="player-avatar"></div>
              <p>Player 3</p>
            </div>
          </div>

          <button class="btn">Leave Lobby</button>
        </div> */}
        <ChatRoom />
      </div>
    ),
  };
}
