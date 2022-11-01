/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import {store} from "../app"

import { MonsterSprite } from "../components/MonsterSprite"
import { BombSprite } from "../components/BombSprite";
// import { FireSprite } from "../components/FireSprite";

export function HomeView() {  
    
    let bombDrop = store.state.bomb.drop;
    return {
        template: (
            <div id="home">
                <h1>Hello monster</h1>
                <MonsterSprite/>
                {bombDrop && <BombSprite/>}
                {/* <FireSprite/> */}
            </div>
        )
    }
}