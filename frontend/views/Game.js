/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { MonsterSprite } from "../components/MonsterSprite";
import { store } from "../app";
// import { players } from "../websocket";
export function GameView() {
    let players = store.state.players
    return {
        template: (
            <div id="home">
                <h1>Game View</h1>
                {players.map(player => <MonsterSprite player={player} />)}
            </div >
        )
    }
}
